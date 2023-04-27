package logic

import (
	"context"
	"fmt"
	"log"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryLogic {
	return &ListCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCategoryLogic) ListCategory(in *pb.ListCategoryReq) (*pb.ListCategoryResp, error) {
	list, err := l.svcCtx.CategoryModel.List(l.ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("list: ", list)

	var categoryList []*pb.Category

	for _, c := range list {
		categoryList = append(categoryList, &pb.Category{
			Id:           c.Id,
			CategoryName: c.CategoryName,
		})
	}

	return &pb.ListCategoryResp{
		Category: categoryList,
	}, nil
}
