package graph

// PropertyID is the integer ID of a property in a schema with 0 indicating undefined
type PropertyID int64

// SchemaResolver resolves properties against a schema
type SchemaResolver interface {
	GetPropertyByURL(url string) Property
	GetPropertyByID(id PropertyID) Property
	GetPropertyID(p Property) PropertyID
	// MinPropertyID indicates the lowest defined PropertyID, usually negative (for built-in properties) or zero
	MinPropertyID() PropertyID
	// MaxPropertyID indicates the highest defined PropertyID, usually positive (for used defined properties)
	MaxPropertyID() PropertyID
}
