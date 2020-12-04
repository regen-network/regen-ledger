package types

import "context"

type HasContext interface {
	Context() context.Context
}
