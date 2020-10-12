run:
	docker-compose up

reset:
	docker-compose down
	docker-compose -f docker-compose.test.yaml down
	docker volume rm products-manager_db
	docker volume rm products-manager_test_db

test:
	docker-compose -f docker-compose.test.yaml run --rm test_products_api
	docker-compose -f docker-compose.test.yaml stop

rebuild:
	docker-compose build
	docker-compose -f docker-compose.test.yaml build

setup:
	docker network create products_manager
