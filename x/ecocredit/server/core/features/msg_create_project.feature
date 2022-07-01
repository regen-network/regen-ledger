Feature: CreateProject

  Projects can be created:
  - when the creator is an approved credit class issuer
  - when the non-empty reference id is unique within the scope of the credit class
  - ...

  Rule: A project id is always unique

    Scenario: multiple projects from the same credit class
      Given a credit type with abbreviation "A"
      And a credit class with class id "A01" and issuer alice
      And a project with project id "A01-001"
      When alice attempts to create a project with class id "A01"
      Then expect project with project id "A01-002"

    Scenario: multiple projects from different credit classes
      Given a credit type with abbreviation "A"
      And a credit type with abbreviation "B"
      And a credit class with class id "A01" and issuer alice
      And a credit class with class id "B01" and issuer alice
      And a project with project id "A01-001"
      And a project with project id "B01-001"
      When alice attempts to create a project with class id "B01"
      Then expect project with project id "B01-002"

    Scenario Outline: project id is formatted correctly
      Given a credit type with abbreviation "A"
      And a credit class with class id "A01" and issuer alice
      And a project sequence "<next_sequence>" for credit class "A01"
      When alice attempts to create a project with class id "A01"
      Then expect project with project id "<project_id>"

      Examples:
        | next_sequence | project_id  |
        | 1             | A01-001     |
        | 2             | A01-002     |
        | 10            | A01-010     |
        | 100           | A01-100     |
        | 1000          | A01-1000    |
        | 10000         | A01-10000   |

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
