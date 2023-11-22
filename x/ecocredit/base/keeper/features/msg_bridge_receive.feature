Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - when the credit class exists
  - when the issuer is an approved credit class issuer
  - when the issuer is the issuer of the existing credit batch
  - when a project from the same class with the same reference id does not exist
  - when a batch from the same class with the same contract does not exist
  - the project is added using the information provided
  - the credit batch is added using the information provided
  - the recipient batch balance is updated
  - the batch supply is updated

  # see Msg/CreateProject, Msg/CreateBatch, and Msg/MintBatchCredits for additional tests

  Rule: The credit class must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: credit class exists
      When alice attempts to bridge credits with class id "C01"
      Then expect no error

    Scenario: credit class does not exist
      When alice attempts to bridge credits with class id "C02"
      Then expect the error "credit class with id C02: not found"

  Rule: The issuer must be an allowed credit class issuer

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: the issuer is not an allowed credit class issuer
      When alice attempts to bridge credits with class id "C01"
      Then expect no error

    Scenario: the issuer is an allowed credit class issuer
      When bob attempts to bridge credits with class id "C01"
      Then expect error contains "is not an issuer for the class: unauthorized"

  Rule: The issuer must be the issuer of the existing credit batch

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice
      And the batch contract
      """
      {
        "batch_key": 1,
        "class_key": 1,
        "contract": "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      }
      """

    Scenario: the issuer is the credit batch issuer
      When alice attempts to bridge credits with contract "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      Then expect no error

    Scenario: the issuer is not the credit batch issuer
      When bob attempts to bridge credits with contract "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      Then expect the error "only the account that issued the batch can mint additional credits: unauthorized"

  Rule: A new project is created if a project from the same class with the same reference id does not exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"
      And a project with id "C01-001" and reference id "VCS-001"

    Scenario: a project from the same class with a different reference id
      When alice attempts to bridge credits with class id "C01" and project reference id "VCS-002"
      Then expect total projects "2"

    Scenario: a project from a different class with the same reference id
      Given a credit class with id "C02" and issuer alice
      When alice attempts to bridge credits with class id "C02" and project reference id "VCS-001"
      Then expect total projects "2"

    Scenario: a project from the same class with the same reference id
      When alice attempts to bridge credits with class id "C01" and project reference id "VCS-001"
      Then expect total projects "1"

  Rule: A new credit batch is created if batch from the same class with the same contract does not exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice

    Scenario: batch contract entry exists
      Given the batch contract
      """
      {
        "batch_key": 1,
        "class_key": 1,
        "contract": "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      }
      """
      When alice attempts to bridge credits with contract "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      Then expect total credit batches "1"

    Scenario: batch contract entry does not exist
      When alice attempts to bridge credits with contract "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      Then expect total credit batches "2"

  Rule: The project is added using the information provided

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: the project properties are added
      When alice attempts to bridge credits with project properties
      """
      {
        "reference_id": "VCS-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "jurisdiction": "US-WA"
      }
      """
      Then expect project properties
      """
      {
        "id": "C01-001",
        "reference_id": "VCS-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "jurisdiction": "US-WA"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The credit batch is added using the information provided

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: the batch properties are added
      When alice attempts to bridge credits with batch properties
      """
      {
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z"
      }
      """
      Then expect batch properties
      """
      {
        "denom": "C01-001-20200101-20210101-001",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The recipient batch balance is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: balance updated from issuance with single item
      When alice attempts to bridge credits to bob with tradable amount "10"
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

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: supply updated from issuance
      When alice attempts to bridge credits to bob with tradable amount "10"
      Then expect batch supply
      """
      {
        "tradable_amount": "10",
        "retired_amount": "0",
        "cancelled_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The OriginTx source must be in the AllowedBridgeChain table

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"

    Scenario: the OriginTx source is in the AllowedBridgeChain table
      When alice attempts to bridge credits with OriginTx Source "polygon"
      Then expect no error

    Scenario: the OriginTx source is not in the AllowedBridgeChain table
      When alice attempts to bridge credits with OriginTx Source "solana"
      Then expect the error "solana is not an authorized bridge source: unauthorized"

  Rule: Event is emitted

    Scenario: EventBridgeReceive is emitted
      Given a credit type with abbreviation "C"
      And a credit class with id "C01" and issuer alice
      And allowed bridge chain "polygon"
      And a project with id "C01-001"
      And a credit batch with denom "C01-001-20200101-20210101-001" and issuer alice
      And the batch contract
      """
      {
        "batch_key": 1,
        "class_key": 1,
        "contract": "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      }
      """
      When alice attempts to bridge credits with contract "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
      Then expect event with properties
      """
      {
        "project_id": "C01-001",
        "batch_denom": "C01-001-20200101-20210101-001",
        "amount": "10",
        "origin_tx": {
          "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
          "source": "polygon",
          "contract": "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606"
        }
      }
      """
