package user

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user删除
func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserService.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return
}
