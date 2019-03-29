/*
Package binary defines an efficient binary serialization format of the Regen Ledger graph data structures.

Grammar

The grammar of this format can be expressed as follows in a BNF-like notation:

	File = FileVersion RootNode
	FileVersion = 0x0 as uint64 varint
	OptRootNode = 0x0 | 0x1 NodeProperties
	Node = NodeID NodeProperties
	NodeID = 0x0 <types.GeoAddress as varint length-prefixed bytes> |
			 0x1 <AccAddressID as varint length-prefixed bytes> |
             0x2 <HashID fragment varint length-prefixed string>
	NodeProperties = <varint property count> Property*
	Property = 0x0 PropertyID Value
	Value = (based on property type)
		varint length-prefixed string
        float64 as 8 bytes
		bool as a single byte
		or a varint length-prefixed list of one of the above types

Notes

Only types.GeoAddress, AccAddressID, and HashID are currently accepted as node ID's by this format.
*/
package binary
