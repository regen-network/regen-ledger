/*
Package schema defines a module for defining on-chain schema definitions. The rationale for on-chain schema definitions
is that a blockchain can enforce schema definition rules that force backwards compatibility. Specifically if a
schema property or class is defined to have one meaning at one point, that same identifier can't be redefined with
an incompatible meaning in the future. Having consensus around the meaning of schema identifiers provides a basis
for which applications developers can safely write interoperable apps that don't break arbitrarily. In addition to
providing this basis for consistency, the schema module also aims to provide clear pathways for upgrading schemas
and defining schemas which mix and match different parts of other schemas in order to facilitate a new generation
of interoperable apps and data, most importantly within the ecological domain.
*/
package schema
