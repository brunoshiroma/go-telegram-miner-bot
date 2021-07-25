# Telegram bot for get miners infos
Simple bot for telegram, for get stats of your miners

## API of miners
### excavator
https://github.com/nicehash/excavator/tree/master/api#-algorithmlist

### ethminer
https://github.com/ethereum-mining/ethminer/blob/master/docs/API_DOCUMENTATION.md#miner_getstatdetail

### For more syntax of response.jsonPath
https://github.com/tidwall/gjson

### Executing
Rename the file `miners_bot_sample.json` to `miners_bot.json`, configure the .json file  
and then simple run the program
```bash
go-telegram-miner-bot
```

### Compiling
To compile, use [golang 1.16.x](https://golang.org/dl/)

### Create a new bot
To create a new bot and get a token [doc](docs/telegram_bot_register.md)