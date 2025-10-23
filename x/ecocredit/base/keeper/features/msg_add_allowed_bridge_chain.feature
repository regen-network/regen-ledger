Feature: Msg/AddAllowedBridgeChain

  A chain can be added to the list of allowed bridge chains:
  - when the authority is the governance account
  - when the chain name doesn't already exist in the state

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