Feature: BasketFee

  Scenario: a valid basket fee
    Given the basket fee
    """
    {
      "fee": {
        "denom": "uregen",
        "amount": "20000000"
      }
    }
    """
    When the basket fee is validated
    Then expect no error

  Scenario: an error is returned if basket fee denom is empty
    Given the basket fee
    """
    {
      "fee": {}
    }
    """
    When the basket fee is validated
    Then expect the error "fee: denom cannot be empty: parse error"

  Scenario: an error is returned if basket fee amount is empty
    Given the basket fee
    """
    {
      "fee": {
        "denom": "uregen"
      }
    }
    """
    When the basket fee is validated
    Then expect the error "fee: amount cannot be empty: parse error"

  Scenario: an error is returned if basket fee denom is not formatted
    Given the basket fee
    """
    {
      "fee": {
        "denom": "1",
        "amount": "20000000"
      }
    }
    """
    When the basket fee is validated
    Then expect the error "fee: invalid denom: 1: parse error"

  Scenario: an error is returned if basket fee amount is negative
    Given the basket fee
    """
    {
      "fee": {
        "denom": "uregen",
        "amount": "-20000000"
      }
    }
    """
    When the basket fee is validated
    Then expect the error "fee: negative coin amount: -20000000: parse error"
