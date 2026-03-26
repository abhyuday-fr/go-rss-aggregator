#start database container
db:
	docker start rssagg-db

db-down:
	docker stop rssagg-db

# Run database migrations up (apply changes)
migrate-up:
	goose postgres "postgres://postgres:1709@localhost:5432/rssagg" up

# Run database migrations down (rollback changes)
migrate-down:
	goose postgres "postgres://postgres:1709@localhost:5432/rssagg" down