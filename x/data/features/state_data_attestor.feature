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
