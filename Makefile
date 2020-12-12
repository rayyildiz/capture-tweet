TAG:=1

generate:
	go generate ./...


docker-backend:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/server/tmp/app cmd/server/main.go
	upx cmd/server/tmp/app
	docker build -t rayyildiz.azurecr.io/capturetweet-server:$(TAG) cmd/server
	docker push rayyildiz.azurecr.io/capturetweet-server:$(TAG)
