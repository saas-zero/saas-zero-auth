// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"
	"time"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/bcrypt"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.BaseResp, err error) {
	userResp, err := l.svcCtx.SysUsers.GetUserByUsername(metadata.NewOutgoingContext(l.ctx, metadata.Pairs(
		"x-user-id", "0",
		"x-user-name", "",
		"x-tenant-id", "0",
	)), &apps.UserReq{
		Username: proto.String(req.Username),
	})
	if err != nil {
		return nil, err
	}
	user := userResp.GetData()
	if user == nil {
		return &types.BaseResp{Code: 1, Msg: "用户不存在"}, nil
	}
	if !bcrypt.Verify(req.Password, user.GetPassword()) {
		return &types.BaseResp{Code: 2, Msg: "密码错误"}, nil
	}
	token, err := jwt.Sign(l.svcCtx.Config.JwtSecret, &jwt.Claims{
		UserId:    user.GetId(),
		TenantId:  user.GetTenantId(),
		UserName:  user.GetUsername(),
		RoleCodes: user.GetRoleCodes(),
	}, time.Duration(l.svcCtx.Config.JwtExpire)*time.Second)
	if err != nil {
		return nil, err
	}
	return &types.BaseResp{
		Code: 0,
		Msg:  "登录成功",
		Data: map[string]interface{}{
			"token":    token,
			"userId":   user.GetIdStr(),
			"username": user.GetUsername(),
			"nickname": user.GetNickname(),
		},
	}, nil
}
