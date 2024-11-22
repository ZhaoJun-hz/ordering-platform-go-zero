package menuservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *sysclient.UpdateMenuReq) (*sysclient.UpdateMenuResp, error) {
	// todo: add your logic here and delete this line
	var err error

	// 校验API 是否全部存在
	apis, err := query.SysAPI.WithContext(l.ctx).Where(query.SysAPI.ID.In(in.SelectApi...)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询api 列表失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询api 列表失败 %v, req %v", err, in)
	}
	if len(apis) != len(in.SelectApi) {
		logc.Errorf(l.ctx, "查询api 列表数量不匹配,参数：%+v", in)
		return nil, errors.WithStack(errcode.ApiSelectError)
	}
	// 校验menuId 是否存在
	sysMenu, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.Eq(in.MenuId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Menu不存在,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.WithStack(errcode.MenuNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.MenuId)
	}
	if sysMenu.MenuType != in.MenuType {
		logc.Errorf(l.ctx, "更新 Menu失败,不允许更改菜单类型")
		return nil, errors.WithStack(errcode.MenuUpdateMenuTypeError)
	}

	// 校验父menu 是否存在 且类型是否正确
	if in.ParentMenuId != 0 {
		parentMenu, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.Eq(in.ParentMenuId)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logc.Errorf(l.ctx, "Menu不存在,参数：%d,异常:%s", in.ParentMenuId, err.Error())
			return nil, errors.WithStack(errcode.MenuNotExistError)
		}
		if err != nil {
			logc.Errorf(l.ctx, "查询父 Menu失败,参数：%d,异常:%s", in.ParentMenuId, err.Error())
			return nil, errors.Wrapf(xerr.NewDBErr(), "查询父 Menu失败 %v, param %v", err, in.ParentMenuId)
		}
		if parentMenu.MenuType == "M" && in.MenuType != "C" {
			logc.Errorf(l.ctx, "新增Menu MenuType Error ,parentMenu.MenuType %s in.MenuType %s", parentMenu.MenuType, in.MenuType)
			return nil, errors.WithStack(errcode.MenuTypeError)
		}
		if parentMenu.MenuType == "C" && in.MenuType != "F" {
			logc.Errorf(l.ctx, "新增Menu MenuType Error ,parentMenu.MenuType %s in.MenuType %s", parentMenu.MenuType, in.MenuType)
			return nil, errors.WithStack(errcode.MenuTypeError)
		}
	} else {
		// 没有 父Menu，那么只能是M 目录
		if in.MenuType != "M" {
			logc.Errorf(l.ctx, "新增Menu MenuType Error ,参数：%+v", in)
			return nil, errors.WithStack(errcode.MenuTypeError)
		}
	}

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 添加 menu
	newSysMenu := &model.SysMenu{
		MenuID:          in.MenuId,
		ParentMenuID:    in.ParentMenuId,
		Sort:            in.Sort,
		MenuType:        in.MenuType,
		Path:            in.Path,
		Component:       in.Component,
		Permission:      in.Permission,
		Name:            in.Name,
		Title:           in.Title,
		Icon:            in.Icon,
		HideInMenu:      in.HideInMenu,
		IgnoreKeepAlive: in.IgnoreKeepAlive,
		LinkFlag:        in.LinkFlag,
		Link:            in.Link,
		Disabled:        in.Disabled,
		DeletedAt:       gorm.DeletedAt{},
	}
	// 更新sysmenu
	_, err = tx.SysMenu.WithContext(l.ctx).Updates(newSysMenu)
	if err != nil {
		logc.Errorf(l.ctx, "更新 SysMenu 失败,参数：%+v,异常:%s", newSysMenu, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "更新 SysMenu 失败 %v, 参数 %v", err, newSysMenu)
	}

	// 删除 sys_menu_api_rule
	_, err = tx.SysMenuAPI.WithContext(l.ctx).Where(query.SysMenuAPI.SysMenuID.Eq(in.MenuId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 sys_menu_api_rule 失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 sys_menu_api_rule 失败 %v, 参数 %d", err, in.MenuId)
	}
	// 添加，重新生成 sys_menu_api_rule
	list := []*model.SysMenuAPI{}
	for _, item := range in.SelectApi {
		list = append(list, &model.SysMenuAPI{
			SysMenuID: sysMenu.MenuID,
			SysAPIID:  item,
		})
	}
	err = tx.SysMenuAPI.WithContext(l.ctx).CreateInBatches(list, len(list))
	if err != nil {
		logc.Errorf(l.ctx, "批量创建 SysMenuAPIRule 失败,参数：%+v,异常:%s", list, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "批量创建 SysMenuAPIRule 失败 %v, 参数 %v", err, list)
	}
	return &sysclient.UpdateMenuResp{}, nil
}
