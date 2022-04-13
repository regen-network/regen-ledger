Feature: MsgAnchor

  Scenario: an error is returned if sender is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if sender is not a valid address
    Given the message
    """
    {
      "sender": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if content hash is empty
    Given the message
    """
    {
      "sender": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then expect the error "content hash cannot be empty: invalid request"

  Scenario: no error is returned if sender and content hash are valid
    Given the message
    """
    {
      "sender": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "content_hash": {
        "raw": {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1
        }
      }
    }
    """
    When the message is validated
    Then expect the error ""

  # Note: see ./types_content_hash.feature for content hash validation
