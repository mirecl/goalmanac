export POSTGRES_PASSWORD=666
export POSTGRES_USER=otus
export POSTGRES_DB=otus

export RABBITMQ_DEFAULT_USER=rabbitmq
export RABBITMQ_DEFAULT_PASS=rabbitmq

http:
	go run main.go http -c=./config/config.yaml
grpc:
	go run main.go grpc -c=./config/config.yaml
test:
	go test -cover -timeout 30s ./...
service:
	docker-compose up
	