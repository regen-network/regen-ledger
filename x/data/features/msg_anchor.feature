Feature: MsgAnchor

  Scenario Outline: validate message
    Given a sender of "<sender>"
    And a hash of "<hash>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | sender                                        | hash                                                                                     | error                                                                   |
    |                                               | {"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}} | empty address string is not allowed: invalid address                    |
    | foo                                           | {"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}} | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                                          | hash cannot be empty: invalid request                                   |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | {"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}} |                                                                         |

  # Note: see ./types_content_hash.feature for content hash validation
