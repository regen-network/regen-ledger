Feature: Anchor

  Scenario: data is anchored when the data has not been anchored
    Given a valid content hash
    When alice attempts to anchor the data at block time "2020-01-01"
    Then no error is returned
    And the data id entry exists
    And the data anchor entry exists and the timestamp is equal to "2020-01-01"

  Scenario: data is not updated when the data has already been anchored
    Given a valid content hash
    And alice has anchored the data at block time "2020-01-01"
    When alice attempts to anchor the data at block time "2020-01-02"
    Then no error is returned
    And the data id entry exists
    Then the data anchor entry exists and the timestamp is equal to "2020-01-01"
