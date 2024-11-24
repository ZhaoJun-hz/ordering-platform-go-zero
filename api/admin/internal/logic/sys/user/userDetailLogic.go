package user

import (
	"context"
	"encoding/json"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录获取用户信息
func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail() (resp *types.UserDetailResp, err error) {
	// todo: add your logic here and delete this line
	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	infoResp, err := l.svcCtx.UserService.UserDetail(l.ctx, &sysclient.UserDetailReq{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserDetailResp{
		Avater:   infoResp.Avatar,
		Roles:    infoResp.Roles,
		UserId:   infoResp.UserId,
		Username: infoResp.Username,
		Desc:     infoResp.Desc,
		HomePath: infoResp.HomePath,
	}, nil
	return
}
