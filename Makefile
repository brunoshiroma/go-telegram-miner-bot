build:
	go build cmd/go-telegram-miner-bot/main.go

clean:
	go clean -testcache

clean-all:
	go clean -testcache -modcache

test:
	go test -v -coverprofile cover.out  ./...
	

test-cover: test
	go tool cover -html=cover.out

vet:
	go vet -assign -printf -unmarshal -unreachable -unusedresult -structtag -tests ./...

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.41.1 golangci-lint run -v

tidy:
	go mod tidy