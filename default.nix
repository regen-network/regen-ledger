with import <nixpkgs>{};

buildGoModule rec {
  name = "regen-ledger";

  goPackagePath = "github.com/regen-network/regen-ledger";
  subPackages = [ "cmd/xrnd" "cmd/xrncli" ];

  src = ./.;

  meta = with stdenv.lib; {
    description = "Distributed ledger for planetary regeneration";
    license = licenses.asl20;
    homepage = https://github.com/regen-network/regen-ledger;
  };
}