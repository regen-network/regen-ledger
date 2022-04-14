Feature: MsgDefineResolver

  Rule: only a valid manager is accepted

    Scenario: an error is returned if manager is empty
      Given the message
      """
      {}
      """
      When the message is validated
      Then expect the error "empty address string is not allowed: invalid address"

    Scenario: an error is returned if manager is not a valid address
      Given the message
      """
      {
        "manager": "foo"
      }
      """
      When the message is validated
      Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Rule: only a valid resolver url is accepted

    Scenario: an error is returned if resolver url is empty
      Given the message
      """
      {
        "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
      }
      """
      When the message is validated
      Then expect the error "invalid resolver url: invalid request"

    Scenario: an error is returned if resolver url is not a valid url
      Given the message
      """
      {
        "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
        "resolver_url": "foo"
      }
      """
      When the message is validated
      Then expect the error "invalid resolver url: invalid request"

  Rule: only a valid message is accepted

    Scenario: no error is returned if manager and resolver url are valid
      Given the message
      """
      {
        "manager": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
        "resolver_url": "https://foo.bar"
      }
      """
      When the message is validated
      Then expect no error
