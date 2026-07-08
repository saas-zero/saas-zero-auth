// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"
	"fmt"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-common/pkg/errno"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthVerifyLogic {
	return &OauthVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthVerifyLogic) OauthVerify() (resp *types.BaseResp, err error) {
	tokenStr := GetToken(l.ctx)
	claims, err := jwt.Parse(tokenStr, l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	// Verify token exists in Redis
	key := fmt.Sprintf("token:%s", claims.ID)
	exists, err := l.svcCtx.Redis.Exists(key)
	if err != nil || exists == false {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	return &types.BaseResp{
		Code: errno.Success.Code,
		Msg:  errno.Success.Msg,
		Data: map[string]interface{}{
			"userId":    claims.UserId,
			"tenantId":  claims.TenantId,
			"userName":  claims.UserName,
			"roleCodes": claims.RoleCodes,
		},
	}, nil
}
