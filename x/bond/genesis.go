package bond

const (
	PRECISION uint32 = 2
)

func (m *GenesisState) Validate() error {
	return nil
}

// DefaultGenesisState returns a default bond module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:   DefaultParams(),
		BondInfo: []*BondInfo{},
	}
}
