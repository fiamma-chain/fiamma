package types_test

import (
	"testing"

	"fiamma/testutil/sample"
	"fiamma/x/bitvmstaker/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is Invalid",
			genState: types.DefaultGenesis(),
			valid:    false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				CommitteeAddress: sample.AccAddress(),
				StakerAddresses:  []string{sample.ValAddress()},
				Params:           types.DefaultParams(),

				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
