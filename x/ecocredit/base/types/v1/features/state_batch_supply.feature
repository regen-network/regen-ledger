Feature: BatchSupply

  Scenario: a valid batch supply
    Given the batch supply
    """
    {
      "batch_key": 1,
      "tradable_amount": "100",
      "retired_amount": "100",
      "cancelled_amount": "100"
    }
    """
    When the batch supply is validated
    Then expect no error

  Scenario: a valid batch supply without amounts
    Given the batch supply
    """
    {
      "batch_key": 1
    }
    """
    When the batch supply is validated
    Then expect no error

  Scenario: an error is returned if batch key is empty
    Given the batch supply
    """
    {}
    """
    When the batch supply is validated
    Then expect the error "batch key cannot be zero: parse error"

  Scenario: an error is returned if tradable amount is not positive
    Given the batch supply
    """
    {
      "batch_key": 1,
      "tradable_amount": "-100"
    }
    """
    When the batch supply is validated
    Then expect the error "tradable amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if retired amount is not positive
    Given the batch supply
    """
    {
      "batch_key": 1,
      "retired_amount": "-100"
    }
    """
    When the batch supply is validated
    Then expect the error "retired amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if cancelled amount is not positive
    Given the batch supply
    """
    {
      "batch_key": 1,
      "cancelled_amount": "-100"
    }
    """
    When the batch supply is validated
    Then expect the error "cancelled amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"
