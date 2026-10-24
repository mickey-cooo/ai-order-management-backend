COMPOSE = docker compose -f deploy/local/docker-compose.yml
DIR = services/$(s)

# run loads the service's own .env (docker targets must not inherit it); dev gets it from env_files in .air.toml
ifneq ($(filter run,$(MAKECMDGOALS)),)
-include $(DIR)/.env
export
endif

up:       ; $(COMPOSE) up -d --build
down:     ; $(COMPOSE) down
logs:     ; $(COMPOSE) logs -f $(s)
run:      ; go run ./$(DIR)/cmd
dev:      ; cd $(DIR) && air
test:     ; go test ./$(DIR)/... ./pkg/...
fmt:      ; gofmt -s -w $(DIR) pkg
fmtcheck: ; test -z "$$(gofmt -s -l $(DIR) pkg)"
