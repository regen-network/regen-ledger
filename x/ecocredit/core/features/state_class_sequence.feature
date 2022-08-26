Feature: ClassSequence

  Scenario: a valid class sequence
    Given the class sequence
    """
    {
      "credit_type_abbrev": "C",
      "next_sequence": 1
    }
    """
    When the class sequence is validated
    Then expect no error

  Scenario: an error is returned if credit type is empty
    Given the class sequence
    """
    {}
    """
    When the class sequence is validated
    Then expect the error "credit type abbrev: empty string is not allowed: parse error"

  Scenario: an error is returned if credit type is not formatted
    Given the class sequence
    """
    {
      "credit_type_abbrev": "1"
    }
    """
    When the class sequence is validated
    Then expect the error "credit type abbrev: must be 1-3 uppercase alphabetic characters: parse error"

  Scenario: an error is returned if next sequence is empty
    Given the class sequence
    """
    {
      "credit_type_abbrev": "C"
    }
    """
    When the class sequence is validated
    Then expect the error "next sequence cannot be zero: parse error"
