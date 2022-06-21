Feature: CreateBatch

  Credit batches can be created:
  - when the issuer is an approved credit class issuer
  - the recipient batch balance is updated
  - ...

  Rule: A credit batch denom is always unique

    Scenario: multiple credit batches from the same project
      Given a credit type exists with abbreviation "A"
      And alice has created a credit class with credit type "A"
      And alice has created a project with credit class id "A01"
      And alice has created a credit batch with project id "A01-001"
      When alice creates a credit batch with project id "A01-001"
      Then the credit batch exists with denom "A01-001-20200101-20210101-002"

    Scenario: multiple credit batches from different projects
      Given a credit type exists with abbreviation "A"
      And a credit type exists with abbreviation "B"
      And alice has created a credit class with credit type "A"
      And alice has created a credit class with credit type "B"
      And alice has created a project with credit class id "A01"
      And alice has created a project with credit class id "B01"
      And alice has created a credit batch with project id "A01-001"
      When alice creates a credit batch with project id "B01-001"
      Then the credit batch exists with denom "B01-001-20200101-20210101-001"

  Rule: The recipient batch balance is updated

    Background:
      Given a credit type
      And alice created a credit class
      And alice is a credit class issuer
      And alice created a project

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
      Then expect batch balance for recipient with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
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
      Then expect batch balance for recipient with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
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
      Then expect batch balance for recipient with address "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "escrowed_amount": "0"
      }
      """
      And expect batch balance for recipient with address "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3"
      """
      {
        "tradable_amount": "100",
        "retired_amount": "100",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
