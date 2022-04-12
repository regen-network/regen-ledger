Feature: Anchor

  Background: the message has been validated
    Given alice is the sender
    And the content hash
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1
      }
    }
    """

  Scenario: data is anchored when the data has not been anchored
    When alice attempts to anchor the data at block time "2020-01-01"
    Then the data anchor entry exists and the timestamp is equal to "2020-01-01"

  Scenario: data anchor entry is not updated when the data has already been anchored
    And alice has anchored the data at block time "2020-01-01"
    When alice attempts to anchor the data at block time "2020-01-02"
    Then the data anchor entry exists and the timestamp is equal to "2020-01-01"

  # Note: see ../features/types_content_hash.feature for content hash validation
