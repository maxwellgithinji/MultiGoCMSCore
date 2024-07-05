package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/application"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/infrastructure/tenants"
	tenants_public_http "github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/interface/public/http"
)

func main() {
	log.Println("starting monolith...")

	r := chi.NewRouter()

	initTenantMicroservice(r)

	server := &http.Server{
		Addr:    os.Getenv("MONOLITH_PORT"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to listen and server", "Error", err)
			os.Exit(1)
		}
	}()

	log.Printf("Server is running on %s", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to gracefully shutdown the server", "Error", err)
	}

	log.Println("Server stopped")
}

func initTenantMicroservice(r *chi.Mux) {
	tenantDB := tenants.NewTenantDatabase()
	tenantService := application.NewTenantService(*tenantDB)
	tenants_public_http.AddRoutes(r, tenantService, *tenantDB)
}
