Feature: MsgRegisterResolver

  Scenario Outline: validate message
    Given a manager of "<manager>"
    And a resolver id of "<resolver_id>"
    And data of "<data>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | manager                                       | resolver_id | data                                                                                       | error                                                                   |
    |                                               | 1           | [{"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}}] | empty address string is not allowed: invalid address                    |
    | foo                                           | 1           | [{"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}}] | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | 0           | [{"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}}] | invalid resolver id: invalid request                                    |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | 1           |                                                                                            | data cannot be empty: invalid request                                   |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | 1           | [{"raw": {"hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "digest_algorithm": 1}}] |                                                                         |

  # Note: see ./types_content_hash.feature for content hash validation
