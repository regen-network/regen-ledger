Feature: Msg/CancelSellOrder

  A sell order can be cancelled:
    - when the sell order exists
    - when the user is the owner of the sell order
    - the escrowed credits are converted to tradable credits
    - the sell order is removed from state

  Rule: The sell order must exist

    Scenario: sell order exists
      Given alice created a sell order with id "1"
      When alice attempts to cancel the sell order with id "1"
      Then expect no error

    Scenario: sell order does not exist
      When alice attempts to cancel the sell order with id "1"
      Then expect the error "sell order with id 1: not found: invalid request"

  Rule: The user must be the owner of the sell order

    Scenario: user is the owner
      Given alice created a sell order with id "1"
      When alice attempts to cancel the sell order with id "1"
      Then expect no error

    Scenario: user is not the owner
      Given alice created a sell order with id "1"
      When bob attempts to cancel the sell order with id "1"
      Then expect the error "seller must be the owner of the sell order: unauthorized"

  Rule: The escrowed credits are converted to tradable credits

    Scenario: the credits converted to tradable credits
      Given alice created a sell order with id "1" and quantity "100"
      And alice has the batch balance
      """
      {
        "tradable_amount": "0",
        "escrowed_amount": "100",
        "retired_amount": "0"
      }
      """
      When alice attempts to cancel the sell order with id "1"
      Then expect alice batch balance
      """
      {
        "tradable_amount": "100",
        "escrowed_amount": "0",
        "retired_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The sell order is removed from state

    Scenario: the sell order is removed from state
      Given alice created a sell order with id "1"
      When alice attempts to cancel the sell order with id "1"
      Then expect no sell order with id "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Scenario: EventCancelSellOrder is emitted
      Given alice created a sell order with id "1"
      When alice attempts to cancel the sell order with id "1"
      Then expect event with properties
      """
      {
        "sell_order_id": 1
      }
      """