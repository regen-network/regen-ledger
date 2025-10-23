Feature: Msg/SendFromFeePool

  Background:
    Given recipient "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"

  Rule: Message Validations
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

  Rule: gov authority must be authorized
    Scenario: gov authority is not authorized
      Given authority is set to "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      When funds are sent
      Then expect error contains "unauthorized"

    Scenario: gov authority is authorized
      Given authority is set to the keeper authority
      When funds are sent
      Then expect no error

  Rule: the fee pool must have enough funds to cover the fee
    Background:
      Given authority is set to the keeper authority

    Scenario: fee pool is underfunded
      Given fee pool balance "100foo"
      And send amount "200foo"
      When funds are sent
      Then expect error contains "insufficient funds"

    Scenario: fee pool is well funded
      Given fee pool balance "1000foo"
      And send amount "200foo"
      When funds are sent
      Then expect no error

  Rule: funds are transferred from the fee pool to the recipient
    Scenario: funds are transferred
      Given authority is set to the keeper authority
      Given fee pool balance "1000foo"
      And send amount "200foo"
      When funds are sent
      Then expect no error
      And fee pool balance "800foo"
      And recipient balance "200foo"
