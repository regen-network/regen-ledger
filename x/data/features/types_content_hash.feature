Feature: Types

  Scenario: an error is returned if content hash is empty
    Given a content hash of
    """
    {}
    """
    When the content hash is validated
    Then an error of "content hash must be one of raw type or graph type: invalid request"

  Scenario: an error is returned if raw content hash is empty
    Given a content hash of
    """
    {
      "raw": {}
    }
    """
    When the content hash is validated
    Then an error of "hash cannot be empty: invalid request"

  Scenario: an error is returned if raw content hash digest algorithm is unspecified
    Given a content hash of
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
      }
    }
    """
    When the content hash is validated
    Then an error of "invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if raw content hash digest algorithm is unknown
    Given a content hash of
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then an error of "unknown data.DigestAlgorithm 2: invalid request"

  Scenario: an error is returned if raw content hash length does not match blake2b digest algorithm
    Given a content hash of
    """
    {
      "raw": {
        "hash": [0],
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then an error of "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 1: invalid request"

  Scenario: no error is returned if raw content hash media type is unspecified
    Given a content hash of
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then an error of ""

  Scenario: no error is returned if raw content hash media type is valid
    Given a content hash of
    """
    {
      "raw": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1,
        "media_type": 1
      }
    }
    """
    When the content hash is validated
    Then an error of ""

  Scenario: an error is returned if graph content hash is empty
    Given a content hash of
    """
    {
      "graph": {}
    }
    """
    When the content hash is validated
    Then an error of "hash cannot be empty: invalid request"

  Scenario: an error is returned if graph content hash digest algorithm is unspecified
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
      }
    }
    """
    When the content hash is validated
    Then an error of "invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if graph content hash digest algorithm is unknown
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then an error of "unknown data.DigestAlgorithm 2: invalid request"

  Scenario: an error is returned if graph content hash length does not match blake2b digest algorithm
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0],
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then an error of "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 1: invalid request"

  Scenario: an error is returned if graph content hash canonicalization algorithm is unspecified
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then an error of "invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if graph content hash canonicalization algorithm is unknown
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1,
        "canonicalization_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then an error of "unknown data.GraphCanonicalizationAlgorithm 2: invalid request"

  Scenario: no error is returned if graph content hash merkle tree is unspecified
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then an error of ""

  Scenario: an error is returned if graph content hash merkle tree is unknown
    Given a content hash of
    """
    {
      "graph": {
        "hash": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1,
        "merkle_tree": 1
      }
    }
    """
    When the content hash is validated
    Then an error of "unknown data.GraphMerkleTree 1: invalid request"
