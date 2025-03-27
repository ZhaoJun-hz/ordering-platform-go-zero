package user

import (
	"context"
	"ordering-platform/pkg/encrypt"
	"ordering-platform/pkg/enum/common"
	"ordering-platform/pkg/utils/jwt"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/sysclient"
	"time"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserService.GetUserByUsername(l.ctx, &sysclient.UsernameReq{Username: req.Username})

	if err != nil {

		return nil, err
	}

	if user.Status != int32(common.StatusNormal) {
		return nil, xerr.NewCodeInvalidArgumentError("login.userBanned")
	}

	if !encrypt.BcryptCheck(req.Password, *user.Password) {
		return nil, xerr.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", user.UserId), jwt.WithOption("roleId",
			user.RoleId), jwt.WithOption("deptId", user.DeptId))
	if err != nil {
		return nil, err
	}

	// 将 token 存储到数据库
	expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
	_, err = l.svcCtx.TokenService.CreateToken(l.ctx, &sysclient.CreateTokenReq{
		Status:    int32(common.StatusNormal),
		UserId:    user.UserId,
		Username:  user.Username,
		Token:     token,
		ExpiredAt: expiredAt,
	})

	return &types.LoginResp{
		AccessToken: token,
	}, nil
}
