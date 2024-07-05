package tenants

import (
	"fmt"

	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/domain"
)

type TenantDatabase struct {
	Tenants []domain.Tenant
}

func NewTenantDatabase() *TenantDatabase {
	return &TenantDatabase{[]domain.Tenant{}}
}

func (t *TenantDatabase) AddTenant(tenant *domain.Tenant) error {
	// TODO: check validation on the application layer too
	if err := tenant.Validate(); err != nil {
		return fmt.Errorf("invalid input: %w", err)
	}

	if t.exist(tenant) {
		return fmt.Errorf("tenant %s exist", tenant.Name)
	}

	t.Tenants = append(t.Tenants, *tenant)
	return nil
}

func (t *TenantDatabase) exist(tenant *domain.Tenant) bool {
	for _, v := range t.Tenants {
		if v.ID == tenant.ID ||
			v.Name == tenant.Name {
			return true
		}
	}
	return false
}
