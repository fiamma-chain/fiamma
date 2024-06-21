.PHONY: prove generate-proof send-proof query-tx clean

BINARY ?= fiammad
CHAIN_ID ?= fiamma-testnet-1
MODULE ?= zkpverify
MESSAGE ?= send-task
PROOF_SYSTEM ?= PlonkBn254
PROVER ?= gnark_plonk.go
DIR = ./example/

PROOF_FILE ?= $(DIR)/proof
PUBLIC_INPUTS_FILE ?= $(DIR)/public_input
VERIFYING_KEY_FILE ?= $(DIR)/vk

prove: | generate-proof send-proof clean

generate-proof:
	go run gnark_plonk.go

send-task:
	$(BINARY) tx $(MODULE) $(MESSAGE) $(PROOF_SYSTEM) \
		$$(cat $(PROOF_FILE)) \
		$$(cat $(PUBLIC_INPUTS_FILE)) \
		$$(cat $(VERIFYING_KEY_FILE)) \
		--from alice --chain-id $(CHAIN_ID) --gas 2000000\

query-tx:
	$(BINARY) query tx $(HASH) \

clean:
	rm -f $(DIR)/*
