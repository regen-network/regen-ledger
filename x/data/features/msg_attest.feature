Feature: MsgAttest

  Scenario Outline: attestor must be valid
    Given an attestor of "<attestor>"
    And a valid content hash
    When the message is validated
    Then an error of "<error>"

    Examples:
    | attestor                                      | error                                                                   |
    |                                               | empty address string is not allowed: invalid address                    |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario: hashes cannot be empty
    Given a valid attestor
    And an empty list of content hashes
    When the message is validated
    Then an error of "hashes cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
