Feature: CreateBatch

  Background:
    Given alice has created a credit class

  Scenario: a batch with a custom project id
    Given alice has created a project with id "P1"
    When alice creates a credit batch with project id "P1"
    Then the credit batch exists with denom "C01-P1-20200101-20210101-001"

  Scenario: a batch with an auto-generated project id
    Given alice has created a project with id ""
    When alice creates a credit batch with project id "001"
    Then the credit batch exists with denom "C01-001-20200101-20210101-001"
