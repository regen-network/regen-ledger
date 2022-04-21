Feature: CreateProject

  Background:
    Given alice has created a credit class

  Rule: a custom project id must be unique

    Scenario: a project id is unique
      When alice creates a project with id "P1"
      Then the project exists with id "P1"

    Scenario: a project id is not unique
      Given alice has created a project with id "P1"
      When alice creates a project with id "P1"
      Then expect the error "unique key violation"

  Rule: an auto-generated project id is always unique

    Scenario Outline: a project id is auto-generated based on the project sequence number
      Given the project sequence number is "<next_sequence>"
      When alice creates a project
      Then the project exists with id "<project_id>"

      Examples:
        | next_sequence | project_id  |
        | 1             | 001         |
        | 2             | 002         |
        | 10            | 010         |
        | 100           | 100         |
        | 1000          | 1000        |
        | 10000         | 10000       |
