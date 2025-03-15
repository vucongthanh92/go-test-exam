package http

import (
	"os"

	"github.com/swaggo/swag"
	httpserver "github.com/vucongthanh92/go-base-utils/http/server"
	"github.com/vucongthanh92/go-test-exam/config"
	v1 "github.com/vucongthanh92/go-test-exam/internal/api/http/v1"
)

type Server struct {
	cfg             *config.AppConfig
	productHandler  *v1.ProductHandler
	categoryHandler *v1.CategoryHandler
	supplierHandler *v1.SupplierHandler
	distanceHandler *v1.DistanceHandler
}

func NewServer(
	cfg *config.AppConfig,
	productHandler *v1.ProductHandler,
	categoryHandler *v1.CategoryHandler,
	supplierHandler *v1.SupplierHandler,
	distanceHandler *v1.DistanceHandler,
) *Server {
	return &Server{
		cfg:             cfg,
		productHandler:  productHandler,
		categoryHandler: categoryHandler,
		supplierHandler: supplierHandler,
		distanceHandler: distanceHandler,
	}
}

func (s *Server) Run() {
	config := &httpserver.HttpServerConfig{
		Port:            s.cfg.Http.Port,
		Development:     s.cfg.Http.Development,
		ShutdownTimeout: s.cfg.Http.ShutdownTimeout,
		Resources:       s.cfg.Http.Resources,
		AllowOrigins:    s.cfg.Http.AllowOrigins,
	}
	httpServer, router := httpserver.NewServer(*config)

	// In the future, if we have v2, v3..., we will add at here
	v1.MapRoutes(
		router,
		s.productHandler,
		s.categoryHandler,
		s.supplierHandler,
		s.distanceHandler,
	)
	httpServer.Run()
}

func init() {
	dat, err := os.ReadFile("./docs/swagger.json")
	if err != nil {
		println("error when reading specs, please regenerate swagger")
	}
	spec := &swag.Spec{
		Version:          "1.0",
		BasePath:         "/api/v1/",
		Schemes:          []string{},
		Title:            "Order Service API",
		Description:      "Service for order related api",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(dat),
	}
	swag.Register(spec.InstanceName(), spec)
}
