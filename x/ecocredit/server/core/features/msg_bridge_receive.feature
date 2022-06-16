Feature: BridgeReceive

  Credits can be received by regen from an external source when:
  - the bridge service address is an issuer of the Credit Class the bridged credits belong to.
  - the provided class_id exists
  - the bridge message originates from a unique OriginTx

  Rule: the credit class must exist, and the bridge address must be an issuer
    Scenario: Bridge is an issuer of the class
      Given a valid class id
      And the bridge address "foo" is an issuer
      And a valid bridge msg
      Then the transaction succeeds

    Scenario: Bridge is not an issuer of the class
      Given a valid class id
      And the bridge address "foo" is not an issuer
      And a valid bridge msg
      Then the transaction fails with "is not an issuer for the class: unauthorized"

    Scenario: Class does not exist
      Given an invalid class id
      And the bridge address "foo"
      And a valid bridge msg
      Then the transaction fails with "could not get class"

  Rule: bridged credits with a new reference ID creates a project

    Scenario: a new reference ID creates a project
      Given a valid class id
      And the bridge address "foo" is an issuer
      And a new reference id
      And a valid bridge msg
      Then the transaction succeeds
      And a new project is created
