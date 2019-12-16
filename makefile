http:
	go run main.go http -c=./config/config.yaml

test:
	go test -cover -timeout 30s ./...
	