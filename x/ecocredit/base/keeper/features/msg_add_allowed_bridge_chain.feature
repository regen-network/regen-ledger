Feature: Msg/AddAllowedBridgeChain

  A chain can be added to the list of allowed bridge chains:
  - when the authority is the governance account
  - when the chain name doesn't already exist in the state


  Rule: The authority address is the governance account

    Background:
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"

    Scenario: the authority is the governance account
      When alice attempts to add allowed bridge chain with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "chain_name": "polygon"
      }
      """
      Then expect no error

    Scenario: the authority is not the governance account
      When alice attempts to add allowed bridge chain with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "chain_name": "polygon"
      }
      """
      Then expect the error contains "invalid authority"


  Rule: The chain name must not already exist

    Scenario: there is no chain name in state
      When  alice attempts to add chain name "polygon"
      Then expect no error
      And expect chain name "polygon" to exist

    Scenario: chain name already exists in state
      Given allowed chain name "polygon"
      When  alice attempts to add chain name "polygon"
      Then expect the error contains "could not insert chain name polygon"