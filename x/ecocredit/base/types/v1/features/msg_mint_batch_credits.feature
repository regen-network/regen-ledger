Feature: MsgMintBatchCredits

  Scenario: a valid message
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "origin_tx": {
        "id": "0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1",
        "source": "verra"
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple issuance items
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001",
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

  Scenario: an error is returned if batch denom is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "batch denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if batch denom is not formatted
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"

  Scenario: an error is returned if issuance is empty
   Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001",
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
      "batch_denom": "C01-001-20200101-20210101-001",
      "issuance": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "issuance[0]: recipient: empty address string is not allowed: invalid address"

  # Note: additional validation for batch issuance covered in types_batch_issuance_test.go

  Scenario: an error is returned if origin tx is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001",
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
    When the message is validated
    Then expect the error "origin tx cannot be empty: invalid request"

  Scenario: an error is returned if origin tx id is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
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
      "batch_denom": "C01-001-20200101-20210101-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
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
      "batch_denom": "C01-001-20200101-20210101-001",
      "issuance": [
        {
          "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ],
      "origin_tx": {
        "id": "0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1",
        "source": "verra"
      }
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgMintBatchCredits",
      "value":{
        "batch_denom":"C01-001-20200101-20210101-001",
        "issuance":[
          {
            "recipient":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
            "retired_amount":"100",
            "retirement_jurisdiction":"US-WA",
            "tradable_amount":"100"
          }
        ],
        "issuer":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "origin_tx":{
          "id":"0001-000001-000100-VCS-VCU-003-VER-US-0003-01012020-31122020-1",
          "source":"verra"
        }
      }
    }
    """
