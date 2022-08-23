Feature: DataResolver

  Scenario: a valid data resolver
    Given the data resolver
    """
    {
      "id": "cmVnZW4=",
      "resolver_id": 1
    }
    """
    When the data resolver is validated
    Then expect no error
