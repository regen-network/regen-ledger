Feature: Resolver

  Scenario: a valid resolver
    Given the resolver
    """
    {
      "id": 1,
      "url": "https://regen.network",
      "manager": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the resolver is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the resolver
    """
    {}
    """
    When the resolver is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if url is empty
    Given the resolver
    """
    {
      "id": 1
    }
    """
    When the resolver is validated
    Then expect the error "url cannot be empty: parse error"

  Scenario: an error is returned if url is not formatted
    Given the resolver
    """
    {
      "id": 1,
      "url": "foo"
    }
    """
    When the resolver is validated
    Then expect the error "url: invalid url format: parse error"

  Scenario: an error is returned if manager is empty
    Given the resolver
    """
    {
      "id": 1,
      "url": "https://regen.network"
    }
    """
    When the resolver is validated
    Then expect the error "manager: empty address string is not allowed: parse error"
