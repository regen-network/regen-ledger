Feature: Attest

  Background: the message has been validated
    Given alice is the attestor
    And the content hash
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1
      }
    }
    """

  Scenario: data is attested to when the data has been anchored
    Given alice has anchored the data at block time "2020-01-01"
    When alice attempts to attest to the data at block time "2020-01-02"
    And the data attestor entry exists and the timestamp is equal to "2020-01-02"

  Scenario: data is anchored when the data has not been anchored
    When alice attempts to attest to the data at block time "2020-01-01"
    And the data anchor entry exists and the timestamp is equal to "2020-01-01"

  Scenario: data is attested to when the data has not been anchored
    When alice attempts to attest to the data at block time "2020-01-01"
    And the data attestor entry exists and the timestamp is equal to "2020-01-01"

  Scenario: data attestor entry is not updated when the same address attests to the same data
    And alice has attested to the data at block time "2020-01-01"
    When alice attempts to attest to the data at block time "2020-01-02"
    And the data attestor entry exists and the timestamp is equal to "2020-01-01"

  # Note: see ../features/types_content_hash.feature for content hash validation
