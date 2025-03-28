package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/encrypt"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新user
func (l *UserUpdateLogic) UserUpdate(in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {
	// todo: add your logic here and delete this line
	// role 是否存在
	_, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.RoleID.Eq(in.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.WithStack(errcode.RoleNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Role失败 %v, param %v", err, in.RoleId)
	}
	// dept 是否存在
	_, err = query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.DeptId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Dept不存在,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.WithStack(errcode.DeptNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Dept失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept失败 %v, param %v", err, in.DeptId)
	}
	// user 是否存在
	_, err = query.SysUser.WithContext(l.ctx).Where(query.SysUser.UserID.Eq(in.UserId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "User 不存在,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.WithStack(errcode.UserNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 User失败,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept失败 %v, param %v", err, in.UserId)
	}
	user := &model.SysUser{
		UserID:      in.UserId,
		Status:      in.Status,
		Username:    in.Username,
		Nickname:    in.Nickname,
		Description: in.Description,
		Mobile:      in.Mobile,
		Email:       in.Email,
		Avatar:      in.Avatar,
		DeptID:      in.DeptId,
		RoleID:      in.RoleId,
	}
	if in.Password != "" {
		// password 加密
		passwordEncrypt := encrypt.BcryptEncrypt(in.Password)
		user.Password = passwordEncrypt
	}

	_, err = query.SysUser.WithContext(l.ctx).Updates(user)
	if err != nil {
		logc.Errorf(l.ctx, "更新 user 失败,参数：%+v,异常:%s", user, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "更新 role 失败 %v, req %v", err, user)
	}

	return &sysclient.UserUpdateResp{}, nil
}
