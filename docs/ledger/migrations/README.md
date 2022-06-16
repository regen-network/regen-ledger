# Migration Overview

The migration guides within this section are written for developers writing applications that consume the Regen Ledger API. The first migration guide is for Regen Ledger v4.0:

- [Migration Guide v4.0](v4.0-migration.md)

This first page provides an overview of our release process. In addition, we provide some tips on how application developers can stay up to date and prepare ahead of time.

## Release Process

We recommend application developers begin updating their applications once a beta release for a new major version of Regen Ledger is available. A migration guide specific to each major version will be published at the time or soon after the first beta release has been tagged.

Once the migration guide is published, the guide may need to be updated depending on the outcome of the auditing and testing process but new features will not be added and significant changes will not be made unless they are critical to the official release.

The beta release phase will be at least one week depending on the size of the upgrade and the results of the auditing and testing process. The beta release phase may include more than one beta release and we recommend making updates against the latest beta release tag when made available.

The beta release phase ends once a release candidate is tagged. The release candidate phase will be at least one week and may include more than one release candidate. New features will not be added and significant changes will not be made unless they are critical to the official release.

The release candidate phase ends once an official release is tagged. Once the official release is tagged, a software upgrade proposal is submitted on both [Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet). The voting period is currently one week on Regen Mainnet and one day on Redwood Testnet and the block height at which the upgrade will occur is usually within 24-48 hours after the voting period ends.

Once the software upgrade is executed, applications consuming the Regen Ledger API will need to be updated and deployed using the official release to avoid any potential breakage. 

## Staying Up-To-Date

The time between a beta release and the software upgrade might seem too short depending on the size of an application consuming the Regen Ledger API, the number of features dependent upon the API, and the size of the team building the application.

First and foremost, be sure to reach out and let us know who you are and what you're building so that we are aware of your needs and any concerns you may have. We would also love to hear your ideas and any feature requests you may have. Please reach out to us in our [Discord server](https://discord.gg/BDcBJu3).

...

We also recommend application developers become familiar with running a [local testnet](../get-started/local-testnet.html) so that they can test against the latest changes.
