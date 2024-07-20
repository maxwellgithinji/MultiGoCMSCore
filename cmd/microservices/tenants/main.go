package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/common/cmd"
	tenants_application "github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/application"
	tenants_infrastructure "github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/infrastructure/tenants"
	tenants_public_http "github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/interface/public/http"
)

func main() {
	log.Println("Starting tenants microservice...")

	ctx := cmd.Context()

	r := createTenantsMicroservice()

	server := &http.Server{Addr: os.Getenv("TENANTS_PORT"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("failed to listen and serve tenants microservice", "Error", err)
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("closing tenants microservice...")

	if err := server.Close(); err != nil {
		slog.Error("failed to close tenants microservice", "Error", err)
		panic(err)
	}
}

func createTenantsMicroservice() *chi.Mux {

	tenantDB := tenants_infrastructure.NewTenantDatabase()

	tenantService := tenants_application.NewTenantService(*tenantDB)

	r := cmd.CreateRouter()

	tenants_public_http.AddRoutes(r, tenantService, *tenantDB)

	return r
}
