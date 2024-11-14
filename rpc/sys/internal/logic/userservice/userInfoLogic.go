package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"
	"strconv"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户个人信息
func (l *UserInfoLogic) UserInfo(in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	// todo: add your logic here and delete this line

	//1.根据id查询用户信息
	q := query.SysUser
	info, err := q.WithContext(l.ctx).Where(q.UserID.Eq(in.UserId)).First()

	// 2.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.WithStack(errcode.UserNotExistError)
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户信息 %v, req %v", err, in)
	}

	// 3.查询role
	roleQuery := query.SysRole
	sysRole, err := roleQuery.WithContext(l.ctx).Where(roleQuery.RoleID.Eq(info.RoleID)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "角色不存在,参数：%d,异常:%s", info.RoleID, err.Error())
		return nil, errors.WithStack(errcode.RoleNotExistError)
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%d,异常:%s", info.RoleID, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户信息 %v, req %v", err, in)
	}

	return &sysclient.InfoResp{
		Avatar:   info.Avatar,
		Username: info.Username,
		Roles:    []string{sysRole.RoleName},
		UserId:   strconv.FormatInt(info.UserID, 10),
		Desc:     info.Description,
		HomePath: sysRole.DefaultRouter,
	}, nil
}
