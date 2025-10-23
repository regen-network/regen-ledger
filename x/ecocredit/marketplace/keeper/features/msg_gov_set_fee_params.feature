Feature: Msg/SetFeeParams
  Rule: gov authority must be authorized
    Background:
      Given fee params
      """
      {}
      """

    Scenario: gov authority is not authorized
      Given authority is set to "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      When fee params are set
      Then expect error contains "unauthorized"

    Scenario: gov authority is authorized
      Given authority is set to the keeper authority
      When fee params are set
      Then expect no error

  Rule: Message Validations

    Scenario Outline: validate fee params
      Given authority "<authority>"
      And fee params `<fee_params>`
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | authority                                    | fee_params                      | error                |
        | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw |                                 | fees cannot be nil   |
        | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {}                              |                      |
        | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"-0.1"} | non-negative decimal |
        | regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw | {"buyer_percentage_fee":"0.1"}  |                      |

  Rule: fee params get saved
    Scenario: non-empty fee params
      Given authority is set to the keeper authority
      And fee params
        """
        {
          "buyer_percentage_fee": "0.01",
          "seller_percentage_fee": "0.01"
        }
        """
      When fee params are set
      Then expect no error
      And expect fee params
      """
        {
          "buyer_percentage_fee": "0.01",
          "seller_percentage_fee": "0.01"
        }
      """

    Scenario: empty fee params
      Given authority is set to the keeper authority
      And fee params
        """
        {}
        """
      When fee params are set
      Then expect no error
      And expect fee params
      """
        {}
      """
