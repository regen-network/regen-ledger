# Release Process

This document outlines the process for releasing new versions of Regen Ledger.

## Semantic Versioning

Regen Ledger follows [Semantic Versioning](https://semver.org/):

> 1. MAJOR version when you make incompatible API changes
> 1. MINOR version when you add functionality in a backwards compatible manner
> 1. PATCH version when you make backwards compatible bug fixes

Within the context of a blockchain application, we amend the summary to include the following:

1. MAJOR version when you make incompatible API and client changes
1. MINOR version when you make incompatible state machine changes and backwards compatible API and client changes 
1. PATCH version when you make backwards compatible bug fixes not affecting the API, client, and state machine

## Major Release Process

A major release is an increment of the first number in the version tag (e.g. `v1.1.1` → `v2.0.0`).

A major release has no restrictions on what changes it can or cannot include but frequent incompatible API and client changes should not be introduced lightly. Upgrading a live network to a major release requires a coordinated effort among many stakeholders. For more information about upgrading a live network, see [Upgrade Overview](https://docs.regen.network/validators/migrations/upgrade).

Prior to each major release, at least one beta release and one release candidate must be published. The tag of each pre-release must include a sequence number scoped to the phase of the release. For example, the pre-releases leading up to an official `v2.0.0` release must occur in the following order with the following format:

```
v2.0.0-beta1 → v2.0.0-beta2 → ... → v2.0.0-rc1 → v2.0.0-rc2 → ... → v2.0.0
```

The release process for a major release starts once all changes for the major release have been implemented on the `master` branch. The release process is divided into three phases: beta release, release candidate, and official release.

### Beta Release

- Tag a beta release on the `master` branch and prevent any further changes unrelated to the release.
  - When tagging a beta release, use an annotated git tag (e.g. `git tag -a v2.0.0-beta1`).
- Perform an audit of all new changes and add additional automated tests and test cases if needed.
- Perform a software upgrade on a temporary test network and all manual tests against that network.
- If any issues are found, implement the necessary changes and then start over with a new beta release.
- If no issues were found, move on to the next phase of the release.

### Release Candidate

- Create a new release branch from `master` using the format `release/vX` for the branch name.
- Ensure the release branch is protected so that pushes are only permitted by the release managers.
- Add a backport label with the format `backport/vX` and update the mergify backport integration.
- In the release branch, update `CHANGELOG.md` to reflect the changes for the release candidate.
  - The first release candidate must include all changes since the last official release.
  - All other release candidates must include all changes since the last release candidate.
- In the release branch, tag a release candidate and prevent any further changes unrelated to the release.
  - When tagging a release candidate, use an annotated git tag (e.g. `git tag -a v2.0.0-rc1`).
  - Once a release candidate has been tagged, the `master` branch is no longer restricted to release changes.
- Push the release tag and update the release description to include the changes for the release candidate.
- Perform an audit of all new changes and add additional automated tests and test cases if needed.
- Perform a software upgrade on a temporary test network and all manual tests against that network.
- If any issues are found, implement the necessary changes and then start over with a new release candidate.
  - Any changes made at this stage should be done on `master` and then backported to the release branch.
- If no issues were found, move on to the next phase of the release.

### Official Release

- In the `master` branch, update `CHANGELOG.md` to reflect the changes for the official release.
  - The official release must include all changes since the last official release.
  - The updates should be merged to `master` and then backported to the release branch.
- In the release branch, tag the official release using an annotated git tag (e.g. `git -a v2.0.0`).
- Push the release tag and update the release description to include the changes for the official release.

## Minor Release Process

A minor release is an increment of the second number in the version tag (e.g. `v1.1.1 → v1.2.0`).

**A minor release must not introduce incompatible API and client changes.**

Any changes to a release branch should be done on `master` and then backported to the release branch. The process for a minor release may vary depending on the number and significance of the changes made.

- Perform an audit of all new changes and add additional automated tests and test cases if needed.
- Perform a software upgrade on a temporary test network and all manual tests against that network.
- If any issues are found, implement the necessary changes and then start over.
- In the `master` branch, update `CHANGELOG.md` to reflect the changes for the minor release.
  - The minor release must include all changes since the last release.
  - The updates should be merged to `master` and then backported to the release branch.
- In the release branch, tag the minor release using an annotated git tag (e.g. `git -a v1.2.0`).
- Push the release tag and update the release description to include the changes for the minor release.

## Patch Release Process

A patch release is an increment of the third number in the version tag (e.g. `v1.1.1` → `v1.1.2`).

**A patch release must not introduce incompatible API, client, or state machine changes.**

Any changes to a release branch should be done on `master` and then backported to the release branch. The process for a patch release may vary depending on the number and significance of the changes made.

- Perform an audit of all new changes and add additional automated tests and test cases if needed.
- If any issues are found, implement the necessary changes and then start over.
- In the `master` branch, update `CHANGELOG.md` to reflect the changes for the patch release.
  - The patch release must include all changes since the last release.
  - The updates should be merged to `master` and then backported to the release branch.
- In the release branch, tag the patch release using an annotated git tag (e.g. `git -a v1.1.2`).
- Push the release tag and update the release description to include the changes for the patch release.
