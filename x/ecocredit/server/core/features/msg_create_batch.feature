Feature: CreateBatch

  Credit batches can be created:
  - when the issuer is an approved credit class issuer
  - ...

  Rule: A credit batch denom is always unique

    Background:
      Given a credit type exists with abbreviation "C"
      And alice has created a credit class with credit type "C"
      And alice has created a project with credit class id "C01"

    Scenario: credit classes from the same project
      When alice creates a credit batch with project id "C01-001"
      Then the credit batch exists with denom "C01-001-20200101-20210101-001"

    Scenario: credit classes from different projects
      Given a credit type exists with abbreviation "BIO"
      And alice has created a credit class with credit type "BIO"
      And alice has created a project with credit class id "BIO01"
      And alice has created a credit batch with project id "C01-001"
      When alice creates a credit batch with project id "BIO01-001"
      Then the credit batch exists with denom "BIO01-001-20200101-20210101-001"
