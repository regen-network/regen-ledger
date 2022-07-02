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
regen.ecocredit.v1.Query/Classes
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1.Query/Classes
```

Example Output:

```bash
{
  "classes": [
    {
      "id": "C01",
      "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "creditTypeAbbrev": "C"
    }
  ]
}
```

### Class

The `Class` endpoint allows users to query for information on a credit class.

```bash
regen.ecocredit.v1.Query/Class
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/ClassInfo
```

Example Output:

```bash
{
  "class": {
    "id": "C01",
    "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
    "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
    "creditTypeAbbrev": "C"
  }
}
```

### ClassesByAdmin

The `ClassesByAdmin` endpoint allows users to query for all credit classes by admin.

```bash
regen.ecocredit.v1.Query/ClassesByAdmin
```

Example:

```bash
grpcurl -plaintext \
    -d '{"admin":"regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/ClassesByAdmin
```

Example Output:

```bash
{
  "classes": [
    {
      "id": "C01",
      "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "creditTypeAbbrev": "C"
    }
  ]
}
```

### ClassIssuers

The `ClassIssuers` endpoint allows users to query addresses of the issuers for a credit class.


```bash
regen.ecocredit.v1.Query/ClassIssuers
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/ClassIssuers
```

Example Output:

```bash
{
  "issuers": [
    "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"
  ]
}
```

### Projects

The `Projects` endpoint allows users to query all projects.

```bash
regen.ecocredit.v1.Query/Projects
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1.Query/Projects
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "CD-MN"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE"
    }
  ]
}
```

### ProjectsByClass

The `ProjectsByClass` endpoint allows users to query all projects.

```bash
regen.ecocredit.v1.Query/ProjectsByClass
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}'
    localhost:9090 \
    regen.ecocredit.v1.Query/ProjectsByClass
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "CD-MN"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE"
    }
  ]
}
```

### ProjectsByReferenceId

The `ProjectsByReferenceId` endpoint allows users to query projects by reference id.

```bash
regen.ecocredit.v1.Query/ProjectsByReferenceId
```

Example:

```bash
grpcurl -plaintext \
    -d '{"reference_id":"VERRA01"}'
    localhost:9090 \
    regen.ecocredit.v1.Query/ProjectsByReferenceId
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "CD-MN",
      "reference_id":"VERRA01"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE".
      "reference_id":"VERRA01"
    }
  ]
}
```

### ProjectsByAdmin

The `ProjectsByAdmin` endpoint allows users to query projects by admin.

```bash
regen.ecocredit.v1.Query/ProjectsByAdmin
```

Example:

```bash
grpcurl -plaintext \
    -d '{"admin":"regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"}'
    localhost:9090 \
    regen.ecocredit.v1.Query/ProjectsByAdmin
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "CD-MN",
      "reference_id":"VERRA01"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE".
      "reference_id":"VERRA01"
    }
  ]
}
```

### Project

The `Project` endpoint allows users to query for information on a project.

```bash
regen.ecocredit.v1.Query/Project
```

Example:

```bash
grpcurl -plaintext \
    -d '{"project_id":"C01-001"}'
    localhost:9090 \
    regen.ecocredit.v1.Query/Project
