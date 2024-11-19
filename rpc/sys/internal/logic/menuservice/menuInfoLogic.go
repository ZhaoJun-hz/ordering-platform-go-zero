package menuservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuInfoLogic {
	return &MenuInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuInfoLogic) MenuInfo(in *sysclient.MenuInfoReq) (*sysclient.MenuInfoResp, error) {
	// todo: add your logic here and delete this line

	var err error
	sysMenu, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.Eq(in.MenuId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Menu不存在,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.WithStack(errcode.MenuNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.MenuId)
	}

	sysMenuApis, err := query.SysMenuAPIRule.WithContext(l.ctx).Where(query.SysMenuAPIRule.SysMenuMenuID.Eq(in.MenuId)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 SysMenuAPIRule 失败,参数：%d,异常:%s", in.MenuId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 SysMenuAPIRule 失败 %v, param %v", err, in.MenuId)
	}
	var selectApiIds []int64
	for _, api := range sysMenuApis {
		selectApiIds = append(selectApiIds, api.SysAPIID)
	}

	return &sysclient.MenuInfoResp{
		MenuId:          sysMenu.MenuID,
		MenuType:        sysMenu.MenuType,
		Title:           sysMenu.Title,
		Sort:            sysMenu.Sort,
		ParentMenuId:    sysMenu.ParentMenuID,
		Icon:            sysMenu.Icon,
		Name:            sysMenu.Name,
		Component:       sysMenu.Component,
		Path:            sysMenu.Path,
		Permission:      sysMenu.Permission,
		HideInMenu:      sysMenu.HideInMenu,
		IgnoreKeepAlive: sysMenu.IgnoreKeepAlive,
		LinkFlag:        sysMenu.LinkFlag,
		Link:            sysMenu.Link,
		Disabled:        sysMenu.Disabled,
		SelectApi:       selectApiIds,
	}, nil
}
