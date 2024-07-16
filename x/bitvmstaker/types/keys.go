package types

const (
	// ModuleName defines the module name
	ModuleName = "bitvmstaker"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bitvmstaker"
)

var (
	ParamsKey           = []byte{0x01} // key prefix for the parameters
	CommitteeAddressKey = []byte{0x02} // key prefix for the committee address
	StakerCountKey      = []byte{0x03} // key prefix for the staker count
	StakerKey           = []byte{0x04} // key prefix for the staker

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
