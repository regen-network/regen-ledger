Feature: MsgBurnRegen

  Scenario: a valid message with no reason
    Given the message
    """
    {
      "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "amount":"1000000000"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with a reason
    Given the message
    """
    {
      "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "amount":"1000000000",
      "reason":"for selling credits"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: invalid burner
    Given the message
    """
    {
      "burner": "foobar",
      "amount":"1000000000"
    }
    """
    When the message is validated
    Then expect error contains "invalid bech32"

  Scenario: invalid amount
    Given the message
    """
    {
      "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "amount":"foo"
    }
    """
    When the message is validated
    Then expect error contains "invalid amount"

  Scenario: negative amount
    Given the message
    """
      {
        "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "amount":"-1000000000"
      }
    """
    When the message is validated
    Then expect error contains "amount must be positive"

  Scenario: reason too long
    Given the message
    """
      {
        "burner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "amount":"1000000000",
        "reason":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor."
      }
    """
    When the message is validated
    Then expect error contains "at most 256"
