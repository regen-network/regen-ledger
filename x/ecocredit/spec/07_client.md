# Client

## CLI

A user can query and interact with the `ecocredit` module using the CLI.

### Query

The `query` commands allow users to query `ecocredit` state.

```bash
regen query ecocredit --help
```

#### balance

The `balance` command allows users to query the tradable and retired balances of a given credit batch.

```bash
regen query ecocredit balance [batch_denom] [account] [flags]
```

Example:

```bash
regen query ecocredit balance C01-20200101-20210101-001 regen1..
```

Example Output:

```bash
retired_amount: "20"
tradable_amount: "30"
```

#### batch-info

The `batch-info` command allows users to query information for a given credit batch.

```bash
regen query ecocredit batch-info [batch_denom] [flags]
```

Example:

```bash
regen query ecocredit batch-info C01-20200101-20210101-001
```

Example Output:

```bash
info:
  amount_cancelled: "10"
  batch_denom: C01-20200101-20210101-001
  end_date: "2021-01-01T00:00:00Z"
  metadata: cmVnZW4=
  project_id: P01
  start_date: "2020-01-01T00:00:00Z"
  total_amount: "3.0"
```

#### batches

The `batches` command allows users to query all credit batches for a given project.

```bash
regen query ecocredit batches [project_id] [flags]
```

Example:

```bash
regen query ecocredit batches P01
```

Example Output:

```bash
batches:
- amount_cancelled: "10"
  batch_denom: C01-20200101-20210101-001
  end_date: "2021-01-01T00:00:00Z"
  metadata: cmVnZW4=
  project_id: P01
  start_date: "2020-01-01T00:00:00Z"
  total_amount: "3.0"
pagination:
  next_key: null
  total: "0"
```

#### class-info

The `class-info` command allows users to query information for a given credit class.

```bash
regen query ecocredit class-info [flags]
```

Example:

```bash
regen query ecocredit class-info C01
```

Example Output:

```bash
info:
  admin: regen1..
  class_id: C01
  credit_type:
    abbreviation: C
    name: carbon
    precision: 6
    unit: metric ton CO2 equivalent
  issuers:
  - regen1..
  metadata: cmVnZW4=
  num_batches: "1"
```

#### classes

The `classes` command allows users to query all credit classes.

```bash
regen query ecocredit classes [flags]
```

Example:

```bash
regen query ecocredit classes
```

Example Output:

```bash
classes:
- admin: regen1..
  class_id: C01
  credit_type:
    abbreviation: C
    name: carbon
    precision: 6
    unit: metric ton CO2 equivalent
  issuers:
  - regen1..
  metadata: cmVnZW4=
  num_batches: "0"
- admin: regen1..
  class_id: C02
  credit_type:
    abbreviation: C
    name: carbon
    precision: 6
    unit: metric ton CO2 equivalent
  issuers:
  - regen1..
  metadata: cmVnZW4=
  num_batches: "0"
pagination:
  next_key: null
  total: "0"
```

#### project-info

The `project-info` command allows users to query information for a given project.

```bash
regen query ecocredit project-info [project_id] [flags]
```

Example:

```bash
regen query ecocredit project-info P01
```

Example Output:

```bash
info:
  class_id: C01
  issuer: regen1..
  metadata: cmVnZW4=
  project_id: P01
  project_location: YZ
```

#### projects

The `projects` command allows users to query all projects for a given credit class.

```bash
regen query ecocredit projects [class_id] [flags]
```

Example:

```bash
regen query ecocredit projects C01
```

Example Output:

```bash
pagination:
  next_key: null
  total: "0"
projects:
- class_id: C01
  issuer: regen1..
  metadata: cmVnZW4=
  project_id: P01
  project_location: YZ
```

#### supply

The `supply` command allows users to query the tradable and retired supply of a given credit batch.

```bash
regen query ecocredit supply [batch_denom] [flags]
```

Example:

```bash
regen query ecocredit supply C01-20200101-20210101-001
```

Example Output:

```bash
retired_supply: "20"
tradable_supply: "30"
```

#### types

