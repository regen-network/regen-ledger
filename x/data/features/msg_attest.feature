Feature: MsgAttest

  Scenario: an error is returned if attestor is empty
    Given a message of
    """
    {}
    """
    When the message is validated
    Then an error of "empty address string is not allowed: invalid address"

  Scenario: an error is returned if attestor is not a valid address
    Given a message of
    """
    {
      "attestor": "foo"
    }
    """
    When the message is validated
    Then an error of "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if hash is empty
    Given a message of
    """
    {
      "attestor": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then an error of "hashes cannot be empty: invalid request"

  Scenario: no error is returned if attestor and hash are valid
    Given a message of
    """
    {
      "attestor": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "hashes": [
        {
          "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
          "digest_algorithm": 1,
          "canonicalization_algorithm": 1
        }
      ]
    }
    """
    When the message is validated
    Then an error of ""

  # Note: see ./types_content_hash.feature for content hash validation
