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

  Scenario: an error is returned if bank denom is not formatted
    Given the allowed denom
    """
    {
      "bank_denom": "1"
    }
    """
    When the allowed denom is validated
    Then expect the error "bank denom: invalid denom: 1: parse error"

  Scenario: an error is returned if display denom is empty
    Given the allowed denom
    """
    {
      "bank_denom": "uregen"
    }
    """
    When the allowed denom is validated
    Then expect the error "display denom cannot be empty: parse error"

  Scenario: an error is returned if display denom is not formatted
    Given the allowed denom
    """
    {
      "bank_denom": "uregen",
      "display_denom": "1"
    }
    """
    When the allowed denom is validated
    Then expect the error "display denom: invalid denom: 1: parse error"

  Scenario: an error is returned if exponent does not have a standard prefix
    Given the allowed denom
    """
    {
      "bank_denom": "uregen",
      "display_denom": "regen",
      "exponent": 4
    }
    """
    When the allowed denom is validated
    Then expect the error "exponent must be one of [0 1 2 3 6 9 12 15 18 21 24]: parse error"
