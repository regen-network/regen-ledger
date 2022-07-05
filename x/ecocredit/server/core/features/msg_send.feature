Feature: Msg/Send

  Credits can be sent to another account:
  - the sender credit balance is updated
  - the recipient credit balance is updated
  - the batch supply is updated
  - ...

  Rule: The sender balance is updated

    Background:
      Given a credit batch

    Scenario: the sender tradable balance is updated
      Given alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to send credits to bob with tradable amount "10"
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    Scenario: the sender retired balance is updated
      Given alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to send credits to bob with retired amount "10"
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The recipient balance is updated

    Background:
      Given a credit batch
      And alice owns tradable credit amount "10"

    Scenario: the recipient tradable balance is updated
      Given bob has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to send credits to bob with tradable amount "10"
      Then expect bob batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """

    Scenario: the recipient retired balance is updated
      Given bob has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to send credits to bob with retired amount "10"
      Then expect bob batch balance
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The batch supply is updated

    Scenario: the batch supply is updated
      Given a credit batch
      And alice owns tradable credit amount "10"
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to send credits to bob with retired amount "10"
      Then expect batch supply
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0",
        "cancelled_amount": "0"
      }
      """

    Scenario: the batch supply is not updated
      Given a credit batch
      And alice owns tradable credit amount "10"
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to send credits to bob with retired amount "10"
      Then expect batch supply
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0",
        "cancelled_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
