Feature: MsgCreate

  Scenario: a valid message
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with description
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "description": "Nature-Based Carbon Token",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario Outline: a valid message with disable auto-retire
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "disable_auto_retire": <disable-auto-retire>,
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ]
    }
    """
    When the message is validated
    Then expect no error

    Examples:
      | description          | disable-auto-retire |
      | auto-retire enabled  | false               |
      | auto-retire disabled | true                |

  Scenario Outline: a valid message with data criteria
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "date_criteria": <date-criteria>
    }
    """
    When the message is validated
    Then expect no error

    Examples:
      | description        | date-criteria                              |
      | minimum start date | {"min_start_date": "2012-01-01T00:00:00Z"} |
      | start date window  | {"start_date_window": "315360000s"}        |
      | years in the past  | {"years_in_the_past": 10}                  |

  Scenario: a valid message with fee
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "fee": [
        {
          "denom": "uregen",
          "amount": "20000000"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if curator is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "malformed curator address: empty address string is not allowed: invalid address"

  Scenario: an error is returned if curator is not a bech32 address
    Given the message
    """
    {
      "curator": "foo"
    }
    """
    When the message is validated
    Then expect the error "malformed curator address: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if name is empty
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "name cannot be empty: invalid request"

  Scenario: an error is returned if name does not start with an alphabetic character
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "1CT"
    }
    """
    When the message is validated
    Then expect the error "name must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: invalid request"

  Scenario: an error is returned if name includes non-alphanumeric characters
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "C T"
    }
    """
    When the message is validated
    Then expect the error "name must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: invalid request"

  Scenario: an error is returned if name length is less than three characters
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "CT"
    }
    """
    When the message is validated
    Then expect the error "name must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: invalid request"

  Scenario: an error is returned if name length is greater than eight characters
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "CARBONTOKEN"
    }
    """
    When the message is validated
    Then expect the error "name must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: invalid request"

  Scenario: an error is returned if description length is greater than 256 characters
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."
    }
    """
    When the message is validated
    Then expect the error "description length cannot be greater than 256 characters: invalid request"

  Scenario: an error is returned if credit type abbreviation is empty
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT"
    }
    """
    When the message is validated
    Then expect the error "credit type abbreviation cannot be empty: parse error: invalid request"

  Scenario: an error is returned if credit type abbreviation is not formatted
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "foobar"
    }
    """
    When the message is validated
    Then expect the error "credit type abbreviation must be 1-3 uppercase latin letters: got foobar: parse error: invalid request"

  Scenario: an error is returned if allowed credit classes is empty
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C"
    }
    """
    When the message is validated
    Then expect the error "allowed classes cannot be empty: invalid request"

  Scenario: an error is returned if an allowed credit class is empty
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        ""
      ]
    }
    """
    When the message is validated
    Then expect the error "allowed_classes[0] is not a valid class ID: class id cannot be empty: parse error: invalid request"

  Scenario: an error is returned if an allowed credit class is not formatted
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "foo"
      ]
    }
    """
    When the message is validated
    Then expect the error "allowed_classes[0] is not a valid class ID: class ID didn't match the format: expected A00, got foo: parse error: invalid request"

  Scenario Outline: an error is returned if more than one data criteria is provided
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "date_criteria": <date-criteria>
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
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "date_criteria": {
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
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "date_criteria": {
        "start_date_window": "23h"
      }
    }
    """
    When the message is validated
    Then expect the error "invalid date criteria: start_date_window must be at least 1 day: invalid request"

  Scenario: an error is returned if fee denom is empty
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "fee": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "invalid denom: "

  Scenario: an error is returned if fee denom is not formatted
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "fee": [
        {
          "denom": "1"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "invalid denom: 1"

  Scenario: an error is returned if fee amount is not positive
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "fee": [
        {
          "denom": "uregen",
          "amount": "-1"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "coin -1uregen amount is not positive"

  Scenario: an error is returned if fee length is greater than 1
    Given the message
    """
    {
      "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "allowed_classes": [
        "C01"
      ],
      "fee": [
        {
          "denom": "uregen",
          "amount": "20000000"
        },
        {
          "denom": "uatom",
          "amount": "20000000"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "more than one fee is not allowed: invalid request"
