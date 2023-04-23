package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTopicLogic {
	return &DelTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTopicLogic) DelTopic(in *pb.DelTopicReq) (*pb.DelTopicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelTopicResp{}, nil
}
