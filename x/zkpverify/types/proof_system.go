package types

import (
	"encoding/json"
	"fmt"
)

type ProofSystemId uint64

const (
	PlonkBn254 ProofSystemId = iota
	PlonkBls12_381
	Groth16Bn254
	Groth16Bn254_BTC
	SP1
)

func (t *ProofSystemId) String() string {
	return [...]string{"PlonkBn254", "PlonkBls12_381", "Groth16Bn254", "Groth16Bn254_BTC", "SP1"}[*t]
}

func ProofSystemIdFromString(proofSystem string) (ProofSystemId, error) {
	switch proofSystem {
	case "PlonkBn254":
		return PlonkBn254, nil
	case "PlonkBls12_381":
		return PlonkBls12_381, nil
	case "Groth16Bn254":
		return Groth16Bn254, nil
	case "Groth16Bn254_BTC":
		return Groth16Bn254_BTC, nil
	case "SP1":
		return SP1, nil
	}
	return 0, fmt.Errorf("unknown proof type: %s", proofSystem)
}

func ProofSystemIdToString(ProofSystemId ProofSystemId) (string, error) {
	switch ProofSystemId {
	case PlonkBn254:
		return "PlonkBn254", nil
	case PlonkBls12_381:
		return "PlonkBls12_381", nil
	case Groth16Bn254:
		return "Groth16Bn254", nil
	case Groth16Bn254_BTC:
		return "Groth16Bn254_BTC", nil
	case SP1:
		return "SP1", nil
	}
	return "", fmt.Errorf("unknown proof type id : %d", ProofSystemId)
}

func (t *ProofSystemId) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = ProofSystemIdFromString(s)
	return err
}

func (t *ProofSystemId) Unmarshal(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = ProofSystemIdFromString(s)
	return err
}

func (t ProofSystemId) MarshalJSON() ([]byte, error) {
	// Check if the enum value has a corresponding string representation
	if str, ret := ProofSystemIdToString(t); ret == nil {
		// If yes, marshal the string representation
		return json.Marshal(str)
	}
	// If not, return an error
	return nil, fmt.Errorf("invalid ProofSystemId value: %d", t)
}
