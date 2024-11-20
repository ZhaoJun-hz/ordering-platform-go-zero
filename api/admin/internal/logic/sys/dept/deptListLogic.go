package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// dept列表
func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req *types.DeptListReq) (resp *types.DeptListResp, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.DeptService.DeptList(l.ctx, &sysclient.DeptListReq{})
	if err != nil {
		return nil, err
	}
	var result []*types.DeptInfoData
	for _, item := range list.Data {
		result = append(result, &types.DeptInfoData{
			DeptId:       item.DeptId,
			ParentDeptId: item.ParentDeptId,
			DeptName:     item.DeptName,
			Sort:         item.Sort,
			Leader:       item.Leader,
			Phone:        item.Phone,
			Email:        item.Email,
			Status:       item.Status,
			CreateTime:   item.CreateTime,
		})
	}

	return &types.DeptListResp{Data: result}, nil
}
