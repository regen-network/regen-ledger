Feature: MsgCreateClass

  Scenario: a valid message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "uregen",
        "amount": "20000000"
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple issuers
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "uregen",
        "amount": "20000000"
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message without fee
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C"
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

  Scenario: an error is returned if issuers is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "issuers cannot be empty: invalid request"

  Scenario: an error is returned if issuer is not a bech32 address
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "foo"
      ]
    }
    """
    When the message is validated
    Then expect the error "issuers[0]: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if issuer is a duplicate
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
    }
    """
    When the message is validated
    Then expect the error "issuers[1]: duplicate address regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6: invalid request"

  Scenario: an error is returned if metadata is exceeds 256 characters
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ]
    }
    """
    And metadata with length "257"
    When the message is validated
    Then expect the error "metadata: max length 256: limit exceeded"

  Scenario: an error is returned if credit type abbreviation is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the message is validated
    Then expect the error "credit type abbrev: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if credit type abbreviation is not formatted
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "foobar"
    }
    """
    When the message is validated
    Then expect the error "credit type abbrev: must be 1-3 uppercase alphabetic characters: parse error: invalid request"

  Scenario: an error is returned if fee denom is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {}
    }
    """
    When the message is validated
    Then expect the error "fee denom cannot be empty: invalid request"

  Scenario: an error is returned if fee denom is not a valid bank denom
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "1"
      }
    }
    """
    When the message is validated
    Then expect the error "invalid denom: 1: invalid request"

  Scenario: an error is returned if fee amount is empty
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "uregen"
      }
    }
    """
    When the message is validated
    Then expect the error "fee amount cannot be empty: invalid request"

  Scenario: an error is returned if fee amount is not positive
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "uregen",
        "amount": "-20000000"
      }
    }
    """
    When the message is validated
    Then expect the error "fee amount must be positive: insufficient fee"

  Scenario: a valid amino message
    Given the message
    """
    {
      "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "issuers": [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      ],
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C",
      "fee": {
        "denom": "uregen",
        "amount": "20000000"
      }
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen/MsgCreateClass",
      "value":{
        "admin":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "credit_type_abbrev":"C",
        "fee":{
          "amount":"20000000",
          "denom":"uregen"
        },
        "issuers":[
          "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
        ],
        "metadata":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      }
    }
    """
