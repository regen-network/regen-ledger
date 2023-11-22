Feature: Msg/UpdateDateCriteria

  Basket date criteria can be updated:
  - when the basket exists
  - when the authority is the governance account
  - when the basket date criteria is empty
  - when the basket date criteria includes minimum start date
  - when the basket date criteria includes start date window
  - when the basket date criteria includes years in the past

  Rule: The basket must exist

    Scenario: the basket does not exist
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect the error "basket with denom eco.uC.NCT does not exist: not found"

    Scenario: the basket exists
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect no error

  Rule: The authority address is the governance account

    Scenario: the authority is not the governance account
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect the error "invalid authority: expected regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68, got regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s: expected gov account as only signer for proposal message"

    Scenario: the authority is the governance account
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect no error

  Rule: The basket date criteria may be empty

    Scenario: new date criteria empty
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT"
      }
      """
      Then expect no date criteria

  Rule: The basket date criteria may include minimum start date

    Scenario: new date criteria minimum start date
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "min_start_date": "2020-01-01T00:00:00Z"
        }
      }
      """
      Then expect date criteria
      """
      {
        "min_start_date": "2020-01-01T00:00:00Z"
      }
      """

  Rule: The basket date criteria may include start date window

    Scenario: new date criteria start date window
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "start_date_window": "43800h"
        }
      }
      """
      Then expect date criteria
      """
      {
        "start_date_window": "43800h"
      }
      """

  Rule: The basket date criteria may include years in the past

    Scenario: new date criteria years in the past
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect date criteria
      """
      {
        "years_in_the_past": 10
      }
      """

  Rule: Event is emitted

    Scenario: EventUpdateDateCriteria is emitted
      Given the authority address "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      And a basket with denom "eco.uC.NCT"
      When alice attempts to update date criteria with message
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
      """
      Then expect event with properties
      """
      {
        "denom": "eco.uC.NCT"
      }
      """
