Feature: MsgAddAllowedDenom

  Scenario: a valid message
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "bank_denom": "uregen",
      "display_denom": "REGEN",
      "exponent": 6
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if authority is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "invalid authority address: empty address string is not allowed"

  Scenario: an error is returned if authority is not a valid bech32 address
    Given the message
    """
    {
      "authority": "foo"
    }
    """
    When the message is validated
    Then expect the error "invalid authority address: decoding bech32 failed: invalid bech32 string length 3"

  Scenario: an error is returned if bank denom is empty
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "bank denom cannot be empty: parse error: invalid request"
  
  Scenario: an error is returned if bank denom is not valid denom
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "bank_denom": "1"
    }
    """
    When the message is validated
    Then expect the error "bank denom: invalid denom: 1: parse error: invalid request"

  Scenario: an error is returned if bank display denom is empty
    Given the message
    """
    {
      "authority": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "bank_denom": "uregen"
    }
    """
    When the message is validated
    Then expect the error "display denom cannot be empty: parse error: invalid request"
