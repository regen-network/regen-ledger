# Contributing to Regen Ledger

:earth_asia: Welcome! We're glad you're here. Planetary regeneration is a big project, and we can definitely use help building the tools to support regenerative land stewardship. :earth_africa:

## Recommended Reading

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Regen Ledger Docs](https://docs.regen.network)
- [Cosmos SDK Docs](https://docs.cosmos.network/)

## Our Software Development Workflow

We follow an agile methodology and use ZenHub and [GitHub Issues](https://github.com/regen-network/regen-ledger/issues) for ticket tracking. To understand our current priorities and roadmap, check out [GitHub Milestones](https://github.com/regen-network/regen-ledger/milestones). If you are a first time contributor, check out the issues labeled ["good first issue"](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22+) and ["help wanted"](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22) or send us a message in the **#regen-ledger** channel of our [Discord Server](https://discord.gg/regen-network).

### Using GitHub Labels

We use [GitHub Labels](https://github.com/regen-network/regen-ledger/labels) for issues and pull requests. The following provides some general guidelines for using labels.

#### Using Labels With Issues

- each issue should have one `Type` label
- each issue should have one `Scope` label
- each issue should have one `Status` label
- new issues should always start with either `Status: Proposed` or `Status: Bug`
- new issues should be reviewed by core contributors on a daily basis and updated with the appropriate `Status` label
- ...

#### Using Labels With Pull Requests

- `Type`, `Scope`, and `Status` labels are not required for pull requests because pull request titles must be written using semantic commits and therefore the title should already include the type and scope of each pull request and because each pull request should have a corresponding issue with the appropriate `Type`, `Scope`, and `Status` labels applied
- ...

### Using Semantic Commits

#### Pull Request Titles

We use "squash and merge" when merging pull requests, which uses the title of the pull request as the merged commit. For this reason, pull requests titles must follow the format of [semantic commits](https://www.conventionalcommits.org/en/v1.0.0/) and should include the appropriate type and scope, and `!` should be added to the type prefix if the pull request introduces an API or client breaking change.

The type and scope of the pull request should already be defined in the issue that the pull request is addressing. The format of the pull request title is checked during the CI process and the allowed types are defined in [this json file](https://github.com/commitizen/conventional-commit-types/blob/v3.0.0/index.json).

The scope is not required and may be excluded if the pull request does not update any golang code within a go module but the scope should be included and should reflect the location of the go module whenever golang code within a go module is updated. Only one go module should be updated at a time but in some cases multiple go modules may be updated. In the case of multiple go modules being updated, the location of the updated go modules should be separated by a `,` and no spaces.

For pull requests that update proto definitions, the scope should reflect the location of the go module within which the code is being generated (e.g. `x/ecocredit` should be the scope when updating proto definitions in `proto/regen/ecocredit` or `proto/regen/ecocredit/basket`). This is also the location that the dummy implementation will be added for new features and changes will be made if updating existing features. This may change in the future given most of the code is now being generated and placed with the `api` go module.

#### Individual Commits

It is not required that each commit within a pull request uses semantic commits but the first commit should use a semantic commit, which will be used instead of the pull request title if the pull request only has one non-merge commit (this will also auto-populate the pull request title when opening a new pull request).

### Writing Protobuf Definitions

...

### Writing Acceptance Tests

With Regen Ledger, we take a [Behaviour Driven Development (BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development) approach to the design and implementation of features to encourage collaboration among various stakeholders. After the proto definitions for a feature are written, and before the feature is implemented, acceptance tests for the feature should be written using [Gherkin Syntax](https://cucumber.io/docs/gherkin/).

Writing BDD-style tests provide value at three phases of development:

- during the design and specification phase to get alignment on intended behavior, oftentimes with stakeholders who may not be fluent in golang or coding at all
- during the implementation and pull request review phase to ensure tests are being written to test the intended behavior and not just to satisfy code coverage 
- during the documentation phase to provide base-level documentation that acts as a human-readable source of truth for the intended behavior of a feature

With BDD-style tests, the approach is as follows:

- What are you building?
- How will you test it?
- How did you build it?
- How did you build the tests?

...

#### Rules

- "Rule is a synonym for business requirement and acceptance criterion."
- ...

#### Scenarios

- "Scenarios should read like a specification."
- "Scenarios should be thought of as documentation, rather than tests."
- "Scenarios should enable collaboration between business and delivery, not prevent it."
- "Scenarios should support the evolution of the product, rather than obstruct it."
- "Each scenario should only illustrate a single rule."
- "Shorter scenarios are easier to read, understand and maintain."
- ...

#### Steps

Given:

- ...

When:

- ...

Then:

- ...

#### Resources

- [https://cucumber.io/docs/cucumber/](https://cucumber.io/docs/cucumber/)
- [https://leanpub.com/bddbooks-discovery](https://leanpub.com/bddbooks-discovery)
- [https://leanpub.com/bddbooks-formulation](https://leanpub.com/bddbooks-formulation)

### Implementing Features

...

### Reviewing Pull Requests

...

### Writing Documentation


#### Best Practices

- Always double-check for spelling and grammar.
- Avoid using `code` format when writing in plain English.
- Try to express your thoughts in a clear and concise and clean way.
- RFC keywords should be used in technical documentation (uppercase) and are recommended in user documentation (lowercase). The RFC keywords are: "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL". They are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

#### Resources

- [https://developers.google.com/style](https://developers.google.com/style)
- [https://developers.google.com/tech-writing/overview](https://developers.google.com/tech-writing/overview)

#### Auto-Generated Documentation

- Protobuf documentation is auto-generated and served on [Buf Schema Registry](https://buf.build/regen/regen-ledger/docs) using the comments written in the proto files and acts as a source of truth for the API of the application.
- Documentation for each feature is auto-generated and served on [docs.regen.network](https://docs.regen.network) using the feature files written in Gherkin Syntax and acts as a source of truth for the intended behavior of each feature.
- CLI documentation is auto-generated and served on [docs.regen.network](https://docs.regen.network) using [cobra/doc](https://pkg.go.dev/github.com/spf13/cobra/doc) and acts as a source of truth for the CLI commands available when using the `regen` binary.

## The Words of Maya Angelou

> We, unaccustomed to courage
> exiles from delight
> live coiled in shells of loneliness
> until love leaves its high holy temple
> and comes into our sight
> to liberate us into life.
>
> Love arrives
> and in its train come ecstasies
> old memories of pleasure
> ancient histories of pain.
> Yet if we are bold,
> love strikes away the chains of fear
> from our souls.
>
> We are weaned from our timidity
> In the flush of love's light
> we dare be brave
> And suddenly we see
> that love costs all we are
> and will ever be.
> Yet it is only love
> which sets us free.
>
> ― Maya Angelou