The `types` command allows users to query the list of approved credit types.

```bash
regen query ecocredit types [flags]
```

Example:

```bash
regen query ecocredit types
```

Example Output:

```bash
credit_types:
- abbreviation: C
  name: carbon
  precision: 6
  unit: metric ton CO2 equivalent
```

#### sell-order

The `sell-order` command allows users to query information for a given sell order.

```bash
regen query ecocredit sell-order [flags]
```

Example:

```bash
regen query ecocredit sell-order 1
```

Example Output:

```bash
sell_order:
  ask_price:
    amount: "100"
    denom: stake
  batch_denom: C01-20200101-20210101-001
  disable_auto_retire: false
  order_id: "1"
  owner: regen1..
  quantity: "2"
```

#### sell-orders

The `sell-orders` command allows users to query all sell orders.

```bash
regen query ecocredit sell-orders [flags]
```

Example:

```bash
regen query ecocredit sell-orders
```

Example Output:

```bash
pagination:
  next_key: null
  total: "1"
sell_orders:
- ask_price:
    amount: "100"
    denom: stake
  batch_denom: C01-20200101-20210101-001
  disable_auto_retire: false
  order_id: "1"
  owner: regen1..
  quantity: "2"
```

#### sell-orders-by-address

The `sell-orders-by-address` command allows users to query sell orders by owner address.

```bash
regen query ecocredit sell-orders-by-address [address] [flags]
```

Example:

```bash
regen query ecocredit sell-orders-by-address regen1..
```

Example Output:

```bash
pagination:
  next_key: null
  total: "1"
sell_orders:
- ask_price:
    amount: "100"
    denom: stake
  batch_denom: C01-20200101-20210101-001
  disable_auto_retire: false
  order_id: "1"
  owner: regen1..
  quantity: "2"
```

#### sell-orders-by-batch-denom

The `sell-orders-by-batch-denom` command allows users to query sell orders by credit batch denom.

```bash
regen query ecocredit sell-orders-by-batch-denom [batch_denom] [flags]
```

Example:

```bash
regen query ecocredit sell-orders-by-batch-denom C01-20200101-20210101-001
```

Example Output:

```bash
pagination:
  next_key: null
  total: "1"
sell_orders:
- ask_price:
    amount: "100"
    denom: stake
  batch_denom: C01-20200101-20210101-001
  disable_auto_retire: false
  order_id: "1"
  owner: regen1..
  quantity: "2"
```

#### buy-order

The `buy-order` command allows users to query information for a given buy order.

```bash
regen query ecocredit buy-order [flags]
```

Example:

```bash
regen query ecocredit buy-order 1
```

Example Output:

```bash
# not yet implemented
```

#### buy-orders

The `buy-orders` command allows users to query all buy orders.

```bash
regen query ecocredit buy-orders [flags]
```

Example:

```bash
regen query ecocredit buy-orders
```

Example Output:

```bash
# not yet implemented
```

#### buy-orders-by-address

The `buy-orders-by-address` command allows users to query buy orders by buyer address.

```bash
regen query ecocredit buy-orders-by-address [address] [flags]
```

Example:

```bash
regen query ecocredit buy-orders-by-address regen1..
```

Example Output:

```bash
# not yet implemented
```

#### ask-denoms

The `ask-denoms` command allows users to query the list of allowed ask denoms.

```bash
regen query ecocredit ask-denoms [flags]
```

Example:

```bash
regen query ecocredit ask-denoms
```

Example Output:

```bash
# not yet implemented
```

### Transactions

The `tx` commands allow users to interact with the `ecocredit` module.

```bash
regen tx ecocredit --help
```

#### cancel

The `cancel` command allows users to cancel a specified amount of credits.

```bash
regen tx ecocredit cancel [credits] [flags]
```

Example:

```bash
regen tx ecocredit cancel '10 C01-20200101-20210101-001, 0.1 C01-20200101-20210101-002' --from regen1..
```

#### create-batch

The `create-batch` command allows users to issue a new credit batch.

```bash
regen tx ecocredit create-batch [msg-create-batch-json-file] [flags]
```

