#!/bin/bash

source ./variables.sh

# first check binary exists
if [ ! -f $SERVICE_BINARY ]; then
    echo 'service binary missing, run make build first'
fi

# create user if not exists
if id -u "$SERVICE_USER" >/dev/null 2>&1; then
    echo "$SERVICE_USER EXISTS"
else
    useradd $SERVICE_USER -M -s /sbin/nologin
fi

# create dir
if [ ! -d $SERVICE_DIR ]; then
    mkdir -p $SERVICE_DIR
fi

# copy binary
cp $SERVICE_BINARY $SERVICE_DIR

# copy the config
cp $SERVICE_CONFIG $SERVICE_DIR/miners_bot.json

cp go-telegram-miner-bot.service $SERVICE_FILE_SYSTEMD

cp go-telegram-miner-bot.conf $SERVICE_RSYSLOG_CONFIG

systemctl restart rsyslog

echo 'OK'