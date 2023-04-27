package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMixinsnapshotByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMixinsnapshotByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMixinsnapshotByIdLogic {
	return &GetMixinsnapshotByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMixinsnapshotByIdLogic) GetMixinsnapshotById(in *pb.GetMixinsnapshotByIdReq) (*pb.GetMixinsnapshotByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMixinsnapshotByIdResp{}, nil
}
