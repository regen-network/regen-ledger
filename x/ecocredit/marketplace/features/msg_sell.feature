Feature: MsgSell

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          }
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
    # Then expect the error "empty address string is not allowed: invalid request"

  Scenario: an error is returned if owner is not a bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    # Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid request"

  Scenario: an error is returned if orders is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    # Then expect the error "orders cannot be empty: invalid request"

  Scenario: an error is returned if order batch denom is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {}
      ]
    }
    """
    When the message is validated
    # Then expect the error "order batch denom cannot be empty: invalid request"

  Scenario: an error is returned if order batch denom is not formatted
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "invalid denom: expected format A00-000-00000000-00000000-000: parse error"

  Scenario: an error is returned if order quantity is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    # Then expect the error "order quantity cannot be empty"

  Scenario: an error is returned if order quantity is not a positive decimal
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "quantity must be positive decimal: -100: expected a positive decimal, got -100: invalid decimal string"

  Scenario: an error is returned if ask price is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "ask price cannot be empty: invalid request"

  Scenario: an error is returned if ask price denom is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {}
        }
      ]
    }
    """
    When the message is validated
    # Then expect the error "ask price denom cannot be empty: invalid request"

  Scenario: an error is returned if ask price amount is empty
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen"
          }
        }
      ]
    }
    """
    # When the message is validated
    # Then expect the error "ask price amount cannot be empty: invalid request"

  Scenario: an error is returned if ask price amount is not a positive integer
    Given the message
    """
    {
      "owner": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "0"
          }
        }
      ]
    }
    """
    When the message is validated
    # Then expect the error "ask price amount must be positive: invalid request"
