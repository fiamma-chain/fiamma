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

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
