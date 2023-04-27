package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMixinsnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMixinsnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMixinsnapshotLogic {
	return &DelMixinsnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMixinsnapshotLogic) DelMixinsnapshot(in *pb.DelMixinsnapshotReq) (*pb.DelMixinsnapshotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelMixinsnapshotResp{}, nil
}
