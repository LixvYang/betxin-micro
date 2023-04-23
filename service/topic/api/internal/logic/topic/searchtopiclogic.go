package topic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTopicLogic {
	return &SearchTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTopicLogic) SearchTopic(req *types.SearchTopicReq) (resp *types.SearchTopicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
