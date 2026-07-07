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

type OauthPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthPermissionsLogic {
	return &OauthPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthPermissionsLogic) OauthPermissions() (resp *types.BaseResp, err error) {
	if _, err := jwt.Parse(GetToken(l.ctx), l.svcCtx.Config.JwtSecret); err != nil {
		return &types.BaseResp{Code: errno.TokenExpired.Code, Msg: errno.TokenExpired.Msg}, nil
	}
	apiResp, err := l.svcCtx.SysApis.GetApiList(withAuthContext(l.ctx, l.svcCtx.Config.JwtSecret), &apps.ApiPageReq{})
	if err != nil {
		return nil, err
	}
	list := apiResp.GetList()
	perms := make([]string, len(list))
	for i, a := range list {
		perms[i] = a.GetApiPath()
	}
	if perms == nil {
		perms = []string{}
	}
	return &types.BaseResp{
		Code: errno.Success.Code, Msg: errno.Success.Msg, Data: perms,
	}, nil
}
