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
  class_id: C01
  end_date: "2021-01-01T00:00:00Z"
  issuer: regen1..
  metadata: cmVnZW4=
  project_location: AA-BB 12345
  start_date: "2020-01-01T00:00:00Z"
  total_amount: "50"

```

#### batches

The `batches` command allows users to query all credit batches for a given credit class.

```bash
regen query ecocredit batches [class_id] [flags]
```

Example:

```bash
regen query ecocredit batches C01
```

Example Output:

```bash
batches:
- amount_cancelled: "10"
  batch_denom: C01-20200101-20210101-001
  class_id: C01
  end_date: "2021-01-01T00:00:00Z"
  issuer: regen1..
  metadata: cmVnZW4=
  project_location: AA-BB 12345
  start_date: "2020-01-01T00:00:00Z"
  total_amount: "50"
pagination:
  next_key: null
  total: "0"
```

#### class-info

The `class-info` command allows users to query information for a given credit class.

```bash
regen query ecocredit class-info
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

#### gen-batch-json

The `gen-batch-json` command allows users to generate JSON to represent a new credit batch for use with the `create-batch` command.

```bash
regen tx ecocredit gen-batch-json [flags]
```

Example:

```bash
regen tx ecocredit gen-batch-json --class-id C01 --issuances 1 --start-date 2020-01-01 --end-date 2021-01-01 --project-location 'AA-BB 12345' --metadata cmVnZW4=
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
    "end_date": "2021-01-01T00:00:00Z",
    "project_location": "AA-BB 12345"
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