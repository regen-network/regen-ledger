package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/regen-network/regen-ledger/orm/internal/codegen"
)

func main() {
	protogen.Options{}.Run(codegen.ORMPluginRunner)
}
