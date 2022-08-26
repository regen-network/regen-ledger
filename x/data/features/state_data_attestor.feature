Feature: DataAttestor

  Scenario: a valid data attestor
    Given the data attestor
    """
    {
      "id": "cmVnZW4=",
      "attestor": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "timestamp": "2020-01-01T00:00:00Z"
    }
    """
    When the data attestor is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the data attestor
    """
    {}
    """
    When the data attestor is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if attestor is empty
    Given the data attestor
    """
    {
      "id": "cmVnZW4="
    }
    """
    When the data attestor is validated
    Then expect the error "attestor: empty address string is not allowed: parse error"

  Scenario: an error is returned if timestamp is empty
    Given the data attestor
    """
    {
      "id": "cmVnZW4=",
      "attestor": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the data attestor is validated
    Then expect the error "timestamp cannot be empty: parse error"
