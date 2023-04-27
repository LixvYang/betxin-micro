package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMixinsnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMixinsnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMixinsnapshotLogic {
	return &AddMixinsnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------mixinsnapshot-----------------------
func (l *AddMixinsnapshotLogic) AddMixinsnapshot(in *pb.AddMixinsnapshotReq) (*pb.AddMixinsnapshotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMixinsnapshotResp{}, nil
}
