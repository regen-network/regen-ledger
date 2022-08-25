Feature: DataAnchor

  Scenario: a valid data anchor
    Given the data anchor
    """
    {
      "id": "cmVnZW4=",
      "timestamp": "2020-01-01T00:00:00Z"
    }
    """
    When the data anchor is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the data anchor
    """
    {}
    """
    When the data anchor is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if timestamp is empty
    Given the data anchor
    """
    {
      "id": "cmVnZW4="
    }
    """
    When the data anchor is validated
    Then expect the error "timestamp cannot be empty: parse error"
