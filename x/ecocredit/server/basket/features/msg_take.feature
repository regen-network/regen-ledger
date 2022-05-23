Feature: Msg/Take

  Credits can be taken from a basket:
  - when the basket exists
  - when the user token balance is greater than or equal to the token amount
  - when auto-retire is disabled and the user sets retire on take to true
  - when auto-retire is disabled and the user sets retire on take to false
  - when auto-retire is enabled and the user sets retire on take to true
  - the user token balance is updated
  - the basket token supply is updated
  - the user retired credit balance is updated when the user sets retire on take to false
  - the user tradable credit balance is updated when the user sets retire on take to true
  - the basket credit balance is updated
  - the response includes the credits received

  Rule: The basket must exist

    Scenario: basket exists
      Given a basket with denom "eco.C.NCT"
      And alice owns tokens with denom "eco.C.NCT"
      When alice attempts to take credits with basket denom "eco.C.NCT"
      Then expect no error

    Scenario: basket does not exist
      When alice attempts to take credits with basket denom "eco.C.NCT"
      Then expect the error "basket eco.C.NCT not found: not found"

  Rule: The user token balance must be greater than or equal to the token amount

    Background:
      Given a basket

    Scenario Outline: user token balance is greater than or equal to token amount
      Given alice owns basket token amount "<token-balance>"
      When alice attempts to take credits with basket token amount "<token-amount>"
      Then expect no error

      Examples:
        | description     | token-balance  | token-amount |
        | balance greater | 100            | 50           |
        | balance equal   | 100            | 100          |

    Scenario Outline: user token balance is less than token amount
      Given alice owns basket token amount "<token-balance>"
      When alice attempts to take credits with basket token amount "<token-amount>"
      Then expect the error "insufficient balance for basket denom eco.C.NCT: insufficient funds"

      Examples:
        | description  | token-balance  | token-amount |
        | no balance   | 0              | 100          |
        | balance less | 50             | 100          |

  Rule: The user must set retire on take to true if auto-retire is enabled

    Scenario Outline: basket auto-retire disabled
      Given a basket with disable auto retire "true"
      And alice owns basket tokens
      When alice attempts to take credits with retire on take "<retire-on-take>"
      Then expect no error

      Examples:
        | retire-on-take |
        | true           |
        | false          |

    Scenario: basket auto-retire enabled and user sets retire on take to true
      Given a basket with disable auto retire "false"
      And alice owns basket tokens
      When alice attempts to take credits with retire on take "true"
      Then expect no error

    Scenario: basket auto-retire enabled and user sets retire on take to false
      Given a basket with disable auto retire "false"
      And alice owns basket tokens
      When alice attempts to take credits with retire on take "false"
      Then expect the error "can't disable retirement when taking from this basket"

 Rule: The user token balance is updated when credits are taken from the basket

    Scenario: user token balance is updated
      Given a basket
      And alice owns basket token amount "100"
      When alice attempts to take credits with basket token amount "50"
      Then expect alice basket token balance amount "50"

    # no failing scenario - state transitions only occur upon successful message execution

 Rule: The basket token supply is updated when credits are taken from the basket

    Scenario: basket token supply is updated
      Given a basket
      And basket token supply amount "100"
      And alice owns basket token amount "100"
      When alice attempts to take credits with basket token amount "50"
      Then expect basket token supply amount "50"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The user credit balance is updated when credits are taken from the basket

    Scenario Outline: user retired credit balance is updated
      Given a basket with exponent "<exponent>" and disable auto retire "false"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>" and retire on take "true"
      Then expect alice retired credit balance amount "<retired-credits>"
      And expect alice tradable credit balance amount "0"

      Examples:
        | description                        | exponent | token-amount | retired-credits |
        | exponent zero, credits whole       | 0        | 2            | 2               |
        | exponent non-zero, credits whole   | 6        | 2000000      | 2.000000        |
        | exponent non-zero, credits decimal | 6        | 2500000      | 2.500000        |

    Scenario Outline: user tradable credit balance is updated
      Given a basket with exponent "<exponent>" and disable auto retire "true"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>" and retire on take "false"
      Then expect alice tradable credit balance amount "<tradable-credits>"
      And expect alice retired credit balance amount "0"

      Examples:
        | description                        | exponent | token-amount | tradable-credits |
        | exponent zero, credits whole       | 0        | 2            | 2                |
        | exponent non-zero, credits whole   | 6        | 2000000      | 2.000000         |
        | exponent non-zero, credits decimal | 6        | 2500000      | 2.500000         |

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The basket credit balance is updated when credits are taken from the basket

    Scenario Outline: basket credit balance is updated
      Given a basket with exponent "<exponent>" and credit balance "100"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>"
      Then expect basket credit balance amount "<credit-amount>"

      Examples:
        | description                        | exponent | token-amount | credit-amount |
        | exponent zero, credits whole       | 0        | 2            | 98            |
        | exponent non-zero, credits whole   | 6        | 2000000      | 98.000000     |
        | exponent non-zero, credits decimal | 6        | 2500000      | 97.500000     |

    # no failing scenario - state transitions only occur upon successful message execution

 Rule: The message response includes the credits received when credits are taken from the basket

    Scenario Outline: message response includes basket token amount received
      Given a basket with exponent "<exponent>" and credit balance "100"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>"
      Then expect the response
      """
      {
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "<credit-amount>"
          }
        ]
      }
      """

      Examples:
        | description                        | exponent | token-amount | credit-amount |
        | exponent zero, credits whole       | 0        | 2            | 2             |
        | exponent non-zero, credits whole   | 6        | 2000000      | 2.000000      |
        | exponent non-zero, credits decimal | 6        | 2500000      | 2.500000      |

    # no failing scenario - response should always be empty when message execution fails
