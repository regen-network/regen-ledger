Feature: Create group
  Scenario: Create a group from a single key
    Given a public key address
    When a user creates a group with that address
    And a decision threshold of 1
    Then they should get a new group address back
