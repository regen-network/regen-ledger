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

  Scenario: owner cannot be empty
    Given the message
    """
    {}
    """
    When the message is validated
    # Then expect the error "owner cannot be empty: invalid request"
