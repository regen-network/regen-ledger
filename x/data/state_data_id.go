package data

// Validate performs basic validation of the DataID state type
func (m *DataID) Validate() error {
	if len(m.Id) == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if _, err := ParseIRI(m.Iri); err != nil {
		return ErrParseFailure.Wrap(err.Error())
	}

	return nil
}
