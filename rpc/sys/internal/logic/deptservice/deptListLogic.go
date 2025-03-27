package deptservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"ordering-platform/pkg/utils"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/query"

	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sysclient.DeptListReq) (*sysclient.DeptListResp, error) {
	// todo: add your logic here and delete this line
	sysDepts, err := query.SysDept.WithContext(l.ctx).Order(query.SysDept.Sort).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Dept 列表失败, 异常:%s", err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询 Dept 列表失败 %v", err)
	}

	var result []*sysclient.DeptListData
	for _, sysDept := range sysDepts {
		result = append(result, &sysclient.DeptListData{
			DeptId:       sysDept.DeptID,
			ParentDeptId: sysDept.ParentID,
			DeptName:     sysDept.DeptName,
			Sort:         sysDept.Sort,
			Leader:       sysDept.Leader,
			Phone:        sysDept.Phone,
			Email:        sysDept.Email,
			Status:       sysDept.Status,
			CreateTime:   utils.TimeToString(&sysDept.CreatedAt),
		})
	}
	return &sysclient.DeptListResp{
		Data: result,
	}, nil
}
