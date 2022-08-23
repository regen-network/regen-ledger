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

  Scenario: an error is returned if credit type is empty
    Given the market
    """
    {
      "id": 1
    }
    """
    When the market is validated
    Then expect the error "credit type abbreviation cannot be empty: parse error"

  Scenario: an error is returned if credit type is not formatted
    Given the market
    """
    {
      "id": 1,
      "credit_type_abbrev": "1"
    }
    """
    When the market is validated
    Then expect the error "credit type abbreviation must be 1-3 uppercase latin letters: got 1: parse error"

  Scenario: an error is returned if bank denom is empty
    Given the market
    """
    {
      "id": 1,
      "credit_type_abbrev": "C"
    }
    """
    When the market is validated
    Then expect the error "bank denom cannot be empty: parse error"

  Scenario: an error is returned if bank denom is not formatted
    Given the market
    """
    {
      "id": 1,
      "credit_type_abbrev": "C",
      "bank_denom": "1"
    }
    """
    When the market is validated
    Then expect the error "bank denom: invalid denom: 1: parse error"

  Scenario: an error is returned if precision modifier is not zero
    Given the market
    """
    {
      "id": 1,
      "credit_type_abbrev": "C",
      "bank_denom": "uregen",
      "precision_modifier": 1
    }
    """
    When the market is validated
    Then expect the error "precision modifier must be zero: parse error"
