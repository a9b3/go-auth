PORT ?= 5000

all: help

help:
	@echo ""
	@echo "  deps       - Installs dependencies"
	@echo "  dev        - Runs development server     PORT ?= $(PORT)"
	@echo "  test       - Runs tests"
	@echo "  db.start   - Starts the development dbs"
	@echo "  db.stop    - Stops the development dbs"
	@echo ""

deps:
	@dep ensure

dev: deps db.start
	PORT=$(PORT) watcher

test:
	go test -v $$(go list ./... | grep -v /vendor/)

db.start:
	@docker-compose -f docker-compose.yml up -d --remove-orphans

db.stop:
	@docker-compose -f docker-compose.yml down --remove-orphans
