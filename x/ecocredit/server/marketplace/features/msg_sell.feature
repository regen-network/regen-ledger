Feature: MsgSell

  A sell order can be created:
    - when the credit batch exists
    - when the seller owns credits from the credit batch
    - when the seller owns greater than or equal to the quantity of credits
    - when the number of decimal places in quantity is less than or equal to the credit type precision
    - when the ask denom is an allowed denom
    - when the expiration is after the block time
    - the market is created when the credit type and bank denom pair is unique
    - the tradable credits are converted to escrowed credits
    - the sell orders are stored in state
    - the response includes the sell order ids

  Rule: The credit batch must exist

    Background:
      Given a credit type
      And an allowed denom

    Scenario: credit batch exists
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: credit batch does not exist
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "order[0]: batch denom C01-001-20200101-20210101-001: not found: invalid request"

  Rule: The seller must own credits from the credit batch

    Background:
      Given a credit type
      And an allowed denom

    Scenario: seller owns credits
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: seller does not own credits
      Given a credit batch with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "order[0]: credit quantity: 100, tradable balance: 0: insufficient credit balance"

  Rule: The seller must own greater than or equal to the quantity of credits

    Background:
      Given a credit type
      And an allowed denom

    Scenario Outline: seller owns greater than or equal to credit quantity (single sell order)
      Given alice owns credit quantity "100"
      When alice attempts to create a sell order with credit quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | quantity |
        | greater than | 50       |
        | equal to     | 100      |

    Scenario Outline: seller owns greater than or equal to credit quantity (multiple sell orders)
      Given alice owns credit quantity "200"
      When alice attempts to create two sell orders each with credit quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | quantity |
        | greater than | 50       |
        | equal to     | 100      |

    Scenario: seller owns less than credit quantity (single sell order)
      Given alice owns credit quantity "50"
      When alice attempts to create a sell order with credit quantity "100"
      Then expect the error "order[0]: credit quantity: 100, tradable balance: 50: insufficient credit balance"

    Scenario: seller owns less than credit quantity (multiple sell orders)
      Given alice owns credit quantity "150"
      When alice attempts to create two sell orders each with credit quantity "100"
      Then expect the error "order[1]: credit quantity: 100, tradable balance: 50: insufficient credit balance"

  Rule: The number of decimal places in quantity must be less than or equal to the credit type precision

    Background:
      Given a credit type with precision "6"
      And an allowed denom
      And alice owns credits

    Scenario Outline: quantity decimal places less than or equal to precision
      When alice attempts to create a sell order with credit quantity "<quantity>"
      Then expect no error

      Examples:
        | description | quantity   |
        | less than   | 100.12345  |
        | equal to    | 100.123456 |

    Scenario: quantity decimal places more than precision
      When alice attempts to create a sell order with credit quantity "100.1234567"
      Then expect the error "100.1234567 exceeds maximum decimal places: 6"

  Rule: The ask denom must be an allowed denom

    Background:
      Given a credit type
      And alice owns credits

    Scenario: ask denom is allowed
      Given an allowed denom with bank denom "regen"
      When alice attempts to create a sell order with ask price "100regen"
      Then expect no error

    Scenario: ask denom is not allowed
      When alice attempts to create a sell order with ask price "100regen"
      Then expect the error "order[0]: regen is not allowed to be used in sell orders: invalid request"

  Rule: The expiration must be after the block time

    Background:
      Given a block time with timestamp "2020-01-01"
      And a credit type
      And an allowed denom
      And alice owns credits

    Scenario: expiration is after block time
      When alice attempts to create a sell order with expiration "2021-01-01"
      Then expect no error

    Scenario Outline: expiration is before or equal to block time
      When alice attempts to create a sell order with expiration "<expiration>"
      Then expect the error "order[0]: expiration must be in the future: <expiration> 00:00:00 +0000 UTC: invalid request"

    Examples:
      | description | expiration |
      | before      | 2019-01-01 |
      | equal to    | 2020-01-01 |

  Rule: The market is created when the credit type and bank denom pair is unique

    Background:
      Given a credit type
      And an allowed denom with bank denom "regen"
      And alice owns credits with batch denom "C01-001-20200101-20210101-001"

    Scenario: credit type and bank denom pair is unique
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001" and ask denom "regen"
      Then expect market with id "1" and denom "regen"

    Scenario: credit type and bank denom pair is not unique
      Given a market with credit type "C" and bank denom "regen"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001" and ask denom "regen"
      Then expect no market with id "2"

  Rule: The tradable credits are converted to escrowed credits

    Background:
      Given a credit type
      And an allowed denom
      And alice owns credit quantity "100"

    Scenario: the credits are escrowed
      When alice attempts to create a sell order with credit quantity "100"
      Then expect alice tradable credit balance "0"
      And expect alice escrowed credit balance "100"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The sell orders are stored in state

    Background:
      Given a credit type
      And an allowed denom
      And alice owns credits

    Scenario: the sell orders are stored in state
      When alice attempts to create two sell orders
      Then expect sell order with id "1"
      And expect sell order with id "2"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The response includes the sell order ids

    Background:
      Given a credit type
      And an allowed denom
      And alice owns credits

    Scenario: the response includes sell order ids
      When alice attempts to create two sell orders
      Then expect the response
      """
      {
        "sell_order_ids": [1,2]
      }
      """

    # no failing scenario - response should always be empty when message execution fails
