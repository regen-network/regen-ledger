# Concepts

## Credit Type

A credit type is the top-level category of a credit. A credit type is defined by a name (e.g. carbon, biodiversity), an abbreviation (i.e. a unique 1-3 character uppercase abbreviation used in batch denominations), a measurement unit (e.g. kilograms, tons), and a decimal precision.

## Credit Class

A credit class defines a type of credit that is maintained by a credit designer and issued by a credit issuer. A credit class is defined by an ID, a designer, an approved list of credit issuers, a credit type, and optional metadata.

## Credit Designer

A credit designer is the authority responsible for creating a credit class and updating its list of approved issuers as needed. A credit designer is represented by an address.

## Credit Issuer

A credit issuer is the authority responsible for issueing credit batches to project developers based on successful satisfaction of methodology constraints. A credit issuer is represented by an address.

## Credit Batch

Credits are issued in batches. A credit batch refers to a batch of credits issued at a single point in time (usually corresponding to some off or on-chain verification event, and corresponding to some project).

Upon issuance, a credit batch points to a project, and geo-polygon, and mints all credits in that batch to a set of accounts (typically the land steward / project owner).

## Credit

In this design credit batches can be split up into any fractional amount (arbitrary precision decimal) as needed and thus credit batches are the top-level thing issued, but they can be split up as needed. Credits is thus a loose term to describe some quantities of credits potentially of different batches and classes. “One credit” would generally refer to 1.0 units of a given credit batch.

- Credits are represented as a fungible on-chain asset, where on-chain accounts can have a balance in the given credit
- Credits can be issued/minted at any time by a fixed set of “issuers”
- Credits are issued in batches (a batch of credits is hereafter referred to as a credit batch)

a fractional NFT
an on-chain asset
an ecological credit
can issued, traded and retired

## Retirement

Retirement is the state in which a credit can no longer be transferred. In conventional blockchain terminology, this is practically equivalent to the burned state and the word burn may be used in the technical implementation. The main difference is that we still care to actively track the balance of retired credits. Conceptually retiring a credit implies that the holder of a credit is “consuming” the credit as an offset to satisfy voluntary or compliance-related offset commitments.

Credits that are retired cannot be un-retired by either the credit issuer, or credit designer.

...

## Tradable Credits

Tradable credits can be transferred between accounts, by the owner

## Retired Credits

Retired credits cannot be transferred, and cannot be unretired. Credits can be immediately retired on issuance.