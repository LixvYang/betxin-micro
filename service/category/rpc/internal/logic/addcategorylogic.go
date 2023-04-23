package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------category-----------------------
func (l *AddCategoryLogic) AddCategory(in *pb.AddCategoryReq) (*pb.AddCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCategoryResp{}, nil
}
