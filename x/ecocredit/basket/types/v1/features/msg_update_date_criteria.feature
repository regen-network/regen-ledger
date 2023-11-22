Feature: MsgUpdateDateCriteria

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

  Scenario: an error is returned if authority is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "invalid authority address: empty address string is not allowed: invalid address"

  Scenario: an error is returned if authority is not a bech32 address
    Given the message
    """
    {
      "authority": "foo"
    }
    """
    When the message is validated
    Then expect the error "invalid authority address: decoding bech32 failed: invalid bech32 string length 3: invalid address"

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

  Scenario: a valid amino message
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "denom": "eco.uC.NCT",
      "new_date_criteria": {
        "years_in_the_past": 10
      }
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen.basket/MsgUpdateDateCriteria",
      "value":{
        "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT",
        "new_date_criteria": {
          "years_in_the_past": 10
        }
      }
    }
    """
