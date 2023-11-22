Feature: Msg/RemoveAllowedBridgeChain

  A chain can be removed from the list of allowed bridge chains:
  - when the authority is the governance account
  - when the chain name already exists in the state


  Rule: The authority address is the governance account

    Background:
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And the chain name "polygon"

    Scenario: the authority is the governance account
      When alice attempts to remove allowed bridge chain with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "chain_name": "polygon"
      }
      """
      Then expect no error

    Scenario: the authority is not the governance account
      When alice attempts to remove allowed bridge chain with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "chain_name": "polygon"
      }
      """
      Then expect the error contains "invalid authority"


  Rule: The chain name must exist

    Scenario: the chain name exists in state
      Given the chain name "polygon"
      When  alice attempts to remove chain name "polygon"
      Then expect no error
      And expect chain name "polygon" to not exist

  # no failing scenario, ORM delete does not error when the entry does not exist