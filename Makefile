PORT ?= 5000

all: help

help:
	@echo ""
	@echo "  dev              - Runs development server          PORT ?= $(PORT)"
	@echo "  test             - Runs tests"
	@echo "  test.watch       - Runs tests in watch mode"
	@echo ""
	@echo "  (You probably don't need to run these manually.)"
	@echo "  deps             - Installs dependencies"
	@echo "  db.start         - Starts the development dbs"
	@echo "  db.stop          - Stops the development dbs"
	@echo "  db.starttest     - Starts the test dbs"
	@echo "  db.stoptest      - Stops the test dbs"
	@echo "  migrate          - Runs migration"
	@echo "  migrate.test     - Runs migration on test dbs"
	@echo ""

dev: deps db.start migrate
	PORT=$(PORT) watcher

test: db.starttest migrate.test
	go test -v $$(go list ./... | grep -v /vendor/)

test.watch: db.starttest migrate.test
	@fswatch . | (while read; do \
		echo ""; \
		echo "********************************************"; \
		echo "RUNNING TESTS"; \
		echo "********************************************"; \
		echo ""; \
		go test -v $$(go list ./... | grep -v /vendor/); \
	done)

deps:
	@dep ensure

db.start:
	@docker-compose -f docker-compose.yml up -d && sleep 2

db.stop:
	@docker-compose -f docker-compose.yml down

db.starttest:
	@docker-compose -f docker-compose.test.yml up -d && sleep 2

db.stoptest:
	@docker-compose -f docker-compose.test.yml down

migrate:
	@./scripts/migrate.sh -c .env

migrate.test:
	@./scripts/migrate.sh -c .test.env
