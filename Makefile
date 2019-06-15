export PROJECT_NAME := weeny
export PROJECT_PATH := $(shell pwd)

start: stop
	docker-compose -p ${PROJECT_NAME} -f ${PROJECT_PATH}/ops/docker/docker-compose.yml up -d --build

stop:
	docker-compose -p ${PROJECT_NAME} -f ${PROJECT_PATH}/ops/docker/docker-compose.yml down

logs:
	docker-compose -p ${PROJECT_NAME} -f ${PROJECT_PATH}/ops/docker/docker-compose.yml logs -f

test:
	docker exec -w ${PROJECT_PATH}/src ${PROJECT_NAME}_workspace_1 go test -race ./...