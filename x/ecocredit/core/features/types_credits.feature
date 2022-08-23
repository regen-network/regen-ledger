Feature: Credits

  Scenario: a valid credits
    Given the message
    """
    {
      "batch_denom": "C01-001-20200101-20210101-001",
      "amount": "100"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if batch denom is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "batch denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if batch denom is not formatted
    Given the message
    """
    {
      "batch_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "batch denom: expected format [A-Z]{1,3}[0-9]{2,}-[0-9]{3,}-[0-9]{8}-[0-9]{8}-[0-9]{3,}: parse error: invalid request"

  Scenario: an error is returned if amount is empty
    Given the message
    """
    {
      "batch_denom": "C01-001-20200101-20210101-001"
    }
    """
    When the message is validated
    Then expect the error "amount cannot be empty: invalid request"

  Scenario: an error is returned if amount is not a positive decimal
    Given the message
    """
    {
      "batch_denom": "C01-001-20200101-20210101-001",
      "amount": "-100"
    }
    """
    When the message is validated
    Then expect the error "expected a positive decimal, got -100: invalid decimal string"
