include backend/.env

.PHONY: mock clear-db clean reset new-migration run-migrations

TASTEIT_DB_URL = 'postgres://$(db_user):$(db_password)@localhost:5432/$(db_name)?sslmode=disable'
MIGRATION_PATH = 'backend/internal/db/migrations'
TASTEIT_DB_DOCKER_NAME = 'tasteit-db-1'

GAMMA_DB_DOCKER_NAME = 'tasteit-gamma-db-1'
GAMMA_DB_USER= 'gamma'
GAMMA_DB_NAME= 'gamma'

mock: mock_data/mockdata.sql
	mkdir -p backend/static/images/
	cp mock_data/*.png backend/static/images/
	cp mock_data/*.jpg backend/static/images/
	docker exec -i $(TASTEIT_DB_DOCKER_NAME) psql -U $(db_name) $(db_user) < mock_data/mockdata.sql

clear-db:
	echo 'DROP SCHEMA public CASCADE; CREATE SCHEMA public' | docker exec -i $(TASTEIT_DB_DOCKER_NAME) psql -U $(db_name) $(db_user)

clean: clear-db

new-migration:
	migrate -database ${TASTEIT_DB_URL} -path $(MIGRATION_PATH) create -ext sql -dir $(MIGRATION_PATH) $(mig_name_arg)

run-migrations:
	migrate -database $(TASTEIT_DB_URL) -path $(MIGRATION_PATH) up

setup-gamma-client:
	docker exec -i $(GAMMA_DB_DOCKER_NAME) psql -U $(GAMMA_DB_NAME) $(GAMMA_DB_USER) < mock_data/gammaClient.sql
	
reset-setup-db:
	make clean
	make run-migrations
	make mock
	make setup-gamma-client