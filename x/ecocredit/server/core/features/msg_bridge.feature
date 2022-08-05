Feature: Msg/Bridge

  Credits can be bridged to another chain:
  - when a batch contract entry exists
  - the credits are cancelled and the owner balance is updated
  - the credits are cancelled and the total supply is updated

  Rule: The batch contract entry must exist

    Scenario: the batch contract entry exists
      Given a credit batch exists with a batch contract entry
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect no error

    Scenario: the batch contract entry does not exist
      Given a credit batch exists without a batch contract entry
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect the error "only credits previously bridged from another chain are supported: invalid request"

  Rule: The credits are cancelled and the owner balance is updated

    Scenario: the owner balance is updated
      Given a credit batch exists with a batch contract entry
      And alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to bridge credit amount "10" from the credit batch
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The credits are cancelled and the total supply is updated

    Scenario: the total supply is updated
      Given a credit batch exists with a batch contract entry
      And alice owns tradable credit amount "10" from the credit batch
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to bridge credit amount "10" from the credit batch
      Then expect batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "cancelled_amount": "10"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
