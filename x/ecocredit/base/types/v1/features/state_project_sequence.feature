Feature: ProjectSequence

  Scenario: a valid project sequence
    Given the project sequence
    """
    {
      "class_key": 1,
      "next_sequence": 1
    }
    """
    When the project sequence is validated
    Then expect no error

  Scenario: an error is returned if class key is empty
    Given the project sequence
    """
    {}
    """
    When the project sequence is validated
    Then expect the error "class key cannot be zero: parse error"

  Scenario: an error is returned if next sequence is empty
    Given the project sequence
    """
    {
      "class_key": 1
    }
    """
    When the project sequence is validated
    Then expect the error "next sequence cannot be zero: parse error"
