# 03_04_ci_cd_for_container_images
This lesson demonstrates how to set up a CI/CD workflow for container images.

Specifically, this demo implements a delivery workflow that:
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
1. Once the workflow completes, select the *Code* tab.
1. Refresh the page as needed until the package is listed under *Packages*.
1. Select the package and review the details on the package page.

# Fixing Warnings in the Annotations Section
At the time this course was published, the *Publish Docker Container* workflow contained actions based on `Node 12`.

[Node 12 has been out of support since April 2022 and has been deprecated by Github Actions](https://github.blog/changelog/2022-09-22-github-actions-all-actions-will-begin-running-on-node16-instead-of-node12/).

This causes warnings to be written to the Actions UI.

To remove these warnings, the version for actions used in this workflow need to be updated.

Update your workflow to use the following actions with the indicated versions.  Also, check the provided Marketplace page to see if newer versions are available.

|Action Marketplace Page|Action Name and Version|
|--|--|
|[Docker Setup Buildx](https://github.com/marketplace/actions/docker-setup-buildx)|`docker/setup-buildx-action@v2.5.0`|
|[Docker Login](https://github.com/marketplace/actions/docker-login)|`docker/login-action@v2.1.0`|
|[Docker Metadata action](https://github.com/marketplace/actions/docker-metadata-action)|`docker/metadata-action@v4.4.0`|
|[Build and push Docker images](https://github.com/marketplace/actions/build-and-push-docker-images)|`docker/build-push-action@v4.0.0`|
