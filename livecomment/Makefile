GATEWAY_URL=http://localhost:8080/subscribe?video=
DISPATCHER_URL=http://localhost:8081/comment

build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down --remove-orphans

generate-comments:
	GATEWAY_URL=$(GATEWAY_URL) DISPATCHER_URL=$(DISPATCHER_URL) go run . generator
