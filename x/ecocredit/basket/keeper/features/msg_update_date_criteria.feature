Feature: Msg/UpdateDateCriteria

  Basket date criteria can be updated:
  - message validations
  - when the basket exists
  - when the authority is the governance account
  - when the basket date criteria is empty
  - when the basket date criteria includes minimum start date
  - when the basket date criteria includes start date window
  - when the basket date criteria includes years in the past

  Rule: Message Validations

    Scenario: a valid message
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT"
      }
      """
      When the message is validated
      Then expect no error

    Scenario Outline: a valid message with new data criteria
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT",
        "new_date_criteria": <date-criteria>
      }
      """
      When the message is validated
      Then expect no error

      Examples:
        | description        | date-criteria                              |
        | minimum start date | {"min_start_date": "2012-01-01T00:00:00Z"} |
        | start date window  | {"start_date_window": "315360000s"}        |
        | years in the past  | {"years_in_the_past": 10}                  |

    
    Scenario: an error is returned if basket denom is not formatted
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "foo"
      }
      """
      When the message is validated
      Then expect the error "invalid basket denom: expected format eco.<exponent-prefix><credit-type-abbrev>.<name>: parse error: invalid request"

    Scenario Outline: an error is returned if more than one data criteria is provided
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT",
        "new_date_criteria": <date-criteria>
      }
      """
      When the message is validated
      Then expect the error "invalid date criteria: only one of min_start_date, start_date_window, or years_in_the_past must be set: invalid request"

      Examples:
        | description        | date-criteria                                                                 |
        | date and window    | {"min_start_date": "2012-01-01T00:00:00Z", "start_date_window": "315360000s"} |
        | window and years   | {"start_date_window": "315360000s", "years_in_the_past": 10}                  |
        | years and date     | {"years_in_the_past": 10, "min_start_date": "2012-01-01T00:00:00Z"}           |

    Scenario: an error is returned if minimum start date is before 1900-01-01
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "min_start_date": "1899-01-01T00:00:00Z"
        }
      }
      """
      When the message is validated
      Then expect the error "invalid date criteria: min_start_date must be after 1900-01-01: invalid request"

    Scenario: an error is returned if start date window is less than one day
      Given the message
      """
      {
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "start_date_window": "23h"
        }
      }
      """
      When the message is validated
      Then expect the error "invalid date criteria: start_date_window must be at least 1 day: invalid request"

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
