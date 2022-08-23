Feature: AllowedDenom

  Scenario: a valid allowed denom
    Given the allowed denom
    """
    {
      "bank_denom": "uregen",
      "display_denom": "regen",
      "exponent": 6
    }
    """
    When the allowed denom is validated
    Then expect no error

  Scenario: an error is returned if bank denom is empty
    Given the allowed denom
    """
    {}
    """
    When the allowed denom is validated
    Then expect the error "bank denom cannot be empty: parse error"
