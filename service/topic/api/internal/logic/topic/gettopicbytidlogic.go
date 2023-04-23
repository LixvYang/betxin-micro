package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByTidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicByTidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByTidLogic {
	return &GetTopicByTidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicByTidLogic) GetTopicByTid(req *types.GetTopicByTidReq) (resp *types.GetTopicByTidResp, err error) {
	// todo: add your logic here and delete this line

	return
}
