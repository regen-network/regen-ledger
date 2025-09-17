Feature: MsgAnchor

  Scenario: a valid message
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "content_hash": {
        "raw": {
          "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
          "digest_algorithm": 1,
          "file_extension": "bin"
        }
      }
    }
    """
    When the message is validated
    Then expect no error

  
  Scenario: an error is returned if content hash is empty
    Given the message
    """
    {
      "sender": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "content hash cannot be empty: invalid request"

  # Note: see ./types_content_hash.feature for content hash validation
