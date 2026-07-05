// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthUserInfoLogic {
	return &OauthUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthUserInfoLogic) OauthUserInfo() (resp *types.BaseResp, err error) {
	claims, err := jwt.Parse(GetToken(l.ctx), l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: 3, Msg: "token无效或已过期"}, nil
	}
	userResp, err := l.svcCtx.SysUsers.GetUserById(withAuthContext(l.ctx, l.svcCtx.Config.JwtSecret), &apps.IdReq{Id: claims.UserId})
	if err != nil {
		return nil, err
	}
	user := userResp.GetData()
	if user == nil {
		return &types.BaseResp{Code: 1, Msg: "用户不存在"}, nil
	}
	return &types.BaseResp{
		Code: 0, Msg: "success", Data: user,
	}, nil
}
