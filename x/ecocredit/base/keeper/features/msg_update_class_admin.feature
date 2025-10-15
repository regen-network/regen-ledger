Feature: Msg/UpdateClassAdmin

  The admin of a credit class can be updated:
  - message validations
  - when the credit class exists
  - when the admin is the admin of the credit class
  - the credit class admin is updated

  Rule: Message validations

    Scenario: a valid message
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "new_admin": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
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


    Scenario: an error is returned if admin and new admin are the same
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "new_admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "admin and new admin cannot be the same: invalid request"

  Rule: The credit class must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class exists
      When alice attempts to update class admin with class id "C01"
      Then expect no error

    Scenario: the credit class does not exist
      When alice attempts to update class admin with class id "C02"
      Then expect the error "could not get credit class with id C02: not found: invalid request"

  Rule: The admin must be the admin of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the admin is the admin of the credit class
      When alice attempts to update class admin with class id "C01"
      Then expect no error

    Scenario: the admin is not the admin of the credit class
      When bob attempts to update class admin with class id "C01"
      Then expect error contains "is not the admin of credit class C01: unauthorized"

  Rule: The credit class admin is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class admin is updated
      When alice attempts to update class admin with class id "C01" and new admin bob
      Then expect credit class with class id "C01" and admin bob

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class admin is updated
      When alice attempts to update class admin with class id "C01" and new admin bob
      Then expect credit class with class id "C01" and admin bob
      And expect event with properties
      """
      {
        "class_id": "C01"
      }
      """