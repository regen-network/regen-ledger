Feature: CreateBatch

  Credit batches can be created:
  - when the issuer is an approved credit class issuer
  - ...

  Rule: A credit batch denom is always unique

    Scenario: multiple credit batches from the same project
      Given a credit type exists with abbreviation "A"
      And alice has created a credit class with credit type "A"
      And alice has created a project with credit class id "A01"
      And alice has created a credit batch with project id "A01-001"
      When alice creates a credit batch with project id "A01-001"
      Then the credit batch exists with denom "A01-001-20200101-20210101-002"

    Scenario: multiple credit batches from different projects
      Given a credit type exists with abbreviation "A"
      And a credit type exists with abbreviation "B"
      And alice has created a credit class with credit type "A"
      And alice has created a credit class with credit type "B"
      And alice has created a project with credit class id "A01"
      And alice has created a project with credit class id "B01"
      And alice has created a credit batch with project id "A01-001"
      When alice creates a credit batch with project id "B01-001"
      Then the credit batch exists with denom "B01-001-20200101-20210101-001"
