package data

import "fmt"

func (ch ContentHash) Validate() error {
	switch hash := ch.Sum.(type) {
	case *ContentHash_Raw_:
		return hash.Raw.Validate()
	case *ContentHash_Graph_:
		return hash.Graph.Validate()
	default:
		return fmt.Errorf("invalid %T type %T", ch, hash)
	}
}

func (chr ContentHash_Raw) Validate() error {
	err := chr.MediaType.Validate()
	if err != nil {
		return err
	}

	return chr.DigestAlgorithm.Validate()
}

func (chg ContentHash_Graph) Validate() error {
	err := chg.CanonicalizationAlgorithm.Validate()
	if err != nil {
		return err
	}

	err = chg.MerkleTree.Validate()
	if err != nil {
		return err
	}

	return chg.DigestAlgorithm.Validate()
}

func (x MediaType) Validate() error {
	if _, ok := MediaType_name[int32(x)]; !ok {
		return fmt.Errorf("unknown %T %d", x, x)
	}

	return nil
}

func (x DigestAlgorithm) Validate() error {
	if _, ok := DigestAlgorithm_name[int32(x)]; !ok {
		return fmt.Errorf("unknown %T %d", x, x)
	}

	if x == DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED {
		return fmt.Errorf("%s is not a valid value for %T", x, x)
	}

	return nil
}

func (x GraphCanonicalizationAlgorithm) Validate() error {
	if _, ok := GraphCanonicalizationAlgorithm_name[int32(x)]; !ok {
		return fmt.Errorf("unknown %T %d", x, x)
	}

	if x == GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED {
		return fmt.Errorf("%s is not a valid value for %T", x, x)
	}

	return nil
}

func (x GraphMerkleTree) Validate() error {
	if _, ok := GraphMerkleTree_name[int32(x)]; !ok {
		return fmt.Errorf("unknown %T %d", x, x)
	}

	return nil
}
