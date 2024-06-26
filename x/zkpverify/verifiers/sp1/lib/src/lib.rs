use std::slice;
use sp1_sdk::ProverClient;
use lazy_static::lazy_static;

lazy_static! {
    static ref PROVER_CLIENT: ProverClient = ProverClient::new();
}

#[no_mangle]
pub extern "C" fn verify_sp1_proof_ffi(
    proof_bytes: *const u8,
    proof_len: u32,
    elf_bytes: *const u8,
    elf_len: u32,
) -> bool {
    let proof_bytes = unsafe {
        assert!(!proof_bytes.is_null());
        slice::from_raw_parts(proof_bytes, proof_len as usize)
    };

    let elf_bytes = unsafe {
        assert!(!elf_bytes.is_null());
        slice::from_raw_parts(elf_bytes, elf_len as usize)
    };

    if let Ok(proof) = bincode::deserialize(proof_bytes) {
        let (_pk, vk) = PROVER_CLIENT.setup(elf_bytes);
        return PROVER_CLIENT.verify_compressed(&proof, &vk).is_ok();
    }

    false
}
