Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - a new credit batch is created if a batch contract entry does not exist
  - ...

  Rule: A new credit batch is created if a batch contract entry does not exist

    Scenario: batch contract entry does not exist
      Given a credit batch exists with no contract
      When alice attempts to bridge credits from contract "0x0"
      Then expect total credit batches "2"

    Scenario: batch contract entry exists
      Given a credit batch exists with contract "0x0"
      When alice attempts to bridge credits from contract "0x0"
      Then expect total credit batches "1"
