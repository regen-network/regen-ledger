Feature: Msg/MsgUpdateClassFee

  The class fee can be updated:
  - when the authority is a governance account address
  - the class fee is updated

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
