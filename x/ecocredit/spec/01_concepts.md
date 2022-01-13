# Concepts

## Credit Class

A credit class is the primary abstraction for ecosystem service credits and is defined by 5 attributes:
- **credit class ID**: auto-generated, and auto incrementing
- **admin**: The regen address who can update fields/attributes of this credit class
- **issuer list**: The list of regen addresses who are allowed to issue credit batches under this credit class
- **credit type**: The primary indicator for this credit class (e.g. Carbon measured in Tons Equiv. CO2 sequestered)
- **metadata**: A byte array (up to 256 bytes) which can be used to store small amounts of metadata, or a URI that points to an off-chain resource for querying more complete metadata information. This usually would include descriptive information about the credit class's acceptable methodologies for monitoring changes in ecological state.

Once a new credit class is created, credits can be issued at-will in distinct batches by any address in the **issuer list** of the corresponding credit class.

## Credit Type

A credit type is the primary indicator used by the methodology to measure the change or impact resulting from an ecosystem service. A credit type includes a name (e.g. carbon, biodiversity), an abbreviation (a set of 1-3 uppercase characters), a measurement unit (e.g. kilograms, tons), and a decimal precision.

The credit type abbreviation is used to construct the credit class ID. For example, "C1" is the ID for the first credit class that uses the "carbon" credit type where "C" is the credit type abbreviation and "1" is the sequence number of the credit class for the given credit type.

Credit types are listed as an on-chain parameter. Adding a new credit type to the list of approved credit types requires a parameter change proposal.

## Credit Class Creator Allowlist

The ecocredit module supports the option of restricting credit class creation to a permissionned set of addresses. When enabled, this list of approved credit class creators is stored as an on-chain parameter that can only be updated through the governance process. Regen Ledger 2.0 is intended to launch with this allowlist enabled.

For users wanting to experiment with creating their own credit classes on testnets, the Hambach Testnet will support permissionless credit class creation so any user can create new credit classes and test out the ecocredit module's functionality.

## Credit Class Admin

The credit class admin is defined within a credit class as the address with the authority to manage and update the credit class. When a new credit class is created, the admin is always initially set to the address that created the credit class. The credit class admin will have the ability to transfer the admin role to another address, manage the list of credit class issuers, and change credit class metadata.

## Credit Issuers

The credit issuers are defined within a credit class as a list of addresses with the authority to mint new credits and issue credit batches of the corresponding credit class. The list of credit issuers are defined at the time the credit class is created. The credit class admin will be able to manage the list of credit issuers for the credit class that they administer.

## Credit Batch

Credits are issued in batches by credit issuers granted the authority to issue credits for a given credit class. A credit batch refers to a batch of credits issued at a single point in time.

Each credit batch has a unique ID (i.e. denomination) that starts with the abbreviation of the credit type followed by the start date, end date, and batch sequence number. For example, `C01-20190101-20200101-001` would be the first batch issued (`001`) from the first carbon credit class (`C01`) and the reduction of carbon emissions was measured between `20190101` and `20200101`.

A credit batch also includes information about the issuer of the credit batch and the project location, and any additional information can be attached to the metadata field. The credit batch tracks the total number of active credits and the total number of cancelled credits.

## Credits

Credits are a loose term that refers to any fractional amount of a credit batch. Credits are either tradable or retired and each credit batch tracks the number of tradable and retired credits.

## Tradable Credits

Tradable credits are credits that can be transferred by the owner to another account.

## Retired Credits

Retired credits are credits that cannot be transferred between accounts nor can they be unretired. Retired credits are equivalent to burned tokens with the exception that retired credits are actively tracked after being retired. Retiring a credit implies that the holder of a credit is “claiming” the credit as an offset. Credits can be retired upon issuance, retired upon transfer, and retired by the owner of the credits. The retirement location is required upon retirement.

---

![Ecocredit Types](./assets/types.png)

![Credit Class Roles](./assets/roles.png)

![Allowlist Params](./assets/params.png)
