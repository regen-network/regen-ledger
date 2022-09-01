Feature: BatchSequence

  Scenario: a valid batch sequence
    Given the batch sequence
    """
    {
      "project_key": 1,
      "next_sequence": 1
    }
    """
    When the batch sequence is validated
    Then expect no error

  Scenario: an error is returned if batch key is empty
    Given the batch sequence
    """
    {}
    """
    When the batch sequence is validated
    Then expect the error "project key cannot be zero: parse error"

  Scenario: an error is returned if next sequence is empty
    Given the batch sequence
    """
    {
      "project_key": 1
    }
    """
    When the batch sequence is validated
    Then expect the error "next sequence cannot be zero: parse error"
