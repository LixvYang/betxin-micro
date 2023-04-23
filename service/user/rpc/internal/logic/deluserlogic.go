package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/user/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *pb.DelUserReq) (*pb.DelUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserResp{}, nil
}