Example:

```bash
regen tx ecocredit create-batch batch.json --from regen1..
```

#### create-class

The `create-class` command allows users to create a new credit class.

```bash
regen tx ecocredit create-class [issuer[,issuer]*] [credit type name] [metadata] [flags]
```

Example:

```bash
regen tx ecocredit create-class regen1.. carbon cmVnZW4= --from regen1..
```

#### create-project

The `create-project` command allows users to create a new project.

```bash
regen tx ecocredit create-project [class-id] [project-location] [metadata] [flags]
```

Example:

```bash
regen tx ecocredit create-project C01 YZ cmVnZW4= --project-id P01 --from regen1..
```

#### gen-batch-json

The `gen-batch-json` command allows users to generate JSON to represent a new credit batch for use with the `create-batch` command.

```bash
regen tx ecocredit gen-batch-json [flags]
```

Example:

```bash
regen tx ecocredit gen-batch-json --project-id P01 --issuances 1 --start-date 2020-01-01 --end-date 2021-01-01 --metadata cmVnZW4=
```

Example Output:

```bash
{
    "issuer": "",
    "class_id": "C01",
    "issuance": [
        {
            "recipient": "recipient-address",
            "tradable_amount": "tradable-amount",
            "retired_amount": "retired-amount",
            "retirement_location": "retirement-location"
        }
    ],
    "metadata": "cmVnZW4=",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z"
}
```

#### retire

The `retire` command allows users to retire a specified amount of credits.

```bash
regen tx ecocredit retire [credits] [retirement_location] [flags]
```

Example:

```bash
regen tx ecocredit retire '[{batch_denom: "C01-20200101-20210101-001", amount: "10"}]' 'AA-BB 12345' --from regen1..
```

#### send

The `send` command allows users to send credits.

```bash
regen tx ecocredit send [recipient] [credits] [flags]
```

Example:

```bash
regen tx ecocredit send regen1.. '[{batch_denom: "C01-20200101-20210101-001", tradable_amount: "10", retired_amount: "0","retirement_location":"AA-BB 12345"}]' --from regen1..
```

#### update-class-admin

The `update-class-admin` command allows users to update the admin of a credit class.

```bash
regen tx ecocredit update-class-admin [class-id] [admin] [flags]
```

Example:

```bash
regen tx ecocredit update-class-admin C01 regen1.. --from regen1..
```

#### update-class-issuers

The `update-class-issuers` command allows users to update the issuers of a credit class.

```bash
regen tx ecocredit update-class-issuers [class-id] [issuers] [flags]
```

Example:

```bash
regen tx ecocredit update-class-issuers C01 regen1.. --from regen1..
```

#### update-class-metadata

The `update-class-metadata` command allows users to update the metadata of a credit class.

```bash
regen tx ecocredit update-class-metadata [class-id] [metadata] [flags]
```

Example:

```bash
regen tx ecocredit update-class-metadata C01 cmVnZW4= --from regen1..
```

#### sell

The `sell` command allows users to create new sell orders.

```bash
regen tx ecocredit sell [orders] [flags]
```

Example:

```bash
regen tx ecocredit sell '[{batch_denom: "C01-20200101-20210101-001", quantity: "2", ask_price: "100stake", disable_auto_retire: false}]' --from regen1
```

#### update-sell-order

The `update-sell-order` command allows users to update a given sell order.

```bash
regen tx ecocredit update-sell-order [updates] [flags]
```

Example:

```bash
regen tx ecocredit update-sell-orders '[{sell_order_id: 1, new_quantity: "2", new_ask_price: "200stake", disable_auto_retire: false}]' --from regen1
```

#### buy

The `buy` command allows users to create new buy orders.

```bash
regen tx ecocredit buy [orders] [flags]
```

Example:

```bash
regen tx ecocredit buy '[{sell_order_id: "1", quantity: "2", bid_price: "100regen", disable_auto_retire: false}]' --from regen1..
```

## gRPC

A user can query the `ecocredit` module using gRPC endpoints.

### Classes

The `Classes` endpoint allows users to query all credit classes.

