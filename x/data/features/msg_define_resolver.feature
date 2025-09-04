Feature: MsgDefineResolver

  Scenario: a valid message
    Given the message
    """
    {
      "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "resolver_url": "https://foo.bar"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if definer is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if definer is not a bech32 address
    Given the message
    """
    {
      "definer": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if resolver url is empty
    Given the message
    """
    {
      "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "invalid resolver url: invalid request"

  Scenario: an error is returned if resolver url is missing a protocol
    Given the message
    """
    {
      "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "resolver_url": "foo.bar"
    }
    """
    When the message is validated
    Then expect the error "invalid resolver url: invalid request"