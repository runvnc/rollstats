#!/usr/bin/fish

sed -i 's/TxId/TxID/g' rollstats.teal

set ID (goal app create --approval-prog rollstats.teal --clear-prog clear.teal --creator $APP_CREATOR --app-arg "str:init" --global-byteslices 5 --global-ints 5 --local-byteslices 0 --local-ints 0 | tail -1 | rev | cut -f 1 -d' ' | rev)
echo $ID
set -Ux APP_ID $ID
./read.sh

