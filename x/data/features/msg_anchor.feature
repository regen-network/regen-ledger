Feature: MsgAnchor

  Scenario: a valid message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "content_hash": {
        "raw": {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1
        }
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if sender is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if sender is not a bech32 address
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
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "content hash cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
