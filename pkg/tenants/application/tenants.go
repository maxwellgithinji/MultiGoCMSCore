package application

import (
	"fmt"

	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/domain"
	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/infrastructure/tenants"
)

type TenantService struct {
	db tenants.TenantDatabase
}

func NewTenantService(db tenants.TenantDatabase) TenantService {
	return TenantService{db}
}

type AddTenantCommand struct {
	ID      string
	Name    string
	Domain  string
	Phone   string
	Address string
}

func (s TenantService) AddTenant(cmd AddTenantCommand) error {
	tenant, err := domain.NewTenant(mapTenant(cmd))
	if err != nil {
		return fmt.Errorf("%w: cannot initialize tenant", err)
	}

	if err := s.db.AddTenant(tenant); err != nil {
		return fmt.Errorf("%w: cannot add tenant", err)
	}
	return nil
}

func mapTenant(cmd AddTenantCommand) *domain.Tenant {
	tenant := &domain.Tenant{
		ID:      cmd.ID,
		Name:    cmd.Name,
		Domain:  cmd.Domain,
		Phone:   cmd.Phone,
		Address: cmd.Address,
	}
	return tenant
}
