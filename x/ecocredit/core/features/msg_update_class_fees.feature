Feature: MsgUpdateClassFees

  Scenario: a valid message
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "fees": [
        {
          "denom":"uregen",
          "amount":"1000"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error
  
  Scenario: a valid message with multiple fees
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "fees": [
        {
          "denom":"uatom",
          "amount":"1000"
        },
        {
          "denom":"uregen",
          "amount":"1000"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error
  
  Scenario: a valid message without fees
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "fees": []
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

  Scenario: an error is returned if coin denom is invalid
    Given the message
    """
    {
        "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "fees": [
          {
            "denom": "A+B",
            "amount": "1000"
          }
        ]
    }
    """
    When the message is validated
    Then expect the error "invalid denom: A+B"
