# 02_04 CI for Python

# Permissions for Checks
Permissions are needed to update the Actions interface.  Permissions can be added under the `job` section for a specific job, or at the top of a workflow so all jobs assume the permission.

    jobs:
        build:
            runs-on: ubuntu-latest
            permissions:
                checks: write

# JUnit Reporting
Tests can be updated to include JUnit reports.

Actions can be added to publish JUnit reports to the Actions user interface.

For example, a standard call to `pytest` can be modified from:

    - name: Test with pytest
      run: |
        pytest

to:


    - name: Test with pytest
      run: |
        python -m pytest --verbose --junit-xml=junit.xml
    - name: Publish Test Report
      uses: mikepenz/action-junit-report@v3
      if: success() || failure() # always run even if the previous step fails
      with:
        report_paths: '**/junit.xml'
        detailed_summary: true
        include_passed: true

A complete workflow is located here: [./python-ci-workflow.yml](./python-ci-workflow.yml)
