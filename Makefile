
download.deps:
	go mod download

docker.rebuild:
	docker compose up -d --build app

docker.run:
	docker compose up -d

docker.run.db:
	docker compose up -d postgres

docker.down:
	docker compose down


migrate.up:
	migrate -path ./migrations -database "postgres://yks:yksadm@localhost:5432/postgres?sslmode=disable" up

migrate.down:
	migrate -path ./migrations -database "postgres://yks:yksadm@localhost:5432/postgres?sslmode=disable" down

mock.gen.gateway:
	mockgen -source=internal/gateway/gateway.go \
	-destination=internal/gateway/mock/mock_gateway.go

tests.run:
	go test ./internal/...

tests.cover:
	go test -cover ./internal/...
