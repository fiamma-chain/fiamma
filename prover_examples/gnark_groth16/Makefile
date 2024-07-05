.PHONY: prove generate-proof submit-proof query-tx clean

BINARY ?= fiammad
CHAIN_ID ?= fiamma-testnet-1
MODULE ?= zkpverify
MESSAGE ?= submit-proof
PROOF_SYSTEM ?= GROTH16_BN254
DIR = ./example/

PROOF_FILE ?= $(DIR)/proof
PUBLIC_INPUTS_FILE ?= $(DIR)/public_input
VERIFYING_KEY_FILE ?= $(DIR)/vk

prove: | generate-proof submit-proof clean

generate-proof:
	go run gnark_groth16.go

submit-proof:
	$(BINARY) tx $(MODULE) $(MESSAGE) $(PROOF_SYSTEM) \
		$$(cat $(PROOF_FILE)) \
		$$(cat $(PUBLIC_INPUTS_FILE)) \
		$$(cat $(VERIFYING_KEY_FILE)) \
		--from alice --chain-id $(CHAIN_ID) --gas 2000000\

query-tx:
	$(BINARY) query tx $(HASH) \

clean:
	rm -f $(DIR)/*
