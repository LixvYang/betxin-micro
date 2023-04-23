package category

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.GetCategoryReq) (resp *types.GetCategoryResp, err error) {
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorw("error: ", logx.LogField{Key: "CategoryModel.FindOne error", Value: err})
		return nil, err
	}

	return &types.GetCategoryResp{
		Errmsg: types.Errmsg{
			Code:    200,
			Message: "ok",
		},
		Data: types.Category(*category),
	}, nil
}
