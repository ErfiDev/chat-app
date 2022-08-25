build_server:
	docker build -t chat_server .
run_server:
	docker run --rm -p 5000:5000 chat_server

build_client:
	go build -o client cmd/client/main.go
