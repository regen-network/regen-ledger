Feature: Msg/Send

  Credits can be sent to another account:
  - when the credit batch exists
  - when the sender has a tradable credit balance greater than or equal to the amount to send
  - when the decimal places in amount to send does not exceed credit type precision
  - the sender credit balance is updated
  - the recipient credit balance is updated
  - the batch supply is updated

  Rule: The credit batch must exist

    Scenario: the credit batch exists
      Given a credit batch with denom "C01-001-20200101-20210101-001"
      And alice owns tradable credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to send credits to bob with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch does not exist
      When alice attempts to send credits to bob with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "could not get batch with denom C01-001-20200101-20210101-001: not found: invalid request"

  Rule: The sender must have a tradable credit balance greater that or equal to the amount to send

    Background:
      Given a credit batch
      And alice owns tradable credit amount "10"

    Scenario Outline: sender has tradable balance greater than or equal to tradable amount to send
      When alice attempts to send credits to bob with tradable amount "<amount>"
      Then expect no error

      Examples:
        | description  | amount |
        | greater than | 5      |
        | equal to     | 10     |

    Scenario: sender has tradable balance less than tradable amount to send
      When alice attempts to send credits to bob with tradable amount "15"
      Then expect the error "tradable balance: 10, send tradable amount 15: insufficient credit balance"

    Scenario Outline: sender has tradable balance greater than or equal to retired amount to send
      When alice attempts to send credits to bob with retired amount "<amount>"
      Then expect no error

      Examples:
        | description  | amount |
        | greater than | 5      |
        | equal to     | 10     |

    Scenario: sender has retired balance less than retired amount to send
      When alice attempts to send credits to bob with retired amount "15"
      Then expect the error "tradable balance: 10, send retired amount 15: insufficient credit balance"

  Rule: The decimal places in amount to send must not exceed credit type precision

    Background:
      Given a credit type with abbreviation "C" and precision "6"
      And a credit batch from credit class with credit type "C"
      And alice owns tradable credit amount "10"

    Scenario Outline: the decimal places in tradable amount is less than or equal to credit type precision
      When alice attempts to send credits to bob with tradable amount "<amount>"
      Then expect no error

      Examples:
        | description | amount   |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: the decimal places in tradable amount is greater than credit type precision
      When alice attempts to send credits to bob with tradable amount "9.1234567"
      Then expect the error "9.1234567 exceeds maximum decimal places: 6: invalid request"

    Scenario Outline: the decimal places in retired amount is less than or equal to credit type precision
      When alice attempts to send credits to bob with retired amount "<amount>"
      Then expect no error

      Examples:
        | description | amount   |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: the decimal places in retired amount is greater than credit type precision
      When alice attempts to send credits to bob with retired amount "9.1234567"
      Then expect the error "9.1234567 exceeds maximum decimal places: 6: invalid request"

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

  Rule: Events are emitted

    Background:
      Given a credit batch with denom "C01-001-20200101-20210101-001"
      And alice's address "regen1d466m547y09dgs6xvca7uxs5k2m2pmgspa9kal"
      And bobs address "regen10yhlcvh88sux4zmf67udhg5f5z2803z6jm0d25"
      And alice owns tradable credit amount "10"

    Scenario: Event EventRetire is emitted
      When alice attempts to send credits to bob with retired amount "10" from "US-WA"
      Then expect event retire with properties
      """
      {
        "owner": "regen10yhlcvh88sux4zmf67udhg5f5z2803z6jm0d25",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "10",
        "jurisdiction": "US-WA"
      }
      """

    Scenario: Event EventTransfer is emitted
      When alice attempts to send credits to bob with retired amount "10" from "US-WA"
      Then expect event transfer with properties
      """
      {
        "sender": "regen1d466m547y09dgs6xvca7uxs5k2m2pmgspa9kal",
        "recipient": "regen10yhlcvh88sux4zmf67udhg5f5z2803z6jm0d25",
        "batch_denom": "C01-001-20200101-20210101-001",
        "tradable_amount": "0",
        "retired_amount": "10"
      }
      """

    Scenario: Event EventTransfer is emitted
      When alice attempts to send credits to bob with tradable amount "10"
      Then expect event transfer with properties
      """
      {
        "sender": "regen1d466m547y09dgs6xvca7uxs5k2m2pmgspa9kal",
        "recipient": "regen10yhlcvh88sux4zmf67udhg5f5z2803z6jm0d25",
        "batch_denom": "C01-001-20200101-20210101-001",
        "tradable_amount": "10",
        "retired_amount": "0"
      }
      """