Feature: MsgSend

  Scenario: a valid message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple credits
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "tradable_amount": "100"
        },
        {
          "batch_denom": "C01-001-20200101-20210101-002",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if sender is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "sender: empty address string is not allowed: invalid address"

  Scenario: an error is returned if sender is not a bech32 address
    Given the message
    """
    {
      "sender": "foo"
    }
    """
    When the message is validated
    Then expect the error "sender: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if recipient is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "recipient: empty address string is not allowed: invalid address"

  Scenario: an error is returned if recipient is not a bech32 address
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "foo"
    }
    """
    When the message is validated
    Then expect the error "recipient: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if credits is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
    }
    """
    When the message is validated
    Then expect the error "credits cannot be empty: invalid request"

  Scenario: an error is returned if credits batch denom is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: batch denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if credits batch denom is not formatted
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"

  Scenario: an error is returned if credits tradable amount and retired amount are empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "tradable amount or retired amount required: invalid request"

  Scenario: an error is returned if credits tradable amount is a negative decimal
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "tradable_amount": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "expected a non-negative decimal, got -100: invalid decimal string"

  Scenario: an error is returned if credits retired amount is a negative decimal
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "retired_amount": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "expected a non-negative decimal, got -100: invalid decimal string"

  Scenario: an error is returned if credits retired amount is positive and retirement jurisdiction is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "retired_amount": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "retirement jurisdiction: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if credits retired amount is positive and retirement jurisdiction is not formatted
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "retired_amount": "100",
          "retirement_jurisdiction": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "retirement jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "recipient": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "tradable_amount": "100",
          "retired_amount": "100",
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgSend",
      "value":{
        "credits":[
          {
            "batch_denom":"C01-001-20200101-20210101-001",
            "retired_amount":"100",
            "retirement_jurisdiction":"US-WA",
            "tradable_amount":"100"
          }
        ],
        "recipient":"regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "sender":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
    }
    """
