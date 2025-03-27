package tokenservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"ordering-platform/pkg/xerr"
	"ordering-platform/rpc/sys/gen/model"
	"ordering-platform/rpc/sys/gen/query"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *sysclient.CreateTokenReq) (*sysclient.CreateTokenResp, error) {
	// todo: add your logic here and delete this line

	token := &model.SysToken{
		Status:    in.Status,
		UserID:    in.UserId,
		Username:  in.Username,
		Token:     in.Token,
		ExpiredAt: in.ExpiredAt,
	}
	err := query.SysToken.WithContext(l.ctx).Create(token)
	if err != nil {
		logc.Errorf(l.ctx, "创建 Token 失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.Wrapf(xerr.NewDBErr(), "创建 Token  失败 %v, 参数 %+v", err, in)
	}
	return &sysclient.CreateTokenResp{}, nil
}
