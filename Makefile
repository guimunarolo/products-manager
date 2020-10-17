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

build:
	docker-compose build
	docker-compose -f docker-compose.test.yaml build

setup:
	docker network create products_manager

protos:
	protoc -I calculator-service/calculator/protos calculator-service/calculator/protos/calculator.proto --go_out=plugins=grpc:calculator-service/calculator
	python -m grpc_tools.protoc -I products-api/products_api/protos \
		--python_out=products-api/products_api/protos/calculator \
		--grpc_python_out=products-api/products_api/protos/calculator \
		products-api/products_api/protos/calculator.proto
