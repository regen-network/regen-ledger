Feature: MsgBuyDirect

  Scenario: a valid message
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "100"
          },
          "retirement_jurisdiction": "US-WA"
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
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "100"
          },
          "retirement_jurisdiction": "US-WA"
        },
        {
          "sell_order_id": "1",
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "100"
          },
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with disable auto-retire
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "100",
          "bid_price": {
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

  Scenario: an error is returned if buyer is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "buyer cannot be empty: invalid request"

  Scenario: an error is returned if buyer is not a valid bech32 address
    Given the message
    """
    {
      "buyer": "foo"
    }
    """
    When the message is validated
    Then expect the error "buyer is not a valid address: decoding bech32 failed: invalid bech32 string length 3: invalid request"

  Scenario: an error is returned if orders is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
    }
    """
    When the message is validated
    Then expect the error "orders cannot be empty: invalid request"

  Scenario: an error is returned if sell order id is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: sell order id cannot be empty: invalid request"

  Scenario: an error is returned if order quantity is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1
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
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: quantity must be a positive decimal: invalid request"

  Scenario: an error is returned if bid price is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: bid price cannot be empty: invalid request"

  Scenario: an error is returned if bid price denom is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {}
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: bid price: denom cannot be empty: invalid request"

  Scenario: an error is returned if bid price denom is not formatted
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {
            "denom": "foo.bar"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: bid price: invalid denom: foo.bar: invalid request"

  Scenario: an error is returned if bid price amount is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {
            "denom": "regen"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: bid price: amount cannot be empty: invalid request"

  Scenario: an error is returned if bid price amount is not a positive integer
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "-100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: bid price: amount must be a positive integer: invalid request"

  Scenario: an error is returned if disable auto-retire is true and retirement jurisdiction is empty
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: retirement jurisdiction cannot be empty if auto-retire is disabled: invalid request"

  Scenario: an error is returned if disable auto-retire is true and retirement jurisdiction is not formatted
    Given the message
    """
    {
      "buyer": "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
      "orders": [
        {
          "sell_order_id": 1,
          "quantity": "100",
          "bid_price": {
            "denom": "regen",
            "amount": "100"
          },
          "retirement_jurisdiction": "foo"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "orders[0]: invalid jurisdiction: foo, expected format <country-code>[-<region-code>[ <postal-code>]]: parse error: invalid request"
