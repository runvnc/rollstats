#!/bin/bash
echo "Creator: $APP_CREATOR  ID: $APP_ID"
goal app call --from=$APP_CREATOR --app-id $APP_ID --app-arg "str:test"
./read.sh

