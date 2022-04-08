Feature: Define Resolver

  Scenario: resolver is defined and a resolver info entry is created
    Given a valid resolver url
    When a user attempts to define a resolver
    Then the resolver is defined
    And a resolver info entry is created and the manager is equal to the user address

  Scenario: resolver is not defined when a resolver with the same url already exists
    Given a valid resolver url
    And a resolver entry with the same url already exists
    When a user attempts to define a resolver
    Then the resolver is not defined
