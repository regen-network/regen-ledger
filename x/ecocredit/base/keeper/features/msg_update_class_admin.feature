Feature: Msg/UpdateClassAdmin

  The admin of a credit class can be updated:
  - when the credit class exists
  - when the admin is the admin of the credit class
  - the credit class admin is updated

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