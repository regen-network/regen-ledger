Feature: CreateProject

  Projects can be created:
  - when the credit class exists
  - when the admin is an allowed credit class issuer
  - when the non-empty reference id is unique within the scope of the credit class
  - the project sequence is updated
  - the project properties are added
  - the response includes the project id

  Rule: The credit class must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice

    Scenario: the credit class exists
      When alice attempts to create a project with class id "C01"
      Then expect no error

    Scenario: the credit class does not exist
      When alice attempts to create a project with class id "C02"
      Then expect the error "could not get class with id C02: not found: invalid request"

  Rule: The admin must be an allowed credit class issuer

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice

    Scenario: the issuer is an allowed credit class issuer
      When alice attempts to create a project with class id "C01"
      Then expect no error

    Scenario: the issuer is not an allowed credit class issuer
      When bob attempts to create a project with class id "C01"
      Then expect error contains "is not an issuer for the class: unauthorized"

  Rule: A non-empty reference id must be unique within the scope of a credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice

    Scenario: non-empty reference id is unique within credit class
      Given a project with project id "C01-001" and reference id "VCS-001"
      When alice attempts to create a project with class id "C01" and reference id "VCS-002"
      Then expect no error

    Scenario: empty reference id is allowed for multiple projects
      Given a project with project id "C01-001" and reference id ""
      When alice attempts to create a project with class id "C01" and reference id ""
      Then expect no error

    Scenario: non-empty reference id is not unique within credit class
      Given a project with project id "C01-001" and reference id "VCS-001"
      When alice attempts to create a project with class id "C01" and reference id "VCS-001"
      Then expect the error "a project with reference id VCS-001 already exists within this credit class: invalid request"

  Rule: the project sequence is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice
      And a credit class with class id "C02" and issuer alice

    Scenario: the project sequence is updated
      Given a project sequence with class id "C01" and next sequence "1"
      When alice attempts to create a project with class id "C01"
      Then expect project sequence with class id "C01" and next sequence "2"

    Scenario: the project sequence is not updated
      Given a project sequence with class id "C01" and next sequence "1"
      When alice attempts to create a project with class id "C02"
      Then expect project sequence with class id "C01" and next sequence "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the project properties are added

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice

    Scenario: the project properties are added
      When alice attempts to create a project with properties
      """
      {
        "class_id": "C01",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "jurisdiction": "US-WA",
        "reference_id": "VCS-001"
      }
      """
      Then expect project properties
      """
      {
        "id": "C01-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "jurisdiction": "US-WA",
        "reference_id": "VCS-001"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: the response includes the project id

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and issuer alice

    Scenario: the response includes the project id
      When alice attempts to create a project with class id "C01"
      Then expect the response
      """
      {
        "project_id": "C01-001"
      }
      """

    # no failing scenario - response should always be empty when message execution fails
