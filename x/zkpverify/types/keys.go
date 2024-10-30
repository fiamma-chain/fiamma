package types

const (
	// ModuleName defines the module name
	ModuleName = "zkpverify"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

var (
	ParamsKey              = []byte{0x01} // key prefix for the parameters
	ProofDataKey           = []byte{0x02} // key prefix for the proof data
	VerifyResultKey        = []byte{0x03} // key prefix for the verify result
	BitVMChallengeDataKey  = []byte{0x04} // key prefix for the bit vm witness
	PendingProofsIndexKey  = []byte{0x05} // key prefix for the pending proofs
	BlockProposerKey       = []byte{0x06} // key prefix for the block proposer
	DASubmissionQueueKey   = []byte{0x07} // key prefix for the DA submission data queue
	DASubmissionResultsKey = []byte{0x08} // key prefix for the DA submission results
	DASubmitterKey         = []byte{0x09} // key prefix for the DA submitter address
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
