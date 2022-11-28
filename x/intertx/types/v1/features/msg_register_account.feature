Feature: MsgRegisterAccount


  Scenario: a valid message
    Given the message
    """
    {
      "owner": "regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza",
      "connection_id": "channel-5",
      "version": "3"
    }
    """
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

  Scenario: an error is returned if connection id is empty
    Given the message
    """
    {
      "owner": "regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza",
      "version": "5"
    }
    """
    When the message is validated
    Then expect the error "connection_id cannot be empty: invalid request"

  Scenario: a valid amino msg
    Given the message
    """
    {
      "owner": "regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza",
      "connection_id": "channel-5",
      "version": "3"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type": "intertx/MsgRegisterAccount",
      "value":{
        "connection_id": "channel-5",
        "owner": "regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza",
        "version": "3"
      }
    }
    """
