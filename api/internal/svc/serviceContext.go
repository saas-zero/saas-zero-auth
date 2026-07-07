// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/saas-zero/saas-zero-auth/api/internal/config"
	"github.com/saas-zero/saas-zero-basedata/rpc/apps"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	SysUsers   apps.SysUsersClient
	SysTenants apps.SysTenantsClient
	SysMenus   apps.SysMenusClient
	SysApis    apps.SysApisClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	cli := zrpc.MustNewClient(c.BaseDataRpc)
	return &ServiceContext{
		Config:     c,
		SysUsers:   apps.NewSysUsersClient(cli.Conn()),
		SysTenants: apps.NewSysTenantsClient(cli.Conn()),
		SysMenus:   apps.NewSysMenusClient(cli.Conn()),
		SysApis:    apps.NewSysApisClient(cli.Conn()),
	}
}
