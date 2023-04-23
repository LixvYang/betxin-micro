package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByIdLogic {
	return &GetTopicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopicByIdLogic) GetTopicById(in *pb.GetTopicByIdReq) (*pb.GetTopicByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTopicByIdResp{}, nil
}
