# Contributing

_Welcome!_ We're excited you want to take part in the CrowdStrike community!

Please review this document for details regarding getting started with your first contribution, tools
you'll need to install as a developer, and our development and Pull Request process. If you have any
questions, please let us know by posting your question in the [discussion board](https://github.com/CrowdStrike/gql/discussions).

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How you can contribute](#how-you-can-contribute)
- [Pull Requests](#pull-requests)
    - [License](#License)
    - [Breaking changes](#breaking-changes)
    - [Code Coverage](#code-coverage)
    - [Commit Messages](#commit-message-formatting-and-hygiene)
    - [Pull Request template](#pull-request-template)
    - [Approval/Mergin](#approval--merging) 

## Code of Conduct

Please refer to CrowdStrike's general [Code of Conduct](https://opensource.crowdstrike.com/code-of-conduct/)
and [contribution guidelines](https://opensource.crowdstrike.com/contributing/).

## How you can contribute

- See something? Say something! Submit a [bug report](https://github.com/CrowdStrike/gql/issues/new?assignees=&labels=bug%2Ctriage&template=bug.md&title=) to let the community know what you've experienced or found.
    - Please propose new features on the discussion board first.
- Join the [discussion board](https://github.com/CrowdStrike/gql/discussions) where you can:
    - [Interact](https://github.com/CrowdStrike/gql/discussions/categories/general) with other members of the community
    - [Start a discussion](https://github.com/CrowdStrike/gql/discussions/categories/ideas) or submit a [feature request](https://github.com/CrowdStrike/gql/issues/new?assignees=&labels=enhancement%2Ctriage&template=feature_request.md&title=)
    - Provide [feedback](https://github.com/CrowdStrike/gql/discussions/categories/q-a)
    - [Show others](https://github.com/CrowdStrike/gql/discussions/categories/show-and-tell) how you are using `gql` today
- Submit a [Pull Request](#pull-requests)

## Pull Requests

All code changes should be submitted via a Pull Request targeting the `main` branch. We are not assuming
that every merged PR creates a release, so we will not be automatically creating new SemVer tags as
a side effect of merging your Pull Request. Instead, we will manually tag new releases when required.

### License
When you submit code changes, your submissions are understood to be under the same Unlicense [license](LICENSE) that covers the project.
If this is a concern, contact the maintainers before contributing.

### Breaking changes
In an effort to maintain backwards compatibility, we thoroughly unit test every Pull Request for all 
versions of PowerShell we support. These unit tests are intended to catch general programmatic errors, 
possible vulnerabilities and _potential breaking changes_.

Please fully document unit testing performed within your Pull Request. If you did not specify "Breaking Change" on the 
punch list in the description, and the change is identified as possibly breaking, this may delay or prevent approval of your PR.

### Code Coverage

While we feel like achieving and maintaining 100% code coverage is often an untenable goal with
diminishing returns, any changes that reduce code coverage will receive pushback. We don't want
people to spend days trying to bump coverage from 97% to 98%, often at the expense of code clarity,
but that doesn't mean that we're okay with making things worse.

### Commit Message Formatting and Hygiene

We use [_Conventional Commits_](https://www.conventionalcommits.org/en/v1.0.0/) formatting for commit
messages, which we feel leads to a much more informative change history. Please familiarize yourself
with that specification and format your commit messages accordingly.

Another aspect of achieving a clean, informative commit history is to avoid "noise" in commits.
Ideally, condense your changes to a single commit with a well-written _Conventional Commits_ message
before submitting a PR. In the rare case that a single PR is introducing more than one change, each
change should be a single commit with its own well-written message.

### Pull Request template
Please use the pull request template provided, making sure the following details are included in your request:
+ Is this a breaking change?
+ Are all new or changed code paths covered by unit testing?
+ A complete listing of issues addressed or closed with this change.
+ A complete listing of any enhancements provided by this change.
+ Any usage details developers may need to make use of this new functionality.
  - Does additional documentation need to be developed beyond what is listed in your Pull Request?
+ Any other salient points of interest.

### Approval / Merging
All Pull Requests must be approved by at least one maintainer. Once approved, a maintainer will perform the merge and execute any backend
processes related to package deployment. At this time, contributors _do not_ have the ability to merge to the `main` branch.