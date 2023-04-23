package user

import (
	"context"
	"database/sql"

	"github.com/lixvyang/betxin-micro/service/user/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/api/internal/types"
	"github.com/lixvyang/betxin-micro/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		IdentityNumber: req.IdentityNumber,
		FullName:       sql.NullString{String: req.FullName, Valid: true},
		Uid:            req.Uid,
		AvatarUrl:      sql.NullString{String: req.AvatarUrl, Valid: true},
		Biography:      sql.NullString{String: req.Biography, Valid: true},
	})
	if err != nil {
		logx.Errorw("UserModel.Insert failed! ", logx.LogField{Key: "Err: ", Value: err})
		return nil, err
	}
	resp = new(types.CreateUserResp)
	resp.Code = 200
	resp.Message = "Yes"
	resp.Data = "Yes"

	return resp, nil
}
