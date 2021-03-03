build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httptest main.go

buildDocker:
	docker build -t caot1995/httptest .

.PHONY: build buildDocker