Feature: MsgRegisterResolver

  Scenario: a valid message
    Given the message
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_id": 1,
      "content_hashes": [
        {
          "raw": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if manager is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if manager is not a bech32 address
    Given the message
    """
    {
      "manager": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if resolver id is empty
    Given the message
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then expect the error "resolver id cannot be empty: invalid request"

  Scenario: an error is returned if content hashes is empty
    Given the message
    """
    {
      "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "resolver_id": 1
    }
    """
    When the message is validated
    Then expect the error "content hashes cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
