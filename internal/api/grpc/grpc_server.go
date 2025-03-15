package grpc

import (
	grpc "github.com/vucongthanh92/go-base-utils/grpc/server"
	"github.com/vucongthanh92/go-test-exam/config"
)

type Server struct {
	Cfg *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	return &Server{
		Cfg: cfg,
	}
}

func (s *Server) Run() {
	grpcServer, _ := grpc.NewServer(
		grpc.GrpcServerConfig(*s.Cfg.GRPC),
	)
	grpcServer.Run()
}
