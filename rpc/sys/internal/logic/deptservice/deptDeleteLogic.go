package deptservicelogic

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

type DeptDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptDeleteLogic) DeptDelete(in *sysclient.DeptDeleteReq) (*sysclient.DeptDeleteResp, error) {
	// todo: add your logic here and delete this line

	// 校验一下dept 是否存在
	_, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.DeptId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Dept不存在,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.WithStack(errcode.DeptNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询父 Dept失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询父 Dept失败 %v, param %v", err, in.DeptId)
	}

	// 校验一下 dept 是否还有绑定人员，有则不允许删除
	count, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.DeptID.Eq(in.DeptId)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Dept 是否关联用户 失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept 是否关联用户 失败 %v, param %v", err, in.DeptId)
	}
	if count > 0 {
		logc.Errorf(l.ctx, "Dept 关联的有用户,不允许操作")
		return nil, errors.WithStack(errcode.DeptNotEmptyError)
	}
	// 校验一下，部门下是否还有子部门
	deptCount, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.ParentDeptID.Eq(in.DeptId)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Dept 是否有子部门 失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept 是否有子部门 失败 %v, param %v", err, in.DeptId)
	}
	if deptCount > 0 {
		logc.Errorf(l.ctx, "Dept 下有子部门,不允许操作")
		return nil, errors.WithStack(errcode.DeptHaveSubDeptError)
	}
	_, err = query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.DeptId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除 Dept 失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除 Dept 失败 %v, 参数 %d", err, in.DeptId)
	}

	return &sysclient.DeptDeleteResp{}, nil
}
