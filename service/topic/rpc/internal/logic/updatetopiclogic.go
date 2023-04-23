package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTopicLogic) UpdateTopic(in *pb.UpdateTopicReq) (*pb.UpdateTopicResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateTopicResp{}, nil
}
