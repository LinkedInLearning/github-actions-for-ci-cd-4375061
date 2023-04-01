# 02_05 CI for Go

The following workflow updates the workflow suggested by GitHub Actions with the following:

- Adds an `env` block and a value for `GIN_MODE` set to `release`.
- Adds a `permissions` block for the `build` job with a request for `write` permissions for `checks`.
- Updates the `setup-go` action to use the input for `go-version-file: 'go.mod'`
- Adds the action `morphy2k/revive-action` for linting.
  - [Revive Action - GitHub Marketplace](https://github.com/marketplace/actions/revive-action)
  - [Project website](https://github.com/mgechev/revive)
- Adds the action `dominikh/staticcheck-action` for static analysis.
  - [Staticcheck - GitHub Marketplace](https://github.com/marketplace/actions/staticcheck)
  - [Project website](https://staticcheck.io/)


## The Updated Workflow
The updated workflow can be found in the code snippet below and in the repo here: [./go-ci-workflow.yml](./go-ci-workflow.yml).

    name: Go

    env:
      GIN_MODE: release

    on:
      push:
        branches: [ "main" ]
      pull_request:
        branches: [ "main" ]

    jobs:
      build:
        runs-on: ubuntu-latest

        permissions:
          checks: write

        steps:
        - uses: actions/checkout@v3

        - name: Set up Go
          uses: actions/setup-go@v3
          with:
            go-version-file: 'go.mod'
            cache: true

        - name: Revive Action
          uses: morphy2k/revive-action@v2.4.1

        - name: Staticcheck
          uses: dominikh/staticcheck-action@v1.3.0

        - name: Build
          run: |
            go mod download
            go install gotest.tools/gotestsum@latest

        - name: Test
          run: |
            gotestsum --format=standard-verbose --junitfile=junit.xml

        - name: Publish Test Report
          uses: mikepenz/action-junit-report@v3
          if: success() || failure() # always run even if the previous step fails
          with:
            report_paths: '**/junit.xml'
            detailed_summary: true
            include_passed: true


 For more information see the GitHub documentation for [Building and testing Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
