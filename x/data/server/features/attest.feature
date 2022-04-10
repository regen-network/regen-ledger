Feature: Attest

  Scenario: data is attested to when the data has been anchored
    Given a valid content hash
    And alice has anchored the data at block time "2020-01-01"
    When alice attempts to attest to the data at block time "2020-01-02"
    Then an error of ""
    And the data id entry exists
    And the data anchor entry exists and the timestamp is equal to "2020-01-01"
    And the data attestor entry exists and the timestamp is equal to "2020-01-02"

  Scenario: data is attested to and anchored when the data has not been anchored
    Given a valid content hash
    When alice attempts to attest to the data at block time "2020-01-01"
    Then an error of ""
    And the data id entry exists
    And the data anchor entry exists and the timestamp is equal to "2020-01-01"
    And the data attestor entry exists and the timestamp is equal to "2020-01-01"

  Scenario: an attestor entry is not updated when the same address attests to the same data
    Given a valid content hash
    And alice has attested to the data at block time "2020-01-01"
    When alice attempts to attest to the data at block time "2020-01-02"
    Then an error of ""
    And the data id entry exists
    And the data anchor entry exists and the timestamp is equal to "2020-01-01"
    And the data attestor entry exists and the timestamp is equal to "2020-01-01"

  # Note: see ../features/types_content_hash.feature for content hash validation
