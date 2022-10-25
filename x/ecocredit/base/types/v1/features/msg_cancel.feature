Feature: MsgCancel

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        }
      ],
      "reason": "transferring credits to another registry"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple credits
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        },
        {
          "batch_denom": "C01-001-20200101-20210101-002",
          "amount": "100"
        }
      ],
      "reason": "transferring credits to another registry"
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

  Scenario: an error is returned if owner is not a bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    Then expect the error "owner: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if credits is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "credits cannot be empty: invalid request"

  Scenario: an error is returned if credits batch denom is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: batch denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if credits amount is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credits[0]: amount cannot be empty: invalid request"

  # Note: additional validation for credits covered in types_credits.feature

  Scenario: an error is returned if reason is empty
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "reason cannot be empty: invalid request"

  Scenario: an error is returned if reason exceeds 512 characters
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        }
      ]
    }
    """
    And a reason with length "513"
    When the message is validated
    Then expect the error "reason: max length 512: limit exceeded"

  Scenario: a valid amino message
    Given the message
    """
    {
      "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "100"
        }
      ],
      "reason": "transferring credits to another registry"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgCancel",
      "value":{
        "credits":[
          {
            "amount":"100",
            "batch_denom":"C01-001-20200101-20210101-001"
          }
        ],
        "owner":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "reason":"transferring credits to another registry"
      }
    }
    """
