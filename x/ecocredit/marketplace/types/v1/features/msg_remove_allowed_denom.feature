Feature: MsgRemoveAllowedDenom

  Scenario: a valid message
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "denom": "uregen"
    }
    """
    When the message is validated
    Then expect no error

 
  Scenario: an error is returned if denom is empty
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "denom cannot be empty: invalid request"

  Scenario: an error is returned if denom is not valid denom
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "denom": "1"
    }
    """
    When the message is validated
    Then expect the error "denom: invalid denom: 1: invalid request"

