# Concepts

## Ecocredit Core

### Credit Class

A credit class is the primary abstraction for an ecosystem service credit. A credit class is associated with a credit type that represents the primary unit of measurement for the accepted methodologies included within the credit class. Information about the approved methodologies are typically stored off-chain, securely hashed, timestamped, and then included in the metadata field.

A credit class also defines an admin, an address with permission to update the credit class, and a list of approved issuers, addresses with permission to issue credit batches under the credit class.

For more information about the properties of a credit class, see [ClassInfo](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#regen.ecocredit.v1.ClassInfo).

### Project

A project describes the high-level on-chain information for a project implementing the approved methodologies defined within a credit class. Each credit batch is associated with a project, backing each issuance with information about the project.

Over a project's lifecycle, it's expected that many credit batches will be issued at different points in time (e.g. at the conclusion of each monitoring period). To ensure that only legitimate projects are registered on-chain, projects can only be created by an issuer for the given credit class.

For more information about the properties of a project, see [ProjectInfo](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#regen.ecocredit.v1.ProjectInfo).

### Credit Type

A credit type is the primary unit of measurement used by the approved methodologies defined within a credit class. A credit type includes a name (e.g. carbon, biodiversity), an abbreviation (a set of 1-3 uppercase characters), a measurement unit (e.g. kilograms, tons), and a decimal precision.

The credit type abbreviation is used to construct the credit class ID. For example, `C01` is the ID for the first credit class that uses the `carbon` credit type where `C` is the credit type abbreviation and `01` is the sequence number (i.e. the first credit class for the given credit type).

Credit types are listed as an on-chain parameter. Adding a new credit type to the list of allowed credit types requires a parameter change proposal.

For more information about the properties of a credit type, see [CreditType](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#regen.ecocredit.v1.CreditType).

### Credit Class Creator Allowlist

The ecocredit module supports the option of restricting credit class creation to a set of addresses. The option to enable the credit class creator allowlist and the list of allowed credit class creators are both on-chain parameters that can only be updated through parameter change proposals.

### Credit Class Admin

The credit class admin is defined within a credit class as the address with the authority to update the credit class (i.e. to update the admin, the list of approved issuers, and the metadata). When a credit class is created, the admin is initially set to the address that created the credit class.

### Credit Class Issuers

The credit class issuers are defined within a credit class as the addresses with the authority to issue credit batches from the credit class. The list of credit class issuers is defined at the time the credit class is created and only the admin can update the list after the credit class is created.

### Credit Batch

Credits are issued in batches by credit issuers granted the authority to issue credits for a given credit class. A credit batch refers to a batch of credits issued at a single point in time.

Each credit batch has a unique ID (i.e. denomination) that starts with the abbreviation of the credit type followed by the start date, end date, and batch sequence number. For example, `C01-20190101-20200101-001` would be the first batch issued (`001`) from the first carbon credit class (`C01`) and the amount of carbon sequestered was measured between `20190101` and `20200101`.

Each credit batch is associated with an on-chain project, linking information about the on-the-ground project implementing the methodologies defined within the credit class. Additional information about a credit batch can be attached to the metadata field. When credits are issued, they can be issued in a tradable or retired state. The credit batch also tracks the total number of active credits (tradable and retired credits) and the total number of cancelled credits.

For more information about the properties of a credit batch, see [BatchInfo](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#regen.ecocredit.v1.BatchInfo).

### Credits

Credits are issued in credit batches in either a tradable or retired state. The owner of tradable credits can send, retire, or cancel the credits at any time. Tradable credits are only fungible with credits from the same credit batch. Retiring a credit is permanent.

### Tradable Credits

Tradable credits are credits that the owner has full control over. Tradable credits can be transferred by the owner to another account, put into a basket in return for basket tokens (see [basket](#basket-submodule) for more information), or listed for sale in the marketplace, placing the credits in escrow until the sell order is either processed or cancelled (or the amount is updated).

### Retired Credits

Retiring a credit is equivalent to burning a token with the exception that retired credits are actively tracked after being retired. Retiring a credit implies that the owner of the credit is claiming it as an offset. Credits can be retired upon issuance, retired upon transfer, or retired directly by the owner. Credits can also be set to automatically retire when taken from a [basket](#basket-submodule) or sold in the marketplace. Retiring a credit is permanent.

### Cancelled Credits

Cancelled credits are credit that cannot be traded or retired. Credits are cancelled in the event that the credit has moved to another registry.

## Basket Submodule

### Basket

...

### Basket Tokens

...

## Marketplace Submodule

The ecocredit module supports marketplace functionality using an order book model. The order book is an aggregate list of all the open buy and sell orders for ecosystem service credits. Depending on the preference of buyers and sellers, orders can be fully or partially executed and credits can be auto-retired or remain in a tradable state upon execution. In the current implementation of the order book, there is no automatic matching and users have to manually take the orders.

### Sell Order

A sell order is an order to sell ecosystem service credits. Each sell order has a unique ID that is auto-generated using a sequence table. A sell order stores the address of the owner of the credits being sold, the credit batch ID (denomination) of the credits being sold, the quantity of credits being sold, the asking price for each unit of the credit batch, and an option to enable/disable auto-retirement. Each credit unit of the credit batch will be sold for at least the asking price.

### Buy Order

A buy order is an order to buy ecosystem service credits. Like the sell order, each buy order has a unique ID that is auto-generated using a sequence table. A buy order can either be a direct buy order (an order against a specific sell order) or an indirect buy order (an order that can be filled by multiple sell orders that match a filter criteria). A buy order stores the selection (either the sell order id or the filter criteria), the quantity of credits to buy, the bid price for each unit of the credit batch(es), an option to enable/disable auto-retirement, and an option to enable/disable partial fills. A buy order can only successfully disable auto-retirement if the sell-order has disabled auto-retirement.

### Ask Denom

An "ask denom" is a denom that has been approved through a governance process as an accepted denom for listing ecosystem service credits. The "ask denom" includes the denom to allow (the base denom), the denom to display to the user, and an exponent that relates the denom to the display denom.

---

![Ecocredit Types](./assets/types.png)

![Credit Class Roles](./assets/roles.png)

![Allowlist Params](./assets/params.png)
