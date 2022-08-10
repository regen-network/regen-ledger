Feature: MsgUpdateClassMetadata

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "new_metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message removing metadata
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01"
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
    Then expect the error "class id cannot be empty: parse error: invalid request"

  Scenario: an error is returned if class id is not formatted
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "foo"
    }
    """
    When the message is validated
    Then expect the error "class ID didn't match the format: expected A00, got foo: parse error: invalid request"

  Scenario: an error is returned if new metadata exceeds 256 characters
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01"
    }
    """
    And new metadata with length "257"
    When the message is validated
    Then expect the error "metadata: max length 256: limit exceeded"
