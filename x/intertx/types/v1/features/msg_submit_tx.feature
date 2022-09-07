Feature: MsgSubmitTx

  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
      "connection_id": "channel-5",
    }
    """
    And a valid tx for msg
    When the message is validated
    Then expect no error

    