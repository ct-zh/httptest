build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpfile main.go

buildDocker:
	docker build -t caot1995/httpfile .

.PHONY: build buildDocker