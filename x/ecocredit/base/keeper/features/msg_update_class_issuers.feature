Feature: Msg/UpdateClassIssuers

  The issuers of a credit class can be updated:
  - message validation
  - when the credit class exists
  - when the admin is the admin of the credit class
  - the credit class issuers are updated

  Rule: Message validations

    Scenario: a valid message
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "add_issuers": [
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
        ],
        "remove_issuers": [
          "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message without add issuers
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "remove_issuers": [
          "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message without remove issuers
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "add_issuers": [
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: an error is returned if class id is empty
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "class id: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if class id is not formatted
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "foo"
      }
      """
      When the message is validated
      Then expect the error "class id: expected format <credit-type-abbrev><class-sequence>: parse error: invalid request"

    Scenario: an error is returned if new issuers and remove issuers is empty
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01"
      }
      """
      When the message is validated
      Then expect the error "must specify at least one of add_issuers or remove_issuers: invalid request"

    Scenario: an error is returned if new issuer is a duplicate
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "add_issuers": [
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
        ]
      }
      """
      When the message is validated
      Then expect the error "add_issuers[1]: duplicate issuer: invalid request"


    Scenario: an error is returned if remove issuer is a duplicate
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "remove_issuers": [
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
        ]
      }
      """
      When the message is validated
      Then expect the error "remove_issuers[1]: duplicate issuer: invalid request"

    

  Rule: The credit class must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class exists
      When alice attempts to update class issuers with class id "C01"
      Then expect no error

    Scenario: the credit class does not exist
      When alice attempts to update class issuers with class id "C02"
      Then expect the error "could not get credit class with id C02: not found: invalid request"

  Rule: The admin must be the admin of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the admin is the admin of the credit class
      When alice attempts to update class issuers with class id "C01"
      Then expect no error

    Scenario: the admin is not the admin of the credit class
      When bob attempts to update class issuers with class id "C01"
      Then expect error contains "is not the admin of credit class C01: unauthorized"

  Rule: The credit class issuers are updated

    Background:
      Given a credit type with abbreviation "C"

    Scenario: the credit class issuers are added
      And a credit class with class id "C01" admin alice and issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
      """
      When alice attempts to update class issuers with class id "C01" and add issuers
      """
      [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """
      Then expect credit class with class id "C01" and issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """

    Scenario: the credit class issuers are removed
      And a credit class with class id "C01" admin alice and issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """
      When alice attempts to update class issuers with class id "C01" and remove issuers
      """
      [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """
      Then expect credit class with class id "C01" and issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: EventUpdateClassIssuers is emitted
      When alice attempts to update class issuers with class id "C01"
      Then expect event with properties
      """
      {
        "class_id": "C01"
      }
      """