// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"

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
	claims, err := jwt.Parse(GetToken(l.ctx), l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	return &types.BaseResp{
		Code: errno.Success.Code,
		Msg:  errno.Success.Msg,
		Data: map[string]interface{}{
			"userId":   claims.UserId,
			"tenantId": claims.TenantId,
			"userName": claims.UserName,
		},
	}, nil
}
