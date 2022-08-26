Feature: MsgAddCreditType

  Scenario: a valid message
    Given the message
    """
    {
      "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "name":"carbon",
        "unit":"ton",
        "precision":6
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

  Scenario: an error is returned if credit type abbreviation is empty
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"",
        "name":"carbon",
        "unit":"ton",
        "precision":6
      }
    }
    """
    When the message is validated
    Then expect the error "credit type: abbreviation: empty string is not allowed: parse error: invalid request"
  
  Scenario: an error is returned if credit type name is empty
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "name":"",
        "unit":"ton",
        "precision":6
      }
    }
    """
    When the message is validated
    Then expect the error "credit type: name cannot be empty: parse error: invalid request"
  
  Scenario: an error is returned if credit type unit is empty
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "name":"carbon",
        "unit":"",
        "precision":6
      }
    }
    """
    When the message is validated
    Then expect the error "credit type: unit cannot be empty: parse error: invalid request"

  Scenario: an error is returned if credit type precision is not 6
    Given the message
    """
    {
      "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "credit_type": {
        "abbreviation":"C",
        "name":"carbon",
        "unit":"ton",
        "precision":60
      }
    }
    """
    When the message is validated
    Then expect the error "credit type: precision is currently locked to 6: parse error: invalid request"
