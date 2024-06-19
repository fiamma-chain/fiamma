package types

import (
	"encoding/json"
	"fmt"
)

type ProofTypeId uint16

const (
	GnarkPlonk ProofTypeId = iota
	GnarkGroth16
	SP1
)

func (t *ProofTypeId) String() string {
	return [...]string{"GnarkPlonk", "GnarkGroth16", "SP1"}[*t]
}

func ProofTypeIdFromString(proofType string) (ProofTypeId, error) {
	switch proofType {
	case "GnarkPlonk":
		return GnarkPlonk, nil
	case "GnarkGroth16":
		return GnarkGroth16, nil
	case "SP1":
		return SP1, nil
	}

	return 0, fmt.Errorf("unknown proof type: %s", proofType)
}

func ProofTypeIdToString(proofTypeId ProofTypeId) (string, error) {
	switch proofTypeId {
	case GnarkPlonk:
		return "GnarkPlonk", nil
	case GnarkGroth16:
		return "GnarkGroth16", nil
	case SP1:
		return "SP1", nil
	}

	return "", fmt.Errorf("unknown proof type id : %d", proofTypeId)
}

func (t *ProofTypeId) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = ProofTypeIdFromString(s)
	return err
}

func (t ProofTypeId) MarshalJSON() ([]byte, error) {
	// Check if the enum value has a corresponding string representation
	if str, ret := ProofTypeIdToString(t); ret == nil {
		// If yes, marshal the string representation
		return json.Marshal(str)
	}
	// If not, return an error
	return nil, fmt.Errorf("invalid ProofTypeId value: %d", t)
}
