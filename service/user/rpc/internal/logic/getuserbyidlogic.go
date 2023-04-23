package logic

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/user/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserByIdResp{}, nil
}
