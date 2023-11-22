Feature: DataId

  Scenario: a valid data id
    Given the data id
    """
    {
      "id": "cmVnZW4=",
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the data id is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the data id
    """
    {}
    """
    When the data id is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if iri is empty
    Given the data id
    """
    {
      "id": "cmVnZW4="
    }
    """
    When the data id is validated
    Then expect the error "failed to parse IRI: empty string is not allowed: invalid IRI: parse error"

  Scenario: an error is returned if iri is not formatted
    Given the data id
    """
    {
      "id": "cmVnZW4=",
      "iri": "foo"
    }
    """
    When the data id is validated
    Then expect the error "failed to parse IRI foo: regen: prefix required: invalid IRI: parse error"
