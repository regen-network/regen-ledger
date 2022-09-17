Feature: MsgTake

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100",
      "retirement_jurisdiction": "US-WA",
      "retire_on_take": true
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without retirement jurisdiction
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if owner is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid request"

  Scenario: an error is returned if owner is not a bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid request"

  Scenario: an error is returned if basket denom is empty
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "basket denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if basket denom is not formatted
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "basket denom: expected format eco.<exponent-prefix><credit-type-abbrev>.<name>: parse error: invalid request"

  Scenario: an error is returned if amount is empty
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT"
    }
    """
    When the message is validated
    Then expect the error "amount cannot be empty: invalid request"

  Scenario: an error is returned if a amount is not an integer
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100.5"
    }
    """
    When the message is validated
    Then expect the error "100.5 is not a valid integer: invalid request"

  Scenario: an error is returned if a amount is less than zero
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "-100"
    }
    """
    When the message is validated
    Then expect the error "amount must be positive, got -100: invalid request"

  Scenario: an error is returned if retirement jurisdiction is empty and retire on take is true
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100",
      "retire_on_take": true
    }
    """
    When the message is validated
    Then expect the error "retirement jurisdiction cannot be empty if retire on take is true: invalid request"

  Scenario: an error is returned if retirement jurisdiction is not formatted and retire on take is true
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100",
      "retirement_jurisdiction": "foo",
      "retire_on_take": true
    }
    """
    When the message is validated
    Then expect the error "retirement jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

  Scenario: no error is returned if retirement jurisdiction is not formatted and retire on take is false
    Given the message
    """
    {
      "owner": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "basket_denom": "eco.uC.NCT",
      "amount": "100",
      "retirement_jurisdiction": "foo",
      "retire_on_take": false
    }
    """
    When the message is validated
    Then expect no error
