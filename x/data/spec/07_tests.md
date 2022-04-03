# Tests

This document includes acceptance tests for the data module.

### Anchor Data

If a user tries to anchor data and the signer address is invalid, then the transaction fails and the data is not anchored.

- GIVEN - the signer address is invalid
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is empty, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is empty
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash does not conform to the canonical IRI scheme, then the transaction fails and the data is not anchored.

- GIVEN - the content hash does not conform to the canonical IRI scheme
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of graph type and the canonicalization algorithm is unspecified, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of graph type and the canonicalization algorithm is unspecified
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of graph type and the canonicalization algorithm is not supported, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of graph type and the canonicalization algorithm is not supported
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of graph type and the merkle tree is not supported, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of graph type and the merkle tree is not supported
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of graph type and the digest length is not 256 bits, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of graph type and the digest length is not 256 bits
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of raw type and the media type is not supported, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of raw type and the media type is not supported
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

If a user tries to anchor data and the content hash is of raw type and the digest length is not 256 bits, then the transaction fails and the data is not anchored.

- GIVEN - the content hash is of raw type and the digest length is not 256 bits
- WHEN - user tries to anchor data
- THEN - transaction fails, data is not anchored

### Sign Data

If a user tries to sign data and any of the signer addresses are invalid, then the transaction fails and the data is not signed.

- GIVEN - any of the signer addresses are invalid
- WHEN - user tries to sign data
- THEN - transaction fails, data is not signed

If a user tries to sign data and the content hash is empty, then the transaction fails and the data is not signed.

- GIVEN - the content hash is empty
- WHEN - user tries to sign data
- THEN - transaction fails, data is not signed

If a user tries to sign data and the content hash does not exist, then the transaction fails and the data is not signed.

- GIVEN - the content hash does not exist
- WHEN - user tries to sign data
- THEN - transaction fails, data is not signed
