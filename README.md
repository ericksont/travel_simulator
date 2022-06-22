# travel_simulator

## up container
docker compose up -d

## start bash para edição de arquivos
docker exec -it simulator bash

## Start Kafka
cd .docker/kafka/
docker-compose up -d


### Test 2 janelas
docker exec -it kafka_kafka_1 bash

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction

