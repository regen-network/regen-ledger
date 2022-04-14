Feature: Register Resolver

  Background: the message has been validated
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1
      }
    }
    """

  Rule: data is anchored if not already anchored

    Scenario: the data has been anchored
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver at block time "2020-01-01"
      Then the anchor entry exists with timestamp "2020-01-01"

    Scenario: the data has not been anchored
      Given alice has anchored the data at block time "2020-01-01"
      And alice has defined a resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver
      Then the anchor entry exists with timestamp "2020-01-01"

  Rule: data is registered if not already anchored

    Scenario: the data has not been anchored
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver
      Then the data resolver exists

    Scenario: the data has been anchored
      Given alice has anchored the data at block time "2020-01-01"
      And alice has defined a resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver
      And the data resolver exists

  Rule: data is registered if the resolver exists

    Scenario: an error is returned when the resolver does not exist
      When alice attempts to register the data to a resolver with id "1"
      Then expect the error "resolver with id 1 does not exist: not found"

  Rule: data is registered if the resolver exists

  Scenario: an error is returned when a user that is not the manager attempts to register data to the resolver
    And alice has defined a resolver with url "https://foo.bar"
    When bob attempts to register data to the resolver
    Then expect the error "unauthorized resolver manager"
