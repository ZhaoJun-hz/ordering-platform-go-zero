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

type ListMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuLogic {
	return &ListMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMenuLogic) ListMenu(in *sysclient.ListMenuReq) (*sysclient.ListMenuResp, error) {
	// todo: add your logic here and delete this line

	menus, err := query.SysMenu.WithContext(l.ctx).Order(query.SysMenu.Sort).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Menu 列表失败, 异常:%s", err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu 列表失败 %v", err)
	}

	var result []*sysclient.ListMenuData
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
	return &sysclient.ListMenuResp{
		Data: result,
	}, nil
}
