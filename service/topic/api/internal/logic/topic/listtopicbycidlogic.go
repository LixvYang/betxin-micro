package topic

import (
	"context"
	"fmt"

	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopicByCidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTopicByCidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopicByCidLogic {
	return &ListTopicByCidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTopicByCidLogic) ListTopicByCid(req *types.ListTopicByCidReq) (resp *types.ListTopicByCidResp, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.CategoryRPC.ListCategory(l.ctx, &pb.ListCategoryReq{})
	if err != nil {
		panic(err)
	}
	for _, l := range list.Category {
		fmt.Println(*l)
	}
	return
}
