package logic

import (
	"context"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps/system-service"

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
	// 示例：通过用户名获取用户信息
	stream, err := l.svcCtx.BaseDataRpc.GetUserByUsername(l.ctx)
	if err != nil {
		return nil, err
	}

	// 发送请求
	err = stream.Send(&system_service.UserReq{
		Username: "testuser", // 实际应该从 req 中获取
	})
	if err != nil {
		return nil, err
	}

	// 关闭流并接收响应
	err = stream.CloseSend()
	if err != nil {
		return nil, err
	}

	//// 接收用户信息
	//user, err := stream.Recv()
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 处理用户信息和登录逻辑
	//_ = user // 这里可以根据需要处理用户信息
	//
	//// 返回登录响应
	//return &types.OauthLoginResp{
	//	BaseResp: types.BaseResp{
	//		Code: 0,
	//		Msg:  "登录成功",
	//	},
	//	Data: user,
	//}, nil

	return return &types.OauthLoginResp{
			BaseResp: types.BaseResp{
				Code: 0,
				Msg:  "登录成功",
			}
		}, nil
}