package nubitda

import "github.com/rollkit/go-da/test"

// MockNubitDA is a mock implementation of NubitDA
type MockNubitDA struct {
	NubitDA
}

// NewMockNubitDA creates a new instance of MockNubitDA
func NewMockNubitDA() *MockNubitDA {
	return &MockNubitDA{
		NubitDA: NubitDA{
			ns:     []byte("test"),
			client: test.NewDummyDA(),
		},
	}
}
