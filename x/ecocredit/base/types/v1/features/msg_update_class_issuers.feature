Feature: MsgUpdateClassIssuers

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "add_issuers": [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ],
      "remove_issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without add issuers
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "remove_issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without remove issuers
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "add_issuers": [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
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

  Scenario: an error is returned if new issuers and remove issuers is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01"
    }
    """
    When the message is validated
    Then expect the error "must specify at least one of add_issuers or remove_issuers: invalid request"

  Scenario: an error is returned if new issuer is not a bech32 address
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "add_issuers": [
        "foo"
      ]
    }
    """
    When the message is validated
    Then expect the error "add_issuers[0]: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if new issuer is a duplicate
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "add_issuers": [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
    }
    """
    When the message is validated
    Then expect the error "add_issuers[1]: duplicate issuer: invalid request"

  Scenario: an error is returned if remove issuer is not a bech32 address
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "remove_issuers": [
        "foo"
      ]
    }
    """
    When the message is validated
    Then expect the error "remove_issuers[0]: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if remove issuer is a duplicate
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "remove_issuers": [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
    }
    """
    When the message is validated
    Then expect the error "remove_issuers[1]: duplicate issuer: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "class_id": "C01",
      "add_issuers": [
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ],
      "remove_issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgUpdateClassIssuers",
      "value":{
        "add_issuers":[
          "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
        ],
        "admin":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id":"C01",
        "remove_issuers":[
          "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
        ]
      }
    }
    """
