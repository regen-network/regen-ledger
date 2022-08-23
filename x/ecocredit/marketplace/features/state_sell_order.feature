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
