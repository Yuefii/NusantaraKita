DOCKER_COMPOSE_DEV = docker-compose.yaml
DOCKER_COMPOSE_PROD = docker-compose.prod.yaml

# Untuk menjalankan environment development
.PHONY: dev
dev:
	@echo "Starting development environment..."
	docker-compose -f $(DOCKER_COMPOSE_DEV) up --build

# Untuk menghentikan environment development
.PHONY: dev-down
dev-down:
	@echo "Stopping development environment..."
	docker-compose -f $(DOCKER_COMPOSE_DEV) down

# Untuk menjalankan environment production
.PHONY: prod
prod:
	@echo "Starting production environment..."
	docker-compose -f $(DOCKER_COMPOSE_PROD) up --build -d

# Untuk menghentikan environment production
.PHONY: prod-down
prod-down:
	@echo "Stopping production environment..."
	docker-compose -f $(DOCKER_COMPOSE_PROD) down

# Untuk membersihkan container yang tidak digunakan
.PHONY: clean
clean:
	@echo "Cleaning up unused containers and images..."
	docker system prune -f

# Untuk melihat log dari container
.PHONY: logs
logs:
	@echo "Viewing logs for development environment..."
	docker-compose -f $(DOCKER_COMPOSE_DEV) logs -f

# Untuk melihat log dari container production
.PHONY: logs-prod
logs-prod:
	@echo "Viewing logs for production environment..."
	docker-compose -f $(DOCKER_COMPOSE_PROD) logs -f
