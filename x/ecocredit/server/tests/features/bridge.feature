Feature: Bridge Integration

  Background:
    Given ecocredit state
    """
    {
      "regen.ecocredit.v1.Class": [
        1,
        {
          "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
          "credit_type_abbrev": "C",
          "id": "C01",
          "key": "1",
          "metadata": "metadata"
        }
      ],
      "regen.ecocredit.v1.ClassIssuer": [
        {
          "class_key": "1",
          "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
        }
      ],
      "regen.ecocredit.v1.CreditType": [
        {
          "abbreviation": "C",
          "name": "carbon",
          "precision": 6,
          "unit": "metric ton CO2 equivalent"
        }
      ]
    }
    """

  Scenario: bridge credits to regen from polygon and from regen to polygon
    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-001",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000001",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000001"
      }
    }
    """
    Then expect no error
    And expect event bridge receive with values
    """
    {
      "project_id": "C01-001",
      "batch_denom": "C01-001-20200101-20210101-001"
    }
    """
    And expect total projects "1"
    And expect project with properties
    """
    {
      "id": "C01-001",
      "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "jurisdiction": "US-WA",
      "reference_id": "VCS-001"
    }
    """
    And expect total credit batches "1"
    And expect credit batch with properties
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "open": true
    }
    """
    And expect batch supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """

    # providing the same origin tx id and source will error

    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-001",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000001",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000001"
      }
    }
    """
    Then expect the error
    """
    credits already issued with tx id: 0x0000000000000000000000000000000000000000000000000000000000000001: invalid request
    """

    # providing a new origin tx id with the same origin tx contract will mint to the existing credit batch

    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-001",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000002",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000001"
      }
    }
    """
    Then expect no error
    And expect event bridge receive with values
    """
    {
      "project_id": "C01-001",
      "batch_denom": "C01-001-20200101-20210101-001"
    }
    """
    And expect total projects "1"
    And expect total credit batches "1"
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "200",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """

    # providing a new origin tx contract will create a new credit batch

    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-001",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000003",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000002"
      }
    }
    """
    Then expect no error
    And expect event bridge receive with values
    """
    {
      "project_id": "C01-001",
      "batch_denom": "C01-001-20200101-20210101-002"
    }
    """
    And expect total projects "1"
    And expect total credit batches "2"
    And expect credit batch with properties
    """
    {
      "denom": "C01-001-20200101-20210101-002",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "open": true
    }
    """
    And expect batch supply with batch denom "C01-001-20200101-20210101-002"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-002"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """

    # providing an existing origin tx contract and a new project reference id will ignore the project reference id

    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-002",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000004",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000002"
      }
    }
    """
    Then expect no error
    And expect event bridge receive with values
    """
    {
      "project_id": "C01-001",
      "batch_denom": "C01-001-20200101-20210101-002"
    }
    """
    And expect total projects "1"
    And expect total credit batches "2"
    And expect batch supply with batch denom "C01-001-20200101-20210101-002"
    """
    {
      "tradable_amount": "200",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-002"
    """
    {
      "tradable_amount": "200",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """

    # providing a new origin tx contract and a new project reference id will create a new project

    When bridge service calls bridge receive with message
    """
    {
      "class_id": "C01",
      "issuer": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "project": {
        "reference_id": "VCS-002",
        "jurisdiction": "US-WA",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      },
      "batch": {
        "recipient": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
        "amount": "100",
        "start_date": "2020-01-01T00:00:00Z",
        "end_date": "2021-01-01T00:00:00Z",
        "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      },
      "origin_tx": {
        "id": "0x0000000000000000000000000000000000000000000000000000000000000005",
        "source": "polygon",
        "contract": "0x0000000000000000000000000000000000000003"
      }
    }
    """
    Then expect no error
    And expect event bridge receive with values
    """
    {
      "project_id": "C01-002",
      "batch_denom": "C01-002-20200101-20210101-001"
    }
    """
    And expect total projects "2"
    And expect project with properties
    """
    {
      "id": "C01-002",
      "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "jurisdiction": "US-WA",
      "reference_id": "VCS-002"
    }
    """
    And expect total credit batches "3"
    And expect batch supply with batch denom "C01-002-20200101-20210101-001"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-002-20200101-20210101-001"
    """
    {
      "tradable_amount": "100",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """

    # credits can be bridged back to the source chain

    When recipient calls bridge with message
    """
    {
      "owner": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
      "target": "polygon",
      "recipient": "0x1000000000000000000000000000000000000000",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "200"
        }
      ]
    }
    """
    Then expect no error
    And expect event bridge with values
    """
    {
      "target": "polygon",
      "recipient": "0x1000000000000000000000000000000000000000",
      "contract": "0x0000000000000000000000000000000000000001",
      "amount": "200"
    }
    """
    And expect batch supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "0",
      "retired_amount": "0",
      "cancelled_amount": "200"
    }
    """
    And expect batch balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "0",
      "retired_amount": "0",
      "escrowed_amount": "0"
    }
    """
