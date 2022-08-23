package data

// Validate performs basic validation of the DataResolver state type
func (m *DataResolver) Validate() error {
	if len(m.Id) == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if m.ResolverId == 0 {
		return ErrParseFailure.Wrap("resolver id cannot be empty")
	}

	return nil
}
