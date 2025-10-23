Feature: Msg/AddAllowedDenom

  An allowed denom can be added:
  - message validations
  - when the authority is a governance account address
  - when the denom does not exist
  - the allowed denom is added

  Rule: Message Validations

    Scenario: a valid message
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "bank_denom": "uregen",
        "display_denom": "REGEN",
        "exponent": 6
      }
      """
      When the message is validated
      Then expect no error

    
    Scenario: an error is returned if bank denom is empty
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      }
      """
      When the message is validated
      Then expect the error "bank denom cannot be empty: parse error: invalid request"
    
    Scenario: an error is returned if bank denom is not valid denom
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "bank_denom": "1"
      }
      """
      When the message is validated
      Then expect the error "bank denom: invalid denom: 1: parse error: invalid request"

    Scenario: an error is returned if bank display denom is empty
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "bank_denom": "uregen"
      }
      """
      When the message is validated
      Then expect the error "display denom cannot be empty: parse error: invalid request"



  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect no error

    Scenario: The authority is not a governance account address
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The allowed denom does not exists

    Scenario: The allowed denom does not exist
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect no error

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
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect the error "bank denom uregen already exists: conflict"

    Scenario: The display denom exists
      Given an allowed denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uatom",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect the error "display denom REGEN already exists: conflict"

  Rule: Event is emitted

    Scenario: EventAllowDenom is emitted
      When alice attempts to add a denom with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "bank_denom":"uregen",
        "display_denom":"REGEN",
        "exponent":6
      }
      """
      Then expect event with properties
      """
      {
        "denom": "uregen"
      }
      """