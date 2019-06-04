let
  xrn_build = pkg:
    with import <nixpkgs>{};
    buildGoModule rec {
      name = "regen-ledger";

      goPackagePath = "github.com/regen-network/regen-ledger";
      subPackages = [ pkg ];

      src = ./.;

      modSha256 = "0z3a62fcfnv40z5kxpv82z9mlmc8i34pbzdmlxvwy2l17y7hys32";

      meta = with stdenv.lib; {
        description = "Distributed ledger for planetary regeneration";
        license = licenses.asl20;
        homepage = https://github.com/regen-network/regen-ledger;
      };
    };
in {
  xrnd = (xrn_build "cmd/xrnd");
  xrncli = (xrn_build "cmd/xrncli");
}
