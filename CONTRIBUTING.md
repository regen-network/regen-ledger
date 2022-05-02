# Contributing to Regen Ledger

:earth_asia: Welcome! We're glad you're here. Planetary regeneration is a big project, and we can definitely use help building the tools to support regenerative land stewardship. :earth_africa:

We follow an agile methodology and use ZenHub and [GitHub Issues](https://github.com/regen-network/regen-ledger/issues) for ticket tracking. To understand our current priorities and roadmap, check out [GitHub Milestones](https://github.com/regen-network/regen-ledger/milestones). If you are a first time contributor, check out the issues labeled "[good first issue](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22+)" and "[help wanted](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)" or send us a message in the **#regen-ledger** channel of our [Discord Server](https://discord.gg/regen-network).

### Table of Contents

- [Workflow Guidelines](#workflow-guidelines)
  - [Submitting Issues](#submitting-issues)
  - [Reviewing Proposals](#reviewing-proposals)
  - [Submitting Pull Requests](#submitting-pull-requests)
  - [Requesting Reviews](#requesting-reviews)
  - [Reviewing Pull Requests](#reviewing-pull-requests)
  - [Using GitHub Labels](#using-github-labels)
  - [Using Semantic Commits](#using-semantic-commits)
- [Coding Guidelines](#coding-guidelines)
  - [Writing Proto Files](#writing-proto-files)
  - [Writing Feature Files](#writing-feature-files)
  - [Writing Golang Code](#writing-golang-code)
  - [Writing Golang Tests](#writing-golang-tests)
- [Documentation Guidelines](#documentation-guidelines)
  - [Writing Guidelines](#writing-guidelines)
  - [Writing Documentation](#writing-documentation)
  - [Writing Specifications](#writing-specifications)

<!--
- [Core Contributor Guidelines](#core-contributor-guidelines)
  - [Meetings](#meetings)
  - [Estimating Issues](#estimating-issues)
-->

### Recommended Reading

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Regen Ledger Docs](https://docs.regen.network)
- [Cosmos SDK Docs](https://docs.cosmos.network/)

### Additional Documentation

- [Release Process](RELEASE_PROCESS.md)
- [Security Policy](SECURITY.md)

## Workflow Guidelines

The following guidelines provide an overview of our workflow.

### Submitting Issues

Please use the GitHub interface and [the provided templates](https://github.com/regen-network/regen-ledger/issues/new/choose) when submitting a new issue.

#### For Bugs

**Do not open up a GitHub issue when reporting a security vulnerability. Reporting security vulnerabilities must go through the process defined within our [security policy](SECURITY.md).**

...

#### For Features

Issues for features should have a clear user story or user story breakdown. The whole development process affects the quality of the work that ends up in a pull request, starting with a user story.

### Reviewing Proposals

...

### Submitting Pull Requests

In addition to the following guidelines, please review our [Coding Guidelines](#coding-guidelines) and/or [Documentation Guidelines](#documentation-guidelines) before opening a pull request.

If you are submitting minor documentation changes, you can use the GitHub editor. If you are making code changes or more significant documentation changes, you will need to fork the repository and push your changes to the forked repository.

...

<!--
Notes:

- Atomic pull requests. One pull request, one concern.
- "One of the bad practices of a Pull Request is changing things that are not concerned with the functionality that is being addressed, like whitespace changes, typo fixes, variable renaming, etc. If those things are not related to the concern of the Pull Request, it should probably be done in a different one."
- "...cleanup doesn't need to be done in the same Pull Request, the important thing is not leaving the codebase in a bad state after finishing the functionality. If you must, refactor the code in a separate Pull Request, and preferably before the actually concerned functionality is developed, because then if there is a need in the near future to revert the Pull Request, the likelihood of code conflict will be lower."
-->

#### Requesting Reviews

...

<!--
Notes:

Two approaches to consider:

1. Assign core contributors who need to review the pull request before it gets merged.
2. Only request reviews from core contributors who need to review the pull request before it gets merged.
-->

### Reviewing Pull Requests

...

<!--
Notes:

What does it mean for a PR to be approved? What does approval include? What does approval not include? Should core contributors have specific roles in the review process? For example, should one core contributor own API design and naming, another own manual testing, another own documentation (tracking inconsistencies, missing documentation, etc.)?

A designated reviewer for each task?

- API Design and Naming
- Tests and Test Coverage
- State Maching Logic
- Manual Testing (if needed)
- Documentation Changes

Pull requests that implement a new command or a new endpoint should always be manually tested. If a pull request is blocked due to manual testing, the manual test required should be added to a tracking issue... the epic for the release? the module readiness checklist? a separate tracking issue specific to manual tests that need to be performed?

Should we have a tracking issue for documentation needs that we add to during the review process? Should documentation updates be included in the code contributions?
-->

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

- `Type`, `Scope`, and `Status` labels are not required for pull requests because pull request titles must be written using semantic commits (i.e. the title should include the type and scope of the pull request) and because each pull request should have a corresponding issue with the appropriate `Type`, `Scope`, and `Status` labels applied
- ...

### Using Semantic Commits

We use [semantic commits](https://www.semanticcommits.org/en/v1.0.0/) for pull request titles and the first commit of a new branch. Semantic commits are not required for individual commits within a pull request. 

- `fix`, `feat`, `refactor`, and `perf` should only be used when updating production code
- `test` should only be used when updating tests
- `style` should only be used when no production code is changed
- `docs` should only be used when updating documentation
- `build`, `ci`, and `chore` should only be used when updating non-production code

To write useful commit message descriptions, it’s good practice to start the description with a verb, such as “add”, “remove”, "update", or “fix”, followed by what was done.

#### Pull Request Titles

We use "squash and merge" when merging pull requests, which uses the title of the pull request as the merged commit. For this reason, pull requests titles must follow the format of semantic commits and should include the appropriate type and scope, and `!` should be added to the type prefix if the pull request introduces an API or client breaking change and will therefore require a major release.

The appropriate type and scope of the pull request should be provided by the labels of the corresponding issue but the scope and type may need to be adjusted during the implementation. If the type and scope need to be adjusted during the implementation, the labels of the corresponding issue should be updated to reflect those changes. The format of the pull request title is verified during the CI process and the allowed type prefixes are defined in [this json file](https://github.com/commitizen/semantic-commit-types/blob/v3.0.0/index.json).

The scope is not required and may be excluded if the pull request does not update any golang code within a go module but the scope should be included and should reflect the location of the go module whenever golang code within a go module is updated. Only one go module should be updated at a time but in some cases multiple go modules may be updated. In the case of multiple go modules being updated, the location of the updated go modules should be separated by a `,` and no spaces.

For pull requests that update proto files, the scope should reflect the location of the go module within which the code will be generated (e.g. `x/ecocredit` should be the scope when updating proto files in `proto/regen/ecocredit` or `proto/regen/ecocredit/basket`). This is also the location of the module in which a dummy implementation will be added when implementing new features and where changes will be made when updating existing features. This use of scope in relation to updating proto files might change in the future.

Examples of pull request titles using semantic commits:

```
docs: add examples to contributing guidelines
feat(x/ecocredit): add query for all credit classes
fix(x/data): update attest to store correct timestamp
refactor(x/ecocredit): update location to jurisdiction
style(x/data): format proto files and fix whitespace
perf(x/ecocredit): move redundant calls outside of for loop
test(x/data): implement acceptance tests for anchoring data
build: bump cosmos-sdk version to latest patch release
ci: add github action workflow for proto lint check
chore: delete test output and add to gitignore
```

Examples of API breaking changes (`!` required):

- ...

Common misconceptions for API breaking changes:

- adding a new message type

Examples of client breaking changes (`!` required):

- ...

Common misconceptions for client breaking changes:

- ...

#### Individual Commits

It is not required that individual commits within a pull request use semantic commits but the first commit should be a semantic commit. The first commit message is used to auto-populate the pull request title when opening a new pull request. If the pull request only has one non-merge commit, the first commit is used to auto-populate the commit message when using "squash and merge".

## Coding Guidelines

The following guidelines provide an overview of our workflow.

### Writing Proto Files

...

### Writing Feature Files

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

<!--
Notes:

simplicity and readability over duplication and redundancy when writing test code
-->

#### Rules

- "Rule is a synonym for business requirement and acceptance criterion."
- ...

#### Background

<!--
Notes:

"We believe that the Background keyword should not be used, because it has a negative effect on the readability of the scenarios."
-->

#### Scenarios

- "The goal of a scenario is to illustrate a rule."
- "Any change to a scenario is effectively a change of the specification."
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

<!--
Notes:

"My advice is to focus on consistency instead of reusability. Once you find a consistent way of phrasing your steps, you will benefit from step reusability anyway."
-->

#### Resources

- [https://cucumber.io/docs/cucumber/](https://cucumber.io/docs/cucumber/)
- [https://leanpub.com/bddbooks-discovery](https://leanpub.com/bddbooks-discovery)
- [https://leanpub.com/bddbooks-formulation](https://leanpub.com/bddbooks-formulation)

### Writing Golang Code

The following are some general guidelines when writing golang code:

- optimize for readability
  - "This is open source software. Consider the people who will read your code, and make it look nice for them."
- use tabs rather than spaces
- end of file should have extra line
- imports should be alphabetical
- imports should be divided into sections with a line between each section
  - native go packages
  - external packages
  - internal packages
- structs declared with more than two properties should be declared with a line for each property
- the first word in a comment should only be capitalized if the comment is a complete sentence or if the first word should be capitalized regardless of its location within the comment (e.g. if a public function name)
- comments should only use a period if the comment is a complete sentence
- when adding a comment to explain code, first consider changing the code to be more self documenting

...

- each message implementation should have its own file and include the full name of the proto message (e.g. `msg_create_class.go`)
- each message server method should have its own file and include the full name of the proto message and the method name should be prefixed with `msg_` to indicate the method is part of the message server (e.g. `msg_create_class.go`)
- each query server method should have its own file and include the full name of the proto message and the method name should be prefixed with `query_` to indicate the method is part of the query server (e.g. `query_class_info.go`)

### Writing Golang Tests

...

## Documentation Guidelines

...

### Writing Guidelines

- Always double-check for spelling and grammar.
- Avoid using `code` format when writing in plain English.
- Try to express your thoughts in a clear and precise way.
- RFC keywords should be used in technical documentation (uppercase) and are recommended in user documentation (lowercase). The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

#### Resources

- [https://developers.google.com/style](https://developers.google.com/style)
- [https://developers.google.com/tech-writing/overview](https://developers.google.com/tech-writing/overview)

### Writing Documentation

Regen Ledger documentation is hosted at [docs.regen.network](https://docs.regen.network) and the site is statically generated using [Vuepress](https://vuepress.vuejs.org/). Each webpage is generated from a markdown file that is either written and maintained within the `docs` directory or the `spec` folder of a module (e.g. `x/ecocredit/spec`), or the markdown file is auto-generated when building the site (see [Auto-Generated Documentation](#auto-generated-documentation)).

To start up a local development server for the documentation site, run the following make command:

```bash
make docs-dev
```

#### Auto-Generated Documentation

- Protobuf documentation is auto-generated and served on [Buf Schema Registry](https://buf.build/regen/regen-ledger/docs). The documentation is auto-generated using the comments provided in the proto files. This documentation acts as the source of truth for the API of the application.
- Documentation for each feature is auto-generated and served on [docs.regen.network](https://docs.regen.network). The documentation is auto-generated using the feature files written in Gherkin Syntax. This documentation acts as the source of truth for the intended behavior of each feature.
- CLI documentation is auto-generated and served on [docs.regen.network](https://docs.regen.network). The documentation is auto-generated using [the cobra command properties](https://pkg.go.dev/github.com/spf13/cobra#Command). This documentation acts as the source of truth for the commands available when using the `regen` binary.

### Writing Specifications

...

<!--
### Meetings

An agenda for each meeting (except for the daily scrum) should be prepared ahead of time and shared with core contributors to keep meetings focused and to the point and to provide an opportunity for other core contributors to add agenda items.

#### Daily Scrum

Core contributors meet on a daily basis for 15 minutes to coordinate activities.

#### Weekly Scrum

Core contributors meet on a weekly basis for 60 minutes to groom the backlog and coordinate activities.

#### Monthly All Hands

Core contributors meet with core contributors from related projects to present updates and coordinate activities.

#### Quarterly Retrospective

Core contributors meet with core contributors from related projects to reflect on processes.

### Estimating Issues

Issues should have estimates and estimates should be justified. Asking someone to justify an estimate has been shown to result in estimates that better compensate for missing information.
-->

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
