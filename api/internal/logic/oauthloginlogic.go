// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	basedata_service "github.com/saas-zero/saas-zero-basedata/rpc/apps"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.OauthLoginResp, err error) {
	// 通过用户名获取用户信息
	userReq := &basedata_service.UserReq{
		Username: req.Username,
	}

	user, err := l.svcCtx.BaseDataRpc.GetUserByUsername(l.ctx, userReq)
	if err != nil {
		return nil, err
	}

	// 验证密码（这里应该有实际的密码验证逻辑）
	// 暂时省略密码验证逻辑，后续可以添加

	// 返回登录响应
	return &types.OauthLoginResp{
		Code: 0,
		Msg:  "登录成功",
		Data: user,
	}, nil
}
