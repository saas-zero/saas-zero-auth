// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/bcrypt"
	"github.com/saas-zero/saas-zero-common/pkg/errno"
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
	authCtx := metadata.NewOutgoingContext(l.ctx, metadata.Pairs(
		"x-user-id", "0",
		"x-user-name", "",
		"x-tenant-id", "0",
	))

	// 1. Look up tenant by code
	tenantResp, err := l.svcCtx.SysTenants.GetTenantByCode(authCtx, &apps.TenantReq{
		Code: proto.String(req.TenantCode),
	})
	if err != nil {
		return &types.BaseResp{Code: errno.TenantNotFound.Code, Msg: errno.TenantNotFound.Msg}, nil
	}
	tenantId := tenantResp.GetData().GetId()
	tenantIdStr := strconv.FormatInt(tenantId, 10)

	// 2. Look up user by tenant + username
	authCtx = metadata.NewOutgoingContext(l.ctx, metadata.Pairs(
		"x-user-id", "0",
		"x-user-name", "",
		"x-tenant-id", tenantIdStr,
	))
	userResp, err := l.svcCtx.SysUsers.GetUserByUsername(authCtx, &apps.UserReq{
		Username: proto.String(req.Username),
	})
	if err != nil {
		return &types.BaseResp{Code: errno.UserNotFound.Code, Msg: errno.UserNotFound.Msg}, nil
	}
	user := userResp.GetData()
	if user == nil {
		return &types.BaseResp{Code: errno.UserNotFound.Code, Msg: errno.UserNotFound.Msg}, nil
	}

	// 3. Verify password
	if !bcrypt.Verify(req.Password, user.GetPassword()) {
		return &types.BaseResp{Code: errno.WrongPassword.Code, Msg: errno.WrongPassword.Msg}, nil
	}

	// 4. Generate JWT
	token, err := jwt.Sign(l.svcCtx.Config.JwtSecret, &jwt.Claims{
		UserId:    user.GetId(),
		TenantId:  tenantId,
		UserName:  user.GetUsername(),
		RoleCodes: user.GetRoleCodes(),
	}, time.Duration(l.svcCtx.Config.JwtExpire)*time.Second)
	if err != nil {
		return nil, err
	}
	return &types.BaseResp{
		Code: errno.Success.Code,
		Msg:  "登录成功",
		Data: map[string]interface{}{
			"token":      token,
			"userId":     user.GetIdStr(),
			"username":   user.GetUsername(),
			"nickname":   user.GetNickname(),
			"tenantId":   tenantIdStr,
			"tenantCode": req.TenantCode,
		},
	}, nil
}
