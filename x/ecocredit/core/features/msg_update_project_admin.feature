Feature: MsgUpdateProjectAdmin

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
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

  Scenario: an error is returned if project id is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "project id cannot be empty: invalid request"

  Scenario: an error is returned if project id is not formatted
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "foo"
    }
    """
    When the message is validated
    Then expect the error "invalid project id: foo: parse error"

  Scenario: an error is returned if new admin is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001"
    }
    """
    When the message is validated
    Then expect the error "new admin: empty address string is not allowed: invalid address"

  Scenario: an error is returned if new admin is not a bech32 address
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "project_id": "C01-001",
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
      "project_id": "C01-001",
      "new_admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "admin and new admin cannot be the same: invalid request"
