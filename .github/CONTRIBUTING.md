# Contributing to terraform-plugin-log

**First:** if you're unsure or afraid of _anything_, just ask
or submit the issue describing the problem you're aiming to solve.

Any bug fix or feature has to be considered in the context of the wider
Terraform ecosystem. This is great, as your contribution can have a big
positive impact, but we have to assess potential negative impact too (e.g.
breaking existing providers which may not use a new feature).

To provide some safety to the wider provider ecosystem, we strictly follow
[semantic versioning](https://semver.org/) and any changes that we consider
breaking changes will only be released as part of major releases. **Note:**
while we're on v0, breaking changes will be accepted during minor releases.

## Table of Contents

 - [I just have a question](#i-just-have-a-question)
 - [I want to report a vulnerability](#i-want-to-report-a-vulnerability)
 - [New Issue](#new-issue)
 - [New Pull Request](#new-pull-request)

## I just have a question

> **Note:** We use GitHub for tracking bugs and feature requests related to
> terraform-plugin-log.

For questions, please see relevant channels at
https://www.terraform.io/community.html

## I want to report a vulnerability

Please disclose security vulnerabilities responsibly by following the procedure
described at https://www.hashicorp.com/security#vulnerability-reporting

## New Issue

We welcome issues of all kinds, including feature requests, bug reports, or
documentation suggestions. Below are guidelines for well-formed issues of each
type.

### Bug Reports

 - [ ] **Test against latest release**: Make sure you test against the latest
   avaiable version of Terraform,
   terraform-plugin-go/terraform-plugin-sdk/terraform-plugin-framework,
   terraform-plugin-mux, and terraform-plugin-log. It is possible we already
   fixed the bug you're experiencing.

 - [ ] **Search for duplicates**: It's helpful to keep bug reports consolidated
   to one thread, so do a quick search on existing bug reports to check if
   anybody else has reported the same thing. You can scope searches by the
   label `bug` to help narrow things down.

 - [ ] **Include steps to reproduce**: Provide steps to reproduce the issue,
   along with code examples (both HCL and Go, where applicable) and/or real
   code, so we can try to reproduce it. Without this, it makes it much harder
   (sometimes impossible) to fix the issue.

### Feature Requests

 - [ ] **Search for possible duplicate requests**: It's helpful to keep
   requests consolidated to one thread, so do a quick search on existing
   requests to check if anybody else has requested the same thing. You can
   scope searches by the label `enhancement` to help narrow things down.

 - [ ] **Include a use case description**: In addition to describing the
   behavior of the feature you'd like to see added, it's helpful to also lay
   out the reason why the feature would be important and how it would benefit
   the wider Terraform ecosystem. A use case in the context of 1 provider is
   good, a use case in the context of many providers is better.

### Documentation Suggestions

 - [ ] **Search for possible duplicate suggestions**: It's helpful to keep
   suggestions consolidated to one thread, so do a quick search on existing
   issues to check if anybody else has suggested the same thing. You can scope
   searches by the label `documentation` to help narrow things down.

 - [ ] **Describe the questions you're hoping the documentation will answer**:
   It's very helpful when writing documentation to have specific questions like
   "how do I default a subsystem logger to off?" in mind. This helps us ensure
   the documentation is targeted, specific, and framed in a useful way.

## New Pull Request

Thank you for contributing!

We are happy to review pull requests without associated issues, but we highly
recommend starting by describing and discussing your problem or feature and
attaching use cases to an issue first before raising a pull request.

- [ ] **Early validation of idea and implementation plan**:
  terraform-plugin-log is complicated enough that there are often
  several ways to implement something, each of which has different implications
  and tradeoffs. Working through a plan of attack with the team before you dive
  into implementation will help ensure that you're working in the right
  direction.

- [ ] **Unit Tests**: It may go without saying, but every new patch should be
  covered by tests wherever possible.

- [ ] **Go Modules**: We use [Go
  Modules](https://github.com/golang/go/wiki/Modules) to manage and version all
  our dependencies. Please make sure that you reflect dependency changes in
  your pull requests appropriately (e.g. `go get`, `go mod tidy` or other
  commands). Where possible it is better to raise a separate pull request with
  just dependency changes as it's easier to review such PR(s).

### Cosmetic changes, code formatting, and typos

In general we do not accept PRs containing only the following changes:

 - Correcting spelling or typos
 - Code formatting, including whitespace
 - Other cosmetic changes that do not affect functionality
 
While we appreciate the effort that goes into preparing PRs, there is always a
tradeoff between benefit and cost. The costs involved in accepting such
contributions include the time taken for thorough review, the noise created in
the git history, and the increased number of GitHub notifications that
maintainers must attend to.

#### Exceptions

We belive that one should "leave the campsite cleaner than you found it", so
you are welcome to clean up cosmetic issues in the neighbourhood when
submitting a patch that makes functional changes or fixes.

## Linting

GitHub Actions workflow bug and style checking is performed via [`actionlint`](https://github.com/rhysd/actionlint).

To run the GitHub Actions linters locally, install the `actionlint` tool, and run:

```shell
actionlint
```

Go code bug and style checking is performed via [`golangci-lint`](https://golangci-lint.run/).

To run the Go linters locally, install the `golangci-lint` tool, and run:

```shell
golangci-lint run ./...
```

## Testing

Code contributions should be supported by unit tests wherever possible.

### GitHub Actions Tests

GitHub Actions workflow testing is performed via [`act`](https://github.com/nektos/act).

To run the GitHub Actions testing locally (setting appropriate event):

```shell
act --artifact-server-path /tmp --env ACTIONS_RUNTIME_TOKEN=test -P ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest pull_request
```

The command options can be added to a `~/.actrc` file:

```text
--artifact-server-path /tmp
--env ACTIONS_RUNTIME_TOKEN=test
-P ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest
```

So they do not need to be specified every invocation:

```shell
act pull_request
```

### Go Unit Tests

Go code unit testing is perfomed via Go's built-in testing functionality.

To run the Go unit testing locally:

```shell
go test ./...
```

This codebase follows Go conventions for unit testing. Some guidelines include:

- **File Naming**: Test files should be named `*_test.go` and usually reside in the same package as the code being tested.
- **Test Naming**: Test functions must include the `Test` prefix and should be named after the function or method under test. An `Example()` function test should be named `TestExample`. A `Data` type `Example()` method test should be named `TestDataExample`.
- **Concurrency**: Where possible, unit tests should be able to run concurrently and include a call to [`t.Parallel()`](https://pkg.go.dev/testing#T.Parallel). Usage of mutable shared data, such as environment variables or global variables that are used with reads and writes, is strongly discouraged.
- **Table Driven**: Where possible, unit tests should be written using the [table driven testing](https://github.com/golang/go/wiki/TableDrivenTests) style.
- **go-cmp**: Where possible, comparison testing should be done via [`go-cmp`](https://pkg.go.dev/github.com/google/go-cmp). In particular, the [`cmp.Diff()`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Diff) and [`cmp.Equal()`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal) functions are helpful.

A common template for implementing unit tests is:

```go
func TestExample(t *testing.T) {
    t.Parallel()
    testCases := map[string]struct{
        // fields to store inputs and expectations
    }{
        "test-description": {
            // fields from above
        },
    }
    for name, testCase := range testCases {
        // Do not omit this next line
        name, testCase := name, testCase
        t.Run(name, func(t *testing.T) {
            t.Parallel()
            // Implement test referencing testCase fields
        })
    }
}
```

## Maintainers Guide

This section is dedicated to the maintainers of this project.

### Releases

Before running a release, the changelog must be constructed from unreleased entries in the `.changelog` directory.

Install the latest version of the [`changelog-build`](https://pkg.go.dev/github.com/hashicorp/go-changelog/cmd/changelog-build) command, if it not already available:

```shell
go install github.com/hashicorp/go-changelog/cmd/changelog-build
```

Run the [`changelog-build`](https://pkg.go.dev/github.com/hashicorp/go-changelog/cmd/changelog-build) command from the root directory of the repository:

```shell
changelog-build -changelog-template .changelog.tmpl -entries-dir .changelog -last-release $(git describe --tags --abbrev=0) -note-template .changelog-note.tmpl -this-release HEAD
```

This will generate a section of Markdown text for the next release. Open the `CHANGELOG.md` file, add a `# X.Y.Z` header as the first line, then add the output from the `changelog-build` command.

Commit, push, create a release Git tag, and push the tag:

```shell
git add CHANGELOG.md
git commit -m "Update CHANGELOG for v1.2.3"
git push
git tag v1.2.3
git push --tags
```

GitHub Actions will pick up the new release tag and kick off the release workflow.
