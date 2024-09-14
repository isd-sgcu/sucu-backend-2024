swagger:
	swag init -d ./internal/interface/handlers -g ../../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal

migrate:
	go run ./pkg/database/migration/migration_script.go