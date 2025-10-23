Feature: Msg/Create

  A basket can be created:
  - message validations
  - when the basket name is unique
  - when the required basket fee is not set and no fee is provided
  - when the basket fee denom matches the required basket fee denom
  - when the basket fee amount is greater than or equal to the required basket fee amount
  - when the admin balance is greater than or equal to the required basket fee amount
  - when the basket includes a credit type that exists
  - when the basket criteria includes credit classes that exist
  - when the basket criteria includes credit classes that match the credit type
  - when the basket criteria includes optional date criteria
  - the user token balance is updated and only the minimum fee is taken
  - the basket denom is formatted with a prefix based on credit type precision
  - the response includes the basket denom

  Rule: Message validations

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

    
    Scenario: an error is returned if name is empty
      Given the message
      """
      {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      }
      """
      When the message is validated
      Then expect the error "name: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if name does not start with an alphabetic character
      Given the message
      """
      {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "name": "1CT"
      }
      """
      When the message is validated
      Then expect the error "name: must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: parse error: invalid request"

    Scenario: an error is returned if name includes non-alphanumeric characters
      Given the message
      """
      {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "name": "C T"
      }
      """
      When the message is validated
      Then expect the error "name: must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: parse error: invalid request"

    Scenario: an error is returned if name length is less than three characters
      Given the message
      """
      {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "name": "CT"
      }
      """
      When the message is validated
      Then expect the error "name: must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: parse error: invalid request"

    Scenario: an error is returned if name length is greater than eight characters
      Given the message
      """
      {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "name": "CARBONTOKEN"
      }
      """
      When the message is validated
      Then expect the error "name: must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: parse error: invalid request"

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
      Then expect the error "credit type abbrev: empty string is not allowed: parse error: invalid request"

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
      Then expect the error "credit type abbrev: must be 1-3 uppercase alphabetic characters: parse error: invalid request"

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
      Then expect the error "allowed classes [0]: empty string is not allowed: parse error: invalid request"

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
      Then expect the error "allowed classes [0]: expected format <credit-type-abbrev><class-sequence>: parse error: invalid request"

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


  Rule: The basket name must be unique

    Background:
      Given a credit type

    Scenario: basket name is unique
      When alice attempts to create a basket with name "NCT"
      Then expect no error

    Scenario: basket name is not unique
      Given a basket with name "NCT"
      When alice attempts to create a basket with name "NCT"
      Then expect the error "basket with name NCT already exists: unique key violation"

  Rule: The basket fee is not required if the required basket fee is not set

    Background:
      Given a credit type

    Scenario: basket fee provided and required basket fee not set
      Given alice has a token balance "20regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee not provided and required basket fee not set
      When alice attempts to create a basket with no fee
      Then expect no error

    # no failing scenario - basket fee is not required if required basket fee is not set

  Rule: The basket fee must match the required basket fee denom

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario: basket fee matches required basket fee denom
      Given required basket fee "20regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee does not match required basket fee denom
      Given required basket fee "20regen"
      When alice attempts to create a basket with fee "20atom"
      Then expect the error "fee must be 20regen, got 20atom: insufficient fee"

    Scenario: basket fee not provided and required basket fee set
      Given required basket fee "20regen"
      When alice attempts to create a basket with no fee
      Then expect the error "fee cannot be empty: must be 20regen: insufficient fee"

  Rule: The basket fee must be greater than or equal to the required basket fee

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario Outline: basket fee is greater than or equal to required basket fee amount
      Given required basket fee "20regen"
      When alice attempts to create a basket with fee "<basket-fee>"
      Then expect no error

      Examples:
        | description  | basket-fee |
        | greater than | 30regen    |
        | equal to     | 20regen    |

    Scenario: basket fee is less than required basket fee amount
      Given required basket fee "20regen"
      When alice attempts to create a basket with fee "10regen"
      Then expect the error "fee must be 20regen, got 10regen: insufficient fee"

  Rule: The admin must have a balance greater than or equal to basket fee amount

    Background:
      Given a credit type
      And required basket fee "20regen"

    Scenario Outline: admin balance is greater than or equal to required basket fee amount
      Given alice has a token balance "<token-balance>"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

      Examples:
        | description  | token-balance |
        | greater than | 30regen       |
        | equal to     | 20regen       |

    Scenario: admin balance is less than required basket fee amount
      Given alice has a token balance "10regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect the error "insufficient balance 10 for bank denom regen: insufficient funds"

  Rule: The basket must include a credit type that exists

    Background:
      Given a credit type with abbreviation "C"

    Scenario: basket credit type exists
      When alice attempts to create a basket with credit type "C"
      Then expect no error

    Scenario: basket credit type does not exist
      When alice attempts to create a basket with credit type "F"
      Then expect the error "could not get credit type with abbreviation F: not found: invalid request"

  Rule: The basket criteria must include a credit class that exists

    Background:
      Given a credit type with abbreviation "C"

    Scenario: basket criteria credit class exists
      Given a credit class with id "C01"
      When alice attempts to create a basket with allowed class "C01"
      Then expect no error

    Scenario: basket criteria credit class does not exist
      When alice attempts to create a basket with allowed class "C01"
      Then expect the error "could not get credit class C01: not found: invalid request"

  Rule: The basket criteria must include a credit class that matches the credit type

    Background:
      Given a credit type with abbreviation "C"

    Scenario: basket criteria credit class matches credit type
      Given a credit class with id "C01"
      When alice attempts to create a basket with credit type "C" and allowed class "C01"
      Then expect no error

    Scenario: basket criteria credit class does not match credit type
      Given a credit class with id "BIO01"
      When alice attempts to create a basket with credit type "C" and allowed class "BIO01"
      Then expect the error "basket specified credit type C, but class BIO01 is of type BIO: invalid request"

  Rule: The basket criteria may include optional start date criteria

    Background:
      Given a credit type

    Scenario: basket criteria minimum start date
      When alice attempts to create a basket with minimum start date "2020-01-01"
      Then expect minimum start date "2020-01-01"

    Scenario: basket criteria start date window
      When alice attempts to create a basket with start date window "43800h"
      Then expect start date window "43800h"

    Scenario: basket criteria years in the past
      When alice attempts to create a basket with years in the past "10"
      Then expect years in the past "10"

  Rule: The user token balance is updated and only the minimum fee is taken

    Background:
      Given a credit type
      And required basket fee "20regen"
      And alice has a token balance "40regen"

    Scenario Outline: user token balance is updated
      When alice attempts to create a basket with fee "<basket-fee>"
      Then expect alice token balance "<token-balance>"

      Examples:
        | description  | basket-fee | token-balance |
        | greater than | 40regen    | 20regen       |
        | equal to     | 20regen    | 20regen       |

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The basket denom is formatted with a prefix based on credit type precision

    Scenario Outline: basket denom is formatted using credit type precision
      Given a credit type with abbreviation "C" and precision "<precision>"
      When alice attempts to create a basket with name "NCT" and credit type "C"
      Then expect the response
      """
      {
        "basket_denom": "<basket-denom>"
      }
      """

      Examples:
        | description | precision | basket-denom |
        | no prefix   | 0         | eco.C.NCT    |
        | d (deci)    | 1         | eco.dC.NCT   |
        | c (centi)   | 2         | eco.cC.NCT   |
        | m (milli)   | 3         | eco.mC.NCT   |
        | u (micro)   | 6         | eco.uC.NCT   |
        | n (nano)    | 9         | eco.nC.NCT   |
        | p (pico)    | 12        | eco.pC.NCT   |
        | f (femto)   | 15        | eco.fC.NCT   |
        | a (atto)    | 18        | eco.aC.NCT   |
        | z (zepto)   | 21        | eco.zC.NCT   |
        | y (yocto)   | 24        | eco.yC.NCT   |

    # no failing scenario - credit type precision should always be a valid SI prefix

  Rule: The message response includes basket denom

    Scenario: message response includes the basket denom
      Given a credit type with abbreviation "C" and precision "6"
      When alice attempts to create a basket with name "NCT"
      Then expect the response
      """
      {
        "basket_denom": "eco.uC.NCT"
      }
      """

    # no failing scenario - response should always be empty when message execution fails

  Rule: Event is emitted

    Scenario: EventCreate is emitted
      Given a credit type with abbreviation "C" and precision "6"
      And alice's address "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      When alice attempts to create a basket with name "NCT"
      Then expect event with properties
      """
      {
        "curator": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "basket_denom": "eco.uC.NCT"
      }
      """