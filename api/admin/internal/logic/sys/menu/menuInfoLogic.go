package menu

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuInfoLogic {
	return &MenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuInfoLogic) MenuInfo(req *types.MenuInfoReq) (resp *types.MenuInfoData, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.MenuService.MenuInfo(l.ctx, &sysclient.MenuInfoReq{
		MenuId: req.MenuId,
	})

	if err != nil {
		return nil, err
	}
	return &types.MenuInfoData{
		MenuId:          result.MenuId,
		MenuType:        result.MenuType,
		Title:           result.Title,
		Sort:            result.Sort,
		ParentMenuId:    result.ParentMenuId,
		Icon:            result.Icon,
		Name:            result.Name,
		Component:       result.Component,
		Path:            result.Path,
		Permission:      result.Permission,
		HideInMenu:      result.HideInMenu,
		IgnoreKeepAlive: result.IgnoreKeepAlive,
		LinkFlag:        result.LinkFlag,
		Link:            result.Link,
		Disabled:        result.Disabled,
		SelectApi:       result.SelectApi,
	}, nil
}
