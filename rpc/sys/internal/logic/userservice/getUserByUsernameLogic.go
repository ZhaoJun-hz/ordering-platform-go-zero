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

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户名获取用户
func (l *GetUserByUsernameLogic) GetUserByUsername(in *sysclient.UsernameReq) (*sysclient.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	sysUser, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.Username.Eq(in.Username)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "User 不存在,参数：%s,异常:%s", in.Username, err.Error())
		return nil, errors.WithStack(errcode.UserNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 User 失败,参数：%s,异常:%s", in.Username, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 User 失败 %v, param %v", err, in.Username)
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
		Password:    &sysUser.Password,
	}, nil
}