```bash
regen.ecocredit.v1alpha1.Query/Classes
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/Classes
```

Example Output:

```bash
{
  "classes": [
    {
      "classId": "C01",
      "admin": "regen1..",
      "issuers": [
        "regen1.."
      ],
      "metadata": "cmVnZW4=",
      "creditType": {
        "name": "carbon",
        "abbreviation": "C",
        "unit": "metric ton CO2 equivalent",
        "precision": 6
      },
      "numBatches": "1"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### ClassInfo

The `ClassInfo` endpoint allows users to query for information on a credit class.

```bash
regen.ecocredit.v1alpha1.Query/ClassInfo
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/ClassInfo
```

Example Output:

```bash
{
  "info": {
    "classId": "C01",
    "admin": "regen1..",
    "issuers": [
      "regen1.."
    ],
    "metadata": "cmVnZW4=",
    "creditType": {
      "name": "carbon",
      "abbreviation": "C",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    },
    "numBatches": "1"
  }
}
```

### Batches

The `Batches` endpoint allows users to query for all batches in the given credit class.

```bash
regen.ecocredit.v1alpha1.Query/Batches
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/Batches
```

Example Output:

```bash
{
  "info": {
    "classId": "C01",
    "admin": "regen1..",
    "issuers": [
      "regen1.."
    ],
    "metadata": "cmVnZW4=",
    "creditType": {
      "name": "carbon",
      "abbreviation": "C",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    },
    "numBatches": "1"
  }
}
```

### BatchInfo

The `BatchInfo` endpoint allows users to query for information on a credit batch.

```bash
regen.ecocredit.v1alpha1.Query/BatchInfo
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/BatchInfo
```

Example Output:

```bash
{
  "info": {
    "classId": "C01",
    "batchDenom": "C01-20200101-20210101-001",
    "issuer": "regen1..",
    "totalAmount": "50",
    "metadata": "cmVnZW4=",
    "amountCancelled": "10",
    "startDate": "2020-01-01T00:00:00Z",
    "endDate": "2021-01-01T00:00:00Z",
    "projectLocation": "AA-BB 12345"
  }
}
```

### Balance

The `Balance` endpoint allows users to query the balance (both tradable and retired) of a given credit batch for a given account.

```bash
regen.ecocredit.v1alpha1.Query/Balance
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001", "account":"regen1.."}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/Balance
```

Example Output:

```bash
{
  "tradableAmount": "10",
  "retiredAmount": "30"
}
```

### Supply

The `Supply` endpoint allows users to query the tradable and retired supply of a credit batch.

```bash
regen.ecocredit.v1alpha1.Query/Supply
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/Supply
```

Example Output:

```bash
{
  "tradableSupply": "20",
  "retiredSupply": "30"
}
```

### CreditTypes

The `CreditTypes` endpoint allows users to query the list of allowed types that credit classes can have.

```bash
regen.ecocredit.v1alpha1.Query/CreditTypes
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/CreditTypes
```

Example Output:

```bash
{
  "creditTypes": [
    {
      "name": "carbon",
      "abbreviation": "C",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    }
  ]
}
```

### SellOrder

The `SellOrder` endpoint allows users to query for information on a sell order.

```bash
regen.ecocredit.v1alpha1.Query/SellOrder
```

Example:

```bash
grpcurl -plaintext \
    -d '{"sell_order_id": "1"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/SellOrder
```

Example Output:

```bash
{
  "sellOrder": {
    "orderId": "1",
    "owner": "regen1..",
    "batchDenom": "C01-20200101-20210101-001",
    "quantity": "2",
    "askPrice": {
      "denom": "stake",
      "amount": "100"
    }
  }
}
```

### SellOrders

The `SellOrders` endpoint allows users to query all sell orders.

```bash
regen.ecocredit.v1alpha1.Query/SellOrders
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/SellOrders
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "orderId": "1",
      "owner": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askPrice": {
        "denom": "stake",
        "amount": "100"
      }
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### SellOrdersByAddress

The `SellOrdersByAddress` endpoint allows users to query sell orders by owner address.

