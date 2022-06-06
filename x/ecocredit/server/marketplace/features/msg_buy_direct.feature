Feature: Msg/BuyDirect

  Credits can be bought directly:
  - when the sell order exists
  - when the bid denom matches the sell denom
  - when the buyer has a bank balance greater than or equal to the total cost
  - when the buyer provides a bid price greater than or equal to the ask price
  - when the buyer provides a quantity less than or equal to the sell order quantity
  - when the number of decimal places in quantity is less than or equal to the credit type precision
  - the buyer cannot disable auto-retire when auto-retire is enabled for the sell order
  - the sell order is removed when the sell order is filled
  - the sell order quantity is updated when the sell order is not filled
  - the seller bank balance is updated
  - the buyer bank balance is updated
  - the seller batch balance is updated
  - the buyer batch balance is updated
  - the batch supply is updated when the credits are auto-retired

  Rule: The sell order must exist

    Background:
      Given a credit type

    Scenario: the sell order exists
      Given alice created a sell order with id "1"
      When bob attempts to buy credits with sell order id "1"
      Then expect no error

    Scenario: the sell order does not exist
      When bob attempts to buy credits with sell order id "1"
      Then expect the error "orders[0]: sell order with id 1: not found"

  Rule: The bid denom must match the sell denom

    Background:
      Given a credit type
      And alice created a sell order with ask denom "regen"
      And bob has a bank balance with denom "regen"

    Scenario: bid denom matches sell denom
      When bob attempts to buy credits with bid denom "regen"
      Then expect no error

    Scenario: bid denom does not match sell denom
      When bob attempts to buy credits with bid denom "atom"
      Then expect the error "orders[0]: bid price denom: atom, ask price denom: regen: invalid request"

  Rule: The buyer must have a bank balance greater than or equal to the total cost

    Background:
      Given a credit type

    Scenario Outline: buyer bank balance is greater than or equal to total cost (single buy order)
      Given alice created a sell order with quantity "10" and ask amount "10"
      And bob has a bank balance with amount "<balance-amount>"
      When bob attempts to buy credits with quantity "10" and bid amount "10"
      Then expect no error

      Examples:
        | description  | balance-amount |
        | greater than | 200            |
        | equal to     | 100            |

    Scenario Outline: buyer bank balance is greater than or equal to total cost (multiple buy orders)
      Given alice created two sell orders each with quantity "10" and ask amount "10"
      And bob has a bank balance with amount "<balance-amount>"
      When bob attempts to buy credits in two orders each with quantity "10" and bid amount "10"
      Then expect no error

      Examples:
        | description  | balance-amount |
        | greater than | 400            |
        | equal to     | 200            |

    Scenario: buyer bank balance is less than total cost (single buy order)
      Given alice created a sell order with quantity "10" and ask amount "10"
      And bob has a bank balance with amount "50"
      When bob attempts to buy credits with quantity "10" and bid amount "10"
      Then expect the error "orders[0]: quantity: 10, ask price: 10regen, total price: 100regen, bank balance: 50regen: insufficient funds"

    Scenario: buyer bank balance is less than total cost (multiple buy orders)
      Given alice created two sell orders each with quantity "10" and ask amount "10"
      And bob has a bank balance with amount "150"
      When bob attempts to buy credits in two orders each with quantity "10" and bid amount "10"
      Then expect the error "orders[1]: quantity: 10, ask price: 10regen, total price: 100regen, bank balance: 50regen: insufficient funds"

  Rule: The buyer must provide a bid price greater than or equal to the ask price

    Background:
      Given a credit type
      And alice created a sell order with ask amount "10"
      And bob has a bank balance with amount "100"

    Scenario Outline: bid price greater than or equal to ask price
      When bob attempts to buy credits with quantity "10" and bid amount "<bid-amount>"
      Then expect no error

      Examples:
        | description  | bid-amount |
        | greater than | 20         |
        | equal to     | 10         |

    Scenario: bid price less than ask price
      When bob attempts to buy credits with quantity "10" and bid amount "5"
      Then expect the error "orders[0]: ask price: 10regen, bid price: 5regen, insufficient bid price: invalid request"

  Rule: The buyer must provide a quantity less than or equal to the sell order quantity

    Background:
      Given a credit type
      And alice created a sell order with quantity "10"
      And bob has a bank balance with amount "150"

    Scenario Outline: quantity less than or equal to sell order quantity
      When bob attempts to buy credits with quantity "<quantity>"
      Then expect no error

      Examples:
        | description | quantity |
        | less than   | 5        |
        | equal to    | 10       |

    Scenario: quantity more than sell order quantity
      When bob attempts to buy credits with quantity "15"
      Then expect the error "orders[0]: requested quantity: 15, sell order quantity 10: invalid request"

  Rule: The number of decimal places in quantity must be less than or equal to the credit type precision

    Background:
      Given a credit type with precision "6"

    Scenario Outline: quantity decimal places less than or equal to precision
      Given alice created a sell order with quantity "<quantity>"
      When bob attempts to buy credits with quantity "<quantity>"
      Then expect no error

      Examples:
        | description | quantity |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: quantity decimal places more than precision
      Given alice created a sell order with quantity "9.1234567"
      When bob attempts to buy credits with quantity "9.1234567"
      Then expect the error "orders[0]: decimal places exceeds precision: quantity: 9.1234567, credit type precision: 6: invalid request"

  Rule: The buyer cannot disable auto-retire if the sell order has auto-retire enabled

    Background:
      Given a credit type

    Scenario Outline: auto retire not required
      Given alice created a sell order with disable auto retire "true"
      When bob attempts to buy credits with disable auto retire "<disable-auto-retire>"
      Then expect no error

      Examples:
        | disable-auto-retire |
        | true                |
        | false               |

    Scenario: auto retire required and buyer enables
      Given alice created a sell order with disable auto retire "false"
      When bob attempts to buy credits with disable auto retire "false"
      Then expect no error

    Scenario: auto retire required and buyer disables
      Given alice created a sell order with disable auto retire "false"
      When bob attempts to buy credits with disable auto retire "true"
      Then expect the error "orders[0]: cannot disable auto-retire for a sell order with auto-retire enabled: invalid request"

  Rule: The sell order is removed when the sell order is filled

    Background:
      Given a credit type
      And alice created a sell order with quantity "10"

    Scenario: the sell order is removed
      When bob attempts to buy credits with quantity "10"
      Then expect no sell order with id "1"

    Scenario: the sell order is not removed
      When bob attempts to buy credits with quantity "5"
      Then expect sell order with id "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The sell order quantity is updated when the sell order is not filled

    Background:
      Given a credit type

    Scenario:
      Given alice created a sell order with quantity "20"
      When bob attempts to buy credits with quantity "10"
      Then expect sell order with quantity "10"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the buyer bank balance is updated

    Background:
      Given a credit type

    Scenario: buyer bank balance updated
      Given alice created a sell order with quantity "10" and ask price "10regen"
      And bob has the bank balance "100regen"
      When bob attempts to buy credits with quantity "10" and bid price "10regen"
      Then expect bob bank balance "0regen"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the seller bank balance is updated

    Background:
      Given a credit type
      And alice created a sell order with quantity "100" and ask price "1regen"
      And alice has bank balance "0regen"

    Scenario: seller bank balance updated
      When bob attempts to buy credits with quantity "100" and bid price "1regen"
      Then expect alice bank balance "100regen"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the buyer batch balance is updated

    Background:
      Given a credit type
      And alice created a sell order with quantity "10" and disable auto retire "true"
      And bob has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    Scenario: buyer batch balance updated with retired credits
      When bob attempts to buy credits with quantity "10" and disable auto retire "false"
      Then expect bob batch balance
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    Scenario: buyer batch balance updated with tradable credits
      When bob attempts to buy credits with quantity "10" and disable auto retire "true"
      Then expect bob batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the seller batch balance is updated

    Background:
      Given a credit type
      And alice created a sell order with quantity "10"
      And alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "10"
      }
      """

    Scenario: seller batch balance updated
      When bob attempts to buy credits with quantity "10"
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the batch supply is updated when the credits are auto-retired

    Background:
      Given a credit type
      And alice created a sell order with quantity "10" and disable auto retire "true"
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10"
      }
      """

    Scenario: batch supply updated
      When bob attempts to buy credits with quantity "10" and disable auto retire "false"
      Then expect batch supply
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0"
      }
      """

    Scenario: batch supply not updated
      When bob attempts to buy credits with quantity "10" and disable auto retire "true"
      Then expect batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
