package main

import (
	"bytes"
	"encoding/hex"
	"io"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	"github.com/consensys/gnark/frontend"
)

// gnark is a zk-SNARK library written in Go. Circuits are regular structs.
// The inputs must be of type frontend.Variable and make up the witness.
type Circuit struct {
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit logic. The compiler then produces a list of constraints
// which must be satisfied (valid witness) in order to create a valid zk-SNARK
func (circuit *Circuit) Define(api frontend.API) error {
	// compute x**3 and store it in the local variable x3.
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)

	// compute x**3 + x + 5 and store it in the local variable res
	res := api.Add(x3, circuit.X, 5)

	// assert that the statement x**3 + x + 5 == y is true.
	api.AssertIsEqual(circuit.Y, res)
	return nil
}

// Defines the circuit that will be proved.
func solution() Circuit {
	return Circuit{
		X: 3,
		Y: 35,
	}
}

//////////////////////
//// DON'T CHANGE ////
//////////////////////

func main() {
	outputDir := "example/"
	var myCircuit Circuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &myCircuit)
	if err != nil {
		panic("circuit compilation error")
	}
	pk, vk, _ := groth16.Setup(ccs)

	circuit := solution()
	fullWitness, _ := frontend.NewWitness(&circuit, ecc.BN254.ScalarField())
	publicWitness, _ := fullWitness.Public()

	proof, _ := groth16.Prove(ccs, pk, fullWitness)

	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		panic("GROTH16 proof not verified")
	}

	serialize(proof, outputDir+"proof")
	serialize(publicWitness, outputDir+"public_input")
	serialize(vk, outputDir+"vk")
}

func serialize[w io.WriterTo](src w, name string) {
	var buffer bytes.Buffer
	_, _ = src.WriteTo(&buffer)

	inner := buffer.Bytes()

	encoded := make([]byte, hex.EncodedLen(len(inner)))
	hex.Encode(encoded, inner)

	_ = os.WriteFile(name, encoded, 0o644)
}
