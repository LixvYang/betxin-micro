package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTopicLogic) DeleteTopic(req *types.DeleteTopicReq) (resp *types.DeleteTopicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
