package category

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/category/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryLogic {
	return &ListCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoryLogic) ListCategory() (resp *types.ListCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
