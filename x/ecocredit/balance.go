package ecocredit

import "github.com/regen-network/regen-ledger/types/math"

type balance interface {
	GetEscrowed() string
	GetTradable() string
	GetRetired() string
}

// GetDecimalsFromBalance takes a balance interface and returns the fields as decimals.
func GetDecimalsFromBalance(b balance) (tradable, retired, escrowed math.Dec, err error) {
	tradable, err = math.NewDecFromString(b.GetTradable())
	if err != nil {
		return
	}
	retired, err = math.NewDecFromString(b.GetRetired())
	if err != nil {
		return
	}
	escrowed, err = math.NewDecFromString(b.GetEscrowed())
	if err != nil {
		return
	}
	return
}
