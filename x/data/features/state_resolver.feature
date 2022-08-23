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
