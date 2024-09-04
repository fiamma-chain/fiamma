package keeper_test

import (
	"context"
	"testing"

	keepertest "fiamma/testutil/keeper"

	"fiamma/x/zkpverify/keeper"

	"fiamma/x/zkpverify/types"

	"github.com/stretchr/testify/require"
)

func TestSubmitProof(t *testing.T) {
	k, ctx := keepertest.ZkpVerifyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	wctx := ctx.WithContext(context.Background())

	tests := []struct {
		name    string
		msg     *types.MsgSubmitProof
		want    *types.MsgSubmitProofResponse
		wantErr bool
	}{
		{
			name: "valid proof",
			msg: &types.MsgSubmitProof{
				Creator:     "creator",
				Proof:       []byte("valid proof"),
				ProofSystem: "GROTH16_BN254_BITVM",
				PublicInput: []byte("valid public input"),
				Vk:          []byte("valid vk"),
				Namespace:   "test",
			},
			want:    &types.MsgSubmitProofResponse{},
			wantErr: false,
		},
		{
			name: "empty proof",
			msg: &types.MsgSubmitProof{
				Creator: "creator",
				Proof:   []byte{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid proof system",
			msg: &types.MsgSubmitProof{
				Creator:     "creator",
				Proof:       []byte("invalid proof"),
				ProofSystem: "999",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.SubmitProof(wctx, tt.msg)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)

			// check if the proof id is valid
			proofId, _ := k.GetProofId(types.ProofData{
				Proof:       tt.msg.Proof,
				ProofSystem: types.ProofSystem_GROTH16_BN254_BITVM,
				Namespace:   "test",
				PublicInput: tt.msg.PublicInput,
				Vk:          tt.msg.Vk,
			})
			// Check if the proof is in the pending proofs index
			isPending := k.IsPendingProof(ctx, proofId[:])
			require.True(t, isPending)

			verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
			require.True(t, found)
			require.NotNil(t, verifyResult)
		})
	}
}

func TestSubmitProofConcurrency(t *testing.T) {
	k, ctx := keepertest.ZkpVerifyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	wctx := ctx.WithContext(context.Background())

	msg := &types.MsgSubmitProof{
		Creator:     "creator",
		Proof:       []byte("valid proof"),
		ProofSystem: "GROTH16_BN254_BITVM",
		PublicInput: []byte("valid public input"),
		Vk:          []byte("valid vk"),
		Namespace:   "test",
	}

	// simulate concurrent submissions
	concurrentSubmissions := 10
	done := make(chan bool)

	for i := 0; i < concurrentSubmissions; i++ {
		go func() {
			_, err := srv.SubmitProof(wctx, msg)
			require.NoError(t, err)
			done <- true
		}()
	}

	// wait for all submissions to complete
	for i := 0; i < concurrentSubmissions; i++ {
		<-done
	}

	// check the final status
	proofId, _ := k.GetProofId(types.ProofData{
		Proof:       msg.Proof,
		ProofSystem: types.ProofSystem_GROTH16_BN254_BITVM,
		Namespace:   "test",
		PublicInput: msg.PublicInput,
		Vk:          msg.Vk,
	})
	isPending := k.IsPendingProof(ctx, proofId[:])
	require.True(t, isPending)

	verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
	require.True(t, found)
	require.NotNil(t, verifyResult)
}
