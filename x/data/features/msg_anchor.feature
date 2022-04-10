Feature: MsgAnchor

  Scenario Outline: sender must be valid
    Given a sender of "<sender>"
    And a valid content hash
    When the message is validated
    Then an error of "<error>"

    Examples:
    | sender                                        | error                                                                   |
    |                                               | empty address string is not allowed: invalid address                    |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario: hash cannot be empty
    Given a valid sender
    And an empty content hash
    When the message is validated
    Then an error of "hash cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
