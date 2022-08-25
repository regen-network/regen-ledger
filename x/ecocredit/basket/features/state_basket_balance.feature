Feature: BasketBalance

  Scenario: a valid basket balance
    Given the basket balance
    """
    {
      "basket_id": 1,
      "batch_denom": "C01-001-20200101-20210101-001",
      "balance": "100",
      "batch_start_date": "2020-01-01T00:00:00Z"
    }
    """
    When the basket balance is validated
    Then expect no error

  Scenario: an error is returned if basket id is empty
    Given the basket balance
    """
    {}
    """
    When the basket balance is validated
    Then expect the error "basket id cannot be zero: parse error"

  Scenario: an error is returned if batch denom is empty
    Given the basket balance
    """
    {
      "basket_id": 1
    }
    """
    When the basket balance is validated
    Then expect the error "batch denom: empty string is not allowed: parse error"

  Scenario: an error is returned if batch denom is not formatted
    Given the basket balance
    """
    {
      "basket_id": 1,
      "batch_denom": "foo"
    }
    """
    When the basket balance is validated
    Then expect the error "batch denom: expected format <project-id>-<start_date>-<end_date>-<batch_sequence>: parse error"

  Scenario: an error is returned if balance is a negative decimal
    Given the basket balance
    """
    {
      "basket_id": 1,
      "batch_denom": "C01-001-20200101-20210101-001",
      "balance": "-100"
    }
    """
    When the basket balance is validated
    Then expect the error "balance: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if batch start date is empty
    Given the basket balance
    """
    {
      "basket_id": 1,
      "batch_denom": "C01-001-20200101-20210101-001",
      "balance": "100"
    }
    """
    When the basket balance is validated
    Then expect the error "batch start date cannot be empty: parse error"
