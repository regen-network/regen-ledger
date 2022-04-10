Feature: Register Resolver

  Scenario: data is registered to the resolver when the data has been anchored
    Given a valid content hash
    And alice has anchored the data
    And alice has defined a resolver
    When alice attempts to register the data to the resolver
    Then an error of ""
    And the data resolver entry exists

  Scenario: data is anchored and registered to the resolver when the data has not been anchored
    Given a valid content hash
    And alice has defined a resolver
    When alice attempts to register the data to the resolver
    Then an error of ""
    And the data anchor entry exists
    And the data resolver entry exists

  Scenario: an error is returned when the provided resolver id does not exist
    Given a valid content hash
    When alice attempts to register the data to a resolver with id "1"
    Then an error of "resolver with id 1 does not exist: not found"

  Scenario: an error is returned when a user that is not the manager attempts to register data to the resolver
    Given a valid content hash
    And alice has defined a resolver
    When bob attempts to register data to the resolver
    Then an error of "unauthorized resolver manager"

  # Note: see ../features/types_content_hash.feature for content hash validation
