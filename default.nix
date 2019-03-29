let
  xrn_build = pkg:
    with import <nixpkgs>{};
    buildGoModule rec {
      name = "regen-ledger";

      goPackagePath = "github.com/regen-network/regen-ledger";
      subPackages = [ pkg ];

      src = ./.;

      modSha256 = "1ja69hd2isndjfsbzk96fhdblyg3k7plss9i8ibsp6ppx16j74ck";

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