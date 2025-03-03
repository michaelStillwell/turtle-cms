PROJ=turtle-cms

DBNAME=turtle_cms.db
DBURL=file:./$(DBNAME)
DBURL_REMOTE=libsql://$(DBNAME)
DBPASS=postgres

VIEWS_DIR=./views
MIGRATIONS_DIR=db/migrations
SCHEMA_FILE=db/schema.sql

# Server

run: sqlc templ
	@DATABASE_URL=$(DBURL) go run ./cmd/server -io

build:
	@if [ -d "bin" ]; then \
		rm -rf "bin"; \
	fi
	@CGO_ENABLED=1 go build -o ./bin/$(PROJ) ./cmd/server

templ:
	@templ fmt $(VIEWS_DIR)
	@templ generate

sqlc:
	@sqlc generate

# Docker

dbuild:
	docker build --no-cache -t $(PROJ) --build-arg PKG=$(PROJ) .

drun: migrate up
	echo 'kill me'
	@echo '' >> logfile
	@docker run --rm \
		-e DATABASE_URL=$(DBURL) \
		--mount type=bind,src=./logfile,dst=/app/logfile \
		--mount type=bind,src=./local.db,dst=/app/local.db \
		--name $(PROJ) $(PROJ)

druni:
	docker run --rm -it $(PROJ)

# Db, local

create-db:
	@touch $(DBNAME)

start-db: create-db

rm-db:
	@if [ -f "$(DBNAME)" ]; then \
		rm -f "$(DBNAME)"; \
	fi

restart-db: rm-db create-db up

# Migration

migrate:
	@atlas migrate diff initial \
		--dir file://$(MIGRATIONS_DIR) \
		--to file://$(SCHEMA_FILE) \
		--dev-url "sqlite://file?mode=memory" 

seed:
	@DATABASE_URL=$(DBURL) go run ./cmd/seed

up:
	@atlas migrate apply --env local

down:
	@echo 'TODO: not sure how this works just yet'
