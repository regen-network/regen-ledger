Feature: MsgAttest

  Scenario Outline: validate message
    Given an attestor of "<attestor>"
    And hashes of "<hashes>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | attestor                                      | hashes                                                                                                             | error                                                                   |
    |                                               | [{"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1, "canonicalization_algorithm": 1}] | empty address string is not allowed: invalid address                    |
    | foo                                           | [{"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1, "canonicalization_algorithm": 1}] | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                                                                    | hashes cannot be empty: invalid request                                 |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | [{"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1, "canonicalization_algorithm": 1}] |                                                                         |

  # Note: see ./types_content_hash.feature for content hash validation
