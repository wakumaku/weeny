export PROJECT_PATH=$(pwd)
./stop.sh
docker-compose -f ops/docker/docker-compose.yml up -d --build
./logs.sh