Feature: Define Resolver

  Scenario: resolver is defined when provided a valid resolver url
    Given a valid resolver url
    When a user attempts to define a resolver
    Then the resolver is defined
    And a resolver info entry is created and the manager is equal to the user address

  Scenario: resolver is not defined when provided an invalid resolver url
    Given an invalid resolver url
    When a user attempts to define a resolver
#    Then the resolver is not defined
