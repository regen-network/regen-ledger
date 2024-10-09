Feature: Msg/UpdateProjectFee
  Rule: the authority can update the project fee
    Scenario Outline: <scenario>
      Given current fee "<cur_fee>"
      When "<authority>" updates the fee to "<new_fee>"
      Then expect error contains "<err>"
      And expect project fee is "<expected_fee>"
      Examples:
        | scenario       | authority | cur_fee | new_fee | expected_fee | err               |
        | gov creates    | gov       |         | 200     | 200          |                   |
        | gov updates    | gov       | 100     | 200     | 200          |                   |
        | gov sets zero  | gov       | 100     | 0       |              |                   |
        | gov sets empty | gov       | 100     |         |              |                   |
        | bob creates    | bob       |         | 200     |              | invalid authority |
        | bob updates    | bob       | 100     | 200     | 100          | invalid authority |
        | bob sets zero  | bob       | 100     | 0       | 100          | invalid authority |
