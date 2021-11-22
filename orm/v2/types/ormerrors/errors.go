package ormerrors

import "github.com/cosmos/cosmos-sdk/types/errors"

var codespace = "orm"

var (
	InvalidTableId                = errors.New(codespace, 1, "invalid or missing table or single id, need a non-zero value")
	MissingPrimaryKey             = errors.New(codespace, 2, "table is missing primary key")
	InvalidKeyFields              = errors.New(codespace, 3, "invalid field definition for key")
	DuplicateKeyField             = errors.New(codespace, 4, "duplicate field in key")
	FieldNotFound                 = errors.New(codespace, 5, "field not found")
	InvalidAutoIncrementKey       = errors.New(codespace, 6, "an auto-increment primary key must specify a single uint64 field")
	InvalidIndexId                = errors.New(codespace, 7, "invalid or missing index id, need a non-zero value")
	DuplicateIndexId              = errors.New(codespace, 8, "duplicate index id")
	PrimaryKeyConstraintViolation = errors.New(codespace, 9, "object with primary key already exists")
	NotFoundOnUpdate              = errors.New(codespace, 10, "can't update object which doesn't exist")
	PrimaryKeyInvalidOnUpdate     = errors.New(codespace, 11, "can't update object with missing or invalid primary key")
	AutoIncrementKeyAlreadySet    = errors.New(codespace, 12, "can't create with auto-increment primary key already set")
)
