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
	res, err := l.svcCtx.CategoryRPC.GetCategoryById(l.ctx, &pb.GetCategoryByIdReq{Id: 1})
	if err != nil {
		logx.Errorw("failed: ", logx.LogField{Value: err, Key: "Error"})
		return nil, err
	}

	fmt.Println(res.Category)
	return
}
