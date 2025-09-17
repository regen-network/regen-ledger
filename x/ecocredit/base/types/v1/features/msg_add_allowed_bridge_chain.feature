Feature: MsgAddAllowedBridgeChain

Scenario: a valid message
Given the message
"""
{
"authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
"chain_name": "polygon"
}
"""
When the message is validated
Then expect no error

Scenario: an error is returned if chain name is empty
Given the message
"""
{
"authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
}
"""
When the message is validated
Then expect the error "chain_name cannot be empty: invalid request"
