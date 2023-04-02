# 02_07 Solution: Develop a CI Workflow

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
    - [python-app.yml](./python-app.yml)

1. Use the GitHub Actions web interface to create a starter workflow.

    From the repo homepage, Select `Actions` -> `Python application` -> `Configure`

    If you are using the solution without creating a new workflow using the step above, move the workflow file [`python-app.yml`](./python-app.yml) to a directory in the repo named `.github/workflows`.

1. Run the workflow and observe the problems the team is referring to.  Note that the `set-output` deprecation warning was present at the time this course was recorded and may be corrected by the time you view this course.

    ![Errors](./images/02-07-Solution-Develop-CI-Workflow-1.png)

1. Review the starter workflow and update any out of date actions to use the latest version.  Note that the course and exercise files refer to the version of `actions/setup-python` in use at the time this course was recorded.  These values may be different by the time you view this course.

    Make the following changes to the workflow file, `python-app.yml`:

    Change:

        - name: Set up Python 3.10
          uses: actions/setup-python@v3

    to:

        - name: Set up Python 3.10
          uses: actions/setup-python@v4.5.0

1. Fix any errors in the code so that the tests pass successfully.

    Make the following changes to the dependency file, `requirements.txt`:

    Change:

        pandas==1.4.0

    to:

        pandas==1.5.3

    Alternatively, you can change the file `test_pandas_version.py` to match the requirements file:

    Change:

        PANDAS_VERSION = "1.5.3"

    to:

        PANDAS_VERSION = "1.4.0"
