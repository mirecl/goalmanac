export POSTGRES_PASSWORD=666
export POSTGRES_USER=otus
export POSTGRES_DB=otus

http:
	go run main.go http -c=./config/config.yaml

test:
	go test -cover -timeout 30s ./...
db:
	docker-compose up
	