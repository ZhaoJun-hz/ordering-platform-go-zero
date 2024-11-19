package menu

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuLogic {
	return &ListMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMenuLogic) ListMenu(req *types.ListMenuReq) (resp *types.ListMenuResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.MenuService.ListMenu(l.ctx, &sysclient.ListMenuReq{})

	if err != nil {
		return nil, err
	}
	var list []*types.ListMenuData
	for _, item := range result.Data {
		list = append(list, &types.ListMenuData{
			MenuId:          item.MenuId,
			MenuType:        item.MenuType,
			Title:           item.Title,
			Sort:            item.Sort,
			ParentMenuId:    item.ParentMenuId,
			Icon:            item.Icon,
			Name:            item.Name,
			Component:       item.Component,
			Path:            item.Path,
			Permission:      item.Permission,
			HideInMenu:      item.HideInMenu,
			IgnoreKeepAlive: item.IgnoreKeepAlive,
			LinkFlag:        item.LinkFlag,
			Link:            item.Link,
			Disabled:        item.Disabled,
		})
	}
	return &types.ListMenuResp{
		Data: list,
	}, nil
}
