package group

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
)

var (
	// AddressLength is the length of all addresses
	// You can modify it in init() before any addresses are calculated,
	// but it must not change during the lifetime of the kvstore
	AddressLength = 20

	// conditionFormat defines the format that a Condition should have.
	// It must have (?s) flags, otherwise it errors when last section contains 0x20 (newline)
	conditionFormat = regexp.MustCompile(`(?s)^([a-zA-Z0-9_\-]{3,8})/([a-zA-Z0-9_\-]{3,8})/(.+)$`)
)

// Condition is a byte array specifying who can authorize an action.
// It has the following format:
//     {extension}/{type}/{data}
// data is binary data that represents an encoded sequence value from the ORM.
type Condition []byte

func NewCondition(ext, typ string, data []byte) Condition {
	pre := fmt.Sprintf("%s/%s/", ext, typ)
	return append([]byte(pre), data...)
}

// Parse will extract the permission sections from the Condition bytes
// and verify it is properly formatted
func (c Condition) Parse() (string, string, []byte, error) {
	chunks := conditionFormat.FindSubmatch(c)
	if len(chunks) == 0 {
		return "", "", nil, errors.Wrapf(ErrInvalid, "condition: %X", []byte(c))

	}
	return string(chunks[1]), string(chunks[2]), chunks[3], nil
}

// Address will convert a Condition into an Address
func (c Condition) Address() sdk.AccAddress {
	return newAddress(c)
}

// Equals checks if two permissions are the same
func (c Condition) Equals(b Condition) bool {
	return bytes.Equal(c, b)
}

// String returns a human readable string.
// We keep the extension and type in ascii and
// hex-encode the binary data
func (c Condition) String() string {
	ext, typ, data, err := c.Parse()
	if err != nil {
		return fmt.Sprintf("Invalid Condition: %X", []byte(c))
	}
	return fmt.Sprintf("%s/%s/%X", ext, typ, data)
}

// Validate returns an error if the Condition is not the proper format
func (c Condition) Validate() error {
	if len(c) == 0 {
		return ErrEmpty
	}
	if !conditionFormat.Match(c) {
		return errors.Wrapf(ErrInvalid, "condition: %X", []byte(c))
	}
	return nil
}

func (c Condition) MarshalJSON() ([]byte, error) {
	if c == nil {
		return []byte(`""`), nil
	}
	return json.Marshal(c.String())
}

func (c *Condition) UnmarshalJSON(raw []byte) error {
	var enc string
	if err := json.Unmarshal(raw, &enc); err != nil {
		return errors.Wrap(err, "cannot decode json")
	}
	return c.deserialize(enc)
}

// deserialize from human readable string.
func (c *Condition) deserialize(source string) error {
	// No value zero the address.
	if len(source) == 0 {
		*c = nil
		return nil
	}

	args := strings.Split(source, "/")
	if len(args) != 3 {
		return errors.Wrap(ErrInvalid, "invalid condition format")
	}
	data, err := hex.DecodeString(args[2])
	if err != nil {
		return errors.Wrapf(ErrInvalid, "malformed condition data: %s", err)
	}
	*c = NewCondition(args[0], args[1], data)
	return nil
}

// newAddress hashes and truncates into the proper size
func newAddress(data []byte) sdk.AccAddress {
	if data == nil {
		return nil
	}
	// h := blake2b.Sum256(data)
	h := sha256.Sum256(data)
	return h[:sdk.AddrLen]
}

// AccountCondition returns a condition to build a group account address.
func AccountCondition(id uint64) Condition {
	return NewCondition("group", "account", orm.EncodeSequence(id))
}
