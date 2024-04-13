// // Copyright (c) 2023 Isaque Veras
// // Use of this source code is governed by MIT style
// // license that can be found in the LICENSE file.

package server

// import (
// 	"net"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
// 	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
// 	gogrpc "google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"

// 	"github.com/isaqueveras/powersso/delivery/grpc/auth"
// 	"github.com/isaqueveras/powersso/middleware"
// 	"github.com/isaqueveras/powersso/oops"
// 	"github.com/isaqueveras/powersso/utils"
// )

// func (s *Server) ServerGRPC() (err error) {
// 	if !s.cfg.Server.StartGRPC {
// 		return
// 	}

// 	var (
// 		listen     net.Listener
// 		serverGRPC = gogrpc.NewServer(
// 			// TODO: grpc.Creds(credentials.NewTLS(nil)),
// 			gogrpc.UnaryInterceptor(
// 				grpcMiddleware.ChainUnaryServer(
// 					middleware.GRPCZap(),
// 					grpcRecovery.UnaryServerInterceptor(
// 						grpcRecovery.WithRecoveryHandler(utils.PanicRecovery),
// 					),
// 				),
// 			),
// 		)
// 	)

// 	s.logg.Info("Server GRPC is running")
// 	if s.cfg.Server.IsModeDevelopment() {
// 		s.logg.Debug("RUNNING IN DEVELOPMENT: REFLECTION ON")
// 		reflection.Register(serverGRPC)
// 	}

// 	go func(server *gogrpc.Server) {
// 		signalOS := make(chan os.Signal, 1)
// 		signal.Notify(signalOS, syscall.SIGINT, syscall.SIGTERM)
// 		for range signalOS {
// 			server.GracefulStop()
// 			return
// 		}
// 	}(serverGRPC)

// 	// TODO: use address in config file
// 	if listen, err = net.Listen("tcp", "localhost:50050"); err != nil {
// 		return oops.Err(err)
// 	}

// 	auth.RegisterAuthenticationServer(serverGRPC, &auth.Server{})

// 	s.group.Go(func() error {
// 		return serverGRPC.Serve(listen)
// 	})

// 	return
// }
