Feature: Msg/Create

  A basket can be created:
  - when the basket name is unique
  - when a minimum basket fee is not set
  - when the basket fee denom matches a minimum basket fee denom
  - when the basket fee amount is greater than or equal to a minimum basket fee amount
  - when the user has a balance greater than or equal to a minimum basket fee amount
  - when the basket includes a credit type that exists
  - when the basket criteria includes credit classes that exist
  - when the basket criteria includes credit classes that match the credit type
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

  Rule: The basket fee is not required if a minimum basket fee is not set

    Background:
      Given a credit type

    Scenario: basket fee provided
      Given alice has a token balance "20regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee not provided
      When alice attempts to create a basket with no fee
      Then expect no error

    # No failing scenario - basket fee is not required if minimum basket fee is not set

  Rule: The basket fee must match a minimum basket fee denom

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario: basket fee matches the denom (single fee)
      Given a minimum basket fee "20regen"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee matches the denom (multiple fees)
      Given minimum basket fees "20regen,20atom"
      When alice attempts to create a basket with fee "20regen"
      Then expect no error

    Scenario: basket fee does not match the denom (single fee)
      Given a minimum basket fee "20regen"
      When alice attempts to create a basket with fee "20atom"
      Then expect the error "minimum fee 20regen, got 20atom: insufficient fee"

    Scenario: basket fee does not match the denom (multiple fees)
      Given minimum basket fees "20regen,20atom"
      When alice attempts to create a basket with fee "20stake"
      Then expect the error "minimum fee one of 20atom,20regen, got 20stake: insufficient fee"

    Scenario: basket fee not provided (single fee)
      Given a minimum basket fee "20regen"
      When alice attempts to create a basket with no fee
      Then expect the error "minimum fee 20regen, got : insufficient fee"

    Scenario: basket fee not provided (multiple fees)
      Given minimum basket fees "20regen,20atom"
      When alice attempts to create a basket with no fee
      Then expect the error "minimum fee one of 20atom,20regen, got : insufficient fee"

  Rule: The basket fee must be greater than or equal to a minimum basket fee

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario Outline: basket fee is greater than or equal to minimum basket fee (single fee)
      Given a minimum basket fee "20regen"
      When alice attempts to create a basket with fee "<basket-fee>"
      Then expect no error

      Examples:
        | description  | basket-fee |
        | greater than | 30regen    |
        | equal to     | 20regen    |

    Scenario Outline: basket fee is greater than or equal to minimum basket fee (multiple fees)
      Given minimum basket fees "20regen,20atom"
      When alice attempts to create a basket with fee "<basket-fee>"
      Then expect no error

      Examples:
        | description  | basket-fee |
        | greater than | 30regen    |
        | equal to     | 20regen    |

    Scenario: basket fee is less than minimum basket fee (single fee)
      Given a minimum basket fee "20regen"
      When alice attempts to create a basket with fee "10regen"
      Then expect the error "minimum fee 20regen, got 10regen: insufficient fee"

    Scenario: basket fee is less than minimum basket fee (multiple fees)
      Given minimum basket fees "20regen,20atom"
      When alice attempts to create a basket with fee "10regen"
      Then expect the error "minimum fee one of 20atom,20regen, got 10regen: insufficient fee"

  Rule: The user must have a balance greater than or equal to the basket fee amount

    Background:
      Given a credit type
      And a minimum basket fee "20regen"

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

  Rule: The user token balance is updated and only the minimum fee is taken

    Background:
      Given a credit type
      And a minimum basket fee "20regen"
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
