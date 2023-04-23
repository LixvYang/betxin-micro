package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------topic-----------------------
func (l *AddTopicLogic) AddTopic(in *pb.AddTopicReq) (*pb.AddTopicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTopicResp{}, nil
}
