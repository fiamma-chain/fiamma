package types

const (
	// ModuleName defines the module name
	ModuleName = "zkpverify"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

var (
	ParamsKey             = []byte{0x01} // key prefix for the parameters
	ProofDataKey          = []byte{0x02} // key prefix for the proof data
	VerifyResultKey       = []byte{0x03} // key prefix for the verify result
	BitVMChallengeDataKey = []byte{0x04} // key prefix for the bit vm witness
	PendingProofsKey      = []byte{0x05} // key prefix for the pending proofs
	BlockProposerKey      = []byte{0x06} // key prefix for the block proposer

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
