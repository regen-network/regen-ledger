Feature: MsgGovSetFeeParams

  Scenario Outline: validate message
    Given authority "<authority>"
    * recipient "<recipient>"
    * amount "<amount>"
    When the message is validated
    Then expect error contains "<error>"

    Examples:
      | authority                                    | recipient                                    | amount | error         |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68 |        | coins         |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68 | 100foo |               |

