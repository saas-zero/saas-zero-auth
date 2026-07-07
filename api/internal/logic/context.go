package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/metadata"
)

type ctxKey string

const tokenKey ctxKey = "auth_token"

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func GetToken(ctx context.Context) string {
	if v, ok := ctx.Value(tokenKey).(string); ok {
		return v
	}
	return ""
}

func ExtractBearerToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}
	return ""
}

func tokenExistsInRedis(rds *redis.Redis, jti string) bool {
	if jti == "" {
		return false
	}
	key := fmt.Sprintf("token:%s", jti)
	exists, err := rds.Exists(key)
	return err == nil && exists
}

func withAuthContext(ctx context.Context, secret string) context.Context {
	token := GetToken(ctx)
	if token == "" {
		return ctx
	}
	claims, err := jwt.Parse(token, secret)
	if err != nil {
		return ctx
	}
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"x-user-id", strconv.FormatInt(claims.UserId, 10),
		"x-user-name", claims.UserName,
		"x-tenant-id", strconv.FormatInt(claims.TenantId, 10),
	))
}
