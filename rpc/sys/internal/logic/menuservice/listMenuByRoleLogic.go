package menuservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMenuByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMenuByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuByRoleLogic {
	return &ListMenuByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMenuByRoleLogic) ListMenuByRole(in *sysclient.ListMenuRoleReq) (*sysclient.ListMenuResp, error) {
	// todo: add your logic here and delete this line
	// 获取 role 绑定的 菜单

	menu := query.SysMenu
	roleMenu := query.SysRoleMenu

	roleMenus, err := roleMenu.WithContext(l.ctx).Where(roleMenu.RoleID.Eq(in.RoleId)).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 RoleMenu 失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 RoleMenu 失败 %v, param %v", err, in.RoleId)
	}
	var menuIds []int64
	for _, tempMenu := range roleMenus {
		menuIds = append(menuIds, tempMenu.MenuID)
	}
	var result []*sysclient.ListMenuData

	if len(menuIds) > 0 {
		menus, err := menu.WithContext(l.ctx).Where(menu.MenuID.In(menuIds...)).Order(query.SysMenu.Sort).Find()
		if err != nil {
			logc.Errorf(l.ctx, "查询 Menu 失败,参数：%v,异常:%s", menuIds, err.Error())
			return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu 失败 %v, menuIds %v", err, menuIds)
		}
		for _, menu := range menus {
			result = append(result, &sysclient.ListMenuData{
				MenuId:          menu.MenuID,
				MenuType:        menu.MenuType,
				Title:           menu.Title,
				Sort:            menu.Sort,
				ParentMenuId:    menu.ParentID,
				Icon:            menu.Icon,
				Name:            menu.Name,
				Component:       menu.Component,
				Path:            menu.Path,
				Permission:      menu.Permission,
				HideInMenu:      menu.HideInMenu,
				IgnoreKeepAlive: menu.IgnoreKeepAlive,
				LinkFlag:        menu.LinkFlag,
				Link:            menu.Link,
				Disabled:        menu.Disabled,
			})
		}
	}

	return &sysclient.ListMenuResp{
		Data: result,
	}, nil
}
