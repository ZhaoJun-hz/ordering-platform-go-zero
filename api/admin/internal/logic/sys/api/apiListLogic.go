package api

import (
	"context"
	"ordering-platform/rpc/sys/sysclient"

	"ordering-platform/api/admin/internal/svc"
	"ordering-platform/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询 api 列表
func NewApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiListLogic {
	return &ApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiListLogic) ApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.ApiService.ListApi(l.ctx, &sysclient.ApiListReq{
		Type:     req.Type,
		Action:   req.Action,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.ApiListData
	for _, item := range result.Data {
		list = append(list, &types.ApiListData{
			Id:         item.Id,
			Handle:     item.Handle,
			Title:      item.Title,
			Path:       item.Path,
			Type:       item.Type,
			Action:     item.Action,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ApiListResp{
		Total: result.Total,
		Data:  list,
	}, nil
}
