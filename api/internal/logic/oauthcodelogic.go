// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthCodeLogic {
	return &OauthCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthCodeLogic) OauthCode() (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
