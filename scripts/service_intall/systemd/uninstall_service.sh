#!/bin/bash

source ./variables.sh

# delete user if exists
if id -u "$SERVICE_USER" >/dev/null 2>&1; then
    userdel $SERVICE_USER
fi

# delete dir if exists
if [ -d $SERVICE_DIR ]; then
    rm -rf $SERVICE_DIR
fi

# delete systemd file
if [ -f $SERVICE_FILE_SYSTEMD ]; then
    rm $SERVICE_FILE_SYSTEMD
fi

 delete rsyslog file
if [ -f $SERVICE_RSYSLOG_CONFIG ]; then
    rm $SERVICE_RSYSLOG_CONFIG
fi

systemctl restart rsyslog

echo 'OK'