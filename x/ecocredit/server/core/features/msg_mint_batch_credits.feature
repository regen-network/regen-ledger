Feature: Msg/MintBatchCredits

  Credits can be dynamically minted to an existing credit batch:
  - when the credit batch exists
  - when the credit batch is open
  - when the issuer is the issuer of the credit batch
  - when the origin tx id and source is unique within the scope of the credit class
  - when the contract is unique within the scope of the credit class
  - the recipient batch balance is updated
  - the batch supply is updated

  Rule: The credit batch must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the credit batch exists
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the credit batch does not exist
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-002"
      Then expect the error "could not get batch with denom C01-001-20200101-20210101-002: not found: invalid request"

  Rule: The credit batch must be open

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"

    Scenario: the credit batch is closed
      Given a credit batch with denom "C01-001-20200101-20210101-001" open "false" and issuer alice
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "credits cannot be minted in a closed batch: invalid request"

    Scenario: the credit batch is open
      Given a credit batch with denom "C01-001-20200101-20210101-001" open "true" and issuer alice
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

  Rule: The issuer must be the issuer of the credit batch

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the issuer is not the credit batch issuer
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: the issuer is the credit batch issuer
      When bob attempts to mint credits with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "only the account that issued the batch can mint additional credits: unauthorized"

  Rule: The origin tx must be unique within the scope of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: the origin tx is not unique within credit class
      Given an origin tx index
      """
      {
        "class_key": 1,
        "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
        "source": "polygon"
      }
      """
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" and origin tx
      """
      {
        "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
        "source": "polygon"
      }
      """
      Then expect the error "credits already issued with tx id: 0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e: invalid request"

    Scenario: the origin tx is unique within credit class
      Given an origin tx index
      """
      {
        "class_key": 2,
        "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
        "source": "polygon"
      }
      """
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" and origin tx
      """
      {
        "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
        "source": "polygon"
      }
      """
      Then expect no error

  Rule: The recipient batch balance is updated

    Scenario: balance updated from issuance with single item
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" recipient bob and tradable amount "10"
      Then expect bob batch balance
      """
      {
        "tradable_amount": "10",
        "retired_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The batch supply is updated

    Scenario: supply updated from issuance
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" and tradable amount "10"
      Then expect batch supply
      """
      {
        "tradable_amount": "10",
        "retired_amount": "0",
        "cancelled_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution
