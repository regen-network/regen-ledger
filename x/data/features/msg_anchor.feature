Feature: MsgAnchor

  Scenario: an error is returned if sender is empty
    Given a message of
    """
    {}
    """
    When the message is validated
    Then an error of "empty address string is not allowed: invalid address"

  Scenario: an error is returned if sender is not a valid address
    Given a message of
    """
    {
      "sender": "foo"
    }
    """
    When the message is validated
    Then an error of "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if hash is empty
    Given a message of
    """
    {
      "sender": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then an error of "hash cannot be empty: invalid request"

  Scenario: no error is returned if sender and hash are valid
    Given a message of
    """
    {
      "sender": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "hash": {
        "raw": {
          "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
          "digest_algorithm": 1
        }
      }
    }
    """
    When the message is validated
    Then an error of ""

  # Note: see ./types_content_hash.feature for content hash validation
