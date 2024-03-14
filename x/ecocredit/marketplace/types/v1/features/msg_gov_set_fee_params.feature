Feature: MsgGovSetFeeParams

  Scenario Outline: validate fee params
    Given authority "<authority>"
    And fee params `<fee_params>`
    When the message is validated
    Then expect error contains "<error>"

    Examples:
      | authority                                    | fee_params                      | error                |
      |                                              | {}                              | address              |
      | foobar                                       | {}                              | bech32               |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw |                                 | fees cannot be nil   |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {}                              |                      |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"-0.1"} | non-negative decimal |
      | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"0.1"}  |                      |

