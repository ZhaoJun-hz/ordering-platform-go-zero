package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
	"ordering-platform/rpc/sys/sysclient"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoData, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.UserService.UserInfo(l.ctx, &sysclient.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoData{
		UserId:      info.UserId,
		Status:      info.Status,
		Username:    info.Username,
		Nickname:    info.Nickname,
		Description: info.Description,
		Mobile:      info.Mobile,
		Email:       info.Email,
		Avatar:      info.Avatar,
		DeptId:      info.DeptId,
		RoleId:      info.RoleId,
		CreateTime:  info.CreateTime,
	}, nil
}
