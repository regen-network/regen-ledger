{ config, pkgs, ... }:                              
let
  xrndCfg = config.services.xrnd;
  xrnrestCfg = config.services.xrnrest;
  xrn = (import ./default.nix);
in
{
  options = {
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
  config = mkIf xrndCfg.enable {
    environment.systemPackages = [ xrn ];
      users.groups.xrn = {};
      users.users.xrnd = {
        isSystemUser = true;
        group = "xrn";
        home = xrndCfg.home;
      };
      systemd.services.xrnd = {
        description = "Regen Ledger Daemon";
        wantedBy = [ "multi-user.target" ];
        after = [ "network.target" ];
        path = [ xrn ];
        preStart = ''
          chown -R xrnd:xrn ${xrndCfg.home}
        '';
        script = ''
          xrnd start --home ${xrndCfg.home}
        '';
        serviceConfig = {
          User = "xrnd";
          Group = "xrn";
          PermissionsStartOnly = true;
        };
      };
  };
}
