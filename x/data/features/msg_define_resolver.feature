Feature: MsgDefineResolver

  Scenario Outline: validate manager address
    Given a manager of "<manager>"
    And a valid resolver url
    When the message is validated
    Then an error of "<error>"

    Examples:
    | manager                                        | error                                                                   |
    | foo                                           | decoding bech32 failed: invalid bech32 string length 3: invalid address |
    | cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27 |                                                                         |

  Scenario Outline: validate resolver url
    Given a resolver url of "<url>"
    And a valid manager address
    When the message is validated
    Then an error of "<error>"

    Examples:
    | url             | error                                 |
    | foo             | invalid resolver url: invalid request |
    | https://foo.bar |                                       |
