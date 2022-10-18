Feature: MsgUpdateClassAdmin

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "new_admin": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if admin is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "admin: empty address string is not allowed: invalid address"

  Scenario: an error is returned if admin is not a bech32 address
    Given the message
    """
    {
      "admin": "foo"
    }
    """
    When the message is validated
    Then expect the error "admin: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if class id is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "class id: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if class id is not formatted
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "foo"
    }
    """
    When the message is validated
    Then expect the error "class id: expected format <credit-type-abbrev><class-sequence>: parse error: invalid request"

  Scenario: an error is returned if new admin is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01"
    }
    """
    When the message is validated
    Then expect the error "new admin: empty address string is not allowed: invalid address"

  Scenario: an error is returned if new admin is not a bech32 address
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "new_admin": "foo"
    }
    """
    When the message is validated
    Then expect the error "new admin: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if admin and new admin are the same
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "new_admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "admin and new admin cannot be the same: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "new_admin": "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgUpdateClassAdmin",
      "value":{
        "admin":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id":"C01",
        "new_admin":"regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      }
    }
    """