```

Example Output:

```bash
{
  "project": {
    "id": "C01-001",
    "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
    "classId": "C01",
    "jurisdiction": "CD-MN"
  }
}
```


### Batches

The `Batches` endpoint allows users to query for all batches in the given credit class.

```bash
regen.ecocredit.v1.Query/Batches
```

Example:

```bash
grpcurl -plaintext \
    -d '{"pagination":{"limit":"2"}}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/Batches
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:25Z"
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:31Z"
    }
  ],
  "pagination": {
    "nextKey": "BwEFAAAC"
  }
}
```

### BatchesByIssuer

The `BatchesByIssuer` endpoint allows users to query for credit batches issued from a given issuer address.

```bash
regen.ecocredit.v1.Query/BatchesByIssuer
```

Example:

```bash
grpcurl -plaintext \
    -d '{"issuer":"regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/BatchesByIssuer
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:25Z"
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:31Z"
    }
  ],
  "pagination": {
    "nextKey": "BwEFAAAC"
  }
}
```

### BatchesByClass

The `BatchesByClass` endpoint allows users to query credit batches issued from a given class.

```bash
regen.ecocredit.v1.Query/BatchesByClass
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/BatchesByClass
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:25Z"
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:31Z"
    }
  ],
  "pagination": {
    "nextKey": "BwEFAAAC"
  }
}
```

### BatchesByProject

The `BatchesByProject` endpoint allows users to query credit batches issued from a given project.

```bash
regen.ecocredit.v1.Query/BatchesByProject
```

Example:

```bash
grpcurl -plaintext \
    -d '{"project_id":"C01-001"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/BatchesByProject
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:25Z"
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "projectId": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "startDate": "2015-01-01T00:00:00Z",
      "endDate": "2015-12-31T00:00:00Z",
      "issuanceDate": "2022-05-06T01:33:31Z"
    }
  ],
  "pagination": {
    "nextKey": "BwEFAAAC"
  }
}
```

### Batch

The `Batch` endpoint allows users to query for information on a credit batch.

```bash
regen.ecocredit.v1.Query/Batch
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-001-20150101-20151231-001"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/Batch
```

Example Output:

```bash
{
  "batch": {
    "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
    "projectId": "C01-001",
    "denom": "C01-001-20150101-20151231-001",
    "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
    "startDate": "2015-01-01T00:00:00Z",
    "endDate": "2015-12-31T00:00:00Z",
    "issuanceDate": "2022-05-06T01:33:25Z"
  }
}
```

### Balance

The `Balance` endpoint allows users to query the balance (tradable, retired and escrowed) of a given credit batch for a given account.

```bash
regen.ecocredit.v1.Query/Balance
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001", "account":"regen1.."}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/Balance
```

Example Output:

```bash
{
  "balance": {
    "address": "regen1qwa9xy0997j5mrc4dxn7jrcvvkpm3uwuldkrmg",
    "batch_denom": "C02-001-20210101-20250101-001",
    "tradable_amount": "92.0",
    "retired_amount": "97",
    "escrowed_amount": "30.0"
  }
}
```

### Balances

The `Balances` endpoint allows users to query the balances (tradable, retired and escrowed) for a given account.

```bash
regen.ecocredit.v1.Query/Balances
```

Example:

```bash
grpcurl -plaintext \
    -d '{"account":"regen1.."}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/Balances
```

Example Output:

```bash
{
  "balances": [
    {
      "address": "regen1qwa9xy0997j5mrc4dxn7jrcvvkpm3uwuldkrmg",
      "batchDenom": "C02-001-20210101-20250101-001",
      "tradableAmount": "92.0",
      "retiredAmount": "97",
      "escrowedAmount": "30.0"
    }
  ],
  "pagination": {
    "nextKey": "BwEJABQDulMR5S+lTY8VaafpDwxlg7jx3AAF"
  }
}
```

### Supply

The `Supply` endpoint allows users to query the tradable and retired supply of a credit batch.

```bash
regen.ecocredit.v1.Query/Supply
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C02-001-20210101-20250101-001"}' \
    localhost:9090 \
    regen.ecocredit.v1.Query/Supply
```

Example Output:

```bash
{
  "tradableSupply": "722",
  "retiredSupply": "158",
  "cancelledAmount": "0"
}
```

### CreditTypes

The `CreditTypes` endpoint allows users to query the list of allowed types that credit classes can have.

```bash
regen.ecocredit.v1.Query/CreditTypes
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.v1.Query/CreditTypes
```

Example Output:

```bash
{
  "creditTypes": [
    {
      "abbreviation": "C",
      "name": "carbon",
      "unit": "metric ton CO2 equivalent",
      "precision": 6
    }
  ]
}
```

### Params

The `Params` endpoint allows users to query ecocredit module params.

```bash
regen.ecocredit.v1.Query/Params
```

Example:

```bash
grpcurl -plaintext localhost:9090 regen.ecocredit.v1.Query/Params
```

Example Output:

```bash
{
  "params": {
    "creditClassFee": [
      {
        "denom": "uregen",
        "amount": "20000000"
      }
    ],
    "basketFee": [
      {
        "denom": "uregen",
        "amount": "1000000000"
      }
    ],
    "allowedClassCreators": [
      "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"
    ]
  }
}
```

### Basket

The `Basket` endpoint allows users to query for information on basket.

```bash
regen.ecocredit.basket.v1.Query/Basket
```

Example:

```bash
grpcurl -plaintext \
    -d '{"basket_denom": "1"}' \
    localhost:9090 \
    regen.ecocredit.basket.v1.Query/Basket
