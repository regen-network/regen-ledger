Feature: BatchBalance

  Scenario: a valid batch balance
    Given the batch balance
    """
    {
      "batch_key": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "tradable_amount": "100",
      "retired_amount": "100",
      "escrowed_amount": "100"
    }
    """
    When the batch balance is validated
    Then expect no error

  Scenario: a valid batch balance without amounts
    Given the batch balance
    """
    {
      "batch_key": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the batch balance is validated
    Then expect no error

  Scenario: an error is returned if batch key is empty
    Given the batch balance
    """
    {}
    """
    When the batch balance is validated
    Then expect the error "batch key cannot be zero: parse error"

  Scenario: an error is returned if address is empty
    Given the batch balance
    """
    {
      "batch_key": 1
    }
    """
    When the batch balance is validated
    Then expect the error "address: empty address string is not allowed: parse error"

  Scenario: an error is returned if tradable amount is not positive
    Given the batch balance
    """
    {
      "batch_key": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "tradable_amount": "-100"
    }
    """
    When the batch balance is validated
    Then expect the error "tradable amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if retired amount is not positive
    Given the batch balance
    """
    {
      "batch_key": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "retired_amount": "-100"
    }
    """
    When the batch balance is validated
    Then expect the error "retired amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if escrowed amount is not positive
    Given the batch balance
    """
    {
      "batch_key": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "escrowed_amount": "-100"
    }
    """
    When the batch balance is validated
    Then expect the error "escrowed amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"
