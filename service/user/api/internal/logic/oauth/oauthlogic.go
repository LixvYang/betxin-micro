package oauth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/lixvyang/betxin-micro/service/user/api/internal/svc"
	"github.com/lixvyang/betxin-micro/service/user/api/pkg/auth"
	"github.com/lixvyang/betxin-micro/service/user/model"
)

type OauthLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	httpRequest  *http.Request
	httpResponse http.ResponseWriter
}

func NewOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *OauthLogic {
	return &OauthLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		httpRequest:  r,
		httpResponse: w,
	}
}

func (l *OauthLogic) Oauth() error {
	// 1. 根据code去获取用户信息
	code := l.httpRequest.URL.Query().Get("code")
	access_token, _, err := mixin.AuthorizeToken(l.ctx, l.svcCtx.Config.Mixin.ClientId, l.svcCtx.Config.Mixin.AppSecret, code, "")
	if err != nil {
		logx.Errorw("mixin.AuthorizeToken err: ", logx.LogField{Key: "err", Value: err.Error()})
		return err
	}

	// 2. 访问用户信息
	userInfo, err := auth.GetUserInfo(access_token)
	if err != nil {
		logx.Errorw("auth.GetUserInfo err: ", logx.LogField{Key: "err", Value: err.Error()})
		return err
	}
	user := model.User{
		AvatarUrl:      sql.NullString{String: userInfo.AvatarURL, Valid: true},
		IdentityNumber: userInfo.UserID,
		SessionId:      sql.NullString{String: userInfo.SessionID, Valid: true},
		Uid:            userInfo.UserID,
		FullName:       sql.NullString{String: userInfo.FullName, Valid: true},
		Biography:      sql.NullString{String: userInfo.Biography, Valid: true},
	}

	// 3. 查询用户是否已经存在
	// 3.1 若不存在则创建
	// 3.2 若存在则更新数据
	_, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, userInfo.UserID)
	if err != sqlx.ErrNotFound {
		// 已经存在
		err = l.svcCtx.UserModel.Update(l.ctx, &user)
		if err != nil {
			logx.Errorw("UserModel.Update err: ", logx.LogField{Key: "err", Value: err.Error()})
			return err
		}
	} else {
		// 创建
		_, err = l.svcCtx.UserModel.Insert(l.ctx, &user)
		if err != nil {
			logx.Errorw("UserModel.Insert failed", logx.LogField{Key: "err", Value: err.Error()})
			return err
		}
	}

	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, "123-456-789")
	if err != nil {
		logx.Errorw("l.getJwtToken failed", logx.LogField{Key: "err", Value: err.Error()})
		return err
	}

	url := fmt.Sprintf("http://%s?token=%s", l.svcCtx.Config.Host, jwtToken)

	http.Redirect(l.httpResponse, l.httpRequest, url, http.StatusPermanentRedirect)
	return nil
}

func (l *OauthLogic) getJwtToken(secretKey string, iat int64, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
