Feature: Msg/Bridge

  Credits can be bridged to another chain:
  - message validations
  - when a batch contract entry exists
  - the credits are cancelled and the owner balance is updated
  - the credits are cancelled and the total supply is updated
  - the bridge target is in the list of allowed bridge chains

  Rule: Message Validations

    Scenario: a valid message
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          }
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message with multiple credits
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "100"
          },
          {
            "batch_denom": "C01-001-20200101-20210101-002",
            "amount": "100"
          }
        ]
      }
      """
      When the message is validated
      Then expect no error



    Scenario: an error is returned if target is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "target cannot be empty: invalid request"

    Scenario: an error is returned if recipient is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon"
      }
      """
      When the message is validated
      Then expect the error "recipient cannot be empty: invalid request"

    Scenario: an error is returned if recipient is not an ethereum address
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "foo"
      }
      """
      When the message is validated
      Then expect the error "recipient must be a valid ethereum address: invalid address"

    Scenario: an error is returned if credits is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d"
      }
      """
      When the message is validated
      Then expect the error "credits cannot be empty: invalid request"

    Scenario: an error is returned if credits batch denom is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {}
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: batch denom: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if credits batch denom is not formatted
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {
            "batch_denom": "foo"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"

    Scenario: an error is returned if credits amount is empty
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: amount cannot be empty: invalid request"

    Scenario: an error is returned if credits amount is not a positive decimal
      Given the message
      """
      {
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "target": "polygon",
        "recipient": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
        "credits": [
          {
            "batch_denom": "C01-001-20200101-20210101-001",
            "amount": "-100"
          }
        ]
      }
      """
      When the message is validated
      Then expect the error "credits[0]: amount: expected a positive decimal, got -100: invalid decimal string"

  Rule: The batch contract entry must exist

    Scenario: the batch contract entry exists
      Given a credit batch exists with a batch contract entry
      And the target is an allowed chain
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect no error

    Scenario: the batch contract entry does not exist
      Given a credit batch exists without a batch contract entry
      And the target is an allowed chain
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch
      Then expect the error "only credits previously bridged from another chain are supported: invalid request"

  Rule: The credits are cancelled and the owner balance is updated

    Scenario: the owner balance is updated
      Given a credit batch exists with a batch contract entry
      And the target is an allowed chain
      And alice has the batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "escrowed_amount": "0"
      }
      """
      When alice attempts to bridge credit amount "10" from the credit batch
      Then expect alice batch balance
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "escrowed_amount": "0"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The credits are cancelled and the total supply is updated

    Scenario: the total supply is updated
      Given a credit batch exists with a batch contract entry
      And the target is an allowed chain
      And alice owns tradable credit amount "10" from the credit batch
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to bridge credit amount "10" from the credit batch
      Then expect batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "0",
        "cancelled_amount": "10"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Scenario: EventBridge is emitted
      Given a credit batch exists
      And the target is an allowed chain
      And batch has batch contract entry with contract address "0x6887246668a3b87f54deb3b94ba47a6f63f32985"
      And alice has address "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And alice owns tradable credit amount "10" from the credit batch
      And the batch supply
      """
      {
        "retired_amount": "0",
        "tradable_amount": "10",
        "cancelled_amount": "0"
      }
      """
      When alice attempts to bridge credit amount "10" from the credit batch to "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"
      Then expect event with properties
      """
      {
        "target": "polygon",
        "recipient": "0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
        "contract": "0x6887246668a3b87f54deb3b94ba47a6f63f32985",
        "amount": "10",
        "owner": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "batch_denom": "C01-001-20200101-20210101-001"
      }
      """

  Rule: The target must be in the AllowedBridgeChain table

    Background:
      Given the target "polygon" is an allowed bridge chain

    Scenario: the bridge chain is allowed
      Given a credit batch exists with a batch contract entry
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch with target "polygon"
      Then expect no error

    Scenario: the bridge chain is not allowed
      Given a credit batch exists with a batch contract entry
      And alice owns tradable credits from the credit batch
      When alice attempts to bridge credits from the credit batch with target "solana"
      Then expect the error "solana is not an authorized bridge target: unauthorized"