# Client

## CLI

A user can query and interact with the `ecocredit` module using the CLI.

### Query


The `query` commands allow users to query `ecocredit` state.

For examples on how to query state using CLI, see the ecocredit module [Query commands](https://docs.regen.network/commands/regen_query_ecocredit.html) documentation.


### Transactions

The `tx` commands allow users to interact with the `ecocredit` module.

For examples on how to submit transactions using CLI, see the ecocredit module [Transaction commands](https://docs.regen.network/commands/regen_tx_ecocredit.html) documentation.

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

The `ProjectsByClass` endpoint allows users to query all projects by credit class.

```bash
regen.ecocredit.v1.Query/ProjectsByClass
```

Example:

```bash
grpcurl -plaintext \
    -d '{"class_id":"C01"}' \
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
    -d '{"reference_id":"VCS-001"}' \
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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE".
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
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
    -d '{"admin":"regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"}' \
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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "classId": "C01",
      "jurisdiction": "KE",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
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
    -d '{"project_id":"C01-001"}' \
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
    "jurisdiction": "CD-MN",
    "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
    "reference_id": "VCS-001"
  }
}
```


### Batches

The `Batches` endpoint allows users to query for all batches.

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
    -d '{"account":"regen1qwa9xy0997j5mrc4dxn7jrcvvkpm3uwuldkrmg"}' \
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
    -d '{"basket_denom": "eco.uC.rNCT"}' \
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
      "curator": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"
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
      "curator": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"
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
      "batchDenom": "C01-001-20190101-20210101-008",
      "balance": "1",
      "batchStartDate": "2019-01-01T00:00:00Z"
    },
    {
      "basketId": "1",
      "batchDenom": "C02-001-20210909-20220101-002",
      "balance": "1",
      "batchStartDate": "2021-09-09T00:00:00Z"
    }
  ],
  "balancesInfo":[
    {
      "batchDenom": "C02-001-20210909-20220101-002",
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
    "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
    "batchDenom": "C01-001-20200101-20210101-001",
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
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batchDenom": "C01-001-20200101-20210101-001",
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
    -d '{"batch_denom": "C01-001-20200101-20210101-001"}' \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/SellOrdersByBatch
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "id": "1",
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batchDenom": "C01-001-20200101-20210101-001",
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
    -d '{"seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"}' \
    localhost:9090 \
    regen.ecocredit.marketplace.v1.Query/SellOrdersBySeller
```

Example Output:

```bash
{
  "sellOrders": [
    {
      "id": "1",
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batchDenom": "C01-001-20200101-20210101-001",
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

### classes-by-admin

The `classes-by-admin` endpoint allows users to query credit classes by admin.

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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    }
  ]
}
```

### projects-by-reference-id

The `projects-by-reference-id` endpoint allows users to query all projects by reference-id.

```bash
/rege/ecocredit/v1/projects-by-reference-id/{reference_id}
/rege/ecocredit/v1/projects/projects/reference-id/{reference_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/projects-by-reference-id/VCS-001
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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "KE",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    }
  ],
  "pagination": {
    "next_key": null,
    "total: 2
  }
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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "KE",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    }
  ],
  "pagination": {
    "next_key": null,
    "total: "2"
  }
}
```

### projects-by-admin

The `projects-by-admin` endpoint allows users to query all projects by project admin.

```bash
/rege/ecocredit/v1/projects-by-admin/{admin}
/rege/ecocredit/v1/projects/projects/admin/{admin}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/projects-by-admin/regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn
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
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    },
    {
      "id": "C01-002",
      "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "class_id": "C01",
      "jurisdiction": "KE",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

### project

The `project` endpoint allows users to query for information on a project.

```bash
/rege/ecocredit/v1/project/{project_id}
/rege/ecocredit/v1/projects/{project_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/projects/C01-001
```

Example Output:

```bash
{
  "project": {
    "id": "C01-001",
    "admin": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
    "class_id": "C01",
    "jurisdiction": "CD-MN",
      "metadata": "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf",
      "reference_id": "VCS-001"
  }
}
```

### batches

The `batches` endpoint allows users to query for all batches.

```bash
/regen/ecocredit/v1/batches
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:25Z",
      "open": false
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:31Z",
      "open": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### batches-by-issuer

The `batches-by-issuer` endpoint allows users to query for credit batches by issuer.

```bash
/regen/ecocredit/v1/batches-by-issuer/{isser}
/regen/ecocredit/v1/batches/issuer/{isser}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches/issuer/regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:25Z",
      "open": false
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:31Z",
      "open": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### batches-by-class

The `batches-by-class` endpoint allows users to query for credit batches issued from a given class.

```bash
/regen/ecocredit/v1/batches-by-class/{class_id}
/regen/ecocredit/v1/batches/class/{class_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches/class/C01
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:25Z",
      "open": false
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:31Z",
      "open": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### batches-by-project

The `batches-by-project` endpoint allows users to query for credit batches issued from a given project.

```bash
/regen/ecocredit/v1/batches-by-project/{project_id}
/regen/ecocredit/v1/batches/project/{project_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches/project/C01-001
```

Example Output:

```bash
{
  "batches": [
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-001",
      "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:25Z",
      "open": false
    },
    {
      "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
      "project_id": "C01-001",
      "denom": "C01-001-20150101-20151231-002",
      "metadata": "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
      "start_date": "2015-01-01T00:00:00Z",
      "end_date": "2015-12-31T00:00:00Z",
      "issuance_date": "2022-05-06T01:33:31Z",
      "open": false
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### batch

The `batch` endpoint allows users to query for information on a credit batch.

```bash
/regen/ecocredit/v1/batches/{batch_denom}
/regen/ecocredit/v1/batche/{batch_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-002
```

Example Output:

```bash
{
  "batch": {
    "issuer": "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
    "project_id": "C01-001",
    "denom": "C01-001-20150101-20151231-001",
    "metadata": "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
    "start_date": "2015-01-01T00:00:00Z",
    "end_date": "2015-12-31T00:00:00Z",
    "issuance_date": "2022-05-06T01:33:25Z",
    "open": false
  }
}
```

### balance

The `balance` endpoint allows users to query the balance (tradable, retired and escrowed) of a given credit batch for a given account.

```bash
/regen/ecocredit/v1/balance/{batch_denom}/{address}
/regen/ecocredit/v1/batches/{batch_denom}/balance/{address}
/regen/ecocredit/v1/balances/{address}/batch/{batch_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/balance/C01-001-20150101-20151231-001/regen1...
```

Example Output:

```bash
{
  "address": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
  "batch_denom": "C01-001-20150101-20151231-001",
  "tradable_amount": "10",
  "retired_amount": "30",
  "escrowed_amount": "10"
}
```

### supply

The `supply` endpoint allows users to query the tradable and retired supply of a credit batch.

```bash
/regen/ecocredit/v1/supply/{batch_denom}
/regen/ecocredit/v1/batches/{batch_denom}/supply
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-001/supply
```

Example Output:

```bash
{
  "tradable_amount": "20",
  "retired_amount": "30",
  "cancelled_amount": "30"
}
```

### credit-types

The `credit-types` endpoint allows users to query the list of allowed types that credit classes can have.

```bash
/regen/ecocredit/v1/credit-types
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/credit-types
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

### params

The `params` endpoint allows users to query ecocredit module params.

```bash
/regen/ecocredit/v1/params
```

Example:

```bash
curl localhost:1317/regen/ecocredit/v1/params
```

Example Output:

```bash
{
  "params": {
    "credit_class_fee": [
      {
        "denom": "uregen",
        "amount": "20000000"
      }
    ],
    "basket_fee": [
      {
        "denom": "uregen",
        "amount": "1000000000"
      }
    ],
    "allowed_class_creators": [
      "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm"
    ],
    "allowlist_enabled": false
  }
}
```



### Basket

The `Basket` endpoint allows users to query for information on basket.

```bash
/regen/ecocredit/basket/v1/basket/{basket_denom}
/regen/ecocredit/basket/v1/baskets/{basket_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/basket/v1/basket/eco.uC.rNCT
```

Example Output:

```bash
{
  "basket": {
      "id": "1",
      "basket_denom": "eco.uC.rNCT",
      "name": "rNCT",
      "disable_auto_retire": false,
      "credit_type_abbrev": "C",
      "date_criteria": {
        "min_start_date": null,
        "start_date_window": "315576000s"
      },
      "exponent": 6
  },
  "basket_info": {
      "basket_denom": "eco.uC.rNCT",
      "name": "rNCT",
      "disable_auto_retire": false,
      "credit_type_abbrev": "C",
      "date_criteria": {
        "min_start_date": null,
        "start_date_window": "315576000s"
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
/regen/ecocredit/basket/v1/baskets
```

Example:

```bash
curl localhost:1317/regen/ecocredit/basket/v1/baskets
```

Example Output:

```bash
{
  "baskets": [
    {
      "id": "1",
      "basket_denom": "eco.uC.rNCT",
      "name": "rNCT",
      "disable_auto_retire": false,
      "credit_type_abbrev": "C",
      "date_criteria": {
        "min_start_date": null,
        "start_date_window": "315576000s"
      },
      "exponent": 6
    }
  ],
  "baskets_info": [
    {
      "basket_denom": "eco.uC.rNCT",
      "name": "rNCT",
      "disable_auto_retire": false,
      "credit_type_abbrev": "C",
      "date_criteria": {
        "min_start_date": null,
        "start_date_window": "315576000s"
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
/regen/ecocredit/basket/v1/basket-balance/{basket_denom}/{batch_denom}
/regen/ecocredit/basket/v1/baskets/{basket_denom}/balances/{batch_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/basket/v1/basket-balance/eco.uC.rNCT/C01-001-20150101-20151231-001
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
/regen/ecocredit/basket/v1/basket-balances/{basket_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/basket/v1/basket-balances/eco.uC.rNCT
```

Example Output:

```bash
{
  "balances": [
    {
      "basket_id": "1",
      "batch_denom": "C01-001-20190101-20210101-008",
      "balance": "1",
      "batch_start_date": "2019-01-01T00:00:00Z"
    },
    {
      "basket_id": "1",
      "batch_denom": "C02-001-20210909-20220101-002",
      "balance": "1",
      "batch_start_date": "2021-09-09T00:00:00Z"
    }
  ],
  "pagination": null,
  "balances_info":[
    {
      "batch_denom": "C02-001-20210909-20220101-002",
      "balance": "1"
    },
    {
      "batch_denom": "C01-001-20190101-20210101-008",
      "balance": "1"
    }
  ]
}
```

### sell-orders

The `sell-orders` endpoint allows users to query all sell orders.

```bash
/regen/ecocredit/marketplace/v1/sell-orders
```

Example:

```bash
curl localhost:1317/regen/ecocredit/marketplace/v1/sell-orders
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "id": "1",
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batch_denom": "C01-001-20200101-20210101-001",
      "quantity": "2",
      "ask_denom": "stake",
      "ask_amount": "100",
      "disable_auto_retire": false,
      "expiration": null
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### sell-orders/seller

The `sell-orders/seller` endpoint allows users to query for all sell orders by owner address.

```bash
/regen/ecocredit/marketplace/v1/sell-orders/seller/{seller}
/regen/ecocredit/marketplace/v1/sell-orders-by-seller/{seller}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/marketplace/v1/sell-orders/seller/regen1....
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "id": "1",
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batch_denom": "C01-001-20200101-20210101-001",
      "quantity": "2",
      "ask_denom": "stake",
      "ask_amount": "100",
      "disable_auto_retire": false,
      "expiration": null
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
/regen/ecocredit/marketplace/v1/sell-orders/batch/{batch_denom}
/regen/ecocredit/marketplace/v1/sell-orders-by-batch/{batch_denom}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/marketplace/v1/sell-orders-by-batch/C01-001-20200101-20210101-001
```

Example Output:

```bash
{
  "sell_orders": [
    {
      "id": "1",
      "seller": "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
      "batch_denom": "C01-001-20200101-20210101-001",
      "quantity": "2",
      "ask_denom": "stake",
      "ask_amount": "100",
      "disable_auto_retire": false,
      "expiration": null
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### sell-order

The `sell-order` endpoint allows users to query for information on a sell order by sell order id.

```bash
/regen/ecocredit/marketplace/v1/sell-orders/{sell_order_id}
/regen/ecocredit/marketplace/v1/sell-order/{sell_order_id}
```

Example:

```bash
curl localhost:1317/regen/ecocredit/marketplace/v1/sell-orders/2
```

Example Output:

```bash
{
  "sell_order": {
    "id": "2",
    "seller": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
    "batch_denom": "C03-001-20200101-20210101-001",
    "quantity": "100",
    "ask_denom": "uregen",
    "ask_amount": "1000000",
    "disable_auto_retire": false,
    "expiration": "2023-12-31T00:00:00Z"
  }
}
```


### allowed-denoms

The `allowed-denoms` endpoint allows users to query all bank denoms allowed to be
used in the marketplace.

```bash
/regen/ecocredit/marketplace/v1/allowed-denoms
```

Example:

```bash
curl localhost:1317/regen/ecocredit/marketplace/v1/allowed-denoms
```

Example Output:

```bash
{
  "allowed_denoms": [
    {
      "bank_denom": "uregen",
      "display_denom": "regen",
      "exponent": 6
    }
  ],
  "pagination": null
}
```
