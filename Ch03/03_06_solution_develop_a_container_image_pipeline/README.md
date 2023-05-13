# 03_06_solution_develop_a_container_image_pipeline
It's time for a challenge!

You’ve just joined a team that’s developing an API for mobile applications.  They’re using containers for the first time and need to set up a CI/CD pipeline to publish container images.

Their last project used an integration workflow for python and the plan is to use that workflow on the new project as well.

The team has asked you to help them set up a repo for the integration workflow so it can be used in the latest project and any new projects that come up in the future.

# SOLUTION
1. Create two new repos.
1. In the first repo, add a workflow for a Python application.  Create a workflow using the starter workflow for *Python applications*.

    From the repo homepage, Select:

    `Actions` -> Search for `Python application` -> `Configure`.

1. Update the workflow so it can be called from another workflow.

    Edit the Python application workflow and remove all triggers.

    Add a trigger for `workflow_call` so the `on` section appears as follows:

        on:
          workflow_call:

1. In the second repo, add the exercise files and then create a workflow using the starter workflow for *Publish Docker Container*.  Update the workflow to call the integration workflow.

    From the repo homepage, Select:

    `Actions` -> `Publish Docker Container` -> `Configure`.

    Create an entry under `jobs` that calls the integration workflow from the first repo.  The entry should have an ID and a `uses` block that:

    - refers to the first repo
    - includes the path `.github/workflows` and the workflow name
    - and a Git reference.

    The complete entry should be similar to:

        integration:
            # REPLACE GITHUB_USERNAME and GITHUB_REPONAME with your own values.
            # This is needed for the workflow_call event to complete successfully.
            uses: GITHUB_USERNAME/GITHUB_REPONAME/.github/workflows/python-app.yml@main

    Additionally, add permissions for the integration workflow.  After the `uses` block, add:

        permissions:
          contents: read

1. Add a job to build and publish the code as a container image once the integration tests are complete.

    This requirement is partially completed in the previous step by adding the starter workflow.

    However, the workflow still needs to be modified so the build and publish steps run _after_ the integration tests.

    After the `build` job, add a `needs` entry with the ID of the integration workflow:

        needs: [integration]

    _NOTE_: *OPTIONAL* If you would like the final docker image to run on a MacOS system using the M1/M2 chips, add `platforms: linux/amd64,linux/arm64
` as an option to the `Build and push Docker image` step.  The entry should look like:

        - name: Build and push Docker image
            id: build-and-push
            uses: docker/build-push-action@ac9327eae2b366085ac7f6a2d02df8aa8ead720a
            with:
            context: .
            push: ${{ github.event_name != 'pull_request' }}
            tags: ${{ steps.meta.outputs.tags }}
            labels: ${{ steps.meta.outputs.labels }}
            cache-from: type=gha
            cache-to: type=gha,mode=max
            platforms: linux/amd64,linux/arm64

1. Commit the workflow and open the *Actions* tab.
1. Observe the workflow.  The `integration/build` job should run first, followed by the `build` job.

    _NOTE: If warning annotations are present in the build log, see the [README file in "03_04_ci_cd_for_container_images"](../03_04_ci_cd_for_container_images/README.md#fixing-warnings-in-the-annotations-section) for details on how to fix the warnings._

1. When the build completes open the main page for the repo by selecting *Code*.
1. Confirm that a container image has been created.
1. Open the page for the container image.
1. Observe the details on the container image page.

# BONUS!  USE THE IMAGE
Congrats on building the image.  Why not use it too!?

Follow these steps:
1. Confirm that `docker` is running on your local system.
    - [Installing Docker Desktop](https://www.docker.com/products/docker-desktop/)
3. Download the image that was built, using the `main` branch as a reference:

        docker pull ghcr.io/GITHUB_USERNAME/GITHUB_REPONAME:main

1. In a terminal, start a container with the `docker run` command, and connecting the local port `3000` to the container port `3000`:

        docker run -p 3000:3000 ghcr.io/GITHUB_USERNAME/GITHUB_REPONAME:main

1. In a browser, open the URL [localhost:3000/](http://localhost:3000/).
1. Follow the instructions to see the API in action!
