# 03_06_solution_develop_a_container_image_pipeline
It's time for a challenge!

You’ve just joined a team that’s developing an API for mobile applications.  They’re using containers for the first time and need to set up a CI/CD pipeline to publish container images.

Their last project used an integration workflow for python and the plan is to use that workflow on the new project as well.

The team has asked you to help them set up a repo for the integration workflow so it can be used in the latest project and any new projects that come up in the future.

# SOLUTION
1. Create two new repos.
1. In the first repo, add a workflow for a Python application.

    Use the supplied workflow,[python-app.yml](./python-app.yml), or create a workflow using the starter workflow for *Python applications*.

    If you use the supplied workflow, move it into the `.github/workflows` directory.

    If you are using a starter workflow, from the repo homepage, Select `Actions` -> `Python application` -> `Configure`.

1. Update the workflow so it can be called from another workflow.

    Edit the Python application workflow and remove all triggers.

    Add a trigger for `workflow_call` so the `on` section appears as follows:

        on:
            workflow_call:

1. In the second repo, add the exercise files and then add a workflow that calls the integration workflow.

    Use a starter workflow. From the repo homepage, Select `Actions` -> `Publish Docker Container`.

    Create an entry under `jobs` that calls the integration workflow from the first repo.  The entry should have an ID and a `uses` block that refers to the first repo, the workflow file located in the `.github/workflows`, and a Git reference.  The entry should be simliar to:

        integration:
            # REPLACE GITHUB_USERNAME and GITHUB_REPONAME with your own values.
            # This is needed for the workflow_call event to complete successfully.
            uses: GITHUB_USERNAME/GITHUB_REPONAME/.github/workflows/python-app.yml@main

    To use the supplied workflow, [docker-publish.yml](./docker-publish.yml), move it into the `.github/workflows` directory.

    In either case, edit the file so that the values for `GITHUB_USERNAME` and `GITHUB_REPONAME` are replaced to use your own username and repo name.

    Additionally, add permissions for the integration workflow.  After the `uses` block, add:

        permissions:
          contents: read
          packages: write

1. Add a job to build and publish the code as a container image once the integration tests are complete.

    This requirement is partially completed in the previous step by adding the supplied workflow or using a starter workflow.

    However, the workflow still needs to be modified so the build and publish steps run _after_ the integration tests.

    After the `build` job, add a `needs` entry with the ID of the integration workflow:

        needs: [integration]

1. Commit the workflow and open the *Actions* tab.
1. Obeserve the workflow.  The `integration/build` job should run first, followed by the `build` job.

    _NOTE: If warning annotations are present in the build log, see the [README file in "03_04_ci_cd_for_container_images"](./README.md03_04_ci_cd_for_container_images) for details on how to fix the warnings._

1. When the build completes open the main page for the repo by selecting *Code*.
1. Confirm that a container image has been created.
1. Open the page for the container image.
1. Observe the details on the container image page.

