Feature: Register Resolver

  Background: the message has been validated
    Given alice is the manager
    And bob is not the manager
    And a content hash of
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1
      }
    }
    """

  Scenario: data is anchored when the data has not been anchored
    And alice has defined a resolver with url "https://foo.bar"
    When alice attempts to register the data to the resolver
    And the data anchor entry exists

  Scenario: data is registered to the resolver when the data has not been anchored
    And alice has defined a resolver with url "https://foo.bar"
    When alice attempts to register the data to the resolver
    And the data resolver entry exists

  Scenario: data is registered to the resolver when the data has been anchored
    Given alice has anchored the data
    And alice has defined a resolver with url "https://foo.bar"
    When alice attempts to register the data to the resolver
    And the data resolver entry exists

  Scenario: an error is returned when the resolver does not exist
    When alice attempts to register the data to a resolver with id "1"
    Then an error of "resolver with id 1 does not exist: not found"

  Scenario: an error is returned when a user that is not the manager attempts to register data to the resolver
    And alice has defined a resolver with url "https://foo.bar"
    When bob attempts to register data to the resolver
    Then an error of "unauthorized resolver manager"

  # Note: see ../features/types_content_hash.feature for content hash validation
