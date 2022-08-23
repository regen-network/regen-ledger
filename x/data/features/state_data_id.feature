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
