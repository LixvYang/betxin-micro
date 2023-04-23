package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCategoryLogic {
	return &SearchCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCategoryLogic) SearchCategory(in *pb.SearchCategoryReq) (*pb.SearchCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchCategoryResp{}, nil
}
