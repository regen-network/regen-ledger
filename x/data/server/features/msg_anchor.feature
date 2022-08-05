Feature: Msg/Anchor

  Background:
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1
      }
    }
    """

  Rule: the data is anchored if the content hash is unique

    Scenario: the data has not been anchored
      When alice attempts to anchor the data at block time "2020-01-01"
      Then the anchor entry exists with timestamp "2020-01-01"

    Scenario: the data has already been anchored by the same address
      Given alice has anchored the data at block time "2020-01-01"
      When alice attempts to anchor the data at block time "2020-01-02"
      Then the anchor entry exists with timestamp "2020-01-01"

    Scenario: the data has already been anchored by a different address
      Given alice has anchored the data at block time "2020-01-01"
      When bob attempts to anchor the data at block time "2020-01-02"
      Then the anchor entry exists with timestamp "2020-01-01"
