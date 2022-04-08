Feature: Anchor

  Scenario: data is anchored when provided a valid content hash
    Given a valid content hash
    When a user attempts to anchor the data
    Then the data is anchored
    And a data id entry is created
    And a data anchor entry is created and the timestamp is equal to the block time

  Scenario: data is not anchored when provided an invalid content hash
    Given an invalid content hash
    When a user attempts to anchor the data
    Then the data is not anchored
