Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - when the credit class exists
  - a new credit batch is created if a batch contract entry does not exist
  - a new project is created if a project from the same class with the same reference id does not exist

  Rule: The credit class must exist

    Scenario: credit class does not exist
      When alice attempts to bridge credits with class id "C01"
      Then expect the error "credit class with id C01: not found"

    Scenario: credit class exists
      Given a credit class with id "C01"
      When alice attempts to bridge credits with class id "C01"
      Then expect no error

  Rule: A new credit batch is created if a batch contract entry does not exist

    Background:
      Given a credit class
      And a project

    Scenario: batch contract entry does not exist
      Given a credit batch with no contract
      When alice attempts to bridge credits with contract "0x0"
      Then expect total credit batches "2"

    Scenario: batch contract entry exists
      Given a credit batch with contract "0x0"
      When alice attempts to bridge credits with contract "0x0"
      Then expect total credit batches "1"

  Rule: A new project is created if a project from the same class with the same reference id does not exist

    Background:
      Given a credit class with id "C01"
      And a project with reference id "VCS-001"

    Scenario: a project from the same class with a different reference id
      When alice attempts to bridge credits with class id "C01" and project reference id "VCS-002"
      Then expect total projects "2"

    Scenario: a project from a different class with the same reference id
      Given a credit class with id "C02"
      When alice attempts to bridge credits with class id "C02" and project reference id "VCS-001"
      Then expect total projects "2"

    Scenario: a project from the same class with the same reference id
      When alice attempts to bridge credits with class id "C01" and project reference id "VCS-001"
      Then expect total projects "1"
