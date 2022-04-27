Feature: CreateProject

  Projects can be created:
  - when the creator is an approved credit class issuer
  - ...

  Rule: A project id is always unique

    Scenario Outline: the project id is auto-generated
      Given a credit type exists with abbreviation "A"
      And alice has created a credit class with credit type "A"
      And the project sequence number is "<next_sequence>"
      When alice creates a project with credit class id "A01"
      Then the project exists with project id "<project_id>"

      Examples:
        | next_sequence | project_id  |
        | 1             | A01-001     |
        | 2             | A01-002     |
        | 10            | A01-010     |
        | 100           | A01-100     |
        | 1000          | A01-1000    |
        | 10000         | A01-10000   |
