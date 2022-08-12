# Contributing

:earth_asia: Welcome! We're glad you're here. Planetary regeneration is a big project, and we can use your help building infrastructure and tools to support regenerative land stewardship. :earth_africa:

We follow an agile methodology and use ZenHub and [GitHub Issues](https://github.com/regen-network/regen-ledger/issues) for ticket tracking. To understand our current priorities and roadmap, check out [GitHub Milestones](https://github.com/regen-network/regen-ledger/milestones). If you're a first time contributor, check out issues labeled "[good first issue](https://github.com/regen-network/regen-ledger/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22+)".

If you have questions about a specific issue or you would like to start working on an issue, please comment in the issue. If you have general questions, feel free to reach out in the **#regen-ledger** channel of our [Discord Server](https://discord.gg/regen-network).

Before you begin your journey, please read [Code of Conduct](CODE_OF_CONDUCT.md). Also, please do your best to search for answers before asking what may seem like common questions, and we will do our best to update our documentation to include the answers.

### Table of Contents

- [General Guidelines](#general-guidelines)
  - [Submitting Issues](#submitting-issues)
  - [Reviewing Proposals](#reviewing-proposals)
  - [Submitting Pull Requests](#submitting-pull-requests)
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

### Recommended Reading

- [Regen Ledger Docs](https://docs.regen.network)
- [Cosmos SDK Docs](https://docs.cosmos.network/)

### Additional Documentation

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Release Process](RELEASE_PROCESS.md)
- [Security Policy](SECURITY.md)

## General Guidelines

The following guidelines cover the basics, including submitting issues and pull requests, reviewing proposals and pull requests, and using labels and semantic commits.

### Submitting Issues

Please use the GitHub interface and [the provided templates](https://github.com/regen-network/regen-ledger/issues/new/choose) when submitting a new issue.

#### For Bugs

***Important: Do not open up a GitHub issue when reporting a security vulnerability. Reporting security vulnerabilities must go through the process defined within our [security policy](SECURITY.md).***

Please make sure you do the following:
- provide a concise description of the issue
- provide a git commit hash or release version
- provide steps on how to reproduce the issue

#### For Features

Issues for features should have a clear user story or user story breakdown. The whole development process affects the quality of the work that ends up in a pull request, starting with a user story.

Please make sure you do the following:
- provide a concise description of the feature
- provide a rationale for including the feature
- provide a description of the requirements

### Reviewing Proposals

When reviewing a feature request, please be conscientious of the proposer and the work they put into submitting the request. If you disagree with the feature or the approach, make sure you provide constructive feedback and clearly explain your rationale.

### Submitting Pull Requests

In addition to the guidelines here, please review [Coding Guidelines](#coding-guidelines) and/or [Documentation Guidelines](#documentation-guidelines) before submitting a pull request.

Ideally a pull request addresses a single issue. In some cases a pull request may address multiple issues or a pull request may not have a corresponding issue. If the necessity of the change is not obvious, please open an issue first.

Another way of framing this, "One pull request, one concern." A pull request to fix a bug should only fix a bug, a pull request to add a new feature should only add a new feature, a pull request to refactor code should only refactor code, etc.

#### Getting Started

Unless you are a core team member or a contributor who has been granted limited write access, you will first need to fork the repository before you can submit a pull request. To fork the repository, [click this link](https://github.com/regen-network/regen-ledger/fork), select the owner, and then click "Create fork".

Once you have forked the repository, you can clone the repository, create a feature branch, and start making changes locally. If you have any questions about how to use Git with GitHub, we recommend checking out [GitHub Git Guides](https://github.com/git-guides). 

If you are making documentation changes to a single page, you can also use the GitHub editor. For example, to edit this document, you would click the edit icon in the top right corner, make the changes, scroll to the bottom, and fill out and submit "Commit changes".

#### Feature Branch
 
The feature branch should follow the format `<author>/<issue-number>-<description>`:

- `author` - your first name or username (should be unique, consistent, and non-ambiguous)
- `issue-number` - the issue number that the feature branch is addressing (if applicable)
- `description` - a very concise description of the changes (should be no more than a few words)

#### Requesting Reviews

Core team members will automatically be assigned to review each pull request upon submission based on [CODEOWNERS](./CODEOWNERS). If the pull request is not ready for an official review, open the pull request as a draft and keep it open as a draft until it is ready. If you would like initial feedback on a draft pull request, leave a comment tagging the individual or @regen-network/ledger-core-dev requesting initial feedback.

Before opening a pull request (either directly or converting from draft), please make sure you have added a description and that you have read through and completed the "Author Checklist" items within the pull request template.

### Reviewing Pull Requests

Before approving a pull request, please make sure you have read through and completed the "Reviewer Checklist" items within the pull request template, or you have checked off the individual items that you have reviewed.

In addition, reviewers should use one of the following keywords in the review comment:

- ACK or LGTM - you reviewed the changes, they look good, and testing does not apply (usually with an approval)
- cACK - "Concept ACK" - you agree with the idea or direction but haven't done a thorough review (usually without an approval)
- utACK - "Untested ACK" - you reviewed the changes but haven't performed any manual tests (usually with an approval)
- tACK - "Tested ACK" - you reviewed the changes and performed the necessary manual tests (usually with an approval)

### Using GitHub Labels

We use [GitHub Labels](https://github.com/regen-network/regen-ledger/labels) for issues and pull requests. The following provides some general guidelines for using labels.

#### Using Labels With Issues

- each issue should have one `Type` label
- each issue should have one `Scope` label
- each issue should have one `Status` label
- new issues should always start with either `Status: Proposed` or `Status: Bug`

#### Using Labels With Pull Requests

- `Type`, `Scope`, and `Status` labels are not required for pull requests because pull request titles must be written using semantic commits (i.e. the title should include the type and scope of the pull request) and because each pull request should have a corresponding issue with the appropriate `Type`, `Scope`, and `Status` labels applied

### Using Semantic Commits

The [semantic commit](https://www.semanticcommits.org/en/v1.0.0/) format is required for pull request titles, which will be used as the commit message with "squash and merge". The first commit of a new branch should be a semantic commit but semantic commits are not required for individual commits. 

General guidelines to keep in mind:

- `build`, `ci`, and `chore` should only be used when updating non-production code
- `fix`, `feat`, `refactor`, and `perf` should only be used when updating production code
- `test` should only be used when updating tests
- `style` should only be used when updating format
- `docs` should only be used when updating documentation

Also, to write useful commit message descriptions, it’s good practice to start with a verb, such as “add”, “remove”, "update", or “fix”.

#### Pull Request Titles

We use "squash and merge" when merging pull requests, which uses the title of the pull request as the merged commit. For this reason, pull requests titles must follow the format of semantic commits and should include the appropriate type and scope, and `!` should be added to the type prefix if the pull request introduces an API or client breaking change and will therefore require a major release.

The appropriate type and scope of the pull request should be provided by the labels of the corresponding issue but the type and scope may need to be adjusted during the implementation. If the type and scope need to be adjusted during the implementation, the label used in the corresponding issue should be updated to reflect those changes. The format of the pull request title is verified during the CI process and the allowed type prefixes are defined in [this json file](https://github.com/commitizen/semantic-commit-types/blob/v3.0.0/index.json).

The scope is not required and may be excluded if the pull request does not update any golang code within a go module but the scope should be included and should reflect the location of the go module whenever golang code within a go module is updated. Only one go module should be updated at a time but in some cases multiple go modules may be updated. In the case of multiple go modules being updated, the location of the updated go modules should be separated by a `,` and no spaces.

For pull requests that update proto files, the scope should reflect the location of the go module within which the code will be generated (e.g. `x/ecocredit` should be the scope when updating proto files in `proto/regen/ecocredit` or `proto/regen/ecocredit/basket`). This is also the location of the go module in which a dummy implementation will be added when implementing new features and where changes will be made when updating existing features. Note, this use of scope in relation to updating proto files may change in the future.

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

#### Individual Commits

It is not required that individual commits within a pull request use semantic commits but the first commit should be a semantic commit. The first commit message is used to auto-populate the pull request title when opening a new pull request. If the pull request only has one non-merge commit, the first commit is used to auto-populate the commit message when using "squash and merge".

## Coding Guidelines

The following guidelines are for writing code.

### Writing Proto Files

[ work in progress... check back soon ]

### Writing Feature Files

[ work in progress... check back soon ]

#### Resources

- [https://cucumber.io/docs/cucumber/](https://cucumber.io/docs/cucumber/)
- [https://leanpub.com/bddbooks-discovery](https://leanpub.com/bddbooks-discovery)
- [https://leanpub.com/bddbooks-formulation](https://leanpub.com/bddbooks-formulation)

### Writing Golang Code

The following are some general guidelines when writing golang code:

- optimize for readability
- use tabs rather than spaces
- end of file should have an extra line
- imports should be alphabetical
- imports should be divided into sections with a line between each section
  - native go packages
  - external packages
  - internal packages
- structs declared with more than two properties should be declared with a line for each property
- the first word in a comment should only be capitalized if the comment is a complete sentence or if the first word should be capitalized regardless of its location within the comment (e.g. if a public function name)
- comments should only use a period if the comment is a complete sentence
- when adding a comment to explain code, first consider changing the code to be more self documenting

#### Module File Names

- each message implementation should have its own file and include the full name of the proto message (e.g. `msg_create_class.go`)
- each message server method should have its own file and include the full name of the proto message and the method name should be prefixed with `msg_` to indicate the method is part of the message server (e.g. `msg_create_class.go`)
- each query server method should have its own file and include the full name of the proto message and the method name should be prefixed with `query_` to indicate the method is part of the query server (e.g. `query_class_info.go`)

### Writing Golang Tests

[ work in progress... check back soon ]

## Documentation Guidelines

The following guidelines are for writing documentation.

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

[ work in progress... check back soon ]

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
