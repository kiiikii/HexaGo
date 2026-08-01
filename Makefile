.PHONY: help db-up db-down frontend backend

help:
	@echo "Available targets:"
	@echo "  db-up    - Build the database in docker"
	@echo "  db-down    - Remove the database in docker"
	@echo "  frontend      - Run the frontend"
	@echo "  backend     - Run the backend"

db-up:
	docker compose up -d

db-down:
	docker compose down

frontend: 
	cd frontend && ng serve

backend: 
	cd backend && go run main.go