Feature: Msg/RemoveAllowedDenom

  An allowed denom can be removed:
  - when the authority is a governance account address
  - when the allowed denom exists
  - the denom is removed

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