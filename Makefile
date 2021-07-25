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

lint:
	go vet -assign -printf -unmarshal -unreachable -unusedresult -structtag -tests ./...