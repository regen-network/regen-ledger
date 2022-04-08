Feature: MsgAnchor

  Scenario Outline: validate sender
    Given a sender of "<sender>"
    And a valid content hash
    When the message is validated
    Then an error of "<error>"

    Examples:
    | sender                                        | error                                                                   |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario Outline: validate raw content hash
    Given a valid sender address
    And a raw content hash of bytes length "<bytes>" and digest algorithm "<digest>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | bytes | digest | error                                                                       |
    | 31    | 1      | expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: invalid request |
    | 32    | 1      |                                                                             |
