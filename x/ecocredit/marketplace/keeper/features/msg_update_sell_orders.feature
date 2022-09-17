Feature: Msg/UpdateSellOrders

  A sell order can be updated:
    - when the sell order exists
    - when the user is the seller of the sell order
    - when the seller owns greater than or equal to the quantity of credits
    - when the number of decimal places in quantity is less than or equal to the credit type precision
    - when the ask denom is an allowed denom
    - when the expiration is after the block time
    - the market is created when the credit type and bank denom pair is unique
    - the escrowed credits are converted to tradable credits when quantity is reduced
    - the tradable credits are converted to escrowed credits when quantity is increased
    - the updated sell orders are stored in state

  Rule: The sell order must exist

    Background:
      Given a credit type
      And an allowed denom

    Scenario: the sell order exists
      Given alice created a sell order with id "1"
      When alice attempts to update the sell order with id "1"
      Then expect no error

    Scenario: the sell order does not exist
      When alice attempts to update the sell order with id "1"
      Then expect the error "updates[0]: sell order with id 1: not found: invalid request"

  Rule: The user must be the seller of the sell order

    Background:
      Given a credit type
      And an allowed denom
      And alice created a sell order

    Scenario: the user is the seller
      When alice attempts to update the sell order
      Then expect no error

    Scenario: the user is not the seller
      When bob attempts to update the sell order
      Then expect the error "updates[0]: seller must be the seller of the sell order: unauthorized"

  Rule: The seller must own greater than or equal to the quantity of credits

    Background:
      Given a credit type
      And an allowed denom

    Scenario Outline: seller owns greater than or equal to credit quantity (single sell order)
      Given alice created a sell order with quantity "50"
      And alice has a batch balance with tradable amount "50" and escrowed amount "50"
      When alice attempts to update the sell order with quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | quantity |
        | greater than | 50       |
        | equal to     | 100      |

    Scenario Outline: seller owns greater than or equal to credit quantity (multiple sell orders)
      Given alice created two sell orders each with quantity "50"
      And alice has a batch balance with tradable amount "100" and escrowed amount "100"
      When alice attempts to update the two sell orders each with quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | quantity |
        | greater than | 50       |
        | equal to     | 100      |

    Scenario: seller owns less than credit quantity (single sell order)
      Given alice created a sell order with quantity "50"
      And alice has a batch balance with tradable amount "0" and escrowed amount "50"
      When alice attempts to update the sell order with quantity "100"
      Then expect the error "updates[0]: credit quantity: 50, tradable balance: 0: insufficient credit balance"

    Scenario: seller owns less than credit quantity (multiple sell orders)
      Given alice created two sell orders each with quantity "50"
      And alice has a batch balance with tradable amount "50" and escrowed amount "100"
      When alice attempts to update the two sell orders each with quantity "100"
      Then expect the error "updates[1]: credit quantity: 50, tradable balance: 0: insufficient credit balance"

  Rule: The number of decimal places in quantity must be less than or equal to the credit type precision

    Background:
      Given a credit type with precision "6"
      And an allowed denom
      And alice created a sell order with quantity "100"

    Scenario Outline: quantity decimal places less than or equal to precision
      When alice attempts to update the sell order with quantity "<quantity>"
      Then expect no error

      Examples:
        | description | quantity   |
        | less than   | 99.12345  |
        | equal to    | 99.123456 |

    Scenario: quantity decimal places more than precision
      When alice attempts to update the sell order with quantity "99.1234567"
      Then expect the error "99.1234567 exceeds maximum decimal places: 6"

  Rule: The ask denom must be an allowed denom

    Background:
      Given a credit type

    Scenario: ask denom is allowed
      Given an allowed denom with bank denom "regen"
      And alice created a sell order with ask denom "regen"
      When alice attempts to update the sell order with ask denom "regen"
      Then expect no error

    Scenario: ask denom is not allowed
      Given an allowed denom with bank denom "atom"
      And alice created a sell order with ask denom "atom"
      When alice attempts to update the sell order with ask denom "regen"
      Then expect the error "updates[0]: regen is not allowed to be used in sell orders: invalid request"

  Rule: The expiration must be after the block time

    Background:
      Given a block time with timestamp "2020-01-01"
      And a credit type
      And an allowed denom
      And alice created a sell order

    Scenario: expiration is after block time
      When alice attempts to update the sell order with expiration "2021-01-01"
      Then expect no error

    Scenario Outline: expiration is before or equal to block time
      When alice attempts to update the sell order with expiration "<expiration>"
      Then expect the error "updates[0]: expiration must be in the future: <expiration> 00:00:00 +0000 UTC: invalid request"

    Examples:
      | description | expiration |
      | before      | 2019-01-01 |
      | equal to    | 2020-01-01 |

  Rule: The market is created when the credit type and bank denom pair is unique

    Background:
      Given a credit type with abbreviation "C"
      And an allowed denom with bank denom "regen"
      And an allowed denom with bank denom "atom"
      And alice created a sell order with batch denom "C01-001-20200101-20210101-001" and ask denom "regen"

    Scenario: credit type and bank denom pair is unique
      When alice attempts to update the sell order with ask denom "atom"
      Then expect market with id "2" and denom "atom"

    Scenario: credit type and bank denom pair is not unique
      When alice attempts to update the sell order with ask denom "regen"
      Then expect no market with id "2"

  Rule: The escrowed credits are converted to tradable credits when the quantity is reduced

    Background:
      Given a credit type
      And an allowed denom
      And alice created a sell order with quantity "100"
      And alice has a batch balance with tradable amount "0" and escrowed amount "100"

    Scenario: the credits are converted to tradable credits
      When alice attempts to update the sell order with quantity "50"
      Then expect alice tradable credit balance "50"
      And expect alice escrowed credit balance "50"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The tradable credits are converted to escrowed credits when the quantity is increased

    Background:
      Given a credit type
      And an allowed denom
      And alice created a sell order with quantity "100"
      And alice has a batch balance with tradable amount "100" and escrowed amount "100"

    Scenario: the credits are converted to escrowed credits
      When alice attempts to update the sell order with quantity "200"
      Then expect alice tradable credit balance "0"
      And expect alice escrowed credit balance "200"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The updated sell orders are stored in state

    Background:
      Given a credit type with abbreviation "C"
      And an allowed denom with bank denom "regen"
      And an allowed denom with bank denom "atom"

    Scenario: the sell order is stored in state (single sell order)
      Given alice created a sell order with the properties
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "10",
        "ask_price": {
          "denom": "regen",
          "amount": "10"
        },
        "disable_auto_retire": true,
        "expiration": "2030-01-01T00:00:00Z"
      }
      """
      And alice has a batch balance with tradable amount "10" and escrowed amount "10"
      When alice attempts to update the sell order with the properties
      """
      {
        "new_quantity": "20",
        "new_ask_price": {
          "denom": "atom",
          "amount": "20"
        },
        "disable_auto_retire": false,
        "new_expiration": "2040-01-01T00:00:00Z"
      }
      """
      Then expect sell order with seller alice and the properties
      """
      {
        "id": 1,
        "ask_amount": "20",
        "expiration": "2040-01-01T00:00:00Z",
        "batch_key": 1,
        "quantity": "20",
        "disable_auto_retire": false,
        "market_id": 2,
        "maker": true
      }
      """

    Scenario: the sell order is stored in state (multiple sell orders)
      Given alice created two sell orders each with the properties
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "10",
        "ask_price": {
          "denom": "regen",
          "amount": "10"
        },
        "disable_auto_retire": true,
        "expiration": "2030-01-01T00:00:00Z"
      }
      """
      And alice has a batch balance with tradable amount "20" and escrowed amount "20"
      When alice attempts to update the two sell orders each with the properties
      """
      {
        "new_quantity": "20",
        "new_ask_price": {
          "denom": "atom",
          "amount": "20"
        },
        "disable_auto_retire": false,
        "new_expiration": "2040-01-01T00:00:00Z"
      }
      """
      Then expect sell order with seller alice and the properties
      """
      {
        "id": 2,
        "ask_amount": "20",
        "expiration": "2040-01-01T00:00:00Z",
        "batch_key": 1,
        "quantity": "20",
        "disable_auto_retire": false,
        "market_id": 2,
        "maker": true
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type
      And an allowed denom

    Scenario: EventUpdateSellOrder is emitted
      Given alice created a sell order with id "1"
      When alice attempts to update the sell order with id "1"
      Then expect event with properties
      """
      {
        "sell_order_id": 1
      }
      """