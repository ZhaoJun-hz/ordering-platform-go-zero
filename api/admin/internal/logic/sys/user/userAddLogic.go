package user

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user添加
func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddReq) (resp *types.UserAddResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserService.UserAdd(l.ctx, &sysclient.UserAddReq{
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
