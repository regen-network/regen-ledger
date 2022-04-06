/*
Package hasher generates a unique binary identifier for a longer piece of binary data
using an efficient, non-cryptographic hash function.

A new Hasher instance can be created with the NewHasher() function. Advanced users can use
the NewHasherWithOptions() function to tweak the underlying parameters, but the defaults
were chosen based on testing and should provide a good balance of performance and storage
efficiency.

Shortened identifiers are generated using the idempotent Hasher.CreateID method.

Using the default algorithm which uses the first 4 bytes of a 64-bit BLAKE2b hash and then
increases the length in the case of collisions. Identifiers will be 4 bytes long in the vast
majority of cases and will sometimes be 5 and rarely 6 bytes long. In some extremely rare
cases (which have not appeared in tests), identifiers may be longer.
*/

package hasher
