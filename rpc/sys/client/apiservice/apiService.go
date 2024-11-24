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
	RoleAddReq       = sysclient.RoleAddReq
	RoleAddResp      = sysclient.RoleAddResp
	RoleDeleteReq    = sysclient.RoleDeleteReq
	RoleDeleteResp   = sysclient.RoleDeleteResp
	RoleInfoReq      = sysclient.RoleInfoReq
	RoleInfoResp     = sysclient.RoleInfoResp
	RoleListData     = sysclient.RoleListData
	RoleListReq      = sysclient.RoleListReq
	RoleListResp     = sysclient.RoleListResp
	RoleUpdateReq    = sysclient.RoleUpdateReq
	RoleUpdateResp   = sysclient.RoleUpdateResp
	UpdateMenuReq    = sysclient.UpdateMenuReq
	UpdateMenuResp   = sysclient.UpdateMenuResp
	UserAddReq       = sysclient.UserAddReq
	UserAddResp      = sysclient.UserAddResp
	UserDeleteReq    = sysclient.UserDeleteReq
	UserDeleteResp   = sysclient.UserDeleteResp
	UserDetailReq    = sysclient.UserDetailReq
	UserDetailResp   = sysclient.UserDetailResp
	UserInfoReq      = sysclient.UserInfoReq
	UserInfoResp     = sysclient.UserInfoResp
	UserListData     = sysclient.UserListData
	UserListReq      = sysclient.UserListReq
	UserListResp     = sysclient.UserListResp
	UserUpdateReq    = sysclient.UserUpdateReq
	UserUpdateResp   = sysclient.UserUpdateResp

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
