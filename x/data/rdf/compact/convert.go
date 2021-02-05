package compact

import "fmt"

func Compact(content []byte, contentType string, resolver InternalIDResolver) (*CompactDataset, error) {
	switch contentType {
	case "application/ld+json":
		return nil, fmt.Errorf("unsupported content type %s", contentType)
	case "application/n-quads":
		return nil, fmt.Errorf("unsupported content type %s", contentType)
	case "application/n-triples":
		return nil, fmt.Errorf("unsupported content type %s", contentType)
	default:
		return nil, fmt.Errorf("unsupported content type %s", contentType)
	}
}
