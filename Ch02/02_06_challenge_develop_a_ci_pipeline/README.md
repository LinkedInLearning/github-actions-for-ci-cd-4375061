# 02_06 Challenge: Develop a CI Workflow

## INTRODUCTION
It’s time for a challenge!

You’re working with a team of data scientists that are just starting out with GitHub Actions.

The team wants to add a continuous integration workflow to their GitHub repo so that all pushes to the main branch are linted using Flake8 and all tests are run using Pytest.

All code in the repo needs to use a specific version of Pandas, a popular Python library. They have code to test for the version but for some reason the test is failing.

They’d also like to find some way to make it easier to summarize the tests being run in the repo.

## REQUIREMENTS
Help the team set up a continuous integration pipeline using a GitHub Actions starter workflow.

1. Start by creating a new repo and adding the exercise files for this challenge.

    - [requirements.txt](./requirements.txt)
    - [test_pandas_version.py](./test_pandas_version.py)

1. Use the GitHub Actions web interface to create a starter workflow.
1. Run the workflow and observe the problems the team is referring to.
1. Fix any errors in the code so that the tests pass successfully.
1. Update the workflow to add a summary of the tests being run. 
    1. Update the workflow so that it has permissions to create checks in the Actions interface.  
    1. Update the call to `pytest` so that it creates a JUnit report named `junit.xml`.
        
            python -m pytest --verbose --junit-xml=junit.xml
            
    1. Add a new step that uses the [JUnit Report Action](https://github.com/marketplace/actions/junit-report-action) from the GitHub Marketplace:

            - name: Publish Test Report
            uses: mikepenz/action-junit-report@v3
            if: success() || failure() # always run even if the previous step fails
            with:
                report_paths: '**/junit.xml'
                detailed_summary: true
                include_passed: true

This challenge should take about fifteen minutes to complete.
