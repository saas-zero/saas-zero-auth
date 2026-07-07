// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-common/pkg/errno"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthRefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthRefreshLogic {
	return &OauthRefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthRefreshLogic) OauthRefresh(req *types.OauthRefreshReq) (resp *types.BaseResp, err error) {
	token := req.Token
	if token == "" {
		token = GetToken(l.ctx)
	}
	claims, err := jwt.Parse(token, l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	if !tokenExistsInRedis(l.svcCtx.Redis, claims.ID) {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	newClaims := &jwt.Claims{
		UserId:       claims.UserId,
		TenantId:     claims.TenantId,
		UserName:     claims.UserName,
		TokenVersion: claims.TokenVersion,
	}
	newToken, err := jwt.Sign(l.svcCtx.Config.JwtSecret, newClaims, time.Duration(l.svcCtx.Config.JwtExpire)*time.Second)
	if err != nil {
		return nil, err
	}
	// Store new token in Redis, delete old
	l.svcCtx.Redis.Del(fmt.Sprintf("token:%s", claims.ID))
	if err := l.svcCtx.Redis.Setex(fmt.Sprintf("token:%s", newClaims.ID), newToken, int(l.svcCtx.Config.JwtExpire)); err != nil {
		return nil, err
	}
	return &types.BaseResp{
		Code: errno.Success.Code,
		Msg:  errno.Success.Msg,
		Data: map[string]string{"token": newToken},
	}, nil
}
