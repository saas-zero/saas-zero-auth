// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"auth-service/api/internal/logic"
	"auth-service/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OauthCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOauthCodeLogic(r.Context(), svcCtx)
		resp, err := l.OauthCode()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
