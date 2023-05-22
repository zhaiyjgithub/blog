.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/authorizer authorizer/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/sendMessage sendMessage/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/worker senderWorker/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
