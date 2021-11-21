package ormgraphql_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/graphql-go/graphql/language/parser"
	"github.com/stretchr/testify/require"

	"github.com/graphql-go/graphql"
)

func Test1(t *testing.T) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"foo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "foo", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
          hello
          test {
           a
			b
}
          foo
sdgsdg
sdgsdgs
qefsdg
		}
	`
	AST, err := parser.Parse(parser.ParseParams{Source: query})
	//vr := graphql.ValidateDocument(&schema, AST, nil)
	//fmt.Printf("%+v", vr)
	require.NoError(t, err)
	params := graphql.ExecuteParams{Schema: schema, AST: AST}
	r := graphql.Execute(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}
