Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - message validations
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

  Rule: Message Validations

    Scenario: a valid message
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "origin_tx": {
    "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
    "source": "polygon",
    "contract": "0x0e65079a29d7793ab5ca500c2d88e60ee99ba606"
    }
    }
    """
    When the message is validated
    Then expect no error


    Scenario: an error is returned if class id is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "class id: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if class id is not formatted
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "foo"
    }
    """
    When the message is validated
    Then expect the error "class id: expected format <credit-type-abbrev><class-sequence>: parse error: invalid request"

    Scenario: an error is returned if project is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01"
    }
    """
    When the message is validated
    Then expect the error "project cannot be empty: invalid request"

    Scenario: an error is returned if project reference id is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {}
    }
    """
    When the message is validated
    Then expect the error "project reference id cannot be empty: invalid request"

    Scenario: an error is returned if project reference id is too long
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01"
    }
    """
    And a project reference id with length "33"
    When the message is validated
    Then expect the error "project reference id: max length 32: limit exceeded"

    Scenario: an error is returned if project jurisdiction is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001"
    }
    }
    """
    When the message is validated
    Then expect the error "project jurisdiction: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if project jurisdiction is not formatted
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "foo"
    }
    }
    """
    When the message is validated
    Then expect the error "project jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

    Scenario: an error is returned if project metadata is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA"
    }
    }
    """
    When the message is validated
    Then expect the error "project metadata cannot be empty: invalid request"

    Scenario: an error is returned if project metadata is too long
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA"
    }
    }
    """
    And project metadata with length "257"
    When the message is validated
    Then expect the error "project metadata: max length 256: limit exceeded"

    Scenario: an error is returned if batch is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    }
    """
    When the message is validated
    Then expect the error "batch cannot be empty: invalid request"


    Scenario: an error is returned if batch amount is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    }
    """
    When the message is validated
    Then expect the error "batch amount cannot be empty: invalid request"

    Scenario: an error is returned if batch amount is not a positive decimal
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "-100"
    }
    }
    """
    When the message is validated
    Then expect the error "batch amount: expected a positive decimal, got -100: invalid decimal string"

    Scenario: an error is returned if batch start date is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100"
    }
    }
    """
    When the message is validated
    Then expect the error "batch start date cannot be empty: invalid request"

    Scenario: an error is returned if batch end date is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z"
    }
    }
    """
    When the message is validated
    Then expect the error "batch end date cannot be empty: invalid request"

    Scenario: an error is returned if batch start date is after batch end date
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2021-01-01T00:00:00Z",
    "end_date": "2020-01-01T00:00:00Z"
    }
    }
    """
    When the message is validated
    Then expect the error "batch start date cannot be after batch end date: invalid request"

    Scenario: an error is returned if batch metadata is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z"
    }
    }
    """
    When the message is validated
    Then expect the error "batch metadata cannot be empty: invalid request"

    Scenario: an error is returned if batch metadata is too long
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z"
    }
    }
    """
    And batch metadata with length "257"
    When the message is validated
    Then expect the error "batch metadata: max length 256: limit exceeded"

    Scenario: an error is returned if origin tx is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    }
    """
    When the message is validated
    Then expect the error "origin tx cannot be empty: invalid request"

    # Note: non-ethereum origin tx id is specific to Msg/BridgeReceive
    Scenario: an error is returned if origin tx id is not an ethereum transaction hash
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "origin_tx": {
    "id": "foo"
    }
    }
    """
    When the message is validated
    Then expect the error "origin tx id must be a valid ethereum transaction hash: invalid request"

    # Note: non-polygon origin tx source is specific to Msg/BridgeReceive
    Scenario: an error is returned if origin tx source is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "origin_tx": {
    "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e"
    }
    }
    """
    When the message is validated
    Then expect the error "origin tx source cannot be empty: invalid request"

    # Note: non-empty origin tx contract is specific to Msg/BridgeReceive
    Scenario: an error is returned if origin tx contract is empty
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "origin_tx": {
    "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
    "source": "polygon"
    }
    }
    """
    When the message is validated
    Then expect the error "origin tx contract cannot be empty: invalid request"

    Scenario: an error is returned if origin tx contract is not a valid ethereum address
    Given the message
    """
    {
    "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "class_id": "C01",
    "project": {
    "reference_id": "VCS-001",
    "jurisdiction": "US-WA",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "batch": {
    "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "amount": "100",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    },
    "origin_tx": {
    "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
    "source": "polygon",
    "contract": "foo"
    }
    }
    """
    When the message is validated
    Then expect the error "origin_tx.contract must be a valid ethereum address: invalid address"



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
