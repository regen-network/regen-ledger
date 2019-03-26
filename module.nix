{ config, pkgs, lib, ... }:

with lib;

let
  xrndCfg = config.services.xrnd;
  xrn = (import ./default.nix);
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
      postgresIndexUrl =
        mkOption {
          type = types.str;
          default = "";
          description = "The URL of the Postgres server to index to. Postgres indexing will be disabled if this is not set.";
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
          path = [ xrn ];
          preStart = ''
            chown -R xrnd:xrn ${xrndCfg.home}
          '';
          script = ''
            xrnd start --moniker ${xrndCfg.moniker} --home ${xrndCfg.home}
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
          path = [ xrn ];
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
  ];
}
