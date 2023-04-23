package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopicLogic {
	return &ListTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTopicLogic) ListTopic() (resp *types.ListTopicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
