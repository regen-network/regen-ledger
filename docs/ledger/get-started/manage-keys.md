# Manage Keys

:::warning
You are responsible for managing your keys. Never share the `mnemonic` used to generate an account that holds tokens of any real value or has sensitive permissions on any network. Make sure you have a secure backup location for your `mnemonic` to recover your account.
:::

## Prerequisites

- [Install Regen](./README.md)

## Configuration

The `regen` binary can store your keys using a few options but only two will be covered here. For more information about the other options, see the documentation for [regen keys](../../commands/regen_keys.md).

The `os` (i.e. "operating system") keyring backend is best for users planning to create an account that will hold tokens of real value or have sensitive permissions. This option uses the same keyring backend as your operating system and stores your encrypted keys in your computer's filesystem.

The `test` keyring backend is best for testing, i.e. when security or recovery is not a concern. When using the `test` keyring backed, the keys are stored in the application's "home" directory.

To check the current configuration, run the following:

```sh
regen config
```

To configure the keyring backend for all commands, run the following:

```sh
regen config keyring-backend [keyring-backend]
```

## Add Key

When you add a key, you are adding a new or existing key to the keyring backend. If you already have an existing key that you would prefer to reuse here, you can use the `--recover` flag. Also, if you have a ledger device, you can use the `--ledger` flag.

To add a key to the keyring backed, run the following command:

```sh
regen keys add [name]
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_keys_add.md).

The `regen` binary is interpreting the key as a `regen` account but the key itself is not specific to the `regen` application. The same key can be used to create an account on another network.

## Show Key

To view your key and account address, run the following command:

```sh
regen keys show [name]
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_keys_show.md).

## List Keys

To view all your keys and account addresses, run the following command:

```sh
regen keys list
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_keys_list.md).
