package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除user
func (l *UserDeleteLogic) UserDelete(in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {
	// todo: add your logic here and delete this line
	_, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.UserID.Eq(in.UserId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.WithStack(errcode.UserNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.UserId)
	}
	_, err = query.SysUser.WithContext(l.ctx).Where(query.SysUser.UserID.Eq(in.UserId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 SysUser 失败,参数：%d,异常:%s", in.UserId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 SysUser 失败 %v, 参数 %d", err, in.UserId)
	}
	return &sysclient.UserDeleteResp{}, nil
}
