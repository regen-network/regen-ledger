Feature: MsgUpdateClassFee

  Scenario: a valid message
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "fee": {
        "denom":"uregen",
        "amount":"1000"
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message removing the fee
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if authority address is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "invalid authority address: empty address string is not allowed"

  Scenario: an error is returned if authority address is not a valid bech32 address
    Given the message
    """
    {
        "authority": "foo"
    }
    """
    When the message is validated
    Then expect the error "invalid authority address: decoding bech32 failed: invalid bech32 string length 3"

  Scenario: an error is returned if fee denom is not a valid bank denom
    Given the message
    """
    {
        "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "fee": {
          "denom": "1",
          "amount": "1000"
        }
    }
    """
    When the message is validated
    Then expect the error "invalid denom: 1: invalid request"

