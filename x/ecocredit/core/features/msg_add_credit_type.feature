Feature: CreditTypes

  Scenario: a valid credit types
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "precision":6,
        "name":"carbon",
        "unit":"ton"
      }
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

  Scenario: an error is returned if credit type is empty
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "credit type cannot be empty: invalid request"

  Scenario: an error is returned if credit type precision is not 6
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "precision":60,
        "name":"carbon",
        "unit":"ton"
      }
    }
    """
    When the message is validated
    Then expect the error "credit type precision is currently locked to 6: invalid request"

  Scenario: an error is returned if credit type abbreviation is empty
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"",
        "precision":6,
        "name":"carbon",
        "unit":"ton"
      }
    }
    """
    When the message is validated
    Then expect the error "credit type abbreviation cannot be empty: parse error"
