// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/isaqueveras/power-sso/internal/utils/grpckit"
)

type ctxKey interface{}

// GRPCZap is a request interceptor for logging all gRPC requests using Zap
func GRPCZap() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var (
			RID ctxKey    = "RID"
			t1  time.Time = time.Now()
		)

		md, _ := metadata.FromIncomingContext(ctx)
		p, _ := peer.FromContext(ctx)

		if len(md.Get("RID")) == 0 {
			rid := uuid.New().String()
			md["rid"] = []string{rid}
			ctx = context.WithValue(ctx, RID, rid)
		}

		resp, err = handler(context.WithValue(ctx, RID, md.Get("RID")[0]), req)

		var (
			statusCode = status.Convert(err)
			latency    = float64(time.Since(t1) / time.Millisecond)
			fields     = []zap.Field{
				zap.Time("date", time.Now()),
				zap.Float64("latency", latency),
				zap.String("path", info.FullMethod),
				zap.String("request_id", md.Get("RID")[0]),
				zap.String("user-agent", md.Get("User-Agent")[0]),
				zap.Int("status_code", int(statusCode.Code())),
				zap.String("status", statusCode.Code().String()),
				zap.String("client_ip", p.Addr.String()),
			}
		)

		if err != nil {
			fields = append(fields, zap.Error(err), zap.String("cause", statusCode.Message()))

			if len(statusCode.Details()) > 0 {
				details := statusCode.Details()[0].(*grpckit.ErrorGRPC)

				fields = append(fields,
					zap.String("cause", statusCode.Message()),
					zap.String("raw_error", details.GetRawError()),
					zap.String("location", details.GetLocation()),
					zap.String("error", details.GetError()),
					zap.Uint64("code", details.GetCode()),
				)
			}

			zap.L().Error("Request handling failure", fields...)
		} else {
			zap.L().Info("Request handling success", fields...)
		}

		return
	}
}
