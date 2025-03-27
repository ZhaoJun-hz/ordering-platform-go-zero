package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/utils"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"
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
func (l *UserInfoLogic) UserInfo(in *sysclient.UserInfoReq) (*sysclient.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	sysUser, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.UserID.Eq(in.UserId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.WithStack(errcode.UserNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.UserId)
	}

	return &sysclient.UserInfoResp{
		UserId:      sysUser.UserID,
		Status:      sysUser.Status,
		Username:    sysUser.Username,
		Nickname:    sysUser.Nickname,
		Description: sysUser.Description,
		Mobile:      sysUser.Mobile,
		Email:       sysUser.Email,
		Avatar:      sysUser.Avatar,
		DeptId:      sysUser.DeptID,
		RoleId:      sysUser.RoleID,
		CreateTime:  utils.TimeToString(&sysUser.CreatedAt),
	}, nil
}
