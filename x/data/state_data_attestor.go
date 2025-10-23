package data

// Validate performs basic validation of the DataAttestor state type
func (m *DataAttestor) Validate() error {
	if len(m.Id) == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if m.Timestamp == nil {
		return ErrParseFailure.Wrapf("timestamp cannot be empty")
	}

	return nil
}
