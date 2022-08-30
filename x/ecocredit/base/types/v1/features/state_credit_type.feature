Feature: CreditType

  Scenario: a valid credit type
    Given the credit type
    """
    {
      "abbreviation": "C",
      "name": "carbon",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    }
    """
    When the credit type is validated
    Then expect no error

  Scenario: an error is returned if abbreviation is empty
    Given the credit type
    """
    {}
    """
    When the credit type is validated
    Then expect the error "abbreviation: empty string is not allowed: parse error"

  Scenario: an error is returned if abbreviation is not formatted
    Given the credit type
    """
    {
      "abbreviation": "1"
    }
    """
    When the credit type is validated
    Then expect the error "abbreviation: must be 1-3 uppercase alphabetic characters: parse error"

  Scenario: an error is returned if name is empty
    Given the credit type
    """
    {
      "abbreviation": "C"
    }
    """
    When the credit type is validated
    Then expect the error "name cannot be empty: parse error"

  Scenario: an error is returned if name exceeds 75 characters
    Given the credit type
    """
    {
      "abbreviation": "C"
    }
    """
    And name with length "76"
    When the credit type is validated
    Then expect the error "credit type name cannot exceed 75 characters: parse error"

  Scenario: an error is returned if unit is empty
    Given the credit type
    """
    {
      "abbreviation": "C",
      "name": "carbon"
    }
    """
    When the credit type is validated
    Then expect the error "unit cannot be empty: parse error"

  Scenario: an error is returned if precision is not 6
    Given the credit type
    """
    {
      "abbreviation": "C",
      "name": "carbon",
      "unit": "metric ton CO2 equivalent",
      "precision": 1
    }
    """
    When the credit type is validated
    Then expect the error "precision is currently locked to 6: parse error"
