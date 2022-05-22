Feature: MsgSell

  A sell order can be created:
    - when the credit batch exists
    - when the seller owns credits from the credit batch
    - when the seller owns greater than or equal to the quantity of credits
    - when the ask denom is an allowed denom
    - when the expiration is after the block time
    - the market is created when a market with credit type and ask denom does not exist
    - the tradable credits are converted to escrowed credits
    - the response includes the sell order id
    - the sell order is stored in state

  Rule: The credit batch must exist

    Background:
      Given an allowed denom

    Scenario: credit batch exists
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: credit batch does not exist
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "batch denom C01-001-20200101-20210101-001: not found: invalid request"

  Rule: The seller must own credits from the credit batch

    Background:
      Given an allowed denom

    Scenario: sell owns credits
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: seller does not own credits
      Given a credit batch with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "not found"

  Rule: The seller must own greater than or equal to the quantity of credits

    Background:
      Given an allowed denom

    Scenario Outline: seller owns greater than or equal to credit quantity
      Given alice owns credit quantity "<balance>"
      When alice attempts to create a sell order with credit quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | balance | quantity |
        | greater than | 100     | 50       |
        | equal to     | 100     | 100      |

    Scenario: seller owns less than credit quantity
      Given alice owns credit quantity "50"
      When alice attempts to create a sell order with credit quantity "100"
      Then expect the error "credit quantity: 100, tradable balance: 50: insufficient credit balance"

  Rule: The ask denom must be an allowed denom

    Background:
      Given alice owns credits

    Scenario: ask denom is allowed
      Given an allowed denom with denom "regen"
      When alice attempts to create a sell order with ask price "100regen"
      Then expect no error

    Scenario: ask denom is not allowed
      When alice attempts to create a sell order with ask price "100regen"
      Then expect the error "regen is not allowed to be used in sell orders: invalid request"

  Rule: The expiration must be after the block time

    Background:
      Given a block time with timestamp "2020-01-01"
      And an allowed denom
      And alice owns credits

    Scenario: expiration is after block time
      When alice attempts to create a sell order with expiration "2021-01-01"
      Then expect no error

    Scenario Outline: expiration is before or equal to block time
      When alice attempts to create a sell order with expiration "<expiration>"
      Then expect error contains "expiration must be in the future"

    Examples:
      | description | expiration |
      | before      | 2019-01-01 |
      | equal to    | 2020-01-01 |

  Rule: The market is created when a market with credit type and ask denom does not exist

    Background:
      Given an allowed denom with denom "regen"
      And alice owns credits with batch denom "C01-001-20200101-20210101-001"

    Scenario: market created
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001" and ask denom "regen"
      Then expect market with id "1" and denom "regen"

    Scenario: market already exists
      Given a market with credit type "C" and bank denom "regen"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001" and ask denom "regen"
      Then expect market with id "2" does not exist

  Rule: The tradable credits are converted to escrowed credits

    Background:
      Given an allowed denom
      And alice owns credit quantity "100"

    Scenario: the credits are escrowed
      When alice attempts to create a sell order with credit quantity "100"
      Then expect alice tradable credit balance "0"
      And expect alice escrowed credit balance "100"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The sell order is stored in state

    Background:
      Given an allowed denom
      And alice owns credits

    Scenario: a single sell order
      When alice attempts to create a sell order
      Then expect sell order with id "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The response includes the sell order id

    Background:
      Given an allowed denom
      And alice owns credits

    Scenario: a single sell order
      When alice attempts to create a sell order
      Then expect the response
      """
      {
        "sell_order_ids": [1]
      }
      """

    # no failing scenario - response should always be empty when message execution fails
