package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/mixinsnapshot/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMixinsnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMixinsnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMixinsnapshotLogic {
	return &SearchMixinsnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMixinsnapshotLogic) SearchMixinsnapshot(in *pb.SearchMixinsnapshotReq) (*pb.SearchMixinsnapshotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchMixinsnapshotResp{}, nil
}
