package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"ordering-platform/api/admin/internal/svc"
)

type UserCodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCodesLogic {
	return &UserCodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCodesLogic) UserCodes() (resp []string, err error) {
	// todo: add your logic here and delete this line
	return
}
