let
  regen_build = pkg:
    with import <nixpkgs>{};
    buildGoModule rec {
      name = "regen-ledger";

      goPackagePath = "github.com/regen-network/regen-ledger";
      subPackages = [ pkg ];

      src = ./.;

      modSha256 = "0f3zj9d3ny5i3y32h7qji7jh1wpjx6fszv3b951jkjjb28xjabjr";

      meta = with stdenv.lib; {
        description = "Distributed ledger for planetary regeneration";
        license = licenses.asl20;
        homepage = https://github.com/regen-network/regen-ledger;
      };
    };
in {
  regen = (regen_build "app/regen");
}
