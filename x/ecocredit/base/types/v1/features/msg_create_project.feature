Feature: MsgCreateProject

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "jurisdiction": "US-WA",
      "reference_id": "VCS-001"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without metadata
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "jurisdiction": "US-WA",
      "reference_id": "VCS-001"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without reference id
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "jurisdiction": "US-WA"
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

  Scenario: an error is returned if metadata is exceeds 256 characters
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01"
    }
    """
    And metadata with length "257"
    When the message is validated
    Then expect the error "metadata: max length 256: limit exceeded"

  Scenario: an error is returned if jurisdiction is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the message is validated
    Then expect the error "jurisdiction: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if jurisdiction is not formatted
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "jurisdiction": "foo"
    }
    """
    When the message is validated
    Then expect the error "jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

  Scenario: an error is returned if reference id is exceeds 32 characters
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "jurisdiction": "US-WA"
    }
    """
    And a reference id with length "33"
    When the message is validated
    Then expect the error "reference id: max length 32: limit exceeded"

  Scenario: a valid amino message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "jurisdiction": "US-WA",
      "reference_id": "VCS-001"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgCreateProject",
      "value":{
        "admin":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id":"C01",
        "jurisdiction":"US-WA",
        "metadata":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
        "reference_id":"VCS-001"
      }
    }
    """
