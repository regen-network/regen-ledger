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

