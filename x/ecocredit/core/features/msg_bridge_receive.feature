Feature: MsgBridgeReceive

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

  Scenario: an error is returned if issuer is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "issuer: empty address string is not allowed: invalid address"

  Scenario: an error is returned if issuer is not a valid bech32 address
    Given the message
    """
    {
      "issuer": "foo"
    }
    """
    When the message is validated
    Then expect the error "issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address"

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

  Scenario: an error is returned if batch recipient is empty
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
      "batch": {}
    }
    """
    When the message is validated
    Then expect the error "batch recipient: empty address string is not allowed: invalid address"

  Scenario: an error is returned if batch recipient is not a valid bech32 address
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
        "recipient": "foo"
      }
    }
    """
    When the message is validated
    Then expect the error "batch recipient: decoding bech32 failed: invalid bech32 string length 3: invalid address"

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
  Scenario: an error is returned if origin tx source is not polygon
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
        "source": "foo"
      }
    }
    """
    When the message is validated
    Then expect the error "origin tx source must be polygon: invalid request"

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

  # Note: additional validation for origin tx covered in origin_tx_test.go
