Feature: Msg/Retire

  Credits can be retired by the owner:
  - message validations
  - when the credit batch exists
  - when the owner has a tradable credit balance greater than or equal to the amount to cancel
  - when the decimal places in amount to cancel does not exceed credit type precision
  - the owner credit balance is updated
  - the batch supply is updated

  Rule: Message validations

    Scenario: a valid message
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ],
        "jurisdiction": "US-WA",
        "reason": "offsetting electricity consumption"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message without reason
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ],
        "jurisdiction": "US-WA"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message with multiple credits
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          },
          {
            "batch_denom": "C01-001-20200101-20210101-002",
            "amount": "100"
          }
        ],
        "jurisdiction": "US-WA",
        "reason": "offsetting electricity consumption"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: an error is returned if credits is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "credits cannot be empty: invalid request"

    Scenario: an error is returned if credits batch denom is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {}
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: batch denom: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if credits amount is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: amount cannot be empty: invalid request"

    # Note: additional validation for credits covered in types_credits.feature

    Scenario: an error is returned if jurisdiction is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "jurisdiction: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if jurisdiction is not formatted
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ],
        "jurisdiction": "foo"
      }
      """
      When the message is validated
      Then expect the error "jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

    Scenario: an error is returned if reason exceeds 512 characters
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ],
        "jurisdiction": "US-WA"
      }
      """
      And reason with length "513"
      When the message is validated
      Then expect the error "reason: max length 512: limit exceeded"

  Rule: The credit batch must exist

    Scenario: the credit batch exists
      Given a credit batch with denom "C01-001-20200101-20210101-001"
      And alice owns tradable credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to retire credits with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch does not exist
      When alice attempts to retire credits with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "could not get batch with denom C01-001-20200101-20210101-001: not found: invalid request"

  Rule: The owner must have a tradable credit balance greater than or equal to the amount to retire

    Background:
      Given a credit batch
      And alice owns tradable credit amount "10"

    Scenario Outline: tradable balance greater than or equal to amount to retire
      When alice attempts to retire credit amount "<amount>"
      Then expect no error

      Examples:
        | description  | amount |
        | greater than | 5      |
        | equal to     | 10     |

    Scenario: tradable balance less than amount to retire
      When alice attempts to retire credit amount "15"
      Then expect the error "tradable balance: 10, retire amount 15: insufficient credit balance"

  Rule: The decimal places in amount to retire must not exceed credit type precision

    Background:
      Given a credit type with abbreviation "C" and precision "6"
      And a credit batch from credit class with credit type "C"
      And alice owns tradable credit amount "10"

    Scenario Outline: the decimal places in amount is less than or equal to credit type precision
      When alice attempts to retire credit amount "<amount>"
      Then expect no error

      Examples:
        | description | amount   |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: the decimal places in amount is greater than credit type precision
      When alice attempts to retire credit amount "9.1234567"
      Then expect the error "9.1234567 exceeds maximum decimal places: 6: invalid request"

  Rule: The owner balance is updated

    Scenario: the owner balance is updated
      Given a credit batch
      And alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to retire credit amount "10"
      Then expect alice batch balance
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
      When alice attempts to retire credit amount "10"
      Then expect batch supply
      """
      {
        "retired_amount": "10",
        "tradable_amount": "0",
        "cancelled_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Scenario: EventRetire is emitted
      Given a credit batch with denom "C01-001-20200101-20210101-001"
      And alice's address "regen15406g34dl5v9780tx2q3vtjdpkdgq4hhegdtm9"
      And alice owns tradable credit amount "10"
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to retire credit amount "10" with jurisdiction "US-WA" and reason "offsetting electricity consumption"
      Then expect event with properties
      """
      {
        "owner": "regen15406g34dl5v9780tx2q3vtjdpkdgq4hhegdtm9",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "10",
        "jurisdiction": "US-WA",
        "reason": "offsetting electricity consumption"
      }
      """
