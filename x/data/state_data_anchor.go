package data

// Validate performs basic validation of the DataAnchor state type
func (m *DataAnchor) Validate() error {
	if len(m.Id) == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if m.Timestamp.Unix() <= 0 {
		return ErrParseFailure.Wrapf("timestamp must be well defined")
	}

	return nil
}
