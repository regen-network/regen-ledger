Feature: Anchor

  Scenario: raw data is anchored
    Given a raw data content hash
    When a user attempts to anchor the data
    Then the data is anchored
    And the anchored data is equal to the data submitted

  Scenario: graph data is anchored
    Given a graph data content hash
    When a user attempts to anchor the data
    Then the data is anchored
    And the anchored data is equal to the data submitted
