Feature: MsgGovSetFeeParams

  Scenario Outline: validate fee params
    Given authority "<authority>"
    And fee params `<fee_params>`
    When the message is validated
    Then expect error <error>

    Examples:
      | authority                                    | fee_params                      | error |
      |                                              | {}                              | true  |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {}                              | false |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"-0.1"} | true  |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"0.1"}  | false |

