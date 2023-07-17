server:
	go run cmd/api/main.go
build-app:
	docker build -f test.Dockerfile -t test-project .