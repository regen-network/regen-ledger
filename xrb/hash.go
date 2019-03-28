package xrb

import (
	"bufio"
	"bytes"
	"strconv"
)

type hasher struct {
	*bufio.Writer
	buf *bytes.Buffer
}

func newHasher() *hasher {
	b := new(bytes.Buffer)
	return &hasher{bufio.NewWriter(b), b}
}

func (ha *hasher) hashGraph(g Graph) {
	panic("TODO")
}

func (ha *hasher) write(x string) {
	_, err := ha.WriteString(x)
	if err != nil {
		panic(err)
	}
}

func (ha *hasher) writeStringLiteral(x string, typeIri string, lang string) {
	ha.write("\"")
	ha.write(strconv.Quote(x))
	ha.write("\"")
	if typeIri != "" {
		if lang != "" {
			panic("cannot specify both a data type IRI and a language tag")
		}
		ha.write("^^")
		ha.writeIRI(typeIri)
	} else if lang != "" {
		ha.write("@")
		ha.write(lang)
	}
}

func (ha *hasher) writeIRI(x string) {
	ha.write("<")
	ha.write(x)
	ha.write("> ")
}

func (ha *hasher) finishLine() {
	ha.write(".\n")
}

func (ha *hasher) hash() []byte {
	panic("TODO")
}
