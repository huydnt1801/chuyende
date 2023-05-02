MYSQL_DATABASE_URI = mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@$(MYSQL_HOST):$(MYSQL_PORT)/$(MYSQL_DATABASE)

migrate:
	@echo \# make migrate name="$(name)"
	go run cmd/migrate/main.go create "$(name)"

migrate-up:
	atlas migrate apply \
  		--dir "file://migrations" \
  		--url $(MYSQL_DATABASE_URI)

generate-ent:
	go generate ./internal/ent