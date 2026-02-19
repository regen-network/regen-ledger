# wasmvm library hashes for v1.5.0
{ pkgs, system }:

{
  libwasmvm = pkgs.fetchurl (
    if system == "aarch64-linux" then {
      url = "https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvm_muslc.aarch64.a";
      sha256 = "2687afbdae1bc6c7c8b05ae20dfb8ffc7ddc5b4e056697d0f37853dfe294e913";
    } else if system == "x86_64-linux" then {
      url = "https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvm_muslc.x86_64.a";
      sha256 = "465e3a088e96fd009a11bfd234c69fb8a0556967677e54511c084f815cf9ce63";
    } else if system == "aarch64-darwin" then {
      url = "https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvmstatic_darwin.a";
      sha256 = "e45a274264963969305ab9b38a992dbc4401ae97252c7d59b217740a378cb5f2";
    } else throw "Unsupported system: ${system}"
  );
}
