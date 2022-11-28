Feature: Msg/UpdateBatchMetadata

  The metadata of a credit batch can be updated:
  - when the credit batch exists
  - when the credit batch is open
  - when the issuer is the issuer of the credit batch
  - the credit batch metadata is updated

  Rule: The credit batch must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with batch denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the credit batch exists
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch does not exist
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-002"
      Then expect the error "could not get credit batch with denom C01-001-20200101-20210101-002: not found: invalid request"

  Rule: The credit batch must be open

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"

    Scenario: the credit batch is open
      Given a credit batch with batch denom "C01-001-20200101-20210101-001" issuer alice and open "true"
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch is not open
      Given a credit batch with batch denom "C01-001-20200101-20210101-001" issuer alice and open "false"
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "credit batch C01-001-20200101-20210101-001 is not open: unauthorized"

  Rule: The issuer must be the issuer of the credit batch

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with batch denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the issuer is the issuer of the credit batch
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the issuer is not the issuer of the credit batch
      When bob attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001"
      Then expect error contains "is not the issuer of credit batch C01-001-20200101-20210101-001: unauthorized"

  Rule: The credit batch metadata is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with batch denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the credit batch metadata is updated
      When alice attempts to update batch metadata with batch denom "C01-001-20200101-20210101-001" and new metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """
      Then expect credit batch with batch denom "C01-001-20200101-20210101-001" and metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """

      # no failing scenario - state transitions only occur upon successful message execution

  Rule: The event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with batch denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: EventUpdateBatchMetadata is emitted
      When alice updates the batch metadata
      Then expect event with properties
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001"
      }
      """
