Feature: Msg/MsgUpdateClassFee

  The class fee can be updated:
  - message validation
  - when the authority is a governance account address
  - the class fee is updated

  Rule: Message validation

    Scenario: a valid message
      Given the message
      """
      {
        "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "fee": {
          "denom":"uregen",
          "amount":"1000"
        }
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message removing the fee
      Given the message
      """
      {
        "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: an error is returned if fee denom is not a valid bank denom
      Given the message
      """
      {
          "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "fee": {
            "denom": "1",
            "amount": "1000"
          }
      }
      """
      When the message is validated
      Then expect the error "invalid denom: 1: invalid request"

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to update class fee with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """
      Then expect no error

    Scenario: the authority is not a governance account
      When alice attempts to update class fee with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The class fee is updated

    Scenario: the class fee is updated
      When alice attempts to update class fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """
      Then expect class fee with properties
      """
      {
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """

    Scenario: the class fee is removed
      When alice attempts to update class fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      }
      """
      Then expect class fee with properties
      """
      {
        "fee": {}
      }
      """

    Scenario: the class fee is removed when amount is zero
      When alice attempts to update class fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fee":{
          "denom": "uregen",
          "amount": "0"
        }
      }
      """
      Then expect class fee with properties
      """
      {
        "fee": {}
      }
      """
