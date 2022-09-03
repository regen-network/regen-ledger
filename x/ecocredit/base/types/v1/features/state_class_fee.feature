Feature: ClassFee

  Scenario: a valid class fee
    Given the class fee
    """
    {
      "fee": {
        "denom": "uregen",
        "amount": "20000000"
      }
    }
    """
    When the class fee is validated
    Then expect no error

  Scenario: an error is returned if class fee denom is empty
    Given the class fee
    """
    {
      "fee": {}
    }
    """
    When the class fee is validated
    Then expect the error "fee: denom cannot be empty: parse error"

  Scenario: an error is returned if class fee amount is empty
    Given the class fee
    """
    {
      "fee": {
        "denom": "uregen"
      }
    }
    """
    When the class fee is validated
    Then expect the error "fee: amount cannot be empty: parse error"

  Scenario: an error is returned if class fee denom is not formatted
    Given the class fee
    """
    {
      "fee": {
        "denom": "1",
        "amount": "20000000"
      }
    }
    """
    When the class fee is validated
    Then expect the error "fee: invalid denom: 1: parse error"

  Scenario: an error is returned if class fee amount is negative
    Given the class fee
    """
    {
      "fee": {
        "denom": "uregen",
        "amount": "-20000000"
      }
    }
    """
    When the class fee is validated
    Then expect the error "fee: negative coin amount: -20000000: parse error"
