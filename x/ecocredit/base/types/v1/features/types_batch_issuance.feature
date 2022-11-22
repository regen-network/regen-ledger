Feature: BatchIssuance

  Scenario: a valid batch issuance
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "tradable_amount": "100",
      "retired_amount": "100",
      "retirement_jurisdiction": "US-WA"
    }
    """
    When the batch issuance is validated
    Then expect no error

  Scenario: a valid batch issuance without tradable amount
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "100",
      "retirement_jurisdiction": "US-WA"
    }
    """
    When the batch issuance is validated
    Then expect no error

  Scenario: a valid batch issuance without retired amount
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "tradable_amount": "100"
    }
    """
    When the batch issuance is validated
    Then expect no error

  Scenario: an error is returned if issuance recipient is empty
    Given the batch issuance
    """
    {}
    """
    When the batch issuance is validated
    Then expect the error "recipient: empty address string is not allowed: invalid address"

  Scenario: an error is returned if issuance recipient is not a bech32 address
    Given the batch issuance
    """
    {
      "recipient": "foo"
    }
    """
    When the batch issuance is validated
    Then expect the error "recipient: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if issuance tradable amount and retired amount are empty
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the batch issuance is validated
    Then expect the error "tradable amount or retired amount required: invalid request"

  Scenario: an error is returned if issuance tradable amount is a negative decimal
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "tradable_amount": "-100"
    }
    """
    When the batch issuance is validated
    Then expect the error "tradable amount: expected a non-negative decimal, got -100: invalid decimal string"

  Scenario: an error is returned if issuance retired amount is a negative decimal
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "-100"
    }
    """
    When the batch issuance is validated
    Then expect the error "retired amount: expected a non-negative decimal, got -100: invalid decimal string"

  Scenario: an error is returned if issuance retired amount is positive and retirement jurisdiction is empty
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "100"
    }
    """
    When the batch issuance is validated
    Then expect the error "retirement jurisdiction: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if issuance retired amount is positive and retirement jurisdiction is not formatted
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "100",
      "retirement_jurisdiction": "foo"
    }
    """
    When the batch issuance is validated
    Then expect the error "retirement jurisdiction: expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"

  Scenario: an error is returned if issuance retired amount is positive and retirement reason exceeds 512 characters
    Given the batch issuance
    """
    {
      "recipient": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "retired_amount": "100",
      "retirement_jurisdiction": "US-WA"
    }
    """
    And retirement reason with length "513"
    When the batch issuance is validated
    Then expect the error "retirement reason: max length 512: limit exceeded"
