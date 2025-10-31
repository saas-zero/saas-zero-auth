// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"auth-service/api/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	SystemserviceRpc sysapis.SysApis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
