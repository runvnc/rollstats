#!/bin/bash
rawo=$(goal app read --app-id $APP_ID --global | jq '.')
#echo $rawo
#goal app read --app-id $APP_ID --global | jq 'del(.dbg)'
echo $rawo | jq -r '.dbg.tb | split("\r") | .[]'
