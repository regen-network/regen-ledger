Feature: MsgRegisterResolver

  Scenario Outline: manager must be valid
    Given a manager of "<manager>"
    And a valid resolver id
    And a valid list of data
    When the message is validated
    Then an error of "<error>"

    Examples:
    | manager                                       | error                                                                   |
    |                                               | empty address string is not allowed: invalid address                    |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario Outline: resolver id must be valid
    Given a valid manager
    And a resolver id of "<id>"
    And a valid list of data
    When the message is validated
    Then an error of "<error>"

    Examples:
    | id | error                                |
    | 0  | invalid resolver id: invalid request |
    | 1  |                                      |

  Scenario: data cannot be empty
    Given a valid manager
    And a valid resolver id
    And an empty list of data
    When the message is validated
    Then an error of "data cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
