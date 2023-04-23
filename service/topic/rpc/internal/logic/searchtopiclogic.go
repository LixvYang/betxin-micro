package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTopicLogic {
	return &SearchTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTopicLogic) SearchTopic(in *pb.SearchTopicReq) (*pb.SearchTopicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchTopicResp{}, nil
}
