Feature: Msg/Create

  A basket can be created:
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