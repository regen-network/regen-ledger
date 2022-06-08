package core

import (
	"fmt"
)

func (m OriginTx) Validate() error {
	if len(m.Id) == 0 {
		return fmt.Errorf("invalid OriginTx: no id")
	}
	if len(m.Source) == 0 {
		return fmt.Errorf("invalid OriginTx: no source")
	}
	return nil
}
