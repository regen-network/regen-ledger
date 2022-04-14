Feature: MsgAttest

  Rule: only a valid attestor is accepted

    Scenario: an error is returned if attestor is empty
      Given the message
      """
      {}
      """
      When the message is validated
      Then expect the error "empty address string is not allowed: invalid address"

    Scenario: an error is returned if attestor is not a valid address
      Given the message
      """
      {
        "attestor": "foo"
      }
      """
      When the message is validated
      Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Rule: only a valid content hash is accepted

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

  Rule: only a valid message is accepted

    Scenario: no error is returned if attestor and content hashes are valid
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

