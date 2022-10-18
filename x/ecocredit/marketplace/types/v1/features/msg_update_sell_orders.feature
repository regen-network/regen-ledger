Feature: MsgUpdateSellOrders

  Scenario: a valid message
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
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
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
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
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "regen",
            "amount": "100"
          },
          "new_expiration": "2030-01-01T00:00:00Z"
        }
      ]
    }
    """
    When the message is validated
    Then expect no error

  Scenario: a valid message with multiple updates
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "regen",
            "amount": "100"
          }
        },
        {
          "sell_order_id": 2,
          "new_quantity": "100",
          "new_ask_price": {
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

  Scenario: an error is returned if updates is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "updates cannot be empty: invalid request"

  Scenario: an error is returned if update sell order id is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {}
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: sell order id cannot be empty: invalid request"

  Scenario: an error is returned if update new quantity is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new quantity cannot be empty: invalid request"

  Scenario: an error is returned if update new quantity is not a positive decimal
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "-100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new quantity must be a positive decimal: invalid request"

  Scenario: an error is returned if ask price is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100"
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new ask price cannot be empty: invalid request"

  Scenario: an error is returned if update new ask price denom is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {}
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new ask price: denom cannot be empty: invalid request"

  Scenario: an error is returned if ask price denom is not formatted
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "foo#bar"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new ask price: invalid denom: foo#bar: invalid request"

  Scenario: an error is returned if ask price amount is empty
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "regen"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new ask price: amount cannot be empty: invalid request"

  Scenario: an error is returned if ask price amount is not a positive integer
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "regen",
            "amount": "-100"
          }
        }
      ]
    }
    """
    When the message is validated
    Then expect the error "updates[0]: new ask price: amount must be a positive integer: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
      "seller": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
      "updates": [
        {
          "sell_order_id": 1,
          "new_quantity": "100",
          "new_ask_price": {
            "denom": "regen",
            "amount": "100"
          }
        }
      ]
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen.marketplace/MsgUpdateSellOrders",
      "value":{
        "seller":"regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "updates":[
          {
            "new_ask_price":{
              "amount":"100",
              "denom":"regen"
            },
          "new_quantity":"100",
          "sell_order_id":"1"
          }
        ]
      }
    }
    """
