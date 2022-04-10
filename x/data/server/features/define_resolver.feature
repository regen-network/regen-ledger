Feature: Define Resolver

  Scenario: a resolver is defined when a resolver is unique
    Given a valid resolver url
    When alice attempts to define the resolver
    Then no error is returned
    And the resolver info entry exists and alice is the manager

  Scenario: an error is returned when a resolver with the same url has been defined
    Given a valid resolver url
    And alice has defined the resolver
    When alice attempts to define the resolver
    Then an error is returned
