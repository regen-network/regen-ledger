# Concepts

## Credit Type

A credit type is the primary indicator for the methodology defined within a credit class. A credit type includes a name (e.g. carbon, biodiversity), an abbreviation (a set of 1-3 uppercase characters), a measurement unit (e.g. kilograms, tons), and a decimal precision.

The credit type abbreviation is used to construct the credit class ID. For example, "C1" is the ID for the first credit class that uses the "carbon" credit type where "C" is the credit type abbreviation and "1" is the sequence number of the credit class for the given credit type.

Credit types are listed as an on-chain parameter. Adding a new credit type to the list of approved credit types requires a parameter change proposal.

## Credit Class

A credit class defines the class of an ecosystem service credit. A credit class includes metadata that defines a methodology for measuring and monitoring changes in ecological state. The credit class metadata defines the structure, procedures, and requirements of the methodology.

Each credit class is associated with a single credit type, which is the primary indicator for an ecosystem service methodology. Each credit class includes a credit class admin (the creator and maintainer of the credit class) and a list of approved credit issuers.

## Credit Class Creator

A credit class creator is an address with the authority to create credit classes. A credit class creator is granted the authority to create credit classes by being on the list of approved credit class creators (an on-chain parameter that can only be updated through the governance process). A credit class creator creates a credit class by executing a transaction. A credit class creator becomes the credit class admin upon creation of the credit class and is therein referred to as the credit class admin.

## Credit Class Admin

The credit class admin is defined within a credit class. A credit class admin is the address with the authority to manage the credit class. In the initial version of the ecocredit module, the credit class admin simply represents the creator of the credit class. In the next version of the ecocredit module, the credit class admin will have the ability to manage the list of credit class issuers and potentially other properties of the credit class.

## Credit Issuers

The credit issuers are defined within a credit class. Credit issuers are a list of addresses with the authority to issue credit batches from the credit class. The list of credit issuers are defined at the time the credit class is created. In the next version of the ecocredit module, the credit class admin will have the ability to manage the list of credit issuers for the credit class that they administer.

## Credit Batch

Credits are issued in batches by credit issuers granted the authority to issue credits for a given credit class. A credit batch refers to a batch of credits issued at a single point in time. Each credit batch has a unique ID (i.e. denomination) that starts with the abbreviation of the credit type followed by the start date, end date, and batch sequence number.

A credit batch includes information about the issuer of the credit batch and the project location, and any additional information can be attached to the metadata field. The credit batch tracks the total number of active credits and the total number of cancelled credits.

## Credits

Credits are a loose term that refers to any fractional amount of a credit batch. Credits are issued in credit batches by an approved credit issuer within a credit class. Credit denominations are defined by the credit batch and each credit batch provides a unique denomination. Credits are either tradable or retired and each credit batch tracks the number of tradable and retired credits.

## Tradable Credits

Tradable credits are credits that can be transferred by the owner to another account.

## Retired Credits

Retired credits are credits that cannot be transferred between accounts nor can they be unretired. Retired credits are equivalent to burned tokens with the exception that retired credits are actively tracked after being retired. Retiring a credit implies that the holder of a credit is “consuming” the credit as an offset. Credits can be retired upon issuance, retired upon transfer, and retired by the owner of the credits.
