build:
	cd cmd/glofox && GOOS=linux go build -ldflags="-s -w" -o ../../glofox

run:
	go run cmd/glofox/main.go