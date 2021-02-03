package group

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

func (s GenesisState) Validate() error {
	return nil
}
