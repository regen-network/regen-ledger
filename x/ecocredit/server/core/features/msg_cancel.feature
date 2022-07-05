Feature: Msg/Cancel

  Credits can be cancelled by the owner:
  - the owner credit balance is updated
  - the batch supply is updated
  - ...

  Rule: The owner balance is updated

    Scenario: the owner balance is updated
      Given a credit batch exists
      And alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to cancel credit amount "10"
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The batch supply is updated

    Scenario: the batch supply is updated
      Given a credit batch exists
      And alice owns tradable credit amount "10"
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to cancel credit amount "10"
      Then expect batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "cancelled_amount": "10"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
