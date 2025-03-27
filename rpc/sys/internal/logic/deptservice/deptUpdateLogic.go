package deptservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/errcode"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
	// todo: add your logic here and delete this line
	// 父parent 是否存在
	if in.ParentDeptId != 0 {
		_, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.ParentDeptId)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logc.Errorf(l.ctx, "Dept不存在,参数：%d,异常:%s", in.ParentDeptId, err.Error())
			return nil, errors.WithStack(errcode.DeptNotExistError)
		}
		if err != nil {
			logc.Errorf(l.ctx, "查询父 Dept失败,参数：%d,异常:%s", in.ParentDeptId, err.Error())
			return nil, errors.Wrapf(xerr.NewDBErr(), "查询父 Dept失败 %v, param %v", err, in.ParentDeptId)
		}
	}

	// 修改的dept 是否存在
	_, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.DeptId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Dept不存在,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.WithStack(errcode.DeptNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询 Dept失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept失败 %v, param %v", err, in.DeptId)
	}

	dept := &model.SysDept{
		DeptID:   in.DeptId,
		ParentID: in.ParentDeptId,
		DeptName: in.DeptName,
		Sort:     in.Sort,
		Leader:   in.Leader,
		Phone:    in.Phone,
		Email:    in.Email,
		Status:   in.Status,
	}
	_, err = query.SysDept.WithContext(l.ctx).Updates(dept)
	if err != nil {
		logc.Errorf(l.ctx, "更新 dept 失败,参数：%+v,异常:%s", dept, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "更新 dept 失败 %v, req %v", err, dept)
	}
	return &sysclient.DeptUpdateResp{}, nil
}
