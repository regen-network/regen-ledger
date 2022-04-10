Feature: Types

  Scenario: content hash cannot be empty
    Given an empty content hash
    When the content hash is validated
    Then an error of "content hash must be one of raw type or graph type: invalid request"

  Scenario: raw content hash cannot be empty
    Given an empty raw content hash
    When the content hash is validated
    Then an error of "hash cannot be empty: invalid request"

  Scenario Outline: raw content hash must be valid
    Given a raw content hash of "<length>" "<digest>" "<media>"
    When the content hash is validated
    Then an error of "<error>"

    Examples:
    | length | digest | media | error                                                                       |
    | 16     | 1      | 0     | expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 16: invalid request |
    | 32     | 0      | 0     | invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request  |
    | 32     | 2      | 0     | unknown data.DigestAlgorithm 2: invalid request                             |
    | 32     | 1      | 0     |                                                                             |
    | 32     | 1      | 1     |                                                                             |
    | 32     | 1      | 2     |                                                                             |
    | 32     | 1      | 3     |                                                                             |
    | 32     | 1      | 4     |                                                                             |
    | 32     | 1      | 5     |                                                                             |
    | 32     | 1      | 6     | unknown data.RawMediaType 6: invalid request                                |

  Scenario: graph content hash cannot be empty
    Given an empty graph content hash
    When the content hash is validated
    Then an error of "hash cannot be empty: invalid request"

  Scenario Outline: graph content hash must be valid
    Given a graph content hash of "<length>" "<digest>" "<canon>" "<merkle>"
    When the content hash is validated
    Then an error of "<error>"

    Examples:
    | length | digest | canon | merkle | error                                                                                                     |
    | 16     | 1      | 1     | 0      | expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 16: invalid request                               |
    | 32     | 0      | 0     | 0      | invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request                                |
    | 32     | 2      | 0     | 0      | unknown data.DigestAlgorithm 2: invalid request                                                           |
    | 32     | 1      | 0     | 0      | invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request |
    | 32     | 1      | 1     | 1      | unknown data.GraphMerkleTree 1: invalid request                                                           |
    | 32     | 1      | 1     | 0      |                                                                                                           |
