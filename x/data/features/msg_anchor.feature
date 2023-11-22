Feature: MsgAnchor

  Scenario: a valid message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "content_hash": {
        "raw": {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1
        }
      }
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if sender is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if sender is not a bech32 address
    Given the message
    """
    {
      "sender": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if content hash is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "content hash cannot be empty: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "content_hash": {
        "raw": {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1
        }
      }
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen-ledger/MsgAnchor",
      "value":{
        "content_hash":{
          "raw":{
            "digest_algorithm":1,
            "hash":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
          }
        },
        "sender":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
    }
    """

  # Note: see ./types_content_hash.feature for content hash validation
