Feature: Attest

  Scenario: data is attested to
    Given a graph data content hash
    And the data has been anchored
    When a user attempts to attest to the data
    Then the data is attested to

  Scenario: data is attested to and anchored
    Given a graph data content hash
    And the data has not been anchored
    When a user attempts to attest to the data
    Then the data is attested to
