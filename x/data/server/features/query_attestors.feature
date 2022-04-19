Feature: QueryAttestors

  Background:
    Given the content entry
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """

  Rule: attestor entries can be queried by IRI

    Scenario: a valid query
      Given the attestor entry
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n",
        "timestamp": "2020-01-01T00:00:00.000000000Z"
      }
      """
      And the query by IRI request
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

    Scenario: a valid query with multiple attestors
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
      And the query by IRI request
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

    Scenario: an error is returned if the IRI is empty
      Given the query by IRI request
      """
      {
        "iri": ""
      }
      """
      When the query by IRI is executed
      Then expect the error "not found"

    Scenario: an error is returned if the data entry does not exist
      Given the query by IRI request
      """
      {
        "iri": "foo"
      }
      """
      When the query by IRI is executed
      Then expect the error "not found"

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

  Rule: attestor entries can be queried by content hash

    Scenario: a valid query
      Given the attestor entry
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n",
        "timestamp": "2020-01-01T00:00:00.000000000Z"
      }
      """
      And the query by hash request
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

    Scenario: a valid query with multiple attestors
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
      And the query by hash request
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

    Scenario: an error is returned if the data entry does not exist
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
