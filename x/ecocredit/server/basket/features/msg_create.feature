Feature: MsgCreate

  A basket can be created:
  - when the basket name is unique
  - when the minimum basket fee is not set
  - when the basket fee denom matches the minimum basket fee denom
  - when the basket fee amount is greater than or equal to the minimum basket fee amount
  - when the user has a balance greater than or equal to the minimum basket fee amount
  - when the basket includes a valid credit type
  - when the basket allowed classes includes valid credit classes
  - when the exponent is greater than or equal to the credit type precision
  - the user token balance is updated
  - the response includes the basket denom

  Rule: The basket name must be unique

    Background:
      Given the credit type with abbreviation "C"

    Scenario: basket name is unique
      When alice attempts to create a basket with name "NCT"
      Then expect no error

    Scenario: basket name is not unique
      Given the basket with name "NCT"
      When alice attempts to create a basket with name "NCT"
      Then expect the error "basket with name NCT already exists: unique key violation"

  Rule: The basket fee is not required if the minimum basket fee is not set

    Background:
      Given the credit type with abbreviation "C"
      And alice has a token balance "20regen"

    Scenario: basket fee provided
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee not provided
      When alice attempts to create a basket with no fee
      Then expect no error

    # No failing scenario - basket fee is not required if minimum basket fee is not set

  Rule: The basket fee must match the minimum basket fee denom

    Background:
      Given the basket fee param "20regen"
      And the credit type with abbreviation "C"
      And alice has a token balance "20regen"

    Scenario: basket fee matches the denom
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee does not match the denom
      When alice attempts to create a basket with fee "20atom"
      Then expect the error "minimum fee 20regen, got 20atom: insufficient fee"

    Scenario: basket fee not provided
      When alice attempts to create a basket with no fee
      Then expect the error "minimum fee 20regen, got : insufficient fee"

  Rule: The basket fee must be greater than or equal to the minimum basket fee

    Background:
      Given the basket fee param "20regen"
      And the credit type with abbreviation "C"
      And alice has a token balance "20regen"

    Scenario Outline: basket fee is greater than or equal to minimum basket fee
      When alice attempts to create a basket with fee "<basket-fee>"
      Then expect no error

      Examples:
        | description  | basket-fee |
        | greater than | 30regen    |
        | equal to     | 20regen    |

    Scenario: basket fee is less than minimum basket fee
      When alice attempts to create a basket with fee "10regen"
      Then expect the error "minimum fee 20regen, got 10regen: insufficient fee"

  Rule: The user must have a balance greater than or equal to the basket fee amount

    Background:
      Given the basket fee param "20regen"
      And the credit type with abbreviation "C"

    Scenario Outline: user has greater than or equal to basket fee amount
      Given alice has a token balance "<token-balance>"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

      Examples:
        | description  | token-balance |
        | greater than | 30regen       |
        | equal to     | 20regen       |

    Scenario: user has less than basket fee amount
      Given alice has a token balance "10regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect the error "insufficient balance for bank denom regen: insufficient funds"

  Rule: The basket must include a valid credit type

    Background:
      Given the credit type with abbreviation "C"

    Scenario: basket criteria includes a valid credit class
      When alice attempts to create a basket with credit type "C"
      Then expect no error

    Scenario: basket criteria includes an invalid credit class
      When alice attempts to create a basket with credit type "F"
      Then expect the error "could not get credit type with abbreviation F: not found: invalid request"

  Rule: The basket criteria must include a valid credit class

    Background:
      Given the credit type with abbreviation "C"

    Scenario: basket criteria includes a valid credit class
      Given the credit class with id "C01"
      When alice attempts to create a basket with allowed class "C01"
      Then expect no error

    Scenario: basket criteria includes an invalid credit class
      When alice attempts to create a basket with allowed class "C01"
      Then expect the error "could not get credit class C01: not found: invalid request"

    Scenario: basket criteria includes an invalid credit class
      Given the credit class with id "BIO01"
      When alice attempts to create a basket with allowed class "BIO01"
      Then expect the error "basket specified credit type C, but class BIO01 is of type BIO: invalid request"

  Rule: The basket exponent must be greater than or equal to the credit type precision

    Background:
      Given the credit type with abbreviation "C" and precision "6"

    Scenario Outline: basket exponent is greater than or equal to credit type precision
      When alice attempts to create a basket with name "NCT" and exponent "<exponent>"
      Then expect no error

      Examples:
        | description  | exponent |
        | greater than | 9        |
        | equal to     | 6        |

    Scenario: basket exponent is less than credit type precision
      When alice attempts to create a basket with name "NCT" and exponent "3"
      Then expect the error "exponent 3 must be >= credit type precision 6: invalid request"

  Rule: The user token balance is updated when the basket is created

    Background:
      Given the credit type with abbreviation "C"
      And alice has a token balance "20regen"

    Scenario: user token balance is updated
      When alice attempts to create a basket with fee "20regen"
      Then alice has a token balance "0regen"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The message response includes basket denom when credits are put into the basket

    Background:
      Given the credit type with abbreviation "C"

    Scenario: message response includes the basket denom
      When alice attempts to create a basket with name "NCT" and exponent "6"
      Then expect the response
      """
      {
        "basket_denom": "eco.uC.NCT"
      }
      """

    # no failing scenario - response should always be empty when message execution fails
