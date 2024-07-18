package nubitda

import (
	"context"
	"encoding/hex"

	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
)

type NubitDA struct {
	ns     da.Namespace
	client da.DA
}

func NewNubitDA(config *DAConfig) (*NubitDA, error) {
	cn, err := proxy.NewClient(config.RpcURL, config.AuthKey)
	if err != nil || cn == nil {
		return nil, err
	}
	namespace, err := hex.DecodeString("00000000000000000000000000000000000000000000006669616d6d61")
	if err != nil {
		return nil, err
	}
	return &NubitDA{
		ns:     namespace,
		client: cn,
	}, nil
}

func NewNubitDATest(url string, authKey string, namespace []byte) (*NubitDA, error) {
	cn, err := proxy.NewClient(url, authKey)
	if err != nil || cn == nil {
		return nil, err
	}
	return &NubitDA{
		ns:     namespace,
		client: cn,
	}, nil
}

// SubmitBlob submits the data to the Nubit chain
func (a *NubitDA) SubmitBlobs(ctx context.Context, batchesData [][]byte) ([][]byte, error) {
	id, err := a.client.Submit(ctx, batchesData, -1, a.ns)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// GetBlob gets the data from the Nubit chain
func (a *NubitDA) GetBlobs(ctx context.Context, id [][]byte) ([][]byte, error) {
	blob, err := a.client.Get(context.TODO(), id, a.ns)
	if err != nil {
		return nil, err
	}
	return blob, nil
}

// GetBlobProof gets the data proofs from the Nubit chain

func (a *NubitDA) GetBlobProofs(ctx context.Context, id [][]byte) ([][]byte, error) {
	blob, err := a.client.GetProofs(context.TODO(), id, a.ns)
	if err != nil {
		return nil, err
	}
	return blob, nil
}
