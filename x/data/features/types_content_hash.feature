Feature: Types

  Scenario: a valid raw content hash
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "media_type": 1
      }
    }
    """
    When the content hash is validated
    Then expect no error

  Scenario: a valid graph content hash
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1,
        "merkle_tree": 0
      }
    }
    """
    When the content hash is validated
    Then expect no error

  Scenario: an error is returned if content hash is empty
    Given the content hash
    """
    {}
    """
    When the content hash is validated
    Then expect the error "content hash must be one of raw type or graph type: invalid request"

  Scenario: an error is returned if content hash includes both raw type and graph type
    Given the content hash
    """
    {
      "raw": {},
      "graph": {}
    }
    """
    When the content hash is validated
    Then expect the error "content hash must be one of raw type or graph type: invalid request"

  Scenario: an error is returned if raw content hash is empty
    Given the content hash
    """
    {
      "raw": {}
    }
    """
    When the content hash is validated
    Then expect the error "hash cannot be empty: invalid request"

  Scenario: an error is returned if raw content hash digest algorithm is unspecified
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
      }
    }
    """
    When the content hash is validated
    Then expect the error "invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if raw content hash digest algorithm is unknown
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then expect the error "unknown data.DigestAlgorithm 2: invalid request"

  Scenario: an error is returned if raw content hash length does not match blake2b digest algorithm
    Given the content hash
    """
    {
      "raw": {
        "hash": "AA==",
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then expect the error "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 1: invalid request"

  Scenario: no error is returned if raw content hash media type is unspecified
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then expect no error

  Scenario: an error is returned if graph content hash is empty
    Given the content hash
    """
    {
      "graph": {}
    }
    """
    When the content hash is validated
    Then expect the error "hash cannot be empty: invalid request"

  Scenario: an error is returned if graph content hash digest algorithm is unspecified
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
      }
    }
    """
    When the content hash is validated
    Then expect the error "invalid data.DigestAlgorithm DIGEST_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if graph content hash digest algorithm is unknown
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then expect the error "unknown data.DigestAlgorithm 2: invalid request"

  Scenario: an error is returned if graph content hash length does not match blake2b digest algorithm
    Given the content hash
    """
    {
      "graph": {
        "hash": "AA==",
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then expect the error "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 1: invalid request"

  Scenario: an error is returned if graph content hash canonicalization algorithm is unspecified
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then expect the error "invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request"

  Scenario: an error is returned if graph content hash canonicalization algorithm is unknown
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "canonicalization_algorithm": 2
      }
    }
    """
    When the content hash is validated
    Then expect the error "unknown data.GraphCanonicalizationAlgorithm 2: invalid request"

  Scenario: no error is returned if graph content hash merkle tree is unspecified
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1
      }
    }
    """
    When the content hash is validated
    Then expect no error

  Scenario: an error is returned if graph content hash merkle tree is unknown
    Given the content hash
    """
    {
      "graph": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "canonicalization_algorithm": 1,
        "merkle_tree": 1
      }
    }
    """
    When the content hash is validated
    Then expect the error "unknown data.GraphMerkleTree 1: invalid request"
