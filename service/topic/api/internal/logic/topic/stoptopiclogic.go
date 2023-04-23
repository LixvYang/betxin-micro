package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopTopicLogic {
	return &StopTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopTopicLogic) StopTopic(req *types.StopTopicReq) (resp *types.StopTopicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
