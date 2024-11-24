package user

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user更新
func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserService.UserUpdate(l.ctx, &sysclient.UserUpdateReq{
		UserId:      req.UserId,
		Status:      req.Status,
		Username:    req.Username,
		Password:    req.Password,
		Nickname:    req.Nickname,
		Description: req.Description,
		Mobile:      req.Mobile,
		Email:       req.Email,
		Avatar:      req.Avatar,
		DeptId:      req.DeptId,
		RoleId:      req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return
}
