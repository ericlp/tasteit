.PHONY: mock clear-db clean

mock: mock_data/mockdata.sql
	mkdir -p backend/static/images/
	cp mock_data/*.png backend/static/images/
	cp mock_data/*.jpg backend/static/images/
	docker exec -i tasteit-db-1 psql -U tasteit tasteit < mock_data/mockdata.sql

clear-db:
	echo 'DROP SCHEMA public CASCADE; CREATE SCHEMA public' | docker exec -i tasteit-db-1 psql -U tasteit tasteit

clean: clear-db

