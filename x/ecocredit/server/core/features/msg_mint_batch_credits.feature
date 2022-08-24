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

  Rule: events are emitted

    Background:
      Given ecocredit module address "regen15406g34dl5v9780tx2q3vtjdpkdgq4hhegdtm9"
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice
      And an origin tx with properties
      """
      {
        "id": "0xbca488b181e3dd66db06f0cccf083004c99a078bcaa70001579e465bb833fd67",
        "source": "polygon",
        "contract": "0x00192fb10df37c9fb26829eb2cc623cd1bf599e8",
        "note": "transaction confirmed by bridge service"
      }
      """

    Scenario: Event EventRetire is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with retired amount "10" from "US-WA" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event retire with properties
      """
      {
        "owner": "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "10",
        "jurisdiction": "US-WA"
      }
      """

    Scenario: Event EventMint is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with retired amount "10" from "US-WA" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event mint with properties
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001",
        "retired_amount": "10",
        "tradable_amount": "0"
      }
      """

    Scenario: Event EventMintBatchCredits is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with retired amount "10" from "US-WA" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event mint batch credits with properties
      """
      {
      "batch_denom": "C01-001-20200101-20210101-001",
      "origin_tx": {
        "id": "0xbca488b181e3dd66db06f0cccf083004c99a078bcaa70001579e465bb833fd67",
        "source": "polygon",
        "contract": "0x00192fb10df37c9fb26829eb2cc623cd1bf599e8",
        "note": "transaction confirmed by bridge service"
        }
      }
      """

    Scenario: Event EventTransfer is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with tradable amount "10" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event transfer with properties
      """
      {
        "sender": "regen15406g34dl5v9780tx2q3vtjdpkdgq4hhegdtm9",
        "recipient": "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8",
        "batch_denom": "C01-001-20200101-20210101-001",
        "tradable_amount": "10",
        "retired_amount": ""
      }
      """

    Scenario: Event EventMint is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with tradable amount "10" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event mint with properties
      """
      {
        "batch_denom": "C01-001-20200101-20210101-001",
        "retired_amount": "0",
        "tradable_amount": "10"
      }
      """

    Scenario: Event EventMintBatchCredits is emitted
      When alice attempts to mint credits with batch denom "C01-001-20200101-20210101-001" with tradable amount "10" to "regen1sl2dsfyf2znn48ehwqg28cv3nuglxkx4h7q5l8"
      Then expect event mint batch credits with properties
      """
      {
      "batch_denom": "C01-001-20200101-20210101-001",
      "origin_tx": {
        "id": "0xbca488b181e3dd66db06f0cccf083004c99a078bcaa70001579e465bb833fd67",
        "source": "polygon",
        "contract": "0x00192fb10df37c9fb26829eb2cc623cd1bf599e8",
        "note": "transaction confirmed by bridge service"
        }
      }
      """