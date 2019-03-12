with import <nixpkgs>{};

buildGoPackage rec {
  name = "regen-ledger";

  goPackagePath = "github.com/regen-network/regen-ledger";
  subPackages = [ "cmd/xrnd" "cmd/xrncli" ];

  src = ./.;

  # buildInputs = [ makeWrapper ];
  # binPath = lib.makeBinPath [ ];

  goDeps = ./deps.nix;

  #postInstall = ''
  #  wrapProgram $bin/bin/dep2nix --prefix PATH ':' ${binPath}
  #'';

  meta = with stdenv.lib; {
    description = "Regen Networks's Regen Ledger distributed ledger for planetary regeneration";
    # TODO license = licenses.bsd3;
    homepage = https://github.com/regen-network/regen-ledger;
  };
}