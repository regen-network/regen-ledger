Feature: Msg/Attest

  Background:
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

  Rule: the data is anchored if not already anchored

    Scenario: the data has not been anchored
      When alice attempts to attest to the data at block time "2020-01-01"
      Then the anchor entry exists with timestamp "2020-01-01"

    Scenario: the data has already been anchored
      Given alice has anchored the data at block time "2020-01-01"
      When alice attempts to attest to the data at block time "2020-01-02"
      Then the anchor entry exists with timestamp "2020-01-01"

  Rule: the data is attested to if not already attested to by the same address

    Scenario: the data has not been attested to
      When alice attempts to attest to the data at block time "2020-01-01"
      Then the attestor entry for alice exists with timestamp "2020-01-01"

    Scenario: the data has already been attested to by the same address
      Given alice has attested to the data at block time "2020-01-01"
      When alice attempts to attest to the data at block time "2020-01-02"
      Then the attestor entry for alice exists with timestamp "2020-01-01"

    Scenario: the data has already been attested to by a different address
      Given alice has attested to the data at block time "2020-01-01"
      When bob attempts to attest to the data at block time "2020-01-02"
      Then the attestor entry for bob exists with timestamp "2020-01-02"

  Rule: event is emitted

    Background:
      Given alice's address "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"

    Scenario: EventAttest is emitted
      When alice attempts to attest to the data at block time "2020-01-01"
      Then the anchor entry exists with timestamp "2020-01-01"
      And event is emitted with properties
      """
      {
        "iri": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
        "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
      }
      """