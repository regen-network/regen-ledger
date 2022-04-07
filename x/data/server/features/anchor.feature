Feature: Anchor

  Scenario: raw data is anchored
    Given a raw data content hash
    When a user attempts to anchor the data
    Then the data is anchored
    And the anchored data is equal to the data submitted

  Scenario: graph data is anchored
    Given a graph data content hash
    When a user attempts to anchor the data
    Then the data is anchored
    And the anchored data is equal to the data submitted

  Scenario Outline: anchoring data
    Given a user has content hash <hash>
    When the user attempts to anchor the data
    Then the data is anchored with IRI <iri>

    Examples:

    | hash | iri |
    | 