# Migration Overview

The migration guides within this section are written for developers writing applications that consume the Regen Ledger API.

- [Migration Guide v4.0](v4.0-migration.md)
- [Migration Guide v5.0](v5.0-migration.md)

This first page provides an overview of our release process. In addition, we provide some tips on how application developers can stay up to date and prepare ahead of time.

## Release Process

We recommend application developers begin updating their applications once a beta release for a new major version of Regen Ledger is available. A migration guide specific to each major version will be published at the time or soon after the first beta release has been tagged.

Once the migration guide is published, the guide may need to be updated depending on the outcome of the auditing and testing process. Some features might be added/updated and some changes might be made but the beta release should be relatively stable.

The beta release phase will be at least one week depending on the size of the release and the results of the auditing and testing process. The beta release phase may include more than one beta release and we recommend making updates against the latest beta release when made available.

The beta release phase ends once a release candidate is tagged. The release candidate phase will be at least one week and may include more than one release candidate. New features will not be added and significant changes will not be made unless they are critical to the official release.

The release candidate phase ends once an official release is tagged and a software upgrade proposal is submitted on both [Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet). The voting period is currently one week on Regen Mainnet and one day on Redwood Testnet and the block height at which the upgrade will occur is usually within 24-48 hours after the voting period ends.

Once the software upgrade is executed, applications consuming the Regen Ledger API will need to be deployed with the migrations in order to avoid any potential breakage.

## Staying Up-To-Date

The time between a beta release and the software upgrade might seem a bit short depending on the size of the application consuming the Regen Ledger API, the number of features dependent upon the API, and the size of the team building the application.

First and foremost, be sure to reach out and introduce yourself if you have not already and please share any questions or concerns you may have. We would love to work alongside teams building on top of Regen Ledger to ensure we are providing the best possible developer experience.

Prior to a beta release tag, application developers might want to consider spinning up a single node network with a specific commit that includes a new feature or significant changes that will require a fair amount of work to implement. We would be happy to assist you in setting this up.

Similar to deciding upon a commit and running a single node network on a remote server, application developers might also want to consider spinning up a local testnet as part of their workflow or more specifically for testing. We would be happy to assist you in setting this up as well.

Once a beta release has been tagged, the Regen Ledger team spins up a temporary test network that will run up until the official upgrade of Regen Mainnet and Redwood Testnet. If you would like to test against this temporary test network, please reach out and we can keep you posted on updates.