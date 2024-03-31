Feature: MsgUpdateProjectFee

  Rule: authority is address and signer
    Scenario Outline: validate admin
      Given authority "<auth>"
      And fee "100regen"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | auth                                         | error             |
        |                                              | invalid authority |
        | 0x0                                          | invalid authority |
        | regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 |                   |

    Scenario: authority is signer
      Given authority "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      Then expect GetSigners returns "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"

  Rule: fee is a valid coin or nil
    Scenario: valid coin
      Given authority "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And fee "100regen"
      When the message is validated
      Then expect no error

    Scenario: fee is nil
      Given authority "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And nil fee
      When the message is validated
      Then expect no error

