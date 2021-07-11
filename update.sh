#!/usr/bin/fish

echo $TEAL
sed -i 's/TxId/TxID/g' rollstats.teal

goal app update --app-id $APP_ID --from $APP_CREATOR --approval-prog $APPROVE --clear-prog clear.teal --app-arg "str:init" | tail -1 | rev | cut -f 1 -d' ' | rev
