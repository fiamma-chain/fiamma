.PHONY: prove generate-proof send-proof query-tx clean

BINARY ?= fiammad
CHAIN_ID ?= fiammd-testnet-1
MODULE ?= verify
MESSAGE ?= gnark-plonk
PROVER ?= gnark_plonk.go

PROOF_FILE ?= proof.base64
PUBLIC_INPUTS_FILE ?= public_inputs.base64
VERIFYING_KEY_FILE ?= verifying_key.base64

prove: | generate-proof send-proof clean

generate-proof:
	go run gnark_plonk.go

submit-proof:
	$(BINARY) tx $(MODULE) $(MESSAGE) \
		$$(cat $(PROOF_FILE)) \
		$$(cat $(PUBLIC_INPUTS_FILE)) \
		$$(cat $(VERIFYING_KEY_FILE)) \
		--from alice --chain-id $(CHAIN_ID) \

query-tx:
	$(BINARY) query tx $(HASH) \

clean:
	rm *.base64
