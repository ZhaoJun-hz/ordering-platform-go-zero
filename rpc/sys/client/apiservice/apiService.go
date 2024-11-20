// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package apiservice

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

	ApiService interface {
		// 初始化Api
		InitApi(ctx context.Context, in *InitApiReq, opts ...grpc.CallOption) (*InitApiResp, error)
		ListApi(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error)
	}

	defaultApiService struct {
		cli zrpc.Client
	}
)

func NewApiService(cli zrpc.Client) ApiService {
	return &defaultApiService{
		cli: cli,
	}
}

// 初始化Api
func (m *defaultApiService) InitApi(ctx context.Context, in *InitApiReq, opts ...grpc.CallOption) (*InitApiResp, error) {
	client := sysclient.NewApiServiceClient(m.cli.Conn())
	return client.InitApi(ctx, in, opts...)
}

func (m *defaultApiService) ListApi(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	client := sysclient.NewApiServiceClient(m.cli.Conn())
	return client.ListApi(ctx, in, opts...)
}