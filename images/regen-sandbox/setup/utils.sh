log_response() {
  jq -r "if .code == 0 then \"INFO:   TxHash: \(.txhash)\" else \"ERROR: (code \(.code)) \(.raw_log)\" end"
}

ADDR1=$(regen keys show -a addr1)
ADDR2=$(regen keys show -a addr2)
ADDR3=$(regen keys show -a addr3)
ADDR4=$(regen keys show -a addr4)
ADDR5=$(regen keys show -a addr5)
