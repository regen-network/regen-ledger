Feature: MsgDefineResolver

  Scenario: an error is returned if manager is empty
    Given a message of
    """
    {
      "manager": ""
    }
    """
    When the message is validated
    Then an error of "empty address string is not allowed: invalid address"

  Scenario: an error is returned if manager is not a valid address
    Given a message of
    """
    {
      "manager": "foo"
    }
    """
    When the message is validated
    Then an error of "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if resolver url is empty
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then an error of "invalid resolver url: invalid request"

  Scenario: an error is returned if resolver url is not a valid url
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_url": "foo"
    }
    """
    When the message is validated
    Then an error of "invalid resolver url: invalid request"

  Scenario: no error is returned if manager and resolver url are valid
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_url": "https://foo.bar"
    }
    """
    When the message is validated
    Then an error of ""

  # Note: see ./types_content_hash.feature for content hash validation
