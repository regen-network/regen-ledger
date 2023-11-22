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

  Scenario: an error is returned if id is empty
    Given the data resolver
    """
    {}
    """
    When the data resolver is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if resolver id is empty
    Given the data resolver
    """
    {
      "id": "cmVnZW4="
    }
    """
    When the data resolver is validated
    Then expect the error "resolver id cannot be empty: parse error"
