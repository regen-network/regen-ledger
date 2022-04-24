Feature: CreateProject

  Projects can be created:
  - when the creator is an approved credit class issuer
  - ...

  Rule: A project id is always unique

    Background:
      Given a credit type exists with abbreviation "C"
      And alice has created a credit class with credit type "C"

    Scenario Outline: the project id is auto-generated
      Given the project sequence number is "<next_sequence>"
      When alice creates a project with credit class id "C01"
      Then the project exists with project id "<project_id>"

      Examples:
        | next_sequence | project_id  |
        | 1             | C01-001     |
        | 2             | C01-002     |
        | 10            | C01-010     |
        | 100           | C01-100     |
        | 1000          | C01-1000    |
        | 10000         | C01-10000   |
