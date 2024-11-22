package roleservicelogic

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

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *sysclient.RoleDeleteReq) (*sysclient.RoleDeleteResp, error) {
	// todo: add your logic here and delete this line
	// 校验一下 Role 是否存在
	role, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.RoleID.Eq(in.RoleId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Role不存在,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.WithStack(errcode.RoleNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Menu失败 %v, param %v", err, in.RoleId)
	}
	// 校验一下 Role 下是否还有用户
	count, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.RoleID.Eq(in.RoleId)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Role 是否关联用户 失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Role 是否关联用户 失败 %v, param %v", err, in.RoleId)
	}
	if count > 0 {
		logc.Errorf(l.ctx, "Role 关联的有用户,不允许操作")
		return nil, errors.WithStack(errcode.RoleHaveUserError)
	}
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 删Role_Menu
	_, err = tx.SysRoleMenu.WithContext(l.ctx).Where(query.SysRoleMenu.RoleID.Eq(in.RoleId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 SysRoleMenu 失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 SysRoleMenu 失败 %v, 参数 %d", err, in.RoleId)
	}

	// 删Role
	_, err = tx.SysRole.WithContext(l.ctx).Where(query.SysRole.RoleID.Eq(in.RoleId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 SysRole 失败,参数：%d,异常:%s", in.RoleId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 SysRole 失败 %v, 参数 %d", err, in.RoleId)
	}
	return &sysclient.RoleDeleteResp{
		RoleKey: role.RoleKey,
	}, nil
}
