// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package baseservice

import (
	"context"

	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq          = sysclient.InfoReq
	InfoResp         = sysclient.InfoResp
	InitApiReq       = sysclient.InitApiReq
	InitApiResp      = sysclient.InitApiResp
	InitApiRouteData = sysclient.InitApiRouteData
	LoginReq         = sysclient.LoginReq
	LoginResp        = sysclient.LoginResp

	BaseService interface {
		// 初始化Api
		InitApi(ctx context.Context, in *InitApiReq, opts ...grpc.CallOption) (*InitApiResp, error)
	}

	defaultBaseService struct {
		cli zrpc.Client
	}
)

func NewBaseService(cli zrpc.Client) BaseService {
	return &defaultBaseService{
		cli: cli,
	}
}

// 初始化Api
func (m *defaultBaseService) InitApi(ctx context.Context, in *InitApiReq, opts ...grpc.CallOption) (*InitApiResp, error) {
	client := sysclient.NewBaseServiceClient(m.cli.Conn())
	return client.InitApi(ctx, in, opts...)
}
