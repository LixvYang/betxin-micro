package user

import (
	"context"
	"fmt"

	"github.com/lixvyang/betxin-micro/service/user/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser() (resp *types.User, err error) {
	uid := fmt.Sprintf("%s", l.ctx.Value("uid"))
	user, err := l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
	if err != sqlx.ErrNotFound {
		logx.Errorw("UserModel.FindOneByUid err: ", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	if err != nil {
		logx.Errorw("UserModel.FindOneByUid err: ", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	resp.AvatarUrl = user.AvatarUrl.String
	resp.Biography = user.Biography.String
	resp.FullName = user.FullName.String
	resp.IdentityNumber = user.IdentityNumber
	resp.Uid = user.Uid
	return resp, nil
}
