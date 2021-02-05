package jsongold

import (
	"fmt"

	"github.com/piprate/json-gold/ld"
	_ "github.com/piprate/json-gold/ld"
)

func normalize(dataset *ld.RDFDataset) (string, error) {
	api := ld.NewJsonLdApi()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/n-quads"
	options.Algorithm = "URDNA2015"
	res, err := api.Normalize(dataset, options)
	if err != nil {
		return "", err
	}

	str, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("expect a string, got %T", res)
	}

	return str, nil
}
