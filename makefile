http:
	go run main.go http -c=config.yaml

test:
	go test -v -cover -timeout 30s ./...