package roleservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/common"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {
	// todo: add your logic here and delete this line
	var err error
	// 更新的Role 是否存在
	role, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.RoleID.Eq(in.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.WithStack(errcode.RoleNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.RoleId)
	}
	sysMenus, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuID.In(in.SelectMenus...)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 menu 列表失败,参数：%+v,异常:%s", in.SelectMenus, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 menu 列表失败 %v, req %v", err, in.SelectMenus)
	}
	sysMenuIds := make([]int64, 0)
	for _, menu := range sysMenus {
		sysMenuIds = append(sysMenuIds, menu.MenuID)
	}
	menuApis, err := query.SysMenuAPI.WithContext(l.ctx).Where(query.SysMenuAPI.SysMenuID.In(sysMenuIds...)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 menu_api 列表失败,参数：%+v,异常:%s", sysMenuIds, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 menu_api 列表失败 %v, req %v", err, sysMenuIds)
	}

	apiIds := make([]int64, 0)
	for _, menuApi := range menuApis {
		apiIds = append(apiIds, menuApi.SysAPIID)
	}

	sysAPIS, err := query.SysAPI.WithContext(l.ctx).Where(query.SysAPI.ID.In(apiIds...)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 api 列表失败,参数：%+v,异常:%s", apiIds, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 menu_api 列表失败 %v, req %v", err, apiIds)
	}

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 创建SysRole
	sysRole := &model.SysRole{
		RoleID:        in.RoleId,
		RoleName:      in.RoleName,
		Status:        in.Status,
		Sort:          in.Sort,
		Admin:         false,
		DefaultRouter: in.DefaultRouter,
	}
	_, err = tx.SysRole.WithContext(l.ctx).Updates(sysRole)
	if err != nil {
		logc.Errorf(l.ctx, "更新 role 失败,参数：%+v,异常:%s", sysRole, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "更新 role 失败 %v, req %v", err, sysRole)
	}
	// 删除sys_role_menu
	_, err = tx.SysRoleMenu.WithContext(l.ctx).Where(query.SysRoleMenu.RoleID.Eq(in.RoleId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 SysRoleMenu 失败,参数：%+v,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 SysRoleMenu 失败 %v, req %v", err, in.RoleId)
	}
	// 添加 sys_role_menu
	list := []*model.SysRoleMenu{}
	for _, item := range in.SelectMenus {
		list = append(list, &model.SysRoleMenu{
			RoleID: sysRole.RoleID,
			MenuID: item,
		})
	}
	err = tx.SysRoleMenu.WithContext(l.ctx).CreateInBatches(list, len(list))
	if err != nil {
		logc.Errorf(l.ctx, "批量创建 SysRoleMenu 失败,参数：%+v,异常:%s", list, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "批量创建 SysRoleMenu 失败 %v, 参数 %v", err, list)
	}

	// 需要返回这些菜单关联的Api列表
	var apiList []*sysclient.ApiInfo
	for _, api := range sysAPIS {
		apiList = append(apiList, &sysclient.ApiInfo{
			Id:         api.ID,
			Handle:     api.Handle,
			Title:      api.Title,
			Path:       api.Path,
			Type:       api.Type,
			Action:     api.Action,
			CreateTime: common.TimeToString(&api.CreatedAt),
		})
	}

	return &sysclient.RoleUpdateResp{
		RoleKey: role.RoleKey,
		Data:    apiList,
	}, nil
}
