# 02_06 Challenge: Develop a CI Workflow

## INTRODUCTION
It’s time for a challenge!

You’re working with a team of data scientists that are just starting out with GitHub Actions.

The team wants to add a continuous integration workflow to their GitHub repo so that all pushes to the main branch are linted using Flake8 and all tests are run using Pytest.

They’ve tried using a starter workflow but are seeing warnings about deprecated features.

In addition, the first test in the repo, a check for a library version, is failing.

## REQUIREMENTS
Help the team set up a continuous integration pipeline using a GitHub Actions starter workflow.

1. Start by creating a new repo and adding the exercise files for this challenge.

    - [requirements.txt](./requirements.txt)
    - [test_pandas_version.py](./test_pandas_version.py)

1. Use the GitHub Actions web interface to create a starter workflow.
1. Run the workflow and observe the problems the team is referring to.
1. Review the starter workflow and update any out of date actions to use the latest version.
1. Fix any errors in the code so that the tests pass successfully.

This challenge should take about fifteen minutes to complete.