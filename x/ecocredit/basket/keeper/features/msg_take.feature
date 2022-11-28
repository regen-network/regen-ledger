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

    Background:
      Given a credit type with abbreviation "C" and precision "0"

    Scenario: basket exists
      Given a basket with denom "eco.uC.NCT"
      And alice owns tokens with denom "eco.uC.NCT"
      When alice attempts to take credits with basket denom "eco.uC.NCT"
      Then expect no error

    Scenario: basket does not exist
      When alice attempts to take credits with basket denom "eco.uC.NCT"
      Then expect the error "basket eco.uC.NCT not found: not found"

  Rule: The user token balance must be greater than or equal to the token amount

    Background:
      Given a credit type with abbreviation "C" and precision "6"
      And a basket

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
      Then expect the error "insufficient balance for basket denom eco.uC.NCT: insufficient funds"

      Examples:
        | description  | token-balance  | token-amount |
        | no balance   | 0              | 100          |
        | balance less | 50             | 100          |

  Rule: The user must set retire on take to true if auto-retire is enabled

    Background:
      Given a credit type

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
      Given a credit type
      And a basket
      And alice owns basket token amount "100"
      When alice attempts to take credits with basket token amount "50"
      Then expect alice basket token balance amount "50"

    # no failing scenario - state transitions only occur upon successful message execution

 Rule: The basket token supply is updated when credits are taken from the basket

    Scenario: basket token supply is updated
      Given a credit type
      And a basket
      And basket token supply amount "100"
      And alice owns basket token amount "100"
      When alice attempts to take credits with basket token amount "50"
      Then expect basket token supply amount "50"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The user credit balance is updated when credits are taken from the basket

    Scenario Outline: user retired credit balance is updated
      Given a credit type with abbreviation "C" and precision "<precision>"
      And a basket with credit type "C" and disable auto retire "false"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>" and retire on take "true"
      Then expect alice retired credit balance amount "<retired-credits>"
      And expect alice tradable credit balance amount "0"

      Examples:
        | description                         | precision | token-amount | retired-credits |
        | precision zero, credits whole       | 0         | 2            | 2               |
        | precision non-zero, credits whole   | 6         | 2000000      | 2.000000        |
        | precision non-zero, credits decimal | 6         | 2500000      | 2.500000        |

    Scenario Outline: user tradable credit balance is updated
      Given a credit type with abbreviation "C" and precision "<precision>"
      And a basket with credit type "C" and disable auto retire "true"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>" and retire on take "false"
      Then expect alice tradable credit balance amount "<tradable-credits>"
      And expect alice retired credit balance amount "0"

      Examples:
        | description                         | precision | token-amount | tradable-credits |
        | precision zero, credits whole       | 0         | 2            | 2                |
        | precision non-zero, credits whole   | 6         | 2000000      | 2.000000         |
        | precision non-zero, credits decimal | 6         | 2500000      | 2.500000         |

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The basket credit balance is updated when credits are taken from the basket

    Scenario Outline: basket credit balance is updated
      Given a credit type with abbreviation "C" and precision "<precision>"
      And a basket with credit type "C" and credit balance "100"
      And basket token supply amount "<token-amount>"
      And alice owns basket token amount "<token-amount>"
      When alice attempts to take credits with basket token amount "<token-amount>"
      Then expect basket credit balance amount "<credit-amount>"

      Examples:
        | description                         | precision | token-amount | credit-amount |
        | precision zero, credits whole       | 0         | 2            | 98            |
        | precision non-zero, credits whole   | 6         | 2000000      | 98.000000     |
        | precision non-zero, credits decimal | 6         | 2500000      | 97.500000     |

    # no failing scenario - state transitions only occur upon successful message execution

 Rule: The message response includes the credits received when credits are taken from the basket

    Scenario Outline: message response includes basket token amount received
      Given a credit type with abbreviation "C" and precision "<precision>"
      And a basket with credit type "C" and credit balance "100"
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
        | description                         | precision | token-amount | credit-amount |
        | precision zero, credits whole       | 0         | 2            | 2             |
        | precision non-zero, credits whole   | 6         | 2000000      | 2.000000      |
        | precision non-zero, credits decimal | 6         | 2500000      | 2.500000      |

    # no failing scenario - response should always be empty when message execution fails

  Rule: Events are emitted

    Background:
      Given a credit type with abbreviation "C" and precision "6"
      And a basket with credit type "C" and disable auto retire "true"
      And basket token supply amount "2000000"
      And Alice's address "regen10z82e5ztmrm4pujgummvmr7aqjzwlp6gz8k8xp"
      And Ecocredit module's address "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And Alice owns basket token amount "2000000"

    Scenario: EventTake is emitted
      When alice attempts to take credits with basket token amount "2000000"
      Then expect event take with properties
      """
      {
        "owner": "regen10z82e5ztmrm4pujgummvmr7aqjzwlp6gz8k8xp",
        "basket_denom": "eco.uC.NCT",
        "credits": [
          {"batch_denom": "C01-001-20200101-20210101-001", "amount": "2.000000"}
        ],
        "amount": "2000000"
      }
      """

    Scenario: EventRetire is emitted
      When alice attempts to take credits with basket token amount "2000000" and retire on take "true" from "US-WA" and reason "offsetting electricity consumption"
      Then expect event retire with properties
      """
      {
        "owner": "regen10z82e5ztmrm4pujgummvmr7aqjzwlp6gz8k8xp",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "2.000000",
        "jurisdiction": "US-WA",
        "reason": "offsetting electricity consumption"
      }
      """

    Scenario: EventTransfer is emitted when retire on take is true
      When alice attempts to take credits with basket token amount "2000000" and retire on take "true" from "US-WA"
      Then expect event transfer with properties
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen10z82e5ztmrm4pujgummvmr7aqjzwlp6gz8k8xp",
        "batch_denom": "C01-001-20200101-20210101-001",
        "retired_amount": "2.000000"
      }
      """

    Scenario: EventTransfer is emitted when retire on take is false
      When alice attempts to take credits with basket token amount "2000000" and retire on take "false"
      Then expect event transfer with properties
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen10z82e5ztmrm4pujgummvmr7aqjzwlp6gz8k8xp",
        "batch_denom": "C01-001-20200101-20210101-001",
        "tradable_amount": "2.000000"
      }
      """