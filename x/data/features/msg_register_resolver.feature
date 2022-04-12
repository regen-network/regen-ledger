Feature: MsgRegisterResolver

  Scenario: an error is returned if manager is empty
    Given a message of
    """
    {}
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

  Scenario: an error is returned if resolver id is empty
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then an error of "resolver id cannot be empty: invalid request"

  Scenario: an error is returned if data is empty
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_id": 1
    }
    """
    When the message is validated
    Then an error of "data cannot be empty: invalid request"

  Scenario: no error is returned if manager and resolver id are valid
    Given a message of
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_id": 1,
      "data": [
        {
          "raw": {
            "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            "digest_algorithm": 1
          }
        }
      ]
    }
    """
    When the message is validated
    Then an error of ""

  # Note: see ./types_content_hash.feature for content hash validation
