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
