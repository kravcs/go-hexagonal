build:
	cd cmd/hexa && GOOS=linux go build -ldflags="-s -w" -o ../../hexa

run:
	go run cmd/hexa/main.go