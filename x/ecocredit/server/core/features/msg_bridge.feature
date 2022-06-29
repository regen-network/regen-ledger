Feature: Msg/Bridge

  Credits can be bridged to another chain:
  - when a batch contract entry exists
  - ...

  Rule: The batch contract entry must exist

    Scenario: the batch contract entry exists
      Given a credit batch exists with a batch contract entry
      And alice owns credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect no error

    Scenario: the batch contract entry does not exist
      Given a credit batch exists without a batch contract entry
      And alice owns credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect the error "only credits previously bridged from another chain are supported at this time: invalid request"
