# 03_04_ci_cd_for_container_images
This lesson demonstrates how to set up a CI/CD workflow for container images.

Specificaly, this demo implements a delivery workflow that:
- Reuses a workflow to run integration tests
- Collects metadata
- Builds the image
- Publishes the image to GitHub Packages

## Using the Exercise Files
1. Create a new repo and upload the files for this lesson.
1. Rename the file [python-ci-workflow.yml](./python-ci-workflow.yml) so that its located in the `.github/workflows` directory.

        Completing this step is key to having the workflow run properly!

1. In the repo select the *Actions* tab.
1. Select and configure the workflow *Publish Docker Container*.
1. Under `jobs`, add an ID named “integration” and the keyword `uses` followed by the path to `python-ci-workflow.yml`.

    _NOTE that the path to the workflow starts with `./` and contains the full path the workflow file._

    Add a permissions block followed by `contents: read` and `checks: write`.

    The completed call to the reused workflow should be similar to the following:


        jobs:
            integration:
                uses: ./.github/workflows/python-ci-workflow.yml
                permissions:
                contents: read
                checks: write

1. Under the job with the ID `build`, add `needs`, followed by the ID for the integration job.  This will cause the `build` job to wait for the `integration` job to complete.

        build:
            needs: [integration]

1. Select *Start Commit* and then select *Commit New File*.
1. Select the *Actions* tab.
1. Select the running workflow.
1. Observe the `integration/build` job followed by the `build` job.
1. Observe the updates to the Actions UI as the `integration/build` job completes.
1. Obverse any annotations written to the Actions UI as the `build` job completes.  If there are no annotations refering to `Node.js 12`, you can safely skip this step.

    At the time this course was published, the *Publish Docker Container* workflow contained actions based on `Node.js 12`.  This causes warnings to be written to the Actions UI.

    To remove this warnings, several actions need to be updated to use newer versions.  Update your workflow to use the following actions with the indicated versions:

        uses: docker/setup-buildx-action@v2.5.0
        uses: docker/login-action@v2.1.0
        uses: docker/metadata-action@v4.4.0
        uses: docker/build-push-action@v4.0.0

1. Once the workflow completes, select the *Code* tab.
1. Refresh the page as needed until the package is listed under *Packages*.
1. Select the package and review the details on the package page.