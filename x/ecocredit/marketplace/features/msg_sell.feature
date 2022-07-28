Feature: MsgSell

  Scenario: a valid message
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with disable auto retire
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with expiration
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          },
          "expiration": "2030-01-01T00:00:00Z"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple orders
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          }
        },
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if seller is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "seller cannot be empty: invalid request"

  Scenario: an error is returned if seller is not a valid bech32 address
    Given the message
    """
    {
      "seller": "foo"
    }
    """
    When the message is validated
    Then expect the error "seller is not a valid address: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if orders is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "orders cannot be empty: invalid request"

  Scenario: an error is returned if order batch denom is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: batch denom cannot be empty: parse error: invalid request"

  Scenario: an error is returned if order batch denom is not formatted
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: invalid batch denom: expected format A00-000-00000000-00000000-000: parse error: invalid request"

  Scenario: an error is returned if order quantity is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: quantity cannot be empty: invalid request"

  Scenario: an error is returned if order quantity is not a positive decimal
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: quantity must be a positive decimal: invalid request"

  Scenario: an error is returned if ask price is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: ask price: cannot be empty: invalid request"

  Scenario: an error is returned if ask price denom is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {}
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: ask price: denom cannot be empty: invalid request"

  Scenario: an error is returned if ask price denom is not formatted
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "foo.bar"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: ask price: invalid denom: foo.bar: invalid request"

  Scenario: an error is returned if ask price amount is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: ask price: amount cannot be empty: invalid request"

  Scenario: an error is returned if ask price amount is not a positive integer
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "100",
          "ask_price": {
            "denom": "regen",
            "amount": "-100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: ask price: amount must be a positive integer: invalid request"
