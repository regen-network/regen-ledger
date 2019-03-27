{ config, pkgs, lib, ... }:

with lib;

let
  xrndCfg = config.services.xrnd;
  xrn_build = pkg:
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
  xrnd = (xrn_build "cmd/xrnd");
  xrncli = (xrn_build "cmd/xrncli");
in
{
  options = {
    programs.xrn = {
      enable =
        mkOption {
          type = types.bool;
          default = false;
          description = ''
            Whether to install regen-ledger.
          '';
        };
    };
    services.xrnd = {
      enable =
        mkOption {
          type = types.bool;
          default = false;
          description = ''
            Whether to run xrnd.
          '';
        };
      home =
        mkOption {
          type = types.path;
          default = "/var/xrnd";
          description = ''
            Path to xrnd home folder. Must be created before the service is started.
          '';
        };
      moniker =
        mkOption {
          type = types.str;
          default = "node0";
          description = ''
            The node moniker.
          '';
        };
      restServer =
        mkOption {
          type = types.bool;
          default = false;
          description = ''
            Whether to run the xrncli REST server.
          '';
        };
      enablePostgresIndex =
        mkOption {
          type = types.bool;
          default = false;
          description = "Automatically enable the Postgresql service and index to a database named xrn. Shouldn't be used together with postgresIndexUrl";
        };
      postgresIndexUrl =
        mkOption {
          type = types.str;
          default = "";
          description = "The URL of a Postgresql database to index to. Shouldn't be used together with enablePostgresIndex";
        };
    };
  };
  config = mkMerge [
    (mkIf config.programs.xrn.enable {
      environment.systemPackages = [ xrn ];
    })

    (mkIf xrndCfg.enable {
        users.groups.xrn = {};

        users.users.xrnd = {
          isSystemUser = true;
          group = "xrn";
          home = xrndCfg.home;
        };

        networking.firewall.allowedTCPPorts = [ 26656 ];

        systemd.services.xrnd = {
          description = "Regen Ledger Daemon";
          wantedBy = [ "multi-user.target" ];
          after = [ "network.target" ];
          path = [ xrnd pkgs.jq ];
          preStart = ''
            chown -R xrnd:xrn ${xrndCfg.home}
          '';
          script = ''
            xrnd start --moniker ${xrndCfg.moniker} --home ${xrndCfg.home}
          '';
          postStop = ''
            export UPGRADE_COMMIT=$(jq '.commit' < ${xrndCfg.home}/data/upgrade-info)
            if  [ $UPGRADE_COMMIT != "null" ]; then
              cd /root/regen-ledger
              git clean -f
              git checkout -f $UPGRADE_COMMIT
              nixos-rebuild switch
            fi
          '';
          environment = {
            POSTGRES_INDEX_URL = xrndCfg.postgresIndexUrl;
          };
          serviceConfig = {
            User = "xrnd";
            Group = "xrn";
            PermissionsStartOnly = true;
          };
        };
    })

    (mkIf (xrndCfg.enable && xrndCfg.restServer) {
        users.groups.xrn = {};

        users.users.xrnrest = {
          isSystemUser = true;
          group = "xrn";
        };

        networking.firewall.allowedTCPPorts = [ 1317 ];

        systemd.services.xrnrest = {
          description = "Regen Ledger REST Server";
          wantedBy = [ "multi-user.target" ];
          after = [ "xrnd.service" ];
          path = [ xrncli ];
          script = ''
            xrncli rest-server --trust-node true
          '';
          serviceConfig = {
            User = "xrnrest";
            Group = "xrn";
            PermissionsStartOnly = true;
          };
        };
    })


    (mkIf (xrndCfg.enable && xrndCfg.restServer) {
        services.postgresql = {
            enable = true;
            enableTCPIP = true;
            package = pkgs.postgresql_11;
            extraPlugins = (pkgs.postgis.override { postgresql = pkgs.postgresql_11; });
            authentication = ''
            '';
            initialScript = ''
            '';
        };
    })
  ];
}
