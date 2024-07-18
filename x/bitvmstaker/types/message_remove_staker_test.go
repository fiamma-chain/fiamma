package types

import (
	"testing"

	"fiamma/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveStaker_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveStaker
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemoveStaker{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRemoveStaker{
				Creator: sample.ValAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
