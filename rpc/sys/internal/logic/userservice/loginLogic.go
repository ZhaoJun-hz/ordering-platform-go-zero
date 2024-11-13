package userservicelogic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/encrypt"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"
	"time"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	// todo: add your logic here and delete this line
	q := query.SysUser
	user, err := q.WithContext(l.ctx).Where(q.Username.Eq(in.Username)).First()
	// 1.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.WithStack(errcode.UserPwdError)
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户信息 %v, req %v", err, in)
	}

	// 2.判断密码是否正确
	if !encrypt.BcryptCheck(in.Password, user.Password) {
		logc.Errorf(l.ctx, "用户密码不正确,参数:%s", in.Password)
		return nil, errors.WithStack(errcode.UserPwdError)
	}

	// 3. 校验用户状态
	if *user.Status != 1 {
		logc.Errorf(l.ctx, "用户状态不正确,status:%d", *user.Status)
		return nil, errors.WithStack(errcode.UserStatusError)
	}

	// 4.生成token
	jwtToken, err := l.getJwtToken(user.UserID, user.Username)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewInternalErr(), "生成token失败 err %v", err)
	}

	return &sysclient.LoginResp{
		Id:          user.UserID,
		AccessToken: jwtToken,
	}, nil

}

// 生成jwt的token
func (l *LoginLogic) getJwtToken(userId int64, userName string) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userId"] = userId
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(accessSecret))
}
