CARD_BINARY=cardsApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_cards:
	@echo "Building cards binary..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${CARD_BINARY} ./cmd/cards
	@echo "Done!"
