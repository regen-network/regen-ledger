Feature: MsgAttest

  Scenario: a valid message
    Given the message
     """
     {
       "attestor": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
       "content_hashes": [
         {
           "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
           "digest_algorithm": 1,
           "canonicalization_algorithm": 1
         }
       ]
     }
     """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if attestor is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "empty address string is not allowed: invalid address"

  Scenario: an error is returned if attestor is not a bech32 address
    Given the message
    """
    {
      "attestor": "foo"
    }
    """
    When the message is validated
    Then expect the error "decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if content hashes is empty
    Given the message
    """
    {
      "attestor": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "content hashes cannot be empty: invalid request"
  
  Scenario: a valid amino message
    Given the message
    """
    {
      "attestor": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "content_hashes": [
        {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1,
          "canonicalization_algorithm": 1
        }
      ]
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen-ledger/MsgAttest",
      "value":{
        "attestor":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "content_hashes":[
          {
            "canonicalization_algorithm":1,
            "digest_algorithm":1,
            "hash":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
          }
        ]
      }
    }
    """

  # Note: see ./types_content_hash.feature for content hash validation
