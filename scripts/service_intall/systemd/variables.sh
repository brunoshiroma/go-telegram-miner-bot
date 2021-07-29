#!/bin/sh

SERVICE_DIR=/usr/share/lib/go-telegram-miner-bot
SERVICE_BINARY=../../../go-telegram-miner-bot
SERVICE_USER=telegram-miner-bot
SERVICE_CONFIG=../../../miners_bot_sample.json
SERVICE_RSYSLOG_CONFIG=/etc/rsyslog.d/go-telegram-miner-bot.conf

SERVICE_FILE_SYSTEMD=/etc/systemd/system/go-telegram-miner-bot.service