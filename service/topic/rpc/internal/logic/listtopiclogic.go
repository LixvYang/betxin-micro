package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopicLogic {
	return &ListTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTopicLogic) ListTopic(in *pb.ListTopicReq) (*pb.ListTopicResp, error) {

	return &pb.ListTopicResp{}, nil
}
