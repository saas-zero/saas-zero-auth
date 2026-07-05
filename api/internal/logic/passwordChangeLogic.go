// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/bcrypt"
	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type PasswordChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordChangeLogic {
	return &PasswordChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordChangeLogic) PasswordChange(req *types.PasswordChangeReq) (resp *types.BaseResp, err error) {
	claims, err := jwt.Parse(GetToken(l.ctx), l.svcCtx.Config.JwtSecret)
	if err != nil {
		return &types.BaseResp{Code: 3, Msg: "token无效或已过期"}, nil
	}
	ctx := withAuthContext(l.ctx, l.svcCtx.Config.JwtSecret)
	userResp, err := l.svcCtx.SysUsers.GetUserById(ctx, &apps.IdReq{Id: claims.UserId})
	if err != nil {
		return nil, err
	}
	user := userResp.GetData()
	if user == nil {
		return &types.BaseResp{Code: 1, Msg: "用户不存在"}, nil
	}
	if !bcrypt.Verify(req.OldPassword, user.GetPassword()) {
		return &types.BaseResp{Code: 2, Msg: "旧密码错误"}, nil
	}
	hash, err := bcrypt.Hash(req.NewPassword)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SysUsers.ResetPassword(ctx, &apps.UserReq{
		Id:       proto.Int64(claims.UserId),
		Password: proto.String(hash),
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseResp{Code: 0, Msg: "密码修改成功"}, nil
}