```

Example Output:

```bash
{
  "basket": {
      "id": "1",
      "basketDenom": "eco.uC.rNCT",
      "name": "rNCT",
      "disableAutoRetire": false,
      "creditTypeAbbrev": "C",
      "dateCriteria": {
        "minStartDate": null,
        "startDateWindow": "315576000s"
      },
      "exponent": 6
  },
  "basketInfo": {
      "basketDenom": "eco.uC.rNCT",
      "name": "rNCT",
      "disableAutoRetire": false,
      "creditRypeAbbrev": "C",
      "dateCriteria": {
        "minStartDate": null,
        "startDateWindow": "315576000s"
      },
      "exponent": 6,
      "curator": "regenabc..."
  },
  "classes": ["C01", "C02"]  
}
```

### Baskets

The `Baskets` endpoint allows users to query all basket.

```bash
regen.ecocredit.basket.v1.Query/Baskets
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.basket.v1.Query/Baskets
```

Example Output:

```bash
{
  "baskets": [
    {
      "id": "1",
      "basketDenom": "eco.uC.rNCT",
      "name": "rNCT",
      "disableAutoRetire": false,
      "creditTypeAbbrev": "C",
      "dateCriteria": {
        "minStartDate": null,
        "startDateWindow": "315576000s"
      },
      "exponent": 6
    }
  ],
  "basketsInfo": [
    {
      "basketDenom": "eco.uC.rNCT",
      "name": "rNCT",
      "disableAutoRetire": false,
      "creditTypeAbbrev": "C",
      "dateCriteria": {
        "minStartDate": null,
        "startDateWindow": "315576000s"
      },
      "exponent": 6,
      "curator": "regenabc..."
    }
  ]
}
```

### BasketBalance

The `BasketBalance` endpoint allows users to query the balance of a specific credit batch in the basket.

```bash
regen.ecocredit.basket.v1.Query/BasketBalance
```

Example:

```bash
grpcurl -plaintext \
    -d '{"basket_denom": "eco.uC.rNCT", "batch_denom": "C02-001-20210101-20250101-001"}' \
    localhost:9090 \
    regen.ecocredit.basket.v1.Query/BasketBalance
```

Example Output:

```bash
{
  "balance": "11947.698895"
}
```


### BasketBalances

The `BasketBalances` endpoint allows users to query the balance of each credit batch in the basket.

```bash
regen.ecocredit.basket.v1.Query/BasketBalances
```

Example:

```bash
grpcurl -plaintext \
    -d '{"basket_denom": "eco.uC.rNCT"}' \
    localhost:9090 \
    regen.ecocredit.basket.v1.Query/BasketBalances
```

Example Output:

```bash
{
  "balances": [
    {
      "basketId": "1",
      "batchDenom": "C01-20190101-20210101-008",
      "balance": "1",
      "batchStartDate": "2019-01-01T00:00:00Z"
    },
    {
      "basketId": "1",
      "batchDenom": "C02-20210909-20220101-002",
      "balance": "1",
      "batchStartDate": "2021-09-09T00:00:00Z"
    }
  ],
  "pagination": null,
  "balancesInfo":[
    {
      "batchDenom": "C02-20210909-20220101-002",
      "balance": "1"
    },
    {
      "batchDenom": "C01-20190101-20210101-008",
      "balance": "1"
    }
  ]
}
```




### SellOrder

The `SellOrder` endpoint allows users to query for information on a sell order.

```bash
regen.ecocredit.marketplace.v1.Query/SellOrder
```

Example:

```bash
grpcurl -plaintext \
    -d '{"sell_order_id": "1"}' \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/SellOrder
