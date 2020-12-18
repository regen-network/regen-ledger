package main

import "google.golang.org/protobuf/compiler/protogen"

type Context struct {
}

func (ctx Context) genMessage(file *protogen.File, g *protogen.GeneratedFile, message *protogen.Message) {
	g.P("type ", message.GoIdent, " struct {")
	for _, field := range message.Fields {
		typ := ctx.getFieldType(field)
		g.P(field.GoName, " ", typ)
	}
	g.P("}")
	g.P()
}

func (ctx Context) getFieldType(f *protogen.Field) protogen.GoIdent {
	return f.Desc.
}
