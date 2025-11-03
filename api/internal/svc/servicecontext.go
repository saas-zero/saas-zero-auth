// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/saas-zero/saas-zero-auth/api/internal/config"
	"github.com/saas-zero/saas-zero-basedata/rpc/sysusers"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	BaseDataRpc sysusers.SysUsers
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BaseDataRpc: sysusers.NewSysUsers(zrpc.MustNewClient(c.BaseDataRpc)),
	}
}
