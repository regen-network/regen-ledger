let
  xrn_build = pkg:
    with import <nixpkgs>{};
    buildGoModule rec {
      name = "regen-ledger";

      goPackagePath = "github.com/regen-network/regen-ledger";
      subPackages = [ pkg ];

      src = ./.;

      modSha256 = "0vhif5khib1hsnr540whqvi41aznnk24haq8252d6qglf2jjy01z";

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
