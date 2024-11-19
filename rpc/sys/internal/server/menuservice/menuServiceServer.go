// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package server

import (
	"context"

	"ordering-platform/rpc/sys/internal/logic/menuservice"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"
)

type MenuServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedMenuServiceServer
}

func NewMenuServiceServer(svcCtx *svc.ServiceContext) *MenuServiceServer {
	return &MenuServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MenuServiceServer) AddMenu(ctx context.Context, in *sysclient.AddMenuReq) (*sysclient.AddMenuResp, error) {
	l := menuservicelogic.NewAddMenuLogic(ctx, s.svcCtx)
	return l.AddMenu(in)
}

func (s *MenuServiceServer) UpdateMenu(ctx context.Context, in *sysclient.UpdateMenuReq) (*sysclient.UpdateMenuResp, error) {
	l := menuservicelogic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *MenuServiceServer) DeleteMenu(ctx context.Context, in *sysclient.DeleteMenuReq) (*sysclient.DeleteMenuResp, error) {
	l := menuservicelogic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *MenuServiceServer) MenuInfo(ctx context.Context, in *sysclient.MenuInfoReq) (*sysclient.MenuInfoResp, error) {
	l := menuservicelogic.NewMenuInfoLogic(ctx, s.svcCtx)
	return l.MenuInfo(in)
}

func (s *MenuServiceServer) ListMenu(ctx context.Context, in *sysclient.ListMenuReq) (*sysclient.ListMenuResp, error) {
	l := menuservicelogic.NewListMenuLogic(ctx, s.svcCtx)
	return l.ListMenu(in)
}
