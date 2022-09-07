Feature: MsgSubmitTx

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
      "connection_id": "channel-5"
    }
    """
    And a valid tx for msg
    When the message is validated
    Then expect no error

  Scenario: an error is returned if owner is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "owner cannot be empty: invalid request"

  Scenario: an error is returned if owner is not a valid bech32 address
    Given the message
    """
    {
      "owner": "foo"
    }
    """
    When the message is validated
    Then expect the error "owner: decoding bech32 failed: invalid bech32 string length 3"

  Scenario: an error is returned if  connection id is empty
    Given the message
    """
    {
      "owner": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
    }
    """
    When the message is validated
    Then expect the error "connection_id cannot be empty: invalid request"

  Scenario: an error is returned if msg is empty
    Given the message
    """
    {
      "owner": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
      "connection_id": "channel-5"
    }
    """
    When the message is validated
    Then expect the error "msg cannot be empty: invalid request"