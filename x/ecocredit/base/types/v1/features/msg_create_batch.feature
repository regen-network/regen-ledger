Feature: MsgCreateBatch

  Scenario: a valid message
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple issuance items
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100"
        },
        {
          "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with open
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "open": true
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with origin tx
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "origin_tx": {
        "id": "0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1",
        "source": "verra"
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

  Scenario: an error is returned if issuer is not a bech32 address
    Given the message
    """
    {
      "issuer": "foo"
    }
    """
    When the message is validated
    Then expect the error "issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if project id is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "project id: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if project id is not formatted
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "foo"
    }
    """
    When the message is validated
    Then expect the error "project id: expected format <class-id>-<project-sequence>: parse error: invalid request"

  Scenario: an error is returned if issuance is empty
   Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": []
    }
    """
    When the message is validated
    Then expect the error "issuance cannot be empty: invalid request"

  Scenario: an error is returned if issuance recipient is empty
   Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "issuance[0]: recipient: empty address string is not allowed: invalid address"

  # Note: additional validation for batch issuance covered in types_batch_issuance_test.go

  Scenario: an error is returned if metadata is exceeds 256 characters
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    And metadata with length "257"
    When the message is validated
    Then expect the error "metadata: max length 256: limit exceeded"

  Scenario: an error is returned if start date is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the message is validated
    Then expect the error "start date cannot be empty: invalid request"

  Scenario: an error is returned if end date is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z"
    }
    """
    When the message is validated
    Then expect the error "end date cannot be empty: invalid request"

  Scenario: an error is returned if start date is after end date
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2021-01-01T00:00:00Z",
      "end_date": "2020-01-01T00:00:00Z"
    }
    """
    When the message is validated
    Then expect the error "start date cannot be after end date: invalid request"

  Scenario: an error is returned if origin tx id is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "origin_tx": {}
    }
    """
    When the message is validated
    Then expect the error "origin_tx.id cannot be empty: invalid request"

  Scenario: an error is returned if origin tx source is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "origin_tx": {
        "id": "0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1"
      }
    }
    """
    When the message is validated
    Then expect the error "origin_tx.source cannot be empty: invalid request"

  # Note: additional validation for origin tx covered in types_origin_tx_test.go

  Scenario: a valid amino message
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgCreateBatch",
      "value":{
        "end_date":"2021-01-01T00:00:00Z",
        "issuance":[
          {
            "recipient":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
            "retired_amount":"100",
            "retirement_jurisdiction":"US-WA",
            "tradable_amount":"100"
          }
        ],
        "issuer":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "metadata":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
        "project_id":"C01-001",
        "start_date":"2020-01-01T00:00:00Z"
      }
    }
    """
