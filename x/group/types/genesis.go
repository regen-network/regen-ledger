package types

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func (s GenesisState) Validate() error {
	return s.Params.Validate()
}
