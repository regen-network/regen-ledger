Feature: MsgRegisterResolver

  Scenario: a valid message
    Given the message
    """
    {
      "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "resolver_id": 1,
      "content_hashes": [
        {
          "raw": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1,
            "file_extension": "txt"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if signer is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if signer is not a bech32 address
    Given the message
    """
    {
      "signer": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if resolver id is empty
    Given the message
    """
    {
      "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "resolver id cannot be empty: invalid request"

  Scenario: an error is returned if content hashes is empty
    Given the message
    """
    {
      "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "resolver_id": 1
    }
    """
    When the message is validated
    Then expect the error "content hashes cannot be empty: invalid request"
  
  Scenario: a valid amino message
    Given the message
    """
    {
      "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "resolver_id": 1,
      "content_hashes": [
        {
          "raw": {
            "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "digest_algorithm": 1
          }
        }
      ]
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type": "regen-ledger/MsgRegisterResolver",
      "value": {
        "content_hashes": [
          {
            "raw": {
              "digest_algorithm": 1,
              "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
            }
          }
        ],
        "resolver_id": "1",
        "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
    }
    """

  # Note: see ./types_content_hash.feature for content hash validation
