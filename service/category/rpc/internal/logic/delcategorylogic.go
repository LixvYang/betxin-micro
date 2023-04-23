package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCategoryLogic {
	return &DelCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCategoryLogic) DelCategory(in *pb.DelCategoryReq) (*pb.DelCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelCategoryResp{}, nil
}
