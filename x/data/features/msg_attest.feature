Feature: MsgAttest

  Scenario: a valid message
    Given the message
     """
     {
       "attestor": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
       "content_hashes": [
         {
           "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
           "digest_algorithm": 1,
           "canonicalization_algorithm": 1
         }
       ]
     }
     """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if attestor is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if attestor is not a bech32 address
    Given the message
    """
    {
      "attestor": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if content hashes is empty
    Given the message
    """
    {
      "attestor": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then expect the error "content hashes cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
