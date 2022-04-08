Feature: Register Resolver

  Scenario: data is registered to the resolver when provided a valid content hash that has been anchored
    Given a valid content hash that has been anchored
    And a resolver has been defined by the manager
    When the manager attempts to register data to the resolver
    Then the data is registered to the resolver
#    And a data resolver entry is created

  Scenario: data is registered to the resolver when provided a valid content hash that has not been anchored
    Given a valid content hash that has not been anchored
    And a resolver has been defined by the manager
    When the manager attempts to register data to the resolver
    Then the data is registered to the resolver
#    And a data resolver entry is created

  Scenario: data is not registered to the resolver when another user attempts to register data
    Given a valid content hash that has been anchored
    And a resolver has been defined by the manager
    When another user attempts to register data to the resolver
    Then the data is not registered to the resolver
