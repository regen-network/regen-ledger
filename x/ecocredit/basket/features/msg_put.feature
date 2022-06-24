Feature: MsgPut

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
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

  Scenario: an error is returned if owner is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid request"

  Scenario: an error is returned if owner is not a bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid request"

  Scenario: an error is returned if basket denom is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then expect the error "basket denom cannot be empty: invalid request"

  Scenario: an error is returned if basket denom is not formatted
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "foo is not a valid basket denom: invalid request"

  Scenario: an error is returned if credit list is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT"
    }
    """
    When the message is validated
    Then expect the error "credits cannot be empty: invalid request"

  Scenario: an error is returned if a credit batch denom is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
      "credits": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "credit batch denom cannot be empty: invalid request"

  Scenario: an error is returned if a credit batch denom is not formatted
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
      "credits": [
        {
          "batch_denom": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "invalid batch denom: expected format A00-000-00000000-00000000-000: parse error: invalid request"

  Scenario: an error is returned if a credit amount is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "credit amount cannot be empty: invalid request"

  Scenario: an error is returned if a credit amount is not an integer
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "parse mantissa: foo: invalid decimal string: invalid decimal string: invalid request"

  Scenario: an error is returned if a credit amount is less than zero
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "basket_denom": "eco.uC.NCT",
      "credits": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "amount": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "expected a positive decimal, got -100: invalid decimal string: invalid request"
