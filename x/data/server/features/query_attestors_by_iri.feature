Feature: QueryAttestorsByIri

  Attestors can be queried by IRI:
  - when provided an IRI
  - when the content entry exists
  - whether or not attestor entries exist

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
    And the attestor entry
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
      "attestor": "cosmos1rxpx674qe4j8m4hy8g64xz4rpkj5k9zx2amw8k",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """
    And the content entry
    """
    {
      "iri": "regen:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf",
      "timestamp": "2020-01-01T00:00:00.000000000Z"
    }
    """

  Scenario: Attestor account addresses are returned if attestor entries exist
    Given the request
    """
    {
      "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When the request is executed
    Then the response
    """
    {
      "attestors": [
        "cosmos1rxpx674qe4j8m4hy8g64xz4rpkj5k9zx2amw8k",
        "cosmos1wrsm9y76utae3vl6xhhm2sk2hp5pjdhjt3yx7n"
      ]
    }
    """

  Scenario: Attestor account addresses are not returned if attestor entries do not exist
    Given the request
    """
    {
      "iri": "regen:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
    }
    """
    When the request is executed
    Then the response
    """
    {
      "attestors": null
    }
    """

  Scenario: An error is returned if the IRI is empty
    Given the request
    """
    {}
    """
    When the request is executed
    Then expect the error "IRI cannot be empty: invalid request"

  Scenario: An error is returned if the content entry does not exist
    Given the request
    """
    {
      "iri": "foo"
    }
    """
    When the request is executed
    Then expect the error "not found"
