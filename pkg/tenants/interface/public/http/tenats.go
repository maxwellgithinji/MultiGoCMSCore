package http

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	common_http "github.com/maxwellgithinji/MultiGoCMSCore/pkg/common/http"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/application"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/infrastructure/tenants"
	uuid "github.com/satori/go.uuid"
)

func AddRoutes(router *chi.Mux, service application.TenantService, db tenants.TenantDatabase) {
	resource := tenantResource{service: service, db: db}
	router.Post("/tenant", resource.Post)
}

type tenantResource struct {
	service application.TenantService
	db      tenants.TenantDatabase
}

func (t tenantResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostTenantRequest{}
	if err := render.Decode(r, &req); err != nil {
		slog.Error("failed to decode", "Error", err)
		_ = render.Render(w, r, common_http.ErrBadRequest(err))

	}

	cmd := application.AddTenantCommand{
		ID:      uuid.NewV1().String(),
		Name:    req.Name,
		Domain:  req.Domain,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := t.service.AddTenant(cmd); err != nil {
		slog.Error("failed to add tenant", "Error", err)
		_ = render.Render(w, r, common_http.ErrInternal(err))
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, PostTenantResponse{ID: cmd.ID})
}

type PostTenantRequest struct {
	Name    string `json:"name"`
	Domain  string `json:"domain"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type PostTenantResponse struct {
	ID string
}
