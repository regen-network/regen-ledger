Feature: Market Integration

  - create sell order and buy credits
  - create sell order and buy credits that auto-retire
  - create sell order and update sell order (increase quantity)
  - create sell order and update sell order (decrease quantity)
  - create sell order and cancel sell order
  - create multiple sell orders and buy credits from one sell order
  - create multiple sell orders and buy credits from multiple sell orders
  - create multiple sell orders and update multiple sell orders (increase quantity)
  - create multiple sell orders and update multiple sell orders (decrease quantity)

  Background:
    Given ecocredit state
    """
    {
      "regen.ecocredit.v1.CreditType": [
        {
          "abbreviation": "C",
          "name": "carbon",
          "precision": 6,
          "unit": "metric ton CO2 equivalent"
        }
      ],
      "regen.ecocredit.v1.Class": [
        1,
        {
          "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
          "credit_type_abbrev": "C",
          "id": "C01",
          "key": "1",
          "metadata": "metadata"
        }
      ],
      "regen.ecocredit.v1.Project": [
        1,
        {
          "key": "1",
          "id": "C01-001",
          "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
          "class_key": "1",
          "jurisdiction": "US-WA",
          "metadata": "metadata"
        }
      ],
      "regen.ecocredit.v1.Batch": [
        1,
        {
          "key": "1",
          "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
          "project_key": "1",
          "denom": "C01-001-20200101-20210101-001",
          "metadata": "metadata",
          "start_date": "2020-01-01T00:00:00Z",
          "end_date": "2021-01-01T00:00:00Z",
          "issuance_date": "2022-01-01T00:00:00Z"
        }
      ],
      "regen.ecocredit.v1.BatchBalance": [
        {
          "batch_key": "1",
          "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
          "tradable_amount": "1",
          "retired_amount": "0",
          "escrowed_amount": "0"
        }
      ],
      "regen.ecocredit.v1.BatchSupply": [
        {
          "batch_key": "1",
          "tradable_amount": "1",
          "retired_amount": "0",
          "cancelled_amount": "0"
        }
      ],
      "regen.ecocredit.marketplace.v1.AllowedDenom": [
        {
          "bank_denom": "uregen",
          "display_denom": "regen",
          "exponent": "6"
        }
      ]
    }
    """

  Scenario: create sell order and buy credits
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect total sell orders "1"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When bob buys credits with message
    """
    {
      "buyer": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "0.000001",
          "bid_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event buy
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000000"
      }
    }
    """
    And expect query balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.000001",
        "retired_amount": "0",
        "escrowed_amount": "0"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create sell order and buy credits that auto-retire
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          }
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect total sell orders "1"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": false
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When bob buys credits with message
    """
    {
      "buyer": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "0.000001",
          "bid_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "retirement_jurisdiction": "US-WA"
        }
      ]
    }
    """
    Then expect no error
    And expect event buy
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000000"
      }
    }
    """
    And expect query balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0",
        "retired_amount": "0.000001",
        "escrowed_amount": "0"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "0.999999",
      "retired_amount": "0.000001",
      "cancelled_amount": "0"
    }
    """

  Scenario: create sell order and update sell order (increase quantity)
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect total sell orders "1"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When alice updates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "updates": [
        {
          "sell_order_id": "1",
          "new_quantity": "0.000002",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event update
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create sell order and update sell order (decrease quantity)
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000002",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect total sell orders "1"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000002",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When alice updates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "updates": [
        {
          "sell_order_id": "1",
          "new_quantity": "0.000001",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event update
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create sell order and cancel sell order
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect total sell orders "1"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999999",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When alice cancels sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "sell_order_id": "1"
    }
    """
    Then expect no error
    And expect event cancel
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "1.000000",
        "retired_amount": "0",
        "escrowed_amount": "0.000000"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create multiple sell orders and buy credits from one sell order
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect total sell orders "2"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query sell order with id "2"
    """
    {
      "sell_order": {
        "id": "2",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When bob buys credits with message
    """
    {
      "buyer": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "0.000001",
          "bid_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event buy
    """
    {
      "sell_order_id": "1"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000001"
      }
    }
    """
    And expect query balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.000001",
        "retired_amount": "0",
        "escrowed_amount": "0"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create multiple sell orders and buy credits from multiple sell orders
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect total sell orders "2"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query sell order with id "2"
    """
    {
      "sell_order": {
        "id": "2",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When bob buys credits with message
    """
    {
      "buyer": "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42",
      "orders": [
        {
          "sell_order_id": "1",
          "quantity": "0.000001",
          "bid_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "sell_order_id": "2",
          "quantity": "0.000001",
          "bid_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event buy
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000000"
      }
    }
    """
    And expect query balance with address "regen1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxwj8p42" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.000002",
        "retired_amount": "0",
        "escrowed_amount": "0"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create multiple sell orders and update multiple sell orders (increase quantity)
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000001",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect total sell orders "2"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query sell order with id "2"
    """
    {
      "sell_order": {
        "id": "2",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000001",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When alice updates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "updates": [
        {
          "sell_order_id": "1",
          "new_quantity": "0.000002",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "sell_order_id": "2",
          "new_quantity": "0.000002",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event update
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999996",
        "retired_amount": "0",
        "escrowed_amount": "0.000004"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """

  Scenario: create multiple sell orders and update multiple sell orders (decrease quantity)
    When alice creates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "orders": [
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000002",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "batch_denom": "C01-001-20200101-20210101-001",
          "quantity": "0.000002",
          "ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event sell
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect total sell orders "2"
    And expect query sell order with id "1"
    """
    {
      "sell_order": {
        "id": "1",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000002",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query sell order with id "2"
    """
    {
      "sell_order": {
        "id": "2",
        "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
        "batch_denom": "C01-001-20200101-20210101-001",
        "quantity": "0.000002",
        "ask_denom": "uregen",
        "ask_amount": "1",
        "disable_auto_retire": true
      }
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999996",
        "retired_amount": "0",
        "escrowed_amount": "0.000004"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
    When alice updates sell order with message
    """
    {
      "seller": "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv",
      "updates": [
        {
          "sell_order_id": "1",
          "new_quantity": "0.000001",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        },
        {
          "sell_order_id": "2",
          "new_quantity": "0.000001",
          "new_ask_price": {
            "denom": "uregen",
            "amount": "1"
          },
          "disable_auto_retire": true
        }
      ]
    }
    """
    Then expect no error
    And expect event update
    """
    {
      "sell_order_id": "2"
    }
    """
    And expect query balance with address "regen1q5m97jdcksj24g9enlkjqq75ygt5q6akfm0ycv" and batch denom "C01-001-20200101-20210101-001"
    """
    {
      "balance": {
        "tradable_amount": "0.999998",
        "retired_amount": "0",
        "escrowed_amount": "0.000002"
      }
    }
    """
    And expect query supply with batch denom "C01-001-20200101-20210101-001"
    """
    {
      "tradable_amount": "1",
      "retired_amount": "0",
      "cancelled_amount": "0"
    }
    """
