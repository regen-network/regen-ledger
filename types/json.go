package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CheckDuplicateKey checks duplicate keys in JSON
func CheckDuplicateKey(d *json.Decoder, path []string) error {
	// Get next token from JSON
	t, err := d.Token()
	if err != nil {
		return err
	}

	delim, ok := t.(json.Delim)

	// There's nothing to do for simple values (strings, numbers, bool, nil)
	if !ok {
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]bool)
		for d.More() {
			// Get field key
			t, err := d.Token()
			if err != nil {
				return err
			}

			key := t.(string)
			// Check for duplicates
			if keys[key] {
				return fmt.Errorf("duplicate key %s", key)
			}
			keys[key] = true

			// Check value
			if err := CheckDuplicateKey(d, append(path, key)); err != nil {
				return err
			}
		}
		// Consume trailing }
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := CheckDuplicateKey(d, append(path, strconv.Itoa(i))); err != nil {
				return err
			}
			i++
		}
		// Consume trailing ]
		if _, err := d.Token(); err != nil {
			return err
		}
	}

	return nil
}
