run:
	docker-compose up --build

reset:
	docker-compose down
	docker volume rm products-manager_db

test:
	docker-compose -f docker-compose.test.yaml run test_products_api
	docker-compose -f docker-compose.test.yaml stop
