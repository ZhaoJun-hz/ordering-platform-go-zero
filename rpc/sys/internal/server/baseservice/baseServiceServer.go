// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package server

import (
	"context"

	"ordering-platform/rpc/sys/internal/logic/baseservice"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"
)

type BaseServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedBaseServiceServer
}

func NewBaseServiceServer(svcCtx *svc.ServiceContext) *BaseServiceServer {
	return &BaseServiceServer{
		svcCtx: svcCtx,
	}
}

// 初始化Api
func (s *BaseServiceServer) InitApi(ctx context.Context, in *sysclient.InitApiReq) (*sysclient.InitApiResp, error) {
	l := baseservicelogic.NewInitApiLogic(ctx, s.svcCtx)
	return l.InitApi(in)
}
