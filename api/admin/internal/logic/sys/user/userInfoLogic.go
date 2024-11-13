package user

import (
	"context"
	"encoding/json"
	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	infoResp, err := l.svcCtx.UserService.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		Avater:   infoResp.Avatar,
		Roles:    infoResp.Roles,
		UserId:   infoResp.UserId,
		Username: infoResp.Username,
		Desc:     infoResp.Desc,
		HomePath: infoResp.HomePath,
	}, nil
}
