package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMixinsnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMixinsnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMixinsnapshotLogic {
	return &UpdateMixinsnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMixinsnapshotLogic) UpdateMixinsnapshot(in *pb.UpdateMixinsnapshotReq) (*pb.UpdateMixinsnapshotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMixinsnapshotResp{}, nil
}