```

Example Output:

```bash
{
  "sellOrder": {
    "id": "1",
    "seller": "regen1..",
    "batchDenom": "C01-20200101-20210101-001",
    "quantity": "2",
    "askDenom": "uregen",
    "DisableAutoRetire": false,
    "expiration": null
  }
}
```

### SellOrders

The `SellOrders` endpoint allows users to query all sell orders.

```bash
regen.ecocredit.marketplace.v1.Query/SellOrders
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/SellOrders
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "id": "1",
      "seller": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askDenom": "uregen",
      "DisableAutoRetire": false,
      "expiration": null
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### SellOrdersByBatch

The `SellOrdersByBatch` endpoint allows users to query sell orders by batch denom.

```bash
regen.ecocredit.marketplace.v1.Query/SellOrdersByBatch
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.v1alpha1.Query/SellOrdersByBatch
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "id": "1",
      "seller": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askDenom": "uregen",
      "DisableAutoRetire": false,
      "expiration": null
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### SellOrdersBySeller

The `SellOrdersBySeller` endpoint allows users to query sell orders by seller address.

```bash
regen.ecocredit.marketplace.v1.Query/SellOrdersBySeller
```

Example:

```bash
grpcurl -plaintext \
    -d '{"batch_denom": "C01-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/SellOrdersBySeller
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "id": "1",
      "seller": "regen1..",
      "batchDenom": "C01-20200101-20210101-001",
      "quantity": "2",
      "askDenom": "uregen",
      "DisableAutoRetire": false,
      "expiration": null
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### AllowedDenoms

The `AllowedDenoms` endpoint allows users to query all allowed denoms.

```bash
regen.ecocredit.marketplace.v1.Query/AllowedDenoms
```

Example:

```bash
grpcurl -plaintext \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/AllowedDenoms
```

Example Output:

```bash
{
  "allowedDenoms": [
    {
      "bankDenom": "uregen",
      "displayDenom": "regen",
      "exponent": 6
    }
  ]
}
```

## REST

A user can query the `ecocredit` module using REST endpoints.

### classes

The `classes` endpoint allows users to query all credit classes.

```bash
/regen/ecocredit/v1/classes
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/classes
```

Example Output:

```bash
{
  "classes": [
    {
      "id": "C01",
      "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "creditTypeAbbrev": "C"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### classes/admin

The `classes/admin` endpoint allows users to query credit classes by admin.

```bash
/regen/ecocredit/v1/classes/admin/{admin}
/regen/ecocredit/v1/classes-by-admin/{admin}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/classes/admin/regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm
```

Example Output:

```bash
{
  "classes": [
    {
      "id": "C01",
      "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "credit_type_abbrev": "C"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### class

The `class` endpoint allows users to query information on credit class.

```bash
/regen/ecocredit/v1/class/{class_id}
/regen/ecocredit/v1/classes/{class_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/classes/C01
```

Example Output:

```bash
{
  "class": {
      "id": "C01",
      "admin": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "creditTypeAbbrev": "C"
    }
}
```

### class-issuers

The `class-issuers` endpoint allows users to query addresses of the issuers for a credit class.

```bash
/regen/ecocredit/v1/class-issuers/{class_id}
/regen/ecocredit/v1/classes/{class_id}/issuers
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/classes/C01/issuers
```

Example Output:

```bash
{
  "issuers": [
      "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### projects

The `projects` endpoint allows users to query all projects.

```bash
/regen/ecocredit/v1/projects
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/projects
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "CD-MN",
      "metadata": "metadata",
      "reference_id": "R01"
    }
  ]
}
```

### projects-by-class

The `projects-by-class` endpoint allows users to query all projects within a credit class.

```bash
/rege/ecocredit/v1/projects-by-class/{class_id}
/rege/ecocredit/v1/projects/class/{class_id}
/rege/ecocredit/v1/projects/classes/{class_id}/projects
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/projects-by-class/C01
```

Example Output:

```bash
{
  "projects": [
    {
      "id": "C01-001",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "CD-MN",
      "metadata": "",
      "reference_id": ""
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "KE",
      "metadata": "",
      "reference_id": ""
    }
  ],
  "pagination": null
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
