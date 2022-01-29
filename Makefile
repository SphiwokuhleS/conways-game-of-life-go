.PHONY: run_endpoint_tests create_test_database

run_api_tests: create_test_database run_endpoint_tests remove_test_db

run_endpoint_tests:
	@echo "Running endpoint tests"
	cd api && go test -v

create_test_database:
	@echo "Creating test database"
	chmod +x test_db_script.sh ; ./test_db_script.sh

remove_test_db:
	@echo "Removing test database"
	rm -rf api/conways.db