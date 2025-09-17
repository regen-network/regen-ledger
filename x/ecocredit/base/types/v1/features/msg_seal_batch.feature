Feature: MsgSealBatch

  Scenario: a valid message
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001"
    }
    """
    When the message is validated
    Then expect no error


  Scenario: an error is returned if batch denom is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "batch denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if batch denom is not formatted
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error: invalid request"
