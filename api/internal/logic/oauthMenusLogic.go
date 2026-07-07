// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/errno"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthMenusLogic {
	return &OauthMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthMenusLogic) OauthMenus() (resp *types.BaseResp, err error) {
	claims, err := jwt.Parse(GetToken(l.ctx), l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	if !tokenExistsInRedis(l.svcCtx.Redis, claims.ID) {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	treeResp, err := l.svcCtx.SysMenus.GetMenuTree(withAuthContext(l.ctx, l.svcCtx.Config.JwtSecret), &apps.EmptyReq{})
	if err != nil {
		return nil, err
	}
	return &types.BaseResp{
		Code: errno.Success.Code, Msg: errno.Success.Msg, Data: treeResp.GetData(),
	}, nil
}
