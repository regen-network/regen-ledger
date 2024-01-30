Feature: Types

  Scenario: a valid raw content hash
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "file_extension": "bin"
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
    Then expect an error

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
    Then expect an error

  Scenario: an error is returned if raw content file extension is empty
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
    Then expect an error

  Scenario: an error is returned if graph content hash is empty
    Given the content hash
    """
    {
      "graph": {}
    }
    """
    When the content hash is validated
    Then expect an error

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
    Then expect an error

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
    Then expect an error
