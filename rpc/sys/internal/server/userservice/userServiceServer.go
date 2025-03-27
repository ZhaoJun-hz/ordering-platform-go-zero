// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: sys.proto

package server

import (
	"context"

	"ordering-platform/rpc/sys/internal/logic/userservice"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

// 登录获取用户个人信息
func (s *UserServiceServer) UserDetail(ctx context.Context, in *sysclient.UserDetailReq) (*sysclient.UserDetailResp, error) {
	l := userservicelogic.NewUserDetailLogic(ctx, s.svcCtx)
	return l.UserDetail(in)
}

// 添加user
func (s *UserServiceServer) UserAdd(ctx context.Context, in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	l := userservicelogic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

// 更新user
func (s *UserServiceServer) UserUpdate(ctx context.Context, in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {
	l := userservicelogic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

// 删除user
func (s *UserServiceServer) UserDelete(ctx context.Context, in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {
	l := userservicelogic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

// user详情
func (s *UserServiceServer) UserInfo(ctx context.Context, in *sysclient.UserInfoReq) (*sysclient.UserInfoResp, error) {
	l := userservicelogic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

// user列表
func (s *UserServiceServer) UserList(ctx context.Context, in *sysclient.UserListReq) (*sysclient.UserListResp, error) {
	l := userservicelogic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

// 根据用户名获取用户
func (s *UserServiceServer) GetUserByUsername(ctx context.Context, in *sysclient.UsernameReq) (*sysclient.UserInfoResp, error) {
	l := userservicelogic.NewGetUserByUsernameLogic(ctx, s.svcCtx)
	return l.GetUserByUsername(in)
}
