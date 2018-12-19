{ config, pkgs, lib, ... }:

with lib;

let
  xrndCfg = config.services.xrnd;
  xrnrestCfg = config.services.xrnrest;
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
    };
    services.xrnrest = {
      enable =
        mkOption {
          type = types.bool;
          default = false;
          description = ''
            Whether to run the xrncli REST server.
          '';
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
          serviceConfig = {
            User = "xrnd";
            Group = "xrn";
            PermissionsStartOnly = true;
          };
        };
    })
  ];
}
