package types

const (
	// ModuleName defines the module name
	ModuleName = "zkpverify"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zkpverify"
)

var (
	ParamsKey       = []byte{0x01} // key prefix for the parameters
	ProofDataKey    = []byte{0x02} // key prefix for the proof data
	VerifyResultKey = []byte{0x03} // key prefix for the verify result
	BitVMWitnessKey = []byte{0x04} // key prefix for the bit vm witness

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
