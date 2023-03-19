# # 02_05 CI for Go

## Steps for Golang CI
- Create repo
- add files from exercise files
- click actions tab
- no choices under continuous integration that pop out
- select go under suggested actions
- add `workflow_dispatch` to the triggers
- add the following under "runs_on"...
```
    permissions:
      checks: write
```

- change "Set up Go" to:
```
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version-file: 'go.mod'
        cache: true
```

- change "Build" step to:
```
    - name: Build
      run: |
        go install gotest.tools/gotestsum@latest
        go mod download
```

- change "Test..." call to:
```
gotestsum --format=standard-verbose --junitfile=junit.xml
```

-- add the following job at the end:
```
    - name: Publish Test Report
      uses: mikepenz/action-junit-report@v3
      if: success() || failure() # always run even if the previous step fails
      with:
        report_paths: '**/junit.xml'
        detailed_summary: true
        include_passed: true
```
- commit new file directly to main branch
- go to actions tab
