Feature: Msg/RemoveAllowedDenom

  An allowed denom can be removed:
  - message validations
  - when the authority is a governance account address
  - when the allowed denom exists
  - the denom is removed

  Rule: Message Validations

    Scenario: a valid message
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "uregen"
      }
      """
      When the message is validated
      Then expect no error

  
    Scenario: an error is returned if denom is empty
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      }
      """
      When the message is validated
      Then expect the error "denom cannot be empty: invalid request"

    Scenario: an error is returned if denom is not valid denom
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "1"
      }
      """
      When the message is validated
      Then expect the error "denom: invalid denom: 1: invalid request"


  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      Given an allowed denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      When alice attempts to remove a bank denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom":"uregen"
      }
      """
      Then expect no error
      And expect bank denom is removed "uregen"

    Scenario: The authority is not a governance account address
      When alice attempts to remove a bank denom with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "denom":"uregen"
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The bank denom exists

    Scenario: The bank denom exists
      Given an allowed denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      When alice attempts to remove a bank denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom":"uregen"
      }
      """
      Then expect no error
      And expect bank denom is removed "uregen"

    Scenario: The bank denom does not exist
      When alice attempts to remove a bank denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom":"uregen"
      }
      """
      Then expect the error "allowed denom uregen: not found"

  Rule: Event is emitted

    Scenario: EventRemoveAllowedDenom is emitted
      Given an allowed denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      When alice attempts to remove a bank denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom":"uregen"
      }
      """
      Then expect no error
      And expect event with properties
      """
      {
        "denom": "uregen"
      }
      """