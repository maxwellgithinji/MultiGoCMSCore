package tenants

import (
	"testing"

	"github.com/maxwellgithinji/MultiGoCMSCore/pkg/tenants/domain"
	"github.com/stretchr/testify/assert"
)

func fakeTenant() *domain.Tenant {
	return &domain.Tenant{
		ID:      "1",
		Name:    "name",
		Domain:  "org.com",
		Phone:   "+111-111-1111",
		Address: "123 Acme st",
	}
}
func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		db       TenantDatabase
		input    *domain.Tenant
		expected bool
	}{
		{
			name:     "save save a new tenant",
			db:       TenantDatabase{},
			input:    fakeTenant(),
			expected: false,
		},
		{
			name: "duplicate tenant with same name and id",
			db: TenantDatabase{
				Tenants: []domain.Tenant{*fakeTenant()},
			},
			input:    fakeTenant(),
			expected: true,
		},
		{
			name: "duplicate tenant with same id",
			db: TenantDatabase{
				Tenants: []domain.Tenant{*fakeTenant()},
			},
			input: func() *domain.Tenant {
				t := fakeTenant()
				t.Name = "new name"
				return t
			}(),
			expected: true,
		},
		{
			name: "duplicate tenant with same name",
			db: TenantDatabase{
				Tenants: []domain.Tenant{*fakeTenant()},
			},
			input: func() *domain.Tenant {
				t := fakeTenant()
				t.ID = "2"
				return t
			}(),
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.db.exist(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestAddTenant(t *testing.T) {
	type fields struct {
		Tenants []domain.Tenant
	}
	type args struct {
		tenant *domain.Tenant
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "add empty tenant",
			fields: fields{
				Tenants: []domain.Tenant{},
			},
			args: args{
				tenant: &domain.Tenant{},
			},
			wantErr: true,
		},
		{
			name: "add tenant with empty phone input",
			fields: fields{
				Tenants: []domain.Tenant{},
			},
			args: args{
				tenant: func() *domain.Tenant {
					t := fakeTenant()
					t.Phone = ""
					return t
				}(),
			},
			wantErr: true,
		},
		{
			name: "add tenant in a new database",
			fields: fields{
				Tenants: []domain.Tenant{},
			},
			args: args{
				tenant: fakeTenant(),
			},
			wantErr: false,
		},
		{
			name: "duplicate tenant",
			fields: fields{
				Tenants: []domain.Tenant{*fakeTenant()},
			},
			args: args{
				tenant: fakeTenant(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TenantDatabase{
				Tenants: tt.fields.Tenants,
			}
			if err := tr.AddTenant(tt.args.tenant); (err != nil) != tt.wantErr {
				t.Errorf("TenantDatabase.AddTenant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
