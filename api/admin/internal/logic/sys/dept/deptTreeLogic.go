package dept

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取 Dept tree 结构
func NewDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptTreeLogic {
	return &DeptTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptTreeLogic) DeptTree(req *types.DeptTreeReq) (resp *types.DeptTreeResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.DeptService.DeptList(l.ctx, &sysclient.DeptListReq{})
	if err != nil {
		return nil, err
	}
	var treeData = make(map[int64]*types.DeptTreeData)
	for _, data := range result.Data {
		treeData[data.DeptId] = &types.DeptTreeData{
			DeptId:       data.DeptId,
			Value:        data.DeptId,
			ParentDeptId: data.ParentDeptId,
			DeptName:     data.DeptName,
			Sort:         data.Sort,
			Leader:       data.Leader,
			Phone:        data.Phone,
			Email:        data.Email,
			Status:       data.Status,
			CreateTime:   data.CreateTime,
			Children:     nil,
		}
	}

	var tree []*types.DeptTreeData
	for _, node := range treeData {
		if node.ParentDeptId == 0 {
			tree = append(tree, node)
		} else {
			if parent, found := treeData[node.ParentDeptId]; found {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return &types.DeptTreeResp{
		Data: tree,
	}, nil
}
