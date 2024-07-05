package domain

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Tenant struct {
	ID      string `validate:"required"`
	Name    string `validate:"required"`
	Domain  string `validate:"required"`
	Phone   string `validate:"required"`
	Address string `validate:"required"`
}

func (t Tenant) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return fmt.Errorf("tenant struct validation error: %w", err)
	}
	return nil
}

func NewTenant(tenant *Tenant) (*Tenant, error) {
	return tenant, nil
}
