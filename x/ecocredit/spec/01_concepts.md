# Concepts

## Credit Type

A credit type defines the type of credit and the primary unit of measurement used within a credit class methodology. Each credit class is associated with a single credit type. A credit type is defined by a name (e.g. "carbon", "biodiversity"), an abbreviation (a unique set of 1-3 uppercase characters), a measurement unit (e.g. "kilograms", "tons"), and a decimal precision.

A credit type is like a super-class. For example, all credit classes that use "metric ton CO2 equivalent" as the primary unit of measurement would use the "carbon" credit type. The credit type abbreviation is used to construct the credit class ID (for example, "C1" is the ID for the first credit class that uses the "carbon" credit type where "C" is the credit type abbreviation).

## Credit Class

A credit class includes metadata that defines a methodology for measuring and monitoring changes in ecological state. The credit class metadata defines the structure, procedures, and requirements of the methodology. Each credit class is associated with a single credit type that defines the primary unit of measurement used within the credit class methodology. Each credit class includes a credit admin and a list of approved credit issuers.

## Credit Admin

A credit admin is the authority responsible for maintaining the credit class. In the first version of the ecocredit module, the credit admin is simply the authority who created the credit class. In the next version of the ecocredit module, the credit admin will be responsible for maintaining the list of approved credit issuers within the credit class. A credit admin is represented by an address.

## Credit Issuer

A credit issuer is the authority responsible for issuing credit batches for a given credit class upon the successful satisfaction of methodology constraints. A credit issuer must be listed as a credit issuer within a given credit class in order to issue credit batches from that credit class. A credit issuer is represented by an address.

## Credit Batch

Credits are issued in batches by credit issuers granted the authority to issue credits for a given credit class. A credit batch refers to a batch of credits issued at a single point in time. Each credit batch has a unique ID (i.e. denomination) that starts with the abbreviation of the credit type, followed by the start date, the end date, and the batch sequence number.

## Credits

Credits is a loose term for any fractional amount of a credit batch. Credits are issued in credit batches for a given credit class by an approved credit issuer.

## Tradable Credits

Tradable credits are credits that can be transferred between accounts.

## Retired Credits

Retired credits are credits that cannot be transferred between accounts nor can they be unretired. Retired credits are equivalent to burned tokens with the exception that retired credits are actively tracked after being retired. Retiring a credit implies that the holder of a credit is “consuming” the credit as an offset. Credits can be immediately retired on issuance. 
