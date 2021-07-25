build:
	go build cmd/go-telegram-miner-bot/main.go

clean:
    go clean -testcache

clean-all:
	go clean -testcache -modcache