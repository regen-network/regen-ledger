Feature: Msg/MsgUpdateBasketFee

  The basket fee can be updated:
  - message validation
  - when the authority is a governance account address
  - the basket fee is updated

  Rule: Message Validations

    Scenario: a valid message
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message removing basket fee
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      }
      """
      When the message is validated
      Then expect no error


    Scenario: an error is returned if basket fee denom is not formatted
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
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
      When alice attempts to update basket fee with properties
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
      When alice attempts to update basket fee with properties
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

  Rule: The basket fee is updated

    Scenario: the basket fee is updated
      When alice attempts to update basket fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """
      Then expect no error
      And expect basket fee with properties
      """
      {
        "fee": {
          "denom": "uregen",
          "amount": "1000"
        }
      }
      """

    Scenario: the basket fee is removed
      When alice attempts to update basket fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      }
      """
      Then expect no error
      And expect basket fee with properties
      """
      {
        "fee": {}
      }
      """
    
    Scenario: the basket fee is removed when amount is zero
      When alice attempts to update basket fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fee": {
          "denom": "uregen",
          "amount": "0"
        }
      }
      """
      Then expect no error
      And expect basket fee with properties
      """
      {
        "fee": {}
      }
      """
