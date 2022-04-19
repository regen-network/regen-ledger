Feature: QueryAttestors

  Background:
    Given the content entry
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """

  Rule: one attestor is returned if one attestor entry exists

    Background:
      Given the attestor entry
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n",
        "timestamp": "2020-01-01T00:00:00.000000000Z"
      }
      """

    Scenario: a valid query by IRI when one attestor entry exists
      Given the query by IRI request
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When the query by IRI is executed
      Then the query by IRI response
      """
      {
        "attestors": ["cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"]
      }
      """

    Scenario: a valid query by hash when one attestor entry exists
      Given the query by hash request
      """
      {
        "content_hash": {
          "graph": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1,
            "canonicalization_algorithm": 1
          }
        }
      }
      """
      When the query by hash is executed
      Then the query by hash response
      """
      {
        "attestors": ["cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"]
      }
      """

  Rule: multiple attestors are returned if multiple attestor entries exists

    Background:
      Given the attestor entry
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n",
        "timestamp": "2020-01-01T00:00:00.000000000Z"
      }
      """
      And the attestor entry
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "cosmos1rxpx674qe4j8m4hy8g64xz4rpkj5k9zx2amw8k",
        "timestamp": "2020-01-01T00:00:00.000000000Z"
      }
      """

    Scenario: a valid query by IRI when multiple attestor entries exist
      Given the query by IRI request
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When the query by IRI is executed
      Then the query by IRI response
      """
      {
        "attestors": [
          "cosmos1rxpx674qe4j8m4hy8g64xz4rpkj5k9zx2amw8k",
          "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"

        ]
      }
      """

    Scenario: a valid query by hash when multiple attestor entries exist
      Given the query by hash request
      """
      {
        "content_hash": {
          "graph": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1,
            "canonicalization_algorithm": 1
          }
        }
      }
      """
      When the query by hash is executed
      Then the query by hash response
      """
      {
        "attestors": [
          "cosmos1rxpx674qe4j8m4hy8g64xz4rpkj5k9zx2amw8k",
          "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"

        ]
      }
      """

  Rule: an error is returned if the request is empty

    Scenario: an error is returned if the IRI is empty
      Given the query by IRI request
      """
      {}
      """
      When the query by IRI is executed
      Then expect the error "IRI cannot be empty: invalid request"

    Scenario: an error is returned if the content hash is empty
      Given the query by hash request
      """
      {}
      """
      When the query by hash is executed
      Then expect the error "content hash cannot be empty: invalid request"

  Rule: an error is returned if the content entry does not exist

    Scenario: an error is returned if the content entry does not exist
      Given the query by IRI request
      """
      {
        "iri": "foo"
      }
      """
      When the query by IRI is executed
      Then expect the error "not found"

    Scenario: an error is returned if the content entry does not exist
      Given the query by hash request
      """
      {
        "content_hash": {
          "graph": {
            "hash": "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB=",
            "digest_algorithm": 1,
            "canonicalization_algorithm": 1
          }
        }
      }
      """
      When the query by hash is executed
      Then expect the error "not found"

  Rule: attestors are empty if there are no attestor entries

    Scenario: the attestors are empty if there are no attestor entries
      Given the query by IRI request
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When the query by IRI is executed
      Then the query by IRI response
      """
      {
        "attestors": null
      }
      """

    Scenario: the attestors are empty if there are no attestor entries
      Given the query by hash request
      """
      {
        "content_hash": {
          "graph": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1,
            "canonicalization_algorithm": 1
          }
        }
      }
      """
      When the query by hash is executed
      Then the query by hash response
      """
      {
        "attestors": null
      }
      """
