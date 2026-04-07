ifneq (,$(wildcard ../../.env))
    include ../../.env
    export
endif

#start database container
db:
	docker start rssagg-db

db-down:
	docker stop rssagg-db

# Run database migrations up (apply changes)
migrate-up:
	goose postgres $(DB_URL) up

# Run database migrations down (rollback changes)
migrate-down:
	goose postgres $(DB_URL) down