{ config, pkgs, lib, ... }:

with lib;

let
  xrndCfg = config.services.xrnd;
  xrn_build = (import ./default.nix);
  xrnd = xrn_build.xrnd;
  xrncli = xrn_build.xrncli;
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
      enablePostgres =
        mkOption {
          type = types.bool;
          default = false;
          description = "Automatically enable the Postgresql service and index to a database named xrn. Shouldn't be used together with postgresUrl";
        };
      postgresUrl =
        mkOption {
          type = types.str;
          default = "";
          description = "The URL of a Postgresql database to index to. Shouldn't be used together with enablePostgres";
        };
    };
  };
  config = mkMerge [
    (mkIf config.programs.xrn.enable {
      environment.systemPackages = [ xrncli xrnd ];
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
            if [ -e ${xrndCfg.home}/data/upgrade-info ]; then
              mv ${xrndCfg.home}/data/upgrade-info ${xrndCfg.home}/data/upgrade-info.bak
            fi
          '';
          script = ''
            xrnd start --moniker ${xrndCfg.moniker} --home ${xrndCfg.home}
          '';
          postStop = ''
            if [ -e ${xrndCfg.home}/data/upgrade-info ]; then
              mv ${xrndCfg.home}/data/upgrade-info ${xrndCfg.home}/data/upgrade-info.bak
              export UPGRADE_COMMIT=$(jq '.commit' < ${xrndCfg.home}/data/upgrade-info.bak)
              if  [ $UPGRADE_COMMIT != "null" ]; then
                cd /root/regen-ledger
                git clean -f
                git checkout -f $UPGRADE_COMMIT
                nixos-rebuild --upgrade switch
              fi
            fi
          '';
          environment = {
            POSTGRES_INDEX_URL = if xrndCfg.enablePostgres then "host=/tmp user=xrnd dbname=xrn sslmode=disable" else xrndCfg.postgresUrl;
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
            extraPlugins = [(pkgs.postgis.override { postgresql = pkgs.postgresql_11; })];
            initialScript = pkgs.writeText "backend-initScript" ''
              CREATE USER xrnd; 
              CREATE DATABASE xrn;
              CREATE EXTENSIOn postgis;
              GRANT ALL PRIVILEGES ON DATABASE xrn TO xrnd;
              CREATE USER guest;
	          GRANT SELECT ON ALL TABLES IN SCHEMA public to PUBLIC;
            '';
            authentication = ''
              local all xrnd trust
              host xrn guest 0.0.0.0/0 trust
              host xrn guest ::0/0 trust
            '';
        };
        # Open fire-wall port for production. WARNING don't put this into production validators:
       	networking.firewall.allowedTCPPorts = [ 5432 ];
    })
  ];
}
