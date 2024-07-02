#include <stdbool.h>
#include <stdint.h>

typedef struct {
    char* data;
    uint32_t len;
    bool success;
} VerifyWitness;

VerifyWitness verify_proof_groth16_ffi(unsigned char *vk_buffer, uint32_t vk_len, unsigned char *proof_buffer, uint32_t proof_len, unsigned char *pi_buffer, uint32_t pi_len);

void free_verify_witness(VerifyWitness witness);

bool verify_sp1_proof_ffi(unsigned char *proof_buffer, uint32_t proof_len,
                          unsigned char *elf_buffer, uint32_t elf_len);


