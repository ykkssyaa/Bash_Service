
docker.rebuild:
	docker compose up -d --build app

docker.run:
	docker compose up -d

docker.down:
	docker compose down


migrate.up:
	migrate -path ./migrations -database "postgres://yks:yksadm@localhost:5432/postgres?sslmode=disable" up

migrate.down:
	migrate -path ./migrations -database "postgres://yks:yksadm@localhost:5432/postgres?sslmode=disable" down
