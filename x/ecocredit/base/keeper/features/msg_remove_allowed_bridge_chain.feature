Feature: Msg/RemoveAllowedBridgeChain

  A chain can be removed from the list of allowed bridge chains:
  - message validation
  - when the authority is the governance account
  - when the chain name already exists in the state

  Rule: Message validation

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