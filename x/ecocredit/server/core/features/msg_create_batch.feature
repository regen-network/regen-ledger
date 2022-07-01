Feature: Msg/CreateBatch

  Credits can be issued in batches:
  - when the issuer is an approved credit class issuer
  - when the origin tx is unique within the scope of a credit class
  - the recipient batch balance is updated
  - ...

  Rule: A credit batch denom is always unique

    Scenario: multiple credit batches from the same project
      Given a credit type with abbreviation "A"
      And a credit class with class id "A01" and issuer alice
      And a project with project id "A01-001"
      And a credit batch with denom "A01-001-20200101-20210101-001"
      When alice attempts to create a credit batch with project id "A01-001" start date "2020-01-01" and end date "2021-01-01"
      Then expect credit batch with denom "A01-001-20200101-20210101-002"

    Scenario: multiple credit batches from different projects
      Given a credit type with abbreviation "A"
      And a credit class with class id "A01" and issuer alice
      And a project with project id "A01-001"
      And a credit batch with denom "A01-001-20200101-20210101-001"
      And a credit type with abbreviation "B"
      And a credit class with class id "B01" and issuer alice
      And a project with project id "B01-001"
      When alice attempts to create a credit batch with project id "B01-001" start date "2020-01-01" and end date "2021-01-01"
      Then expect credit batch with denom "B01-001-20200101-20210101-001"

  Rule: The origin tx must be unique within the scope of the credit class

    Background:
      Given a credit type
      And a credit class with issuer alice
      And a project

    Scenario: the origin tx is not unique within credit class
      Given an origin tx index
      """
      {
        "class_key": 1,
        "id": "0x0",
        "source": "polygon"
      }
      """
      When alice attempts to create a credit batch with origin tx
      """
      {
        "id": "0x0",
        "source": "polygon"
      }
      """
      Then expect the error "credits already issued with tx id: 0x0: invalid request"

    Scenario: the origin tx is unique within the credit class
      Given an origin tx index
      """
      {
        "class_key": 2,
        "id": "0x0",
        "source": "polygon"
      }
      """
      When alice attempts to create a credit batch with origin tx
      """
      {
        "id": "0x0",
        "source": "polygon"
      }
      """
      Then expect no error

  Rule: The recipient batch balance is updated

    Background:
      Given a credit type
      And a credit class with issuer alice
      And a project

    Scenario: issuance with single item
      When alice attempts to create a credit batch with the issuance
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

    Scenario: issuance with multiple items and same recipient
      When alice attempts to create a credit batch with the issuance
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

    Scenario: issuance with multiple items and different recipients
      When alice attempts to create a credit batch with the issuance
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
