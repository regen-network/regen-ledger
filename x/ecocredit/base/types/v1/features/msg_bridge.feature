Feature: MsgBridge

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
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
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        },
        {
          "batch_denom": "C01-001-20200101-20210101-002",
          "amount": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if owner is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "owner: empty address string is not allowed: invalid address"

  Scenario: an error is returned if owner is not a valid bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    Then expect the error "owner: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if target is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "target cannot be empty: invalid request"

  Scenario: an error is returned if recipient is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon"
    }
    """
    When the message is validated
    Then expect the error "recipient cannot be empty: invalid request"

  Scenario: an error is returned if recipient is not an ethereum address
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "foo"
    }
    """
    When the message is validated
    Then expect the error "recipient must be a valid ethereum address: invalid address"

  Scenario: an error is returned if credits is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d"
    }
    """
    When the message is validated
    Then expect the error "credits cannot be empty: invalid request"

  Scenario: an error is returned if credits batch denom is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
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
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
      "credits": [
        {
          "batch_denom": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"

  Scenario: an error is returned if credits amount is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: amount cannot be empty: invalid request"

  Scenario: an error is returned if credits amount is not a positive decimal
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "target": "polygon",
      "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: amount: expected a positive decimal, got -100: invalid decimal string"