```bash
regen.ecocredit.v1alpha1.Query/SellOrdersByAddress
```

Example:

```bash
grpcurl -plaintext \
    -d '{"address": "regen1.."}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/SellOrdersByAddress
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "orderId": "1",
      "owner": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askPrice": {
        "denom": "stake",
        "amount": "100"
      }
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### SellOrdersByBatchDenom

The `SellOrdersByBatchDenom` endpoint allows users to query sell orders by credit batch denom.

```bash
regen.ecocredit.v1alpha1.Query/SellOrdersByBatchDenom
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/SellOrdersByBatchDenom
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "orderId": "1",
      "owner": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askPrice": {
        "denom": "stake",
        "amount": "100"
      }
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### BuyOrder

The `BuyOrder` endpoint allows users to query for information on a buy order.

```bash
regen.ecocredit.v1alpha1.Query/BuyOrder
```

Example:

```bash
grpcurl -plaintext \
    -d '{"buy_order_id": "1"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/BuyOrder
```

Example Output:

```bash
# not yet implemented
```

### BuyOrders

The `BuyOrders` endpoint allows users to query all buy orders.

```bash
regen.ecocredit.v1alpha1.Query/BuyOrders
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/BuyOrders
```

Example Output:

```bash
# not yet implemented
```

### BuyOrdersByAddress

The `BuyOrdersByAddress` endpoint allows users to query buy orders by buyer address.

```bash
regen.ecocredit.v1alpha1.Query/BuyOrdersByAddress
```

Example:

```bash
grpcurl -plaintext \
    -d '{"address": "regen1.."}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/BuyOrdersByAddress
```

Example Output:

```bash
# not yet implemented
```

### AllowedAskDenoms

The `AllowedAskDenoms` endpoint allows users to query all allowed ask denoms.

```bash
regen.ecocredit.v1alpha1.Query/AllowedAskDenoms
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/AllowedAskDenoms
```

Example Output:

```bash
# not yet implemented
```

## REST

A user can query the `ecocredit` module using REST endpoints.

### classes

The `classes` endpoint allows users to query all credit classes.

```bash
/regen/ecocredit/v1alpha1/classes
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/classes
```

Example Output:

