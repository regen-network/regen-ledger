/*Package lookup provides a key-value store backed data structure which finds and stores
a short, unique binary identifier for a longer piece of binary data using an efficient,
non-cryptographic hash function.

A new lookup Table can be created with the NewTable() function. Advanced users can use
the NewTableWithOptions() function to tweak the underlying parameters, but the defaults
were chosen based on testing and should provide a good balance of performance and storage
efficiency.

Shortened identifiers are created or retrieved using the idempotent Table.GetOrCreateID method.

Using the default algorithm which uses the first 4 bytes of a 64-bit FNV-1a hash and then
increases the length the the case of collisions, identifiers will be 4 bytes long and retrieved
with a single KV-store lookup in the vast majority of cases and will sometimes be 5 or rarely
6 bytes long and require 2 to 3 total reads. In some rare cases (which have not appeared in tests),
identifiers may be longer or require more lookups.

Given a shortened identifier, the underlying binary data can be retrieved with the Table.GetValue
method.
*/
package lookup
