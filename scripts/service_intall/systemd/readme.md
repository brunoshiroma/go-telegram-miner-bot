# Scripts for install OS with systemd
Install
```bash
#on dir with scripts
sudo ./install_service.sh

#after that you need to configure the json file
nano /usr/share/lib/go-telegram-miner-bot/miners_bot.json

#and finally start the service
sudo systemctl start go-telegram-miner-bot.service

#you can get the logs on
sudo tail -f /var/log/go-telegram-miner-bot/go-telegram-miner-bot.log
```

Uninstall
```bash
#on dir with scripts
sudo ./uninstall_service.sh

#delete the logs, if you want
sudo rm -rf /var/log/go-telegram-miner-bot
```

## Why not use StandardOutput=file:...
It´s dont work on raspian ( not in the time of creation of the scripts)  
More context in [stackoverflox](https://stackoverflow.com/questions/37585758/how-to-redirect-output-of-systemd-service-to-a-file)
