Feature: Market

  Scenario: a valid market
    Given the market
    """
    {
      "id": 1,
      "credit_type_abbrev": "C",
      "bank_denom": "uregen",
      "precision_modifier": 0
    }
    """
    When the market is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the market
    """
    {}
    """
    When the market is validated
    Then expect the error "id cannot be zero: parse error"
