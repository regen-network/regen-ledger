Feature: MsgDefineResolver

  Scenario Outline: validate message
    Given a manager of "<manager>"
    And a resolver url of "<url>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | manager                                       | url             | error                                                                   |
    |                                               | https://foo.bar | empty address string is not allowed: invalid address                    |
    | foo                                           | https://foo.bar | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                 | invalid resolver url: invalid request                                   |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | foo             | invalid resolver url: invalid request                                   |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 | https://foo.bar |                                                                         |