```bash
{
  "classes": [
    {
      "class_id": "C01",
      "admin": "regen1..",
      "issuers": [
        "regen1.."
      ],
      "metadata": "cmVnZW4=",
      "credit_type": {
        "name": "carbon",
        "abbreviation": "C",
        "unit": "metric ton CO2 equivalent",
        "precision": 6
      },
      "num_batches": "1"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### classes

The `classes` endpoint allows users to query for information on a credit class.

```bash
/regen/ecocredit/v1alpha1/classes/{class_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/classes/C01
```

Example Output:

```bash
{
  "info": {
    "class_id": "C01",
    "admin": "regen1..",
    "issuers": [
      "regen1.."
    ],
    "metadata": "cmVnZW4=",
    "credit_type": {
      "name": "carbon",
      "abbreviation": "C",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    },
    "num_batches": "1"
  }
}
```

### batches

The `batches` endpoint allows users to query for all batches in the given credit class.

```bash
/regen/ecocredit/v1alpha1/batches
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/batches?class_id=C01
```

Example Output:

```bash
{
  "batches": [
    {
      "class_id": "C01",
      "batch_denom": "C01-20200101-20210101-001",
      "issuer": "regen1..",
      "total_amount": "50",
      "metadata": "cmVnZW4=",
      "amount_cancelled": "10",
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "project_location": "AA-BB 12345"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### batches

The `batches` endpoint allows users to query for information on a credit batch.

```bash
/regen/ecocredit/v1alpha1/batches/{batch_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/batches/C01-20200101-20210101-001
```

Example Output:

```bash
{
  "info": {
    "class_id": "C01",
    "batch_denom": "C01-20200101-20210101-001",
    "issuer": "regen1..",
    "total_amount": "50",
    "metadata": "cmVnZW4=",
    "amount_cancelled": "10",
    "start_date": "2020-01-01T00:00:00Z",
    "end_date": "2021-01-01T00:00:00Z",
    "project_location": "AA-BB 12345"
  }
}
```

### balance

The `balance` endpoint allows users to query the balance (both tradable and retired) of a given credit batch for a given account.

```bash
/regen/ecocredit/v1alpha1/batches/{batch_denom}/balance/{account}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/batches/C01-20200101-20210101-001/balance/regen1..
```

Example Output:

```bash
{
  "tradable_amount": "10",
  "retired_amount": "30"
}
```

### supply

The `supply` endpoint allows users to query the tradable and retired supply of a credit batch.

```bash
/regen/ecocredit/v1alpha1/batches/{batch_denom}/supply
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/batches/C01-20200101-20210101-001/supply
```

Example Output:

```bash
{
  "tradable_supply": "20",
  "retired_supply": "30"
}
```

### credit-types

The `credit-types` endpoint allows users to query the list of allowed types that credit classes can have.

```bash
/regen/ecocredit/v1alpha1/credit-types
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/credit-types
```

Example Output:

```bash
{
  "credit_types": [
    {
      "name": "carbon",
      "abbreviation": "C",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    }
  ]
}
```

### sell-orders

The `sell-orders` endpoint allows users to query all sell orders.

```bash
/regen/ecocredit/v1alpha1/sell-orders
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/sell-orders
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "order_id": "1",
      "owner": "regen1..",
      "batch_denom": "C01-20200101-20210101-001",
      "quantity": "2",
      "ask_price": {
        "denom": "stake",
        "amount": "100"
      },
      "disable_auto_retire": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### sell-orders/address

The `sell-orders/address` endpoint allows users to query for all sell orders by owner address.

```bash
/regen/ecocredit/v1alpha1/sell-orders/address/{address}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/sell-orders/address/regen1..
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "order_id": "1",
      "owner": "regen1..",
      "batch_denom": "C01-20200101-20210101-001",
      "quantity": "2",
      "ask_price": {
        "denom": "stake",
        "amount": "100"
      },
      "disable_auto_retire": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### sell-orders/batch-denom

The `sell-orders/batch-denom` endpoint allows users to query for all sell orders by credit batch denom.

```bash
/regen/ecocredit/v1alpha1/sell-orders/batch-denom/{batch-denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/sell-orders/batch-denom/C01-20200101-20210101-001
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "order_id": "1",
      "owner": "regen1..",
      "batch_denom": "C01-20200101-20210101-001",
      "quantity": "2",
      "ask_price": {
        "denom": "stake",
        "amount": "100"
      },
      "disable_auto_retire": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### sell-orders/id

The `sell-orders/id` endpoint allows users to query for information on a sell order by sell order id.

```bash
/regen/ecocredit/v1alpha1/sell-orders/id/{sell_order_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/sell-orders/id/1
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "order_id": "1",
      "owner": "regen1..",
      "batch_denom": "C01-20200101-20210101-001",
      "quantity": "2",
      "ask_price": {
        "denom": "stake",
        "amount": "100"
      },
      "disable_auto_retire": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### buy-orders

The `buy-orders` endpoint allows users to query all buy orders.

```bash
/regen/ecocredit/v1alpha1/buy-orders
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/buy-orders
```

Example Output:

```bash
# not yet implemented
```

### buy-orders/address

The `buy-orders/address` endpoint allows users to query for all buy orders by buyer address.

```bash
/regen/ecocredit/v1alpha1/buy-orders/address/{address}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/buy-orders/address/regen1..
```

Example Output:

```bash
# not yet implemented
```

### buy-orders/id

The `buy-orders/id` endpoint allows users to query for information on a buy order by buy order id.

```bash
/regen/ecocredit/v1alpha1/buy-orders/id/{buy_order_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/buy-orders/id/1
```

Example Output:

```bash
# not yet implemented
```

### ask-denoms

The `ask-denoms` endpoint allows users to query all allowed ask denoms.

```bash
/regen/ecocredit/v1alpha1/ask-denoms
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1alpha1/ask-denoms
```

Example Output:

```bash
# not yet implemented
```
