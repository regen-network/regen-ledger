# Contributing to Regen Ledger

:earth_asia: Welcome! We're glad you're here. Planetary regeneration is a big project, and we can definitely use help building the tools to support regenerative land stewardship. :earth_africa:

## Recommended Reading

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Regen Ledger Docs](https://docs.regen.network)
- [Cosmos SDK Docs](https://docs.cosmos.network/)

## Our Software Development Workflow

We follow an agile methodology and use ZenHub and [GitHub Issues](https://github.com/regen-network/regen-ledger/issues) for ticket tracking. To understand our current priorities and roadmap, check out [GitHub Milestones](https://github.com/regen-network/regen-ledger/milestones).

If you are a first time contributor, check out the issues labeled "[good first issue](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22+)" and "[help wanted](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)" or send us a message in the #regen-ledger channel of our [Discord Server](https://discord.gg/regen-network).

### Using GitHub Labels

We use [GitHub Labels](https://github.com/regen-network/regen-ledger/labels) for issues and pull requests. The following are some general guidelines for using labels.

#### Guidelines for using labels with issues:

- each issue should have one `Type` label
- each issue should have one `Scope` label
- each issue should have one `Status` label
- new issues should always start with `Status: Proposed` or `Status: Bug`
- new issues should be reviewed on a daily basis and updated with the appropriate `Status` label
- ...

#### Guidelines for using labels with pull requests:

- `Type`, `Scope`, and `Status` labels are not required for pull requests because we use semantic commits to define the type and scope of each pull request and because each pull request should have a corresponding issue with the appropriate `Type`, `Scope`, and `Status` labels applied
- ...

### Using Semantic Commits

...

### Writing Protobuf Definitions

...

### Writing Acceptance Tests

With Regen Ledger, we take a [Behaviour Driven Development (BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development) approach to the design and implementation of new features to encourage collaboration among various stakeholders.

After the proto definitions for a new feature are written, and before the new feature is implemented, acceptance tests for the new feature should be written using [Gherkin Syntax](https://cucumber.io/docs/gherkin/).

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

## The Words of Maya Angelou

> We, unaccustomed to courage
exiles from delight
live coiled in shells of loneliness
until love leaves its high holy temple
and comes into our sight
to liberate us into life.
>
> Love arrives
and in its train come ecstasies
old memories of pleasure
ancient histories of pain.
Yet if we are bold,
love strikes away the chains of fear
from our souls.
>
> We are weaned from our timidity
In the flush of love's light
we dare be brave
And suddenly we see
that love costs all we are
and will ever be.
Yet it is only love
which sets us free.
>
> ― Maya Angelou
