// Package http_server ...
package http_server

import (
	"context"
	"log/slog"
	"net/http"
	"slices"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"trintech/review/config"
	"trintech/review/pkg/token_util"
)

type middlewareFunc func(http.Handler) http.Handler

type HttpServer struct {
	*http.Server
	cfg *config.Endpoint
}

func NewHttpServer(
	handler func(mux *runtime.ServeMux),
	cfg *config.Endpoint,
	authenticator token_util.JWTAuthenticator,
) *HttpServer {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: false, EmitUnpopulated: true},
			UnmarshalOptions: protojson.UnmarshalOptions{AllowPartial: true},
		}),
		runtime.WithMetadata(MapMetaDataWithBearerToken(authenticator)),
		// runtime.WithErrorHandler(forwardErrorResponse),
	)
	handler(mux)
	middlewares := []middlewareFunc{
		allowCORS,
	}

	slices.Reverse(middlewares)

	var handleR http.Handler = mux
	// for _, handle := range middlewares {
	// 	handleR = handle(handleR)
	// }

	return &HttpServer{
		cfg: cfg,
		Server: &http.Server{
			Addr:    cfg.Address(),
			Handler: handleR,
		},
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	slog.Info("Server listening in port", "port", s.cfg.Port)

	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}

	slog.Info("shutdown http server successful")
	return nil
}
