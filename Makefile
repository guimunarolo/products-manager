run:
	docker-compose up --build

reset:
	docker-compose down
	docker volume rm products-manager_db

test:
	docker-compose -f docker-compose.test.yaml run products_api --build
	docker-compose -f docker-compose.test.yaml stop
