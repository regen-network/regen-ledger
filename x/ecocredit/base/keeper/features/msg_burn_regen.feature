Feature: Msg/Regen

  Scenario: burning regen
    Given the message
    """
    {
      "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "amount":"1000000000",
      "reason":"for selling credits"
    }
    """
    Then expect "1000000000uregen" are sent from "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6" to the ecocredit module
    * expect "1000000000uregen" are burned by the ecocredit module
    When it is executed
    And expect the event is emitted
    """
    {
      "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "amount":"1000000000",
      "reason":"for selling credits"
    }
    """
    And expect no error
