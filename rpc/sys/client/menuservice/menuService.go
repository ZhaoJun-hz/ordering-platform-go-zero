// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package menuservice

import (
	"context"

	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMenuReq       = sysclient.AddMenuReq
	AddMenuResp      = sysclient.AddMenuResp
	ApiInfo          = sysclient.ApiInfo
	ApiListReq       = sysclient.ApiListReq
	ApiListResp      = sysclient.ApiListResp
	DeleteMenuReq    = sysclient.DeleteMenuReq
	DeleteMenuResp   = sysclient.DeleteMenuResp
	DeptAddReq       = sysclient.DeptAddReq
	DeptAddResp      = sysclient.DeptAddResp
	DeptDeleteReq    = sysclient.DeptDeleteReq
	DeptDeleteResp   = sysclient.DeptDeleteResp
	DeptInfoReq      = sysclient.DeptInfoReq
	DeptInfoResp     = sysclient.DeptInfoResp
	DeptListData     = sysclient.DeptListData
	DeptListReq      = sysclient.DeptListReq
	DeptListResp     = sysclient.DeptListResp
	DeptUpdateReq    = sysclient.DeptUpdateReq
	DeptUpdateResp   = sysclient.DeptUpdateResp
	InfoReq          = sysclient.InfoReq
	InfoResp         = sysclient.InfoResp
	InitApiReq       = sysclient.InitApiReq
	InitApiResp      = sysclient.InitApiResp
	InitApiRouteData = sysclient.InitApiRouteData
	ListMenuData     = sysclient.ListMenuData
	ListMenuReq      = sysclient.ListMenuReq
	ListMenuResp     = sysclient.ListMenuResp
	LoginReq         = sysclient.LoginReq
	LoginResp        = sysclient.LoginResp
	MenuInfoReq      = sysclient.MenuInfoReq
	MenuInfoResp     = sysclient.MenuInfoResp
	UpdateMenuReq    = sysclient.UpdateMenuReq
	UpdateMenuResp   = sysclient.UpdateMenuResp

	MenuService interface {
		AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc.CallOption) (*AddMenuResp, error)
		UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuResp, error)
		DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuResp, error)
		MenuInfo(ctx context.Context, in *MenuInfoReq, opts ...grpc.CallOption) (*MenuInfoResp, error)
		ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuResp, error)
	}

	defaultMenuService struct {
		cli zrpc.Client
	}
)

func NewMenuService(cli zrpc.Client) MenuService {
	return &defaultMenuService{
		cli: cli,
	}
}

func (m *defaultMenuService) AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc.CallOption) (*AddMenuResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.AddMenu(ctx, in, opts...)
}

func (m *defaultMenuService) UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultMenuService) DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

func (m *defaultMenuService) MenuInfo(ctx context.Context, in *MenuInfoReq, opts ...grpc.CallOption) (*MenuInfoResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.MenuInfo(ctx, in, opts...)
}

func (m *defaultMenuService) ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.ListMenu(ctx, in, opts...)
}
