Feature: Msg/BridgeReceive

  Credits can be bridged from another chain:
  - credits are issued in a new credit batch when batch contract mapping does not exist
  - credits are minted to an existing credit batch when batch contract mapping exists
  - ...

  Rule: Credits are issued in a new credit batch when batch contract mapping does not exist

    Scenario: batch contract mapping exists

    Scenario: batch contract mapping does not exist

  Rule: Credits are minted to an existing credit batch when batch contract mapping does not exist

    Scenario: the batch contract mapping exists

    Scenario: the batch contract mapping does not exist
