Feature: Attest

  Scenario: data is attested to when provided a valid content hash and the data has been anchored
    Given a valid content hash
    And the data has been anchored at block time "2020-01-01"
    When a user attempts to attest to the data at block time "2020-07-01"
    Then the data is attested to
    And a data id entry is created
    And a data anchor entry is created and the timestamp is equal to "2020-01-01"
    And a data attestor entry is created and the timestamp is equal to "2020-07-01"

  Scenario: data is attested to and anchored when provided a valid content hash and the data has not been anchored
    Given a valid content hash
    And the data has not been anchored
    When a user attempts to attest to the data at block time "2020-07-01"
    Then the data is attested to
    And a data id entry is created
    And a data anchor entry is created and the timestamp is equal to "2020-07-01"
    And a data attestor entry is created and the timestamp is equal to "2020-07-01"

  Scenario: data is not attested when provided an invalid content hash
    Given an invalid content hash
    And the data has not been anchored
    When a user attempts to attest to the data at block time "2020-07-01"
    Then the data is not attested to
