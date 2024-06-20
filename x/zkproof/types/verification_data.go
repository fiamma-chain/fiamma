package types

type VerificationData struct {
	ProvingTypeId   ProofTypeId `json:"proof_type"`
	Proof           []byte      `json:"proof"`
	PubInput        []byte      `json:"pub_input"`
	VerificationKey []byte      `json:"verification_key"`
}
