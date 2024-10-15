package nubitda

import (
	"context"
	"encoding/hex"
	"testing"
	"time"
)

func TestNubitIntegration(t *testing.T) {
	// Test parameters
	const (
		nubitURL      = "http://127.0.0.1:26658"
		authToken     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.PsZV2O2vUZ5BBluCDMTcnzt1hGGFXOVSNZTfGEXw6DY"
		namespaceHex  = "00000000000000000000000000000000000000000000006669616d6d61"
		testData      = "fiamma-testdata"
		submitTimeout = time.Minute
		getTimeout    = 2 * time.Minute
		waitTime      = 10 * time.Second
	)

	// Decode namespace
	namespace, err := hex.DecodeString(namespaceHex)
	if err != nil {
		t.Fatalf("Failed to decode namespace: %v", err)
	}

	// Create NubitDA instance
	nubit, err := NewNubitDATest(nubitURL, authToken, namespace)
	if err != nil {
		t.Fatalf("Failed to create NubitDA instance: %v", err)
	}

	// Submit data
	ctx, cancel := context.WithTimeout(context.Background(), submitTimeout)
	defer cancel()
	id, err := nubit.SubmitBlobs(ctx, [][]byte{[]byte(testData)})
	if err != nil {
		t.Fatalf("Failed to submit data: %v\nURL: %s\nNamespace: %x", err, nubitURL, namespace)
	}
	t.Logf("Submitted data ID: %#x", id)

	// Wait for data processing
	time.Sleep(waitTime)

	// Retrieve data
	ctx, cancel = context.WithTimeout(context.Background(), getTimeout)
	defer cancel()
	returnedData, err := nubit.GetBlobs(ctx, id)
	if err != nil {
		t.Fatalf("Failed to retrieve data: %v", err)
	}

	// Validate returned data
	if len(returnedData) != 1 || string(returnedData[0]) != testData {
		t.Errorf("Returned data mismatch. Expected %q, got %q", testData, returnedData)
	} else {
		t.Logf("Successfully retrieved data: %q", returnedData[0])
	}
}
