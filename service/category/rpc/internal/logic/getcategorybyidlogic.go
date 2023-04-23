package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *pb.GetCategoryByIdReq) (*pb.GetCategoryByIdResp, error) {
	// todo: add your logic here and delete this line
	cate, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		logx.Errorw("error: ", logx.LogField{Value: err, Key: "CategoryModel.FindOne "})
	}

	return &pb.GetCategoryByIdResp{
		Category: &pb.Category{
			Id:           cate.Id,
			CategoryName: cate.CategoryName,
		},
	}, nil
}
