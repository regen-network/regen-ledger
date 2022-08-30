Feature: SellOrder

  Scenario: a valid sell order
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "100",
      "market_id": 1,
      "ask_amount": "100"
    }
    """
    When the sell order is validated
    Then expect no error

  Scenario: a valid sell order with expiration
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "100",
      "market_id": 1,
      "ask_amount": "100",
      "expiration": "2020-01-01T00:00:00Z"
    }
    """
    When the sell order is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the sell order
    """
    {}
    """
    When the sell order is validated
    Then expect the error "id cannot be zero: parse error"

  Scenario: an error is returned if seller is empty
    Given the sell order
    """
    {
      "id": 1
    }
    """
    When the sell order is validated
    Then expect the error "seller: empty address string is not allowed: parse error"

  Scenario: an error is returned if batch key is empty
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the sell order is validated
    Then expect the error "batch key cannot be zero: parse error"

  Scenario: an error is returned if quantity is empty
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1
    }
    """
    When the sell order is validated
    Then expect the error "quantity cannot be empty: parse error"

  Scenario: an error is returned if quantity is not a positive decimal
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "-100"
    }
    """
    When the sell order is validated
    Then expect the error "quantity: expected a non-negative decimal, got -100: invalid decimal string: parse error"

  Scenario: an error is returned if market id is empty
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "100"
    }
    """
    When the sell order is validated
    Then expect the error "market id cannot be zero: parse error"

  Scenario: an error is returned if ask amount is empty
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "100",
      "market_id": 1
    }
    """
    When the sell order is validated
    Then expect the error "ask amount cannot be empty: parse error"

  Scenario: an error is returned if ask amount is a negative decimal
    Given the sell order
    """
    {
      "id": 1,
      "seller": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "batch_key": 1,
      "quantity": "100",
      "market_id": 1,
      "ask_amount": "-100"
    }
    """
    When the sell order is validated
    Then expect the error "ask amount: expected a non-negative decimal, got -100: invalid decimal string: parse error"
