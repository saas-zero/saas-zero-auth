// Code scaffolded by goctl. Safe to edit.

package logic

import (
	"context"
	"strconv"

	"github.com/saas-zero/saas-zero-auth/api/internal/svc"
	"github.com/saas-zero/saas-zero-auth/api/internal/types"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/saas-zero/saas-zero-common/pkg/bcrypt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type PasswordResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordResetLogic {
	return &PasswordResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordResetLogic) PasswordReset(req *types.PasswordResetReq) (resp *types.BaseResp, err error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return &types.BaseResp{Code: 4, Msg: "用户ID格式错误"}, nil
	}
	hash, err := bcrypt.Hash(req.NewPassword)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SysUsers.ResetPassword(withAuthContext(l.ctx, l.svcCtx.Config.JwtSecret), &apps.UserReq{
		Id:       proto.Int64(userId),
		Password: proto.String(hash),
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseResp{Code: 0, Msg: "密码重置成功"}, nil
}
