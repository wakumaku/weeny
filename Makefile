export PROJECT_NAME := weeny
export PROJECT_PATH := $(shell pwd)

start:
	${PROJECT_PATH}/ops/scripts/stop.sh
	${PROJECT_PATH}/ops/scripts/start.sh
	${PROJECT_PATH}/ops/scripts/logs.sh

stop:
	${PROJECT_PATH}/ops/scripts/stop.sh

logs:
	${PROJECT_PATH}/ops/scripts/logs.sh