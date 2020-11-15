#! /bin/bash

# start crond
crond

LOG_DIR="/home/work/log"

if [ ! -d $LOG_DIR ];then
  mkdir -p $LOG_DIR
fi

LOG_FILE=${LOG_DIR}/app-manager.log

/home/work/bin/app-manager 2>&1 | tee "${LOG_FILE}"