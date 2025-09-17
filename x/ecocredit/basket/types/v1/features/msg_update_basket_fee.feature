Feature: MsgUpdateBasketFee

  Scenario: a valid message
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "fee": {
        "denom": "uregen",
        "amount": "1000"
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message removing basket fee
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect no error


  Scenario: an error is returned if basket fee denom is not formatted
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "fee": {
        "denom": "1",
        "amount": "1000"
      }
    }
    """
    When the message is validated
    Then expect the error "invalid denom: 1: invalid request"
