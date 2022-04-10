Feature: MsgDefineResolver

  Scenario Outline: manager must be valid
    Given a manager of "<manager>"
    And a valid resolver url
    When the message is validated
    Then an error of "<error>"

    Examples:
    | manager                                       | error                                                                   |
    |                                               | empty address string is not allowed: invalid address                    |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario Outline: resolver url must be valid
    Given a valid manager
    And a resolver url of "<url>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | url             | error                                 |
    | foo             | invalid resolver url: invalid request |
    | https://foo.bar |                                       |
