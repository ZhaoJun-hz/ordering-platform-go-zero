package menu

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.MenuService.UpdateMenu(l.ctx, &sysclient.UpdateMenuReq{
		MenuId:          req.MenuId,
		MenuType:        req.MenuType,
		Title:           req.Title,
		Sort:            req.Sort,
		ParentMenuId:    req.ParentMenuId,
		Icon:            req.Icon,
		Name:            req.Name,
		Component:       req.Component,
		Path:            req.Path,
		Permission:      req.Permission,
		HideInMenu:      req.HideInMenu,
		IgnoreKeepAlive: req.IgnoreKeepAlive,
		LinkFlag:        req.LinkFlag,
		Link:            req.Link,
		Disabled:        req.Disabled,
		SelectApi:       req.SelectApi,
	})

	if err != nil {
		return nil, err
	}
	return
}
