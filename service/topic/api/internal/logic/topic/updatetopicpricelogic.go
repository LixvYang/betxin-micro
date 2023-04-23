package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicPriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicPriceLogic {
	return &UpdateTopicPriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicPriceLogic) UpdateTopicPrice(req *types.UpdateTopicPriceReq) (resp *types.UpdateTopicPriceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
