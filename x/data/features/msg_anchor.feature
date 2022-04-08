Feature: MsgAnchor

  Scenario Outline: basic message validation
    Given a sender of "<sender>"
    And a content hash of "<hash>"
    When the message is validated
    Then an error of "<error>"

    Examples:
    | sender | hash | error  |
    | foo    | bar  | decoding bech32 failed: invalid bech32 string length 3: invalid address |