# Msg Service

## Anchor

`Anchor` "anchors" a piece of data to the blockchain based on its secure hash, effectively providing a tamper resistant timestamp. The sender in `Anchor` is not attesting to the veracity of the underlying data. They can simply be an intermediary providing timestamp services. `Attest` should be used to create a digital signature attesting to the veracity of some piece of data. 

+++ https://github.com/regen-network/regen-ledger/blob/55d39dc7d9768e5a0bd18c48ca78cb4a02145e81/proto/regen/data/v1alpha2/tx.proto#L44-L53

### Validation:

- `sender` must ba a valid address, and their signature must be present in the transaction
- `hash` must be a valid content hash, either raw data that does not specify a canonical encoding or graph data that conforms to the RDF data model

## Attest

`Attest` allows for signing of an arbitrary piece of data on the blockchain. By attesting to data, the attestor is making a statement about the veracity of the data itself. It is like signing a legal document, meaning that I agree to all conditions and to the best of my knowledge everything is true. When anchoring data, the sender is not attesting to the veracity of the data, they are simply communicating that it exists.

On-chain signatures have the following benefits:
- on-chain identities can be managed using different cryptographic keys that change over time through key rotation practices
- an on-chain identity may represent an organization and through delegation individual members may sign on behalf of the group
- the blockchain transaction envelope provides built-in replay protection and timestamping

`Attest` implicitly calls `Anchor` if the data was not already anchored. `Attest` can be called multiple times for the same content hash with different attestors and those attestors will be appended to the list of attestors.

+++ https://github.com/regen-network/regen-ledger/blob/55d39dc7d9768e5a0bd18c48ca78cb4a02145e81/proto/regen/data/v1alpha2/tx.proto#L64-L77

### Validation:

- `attestors` must be a valid addresses, and their signatures must be present in the transaction
- `hash` must be a valid content hash, either raw data that does not specify a canonical encoding or graph data that conforms to the RDF data model
