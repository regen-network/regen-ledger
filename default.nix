let
  xrn_build = pkg:
    with import <nixpkgs>{};
    buildGoModule rec {
      name = "regen-ledger";

      goPackagePath = "github.com/regen-network/regen-ledger";
      subPackages = [ pkg ];

      src = ./.;

      modSha256 = "0cfb481v5cl7g2klffni4nx1wnd35kc49r0ahs348dr7zk6462dc";

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