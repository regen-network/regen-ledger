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

  Scenario: an error is returned if content hashes is empty
    Given the message
    """
    {
      "attestor": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "content hashes cannot be empty: invalid request"
  
  # Note: see ./types_content_hash.feature for content hash validation
