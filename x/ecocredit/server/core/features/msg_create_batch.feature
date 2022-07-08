Feature: Msg/CreateBatch

  Credits can be issued in batches:
  - when the project exists
  - when the issuer is an allowed credit class issuer
  - when the decimal places in issuance amount does not exceed credit type precision
  - when the origin tx is unique within the scope of a credit class
  - when the contract is unique within the scope of a credit class
  - the recipient batch balance is updated
  - the batch supply is updated
  - the batch sequence is updated
  - the batch properties are added
  - the batch contract mapping is added
  - the response includes the batch denom

  Rule: The project must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the project exists
      When alice attempts to create a batch with project id "C01-001"
      Then expect no error

    Scenario: the project does not exist
      When alice attempts to create a batch with project id "C01-002"
      Then expect the error "could not get project with id C01-002: not found: invalid request"

  Rule: The issuer must be an allowed credit class issuer

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the issuer is an allowed credit class issuer
      When alice attempts to create a batch with project id "C01-001"
      Then expect no error

    Scenario: the issuer is not an allowed credit class issuer
      When bob attempts to create a batch with project id "C01-001"
      Then expect error contains "is not an issuer for the class: unauthorized"

  Rule: The decimal places in issuance amount must not exceed credit type precision

    Background:
      Given a credit type with abbreviation "C" and precision "6"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario Outline: the decimal places in tradable amount is less than or equal to credit type precision
      When alice attempts to create a batch with project id "C01-001" and tradable amount "<amount>"
      Then expect no error

      Examples:
        | description | amount   |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: the decimal places in tradable amount is greater than credit type precision
      When alice attempts to create a batch with project id "C01-001" and tradable amount "9.1234567"
      Then expect the error "9.1234567 exceeds maximum decimal places: 6: invalid request"

    Scenario Outline: the decimal places in retired amount is less than or equal to credit type precision
      When alice attempts to create a batch with project id "C01-001" and retired amount "<amount>"
      Then expect no error

      Examples:
        | description | amount   |
        | less than   | 9.12345  |
        | equal to    | 9.123456 |

    Scenario: the decimal places in retired amount is greater than credit type precision
      When alice attempts to create a batch with project id "C01-001" and retired amount "9.1234567"
      Then expect the error "9.1234567 exceeds maximum decimal places: 6: invalid request"

  Rule: The origin tx must be unique within the scope of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the origin tx is not unique within the credit class
      Given an origin tx index
      """
      {
        "class_key": 1,
        "id": "0x64",
        "source": "polygon"
      }
      """
      When alice attempts to create a batch with project id "C01-001" and origin tx
      """
      {
        "id": "0x64",
        "source": "polygon"
      }
      """
      Then expect the error "credits already issued with tx id: 0x64: invalid request"

    Scenario: the origin tx is unique within the credit class
      Given an origin tx index
      """
      {
        "class_key": 2,
        "id": "0x64",
        "source": "polygon"
      }
      """
      When alice attempts to create a batch with project id "C01-001" and origin tx
      """
      {
        "id": "0x64",
        "source": "polygon"
      }
      """
      Then expect no error

  Rule: The contract must be unique within the scope of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the contract is not unique within credit class
      Given a batch contract
      """
      {
        "batch_key": 2,
        "class_key": 1,
        "contract": "0x40"
      }
      """
      When alice attempts to create a batch with project id "C01-001" and origin tx
      """
      {
        "id": "0x64",
        "source": "polygon",
        "contract": "0x40"
      }
      """
      Then expect the error "credit batch with contract already exists: 0x40: invalid request"

    Scenario: the contract is unique within the credit class
      Given a batch contract
      """
      {
        "batch_key": 2,
        "class_key": 2,
        "contract": "0x40"
      }
      """
      When alice attempts to create a batch with project id "C01-001" and origin tx
      """
      {
        "id": "0x64",
        "source": "polygon",
        "contract": "0x40"
      }
      """
      Then expect no error

  Rule: The recipient batch balance is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: balance updated from issuance with single item
      When alice attempts to create a batch with project id "C01-001" and issuance
      """
      [
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
      """
      Then expect recipient batch balance with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "escrowed_amount": "0"
      }
      """

    Scenario: balance updated from issuance with multiple items and same recipient
      When alice attempts to create a batch with project id "C01-001" and issuance
      """
      [
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        },
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
      """
      Then expect recipient batch balance with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
      """
      {
        "tradable_amount": "200",
        "retired_amount": "200",
        "escrowed_amount": "0"
      }
      """

    Scenario: balance updated from issuance with multiple items and different recipients
      When alice attempts to create a batch with project id "C01-001" and issuance
      """
      [
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        },
        {
          "recipient": "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
      """
      Then expect recipient batch balance with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "escrowed_amount": "0"
      }
      """
      And expect recipient batch balance with address "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3"
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The batch supply is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: supply updated from issuance with single item
      When alice attempts to create a batch with project id "C01-001" and issuance
      """
      [
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
      """
      Then expect batch supply
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "cancelled_amount": "0"
      }
      """

    Scenario: supply updated from issuance with multiple items
      When alice attempts to create a batch with project id "C01-001" and issuance
      """
      [
        {
          "recipient": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        },
        {
          "recipient": "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
      """
      Then expect batch supply
      """
      {
        "tradable_amount": "200",
        "retired_amount": "200",
        "cancelled_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the batch sequence is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"
      And a project with project id "C01-002"

    Scenario: the batch sequence is updated
      Given a batch sequence with project id "C01-001" and next sequence "1"
      When alice attempts to create a batch with project id "C01-001"
      Then expect batch sequence with project id "C01-001" and next sequence "2"

    Scenario: the batch sequence is not updated
      Given a batch sequence with project id "C01-001" and next sequence "1"
      When alice attempts to create a batch with project id "C01-002"
      Then expect batch sequence with project id "C01-001" and next sequence "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the batch properties are added

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the batch properties are added
      When alice attempts to create a batch with properties
      """
      {
        "project_id": "C01-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z"
      }
      """
      Then expect batch properties
      """
      {
        "denom": "C01-001-20200101-20210101-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the batch contract mapping is added

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the batch contract mapping is added
      When alice attempts to create a batch with project id "C01-001" and origin tx
      """
      {
        "id": "0x64",
        "source": "polygon",
        "contract": "0x40"
      }
      """
      Then expect batch contract
      """
      {
        "batch_key": 1,
        "class_key": 1,
        "contract": "0x40"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the response includes the batch denom

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a project with project id "C01-001"

    Scenario: the response includes the batch denom
      When alice attempts to create a batch with project id "C01-001"
      Then expect the response
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001"
      }
      """

    # no failing scenario - response should always be empty when message execution fails
