Feature: Msg/UpdateProjectAdmin

  The admin of a project can be updated:
  - when the project exists
  - when the admin is the admin of the project
  - the project admin is updated

  Rule: The project must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the project exists
      When alice attempts to update project admin with project id "C01-001"
      Then expect no error

    Scenario: the project does not exist
      When alice attempts to update project admin with project id "C01-002"
      Then expect the error "could not get project with id C01-002: not found: invalid request"

  Rule: The admin must be the admin of the project

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the admin is the admin of the project
      When alice attempts to update project admin with project id "C01-001"
      Then expect no error

    Scenario: the admin is not the admin of the project
      When bob attempts to update project admin with project id "C01-001"
      Then expect error contains "is not the admin of project C01-001: unauthorized"

  Rule: The project admin is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the project admin is updated
      When alice attempts to update project admin with project id "C01-001" and new admin bob
      Then expect project with project id "C01-001" and admin bob

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: EventUpdateProjectAdmin is emitted
      When alice attempts to update project admin with project id "C01-001" and new admin bob
      Then expect event with properties
      """
      {
        "project_id": "C01-001"
      }
      """