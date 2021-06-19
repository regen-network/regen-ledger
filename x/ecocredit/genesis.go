package ecocredit

// ValidateGenesis check the given genesis state has no integrity issues.
func (s GenesisState) Validate() error {
	return nil
}

// DefaultGenesisState returns a default ecocredit module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		ClassInfos:       []*ClassInfo{},
		BatchInfos:       []*BatchInfo{},
		IdSeq:            0,
		TradableBalances: []*Balance{},
		RetriedBalances:  []*Balance{},
		TradableSupplies: []*Supply{},
		RetriedSupplies:  []*Supply{},
		Precisions:       []*Precision{},
	}
}
