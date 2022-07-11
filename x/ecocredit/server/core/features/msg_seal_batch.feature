Feature: Msg/SealBatch

  An open credit batch can be sealed:
  - when the credit batch exists
  - when the issuer is the issuer of the credit batch

  Rule: The credit batch must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the credit batch exists
      When alice attempts to seal batch with denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch does not exist
      When alice attempts to seal batch with denom "C01-001-20200101-20210101-002"
      Then expect the error "could not get batch with denom C01-001-20200101-20210101-002: not found: invalid request"

  Rule: The issuer must be the issuer of the credit batch

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the issuer is not the credit batch issuer
      When alice attempts to seal batch with denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the issuer is the credit batch issuer
      When bob attempts to seal batch with denom "C01-001-20200101-20210101-001"
      Then expect the error "only the batch issuer can seal the batch: unauthorized"
