test:
	docker-compose -f docker-compose.test.yaml run products_api
	docker-compose -f docker-compose.test.yaml stop
