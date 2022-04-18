Feature: QueryAttestors

  Background:
    Given the content entry
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """
    And the attestor entry
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "attestor": "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """

  Scenario: valid query response
    Given the query by IRI request
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When the query is executed
    Then the query by IRI response
    """
    {
      "attestors": ["cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"]
    }
    """

  Scenario: attestor entry is not found if it does not exist
    Given the query by IRI request
    """
    {
      "iri": "foo"
    }
    """
    When the query is executed
    Then expect the error "not found"
