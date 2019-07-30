package gateway

import (
	"context"
	"net/http"
)

type GatewayHandler func(ctx context.Context) interface{}

func MakeHandler(h http.Handler) GatewayHandler {
	return func(ctx context.Context) interface{} {
		return ctx
	}
}
