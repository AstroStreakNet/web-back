#/usr/bin/env bash

# create logs dir if it doesn't already exists
mkdir -p logs

LOGFILE=$(date "+%Y-%m-%d-%H%M%S.log")

# run the binary and store output in logs
./webback | tee logs/$LOGFILE

