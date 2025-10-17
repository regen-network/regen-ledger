Feature: Msg/Send

  Credits can be sent to another account:
  - message validation
  - when the credit batch exists
  - when the sender has a tradable credit balance greater than or equal to the amount to send
  - when the decimal places in amount to send does not exceed credit type precision
  - the sender credit balance is updated
  - the recipient credit balance is updated
  - the batch supply is updated

  Rule: Message validation

    Scenario: a valid message
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "tradable_amount": "100",
            "retired_amount": "100",
            "retirement_jurisdiction": "US-WA",
            "retirement_reason": "offsetting electricity consumption"
          }
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message without retirement reason
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "tradable_amount": "100",
            "retired_amount": "100",
            "retirement_jurisdiction": "US-WA"
          }
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message with multiple credits
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "tradable_amount": "100"
          },
          {
            "batch_denom": "C01-001-20200101-20210101-002",
            "retired_amount": "100",
            "retirement_jurisdiction": "US-WA",
            "retirement_reason": "offsetting electricity consumption"
          }
        ]
      }
      """
      When the message is validated
      Then expect no error

  
    Scenario: an error is returned if credits is empty
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      }
      """
      When the message is validated
      Then expect the error "credits cannot be empty: invalid request"

    Scenario: an error is returned if credits batch denom is empty
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {}
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: batch denom: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if credits batch denom is not formatted
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "foo"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"

    Scenario: an error is returned if credits tradable amount and retired amount are empty
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: tradable amount or retired amount required: invalid request"

    Scenario: an error is returned if credits tradable amount is a negative decimal
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "tradable_amount": "-100"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: expected a non-negative decimal, got -100: invalid decimal string"

    Scenario: an error is returned if credits retired amount is a negative decimal
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "retired_amount": "-100"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: expected a non-negative decimal, got -100: invalid decimal string"

    Scenario: an error is returned if credits retired amount is positive and retirement jurisdiction is empty
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "retired_amount": "100"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: retirement jurisdiction: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if credits retired amount is positive and retirement jurisdiction is not formatted
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "retired_amount": "100",
            "retirement_jurisdiction": "foo"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: retirement jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

    Scenario: an error is returned if credits retired amount is positive and retirement reason exceeds 512 characters
      Given the message
      """
      {
        "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "retired_amount": "100",
            "retirement_jurisdiction": "US-WA"
          }
        ]
      }
      """
      And retirement reason with length "513"
      When the message is validated
      Then expect the error "credits[0]: retirement reason: max length 512: limit exceeded"


    Scenario: an error is returned if sender and recipient are the same
      Given the message
      """
      {
        "sender":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "recipient":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "tradable_amount": "100",
            "retired_amount": "100",
            "retirement_jurisdiction": "US-WA",
            "retirement_reason": "offsetting electricity consumption"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "sender and recipient cannot be the same: invalid request"


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
      When alice attempts to send credits to bob with retired amount "10" jurisdiction "US-WA" and reason "offsetting electricity consumption"
      Then expect event retire with properties
      """
      {
        "owner": "regen10yhlcvh88sux4zmf67udhg5f5z2803z6jm0d25",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "10",
        "jurisdiction": "US-WA",
        "reason": "offsetting electricity consumption"
      }
      """

    Scenario: Event EventTransfer is emitted
      When alice attempts to send credits to bob with retired amount "10" and jurisdiction "US-WA"
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