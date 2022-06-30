Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - a new credit batch is created if a batch contract entry does not exist
  - a new project is created if a project with the same reference id does not exist

  Rule: A new credit batch is created if a batch contract entry does not exist

    Scenario: batch contract entry does not exist
      Given a credit batch exists with no contract
      When alice attempts to bridge credits from contract "0x0"
      Then expect total credit batches "2"

    Scenario: batch contract entry exists
      Given a credit batch exists with contract "0x0"
      When alice attempts to bridge credits from contract "0x0"
      Then expect total credit batches "1"

  Rule: A new project is created if a project with the same reference id does not exist

    Background:
      Given a project exists with reference id "VCS-001"

    Scenario: a project with the same reference id does not exist
      When alice attempts to bridge credits with project reference id "VCS-002"
      Then expect total projects "2"

    Scenario: a project with the same reference id exists
      When alice attempts to bridge credits with project reference id "VCS-001"
      Then expect total projects "1"
