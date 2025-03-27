package deptservicelogic

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

type DeptInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptInfoLogic {
	return &DeptInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptInfoLogic) DeptInfo(in *sysclient.DeptInfoReq) (*sysclient.DeptInfoResp, error) {
	// todo: add your logic here and delete this line

	// 校验一下dept 是否存在
	sysDept, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.DeptID.Eq(in.DeptId)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "Dept不存在,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.WithStack(errcode.DeptNotExistError)
	}
	if err != nil {
		logc.Errorf(l.ctx, "查询父 Dept失败,参数：%d,异常:%s", in.DeptId, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询父 Dept失败 %v, param %v", err, in.DeptId)
	}
	return &sysclient.DeptInfoResp{
		DeptId:       sysDept.DeptID,
		ParentDeptId: sysDept.ParentID,
		DeptName:     sysDept.DeptName,
		Sort:         sysDept.Sort,
		Leader:       sysDept.Leader,
		Phone:        sysDept.Phone,
		Email:        sysDept.Email,
		Status:       sysDept.Status,
		CreateTime:   utils.TimeToString(&sysDept.CreatedAt),
	}, nil
}
