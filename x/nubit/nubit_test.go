package nubit

import (
	"context"
	"encoding/hex"
	"testing"
	"time"
)

func TestNubitIntegration(t *testing.T) {

	namespace, err := hex.DecodeString("00000000000000000000000000000000000000000000006669616d6d61")
	if err != nil {
		t.Fatal(err)
	}
	nubit, err := NewNubitDATest("http://127.0.0.1:26658", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.cs2Y8oL1JNGhSTn29Khe_vEmQUB9_JeqI_LnQ2isWS8", namespace)
	if err != nil {
		t.Fatal(err)
	}
	txs := []byte("fiamma-testdata")

	id, err := nubit.SubmitBlobs(context.TODO(), [][]byte{txs})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("id: %#x", id)
	time.Sleep(600 * time.Microsecond)
	returdata, err := nubit.GetBlobs(context.TODO(), id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("returdata: %v", returdata)
}
