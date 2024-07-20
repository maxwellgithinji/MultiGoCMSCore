# Go parameters
GOCMD=go
GORUN=$(GOCMD) run

# Main package paths
MONOLITH_PACKAGE=./cmd/monolith
TENANTS_PACKAGE=./cmd/microservices/tenants

# Run commands
run-monolith:
	$(GORUN) $(MONOLITH_PACKAGE)

run-tenants:
	$(GORUN) $(TENANTS_PACKAGE)

# Run all services
run-all:
	$(MAKE) -j run-monolith run-tenants
	
.PHONY: run-monolith run-tenants run-all