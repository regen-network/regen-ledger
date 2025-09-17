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