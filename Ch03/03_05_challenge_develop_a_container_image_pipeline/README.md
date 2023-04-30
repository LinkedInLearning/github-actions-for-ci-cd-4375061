# 03_05_challenge_develop_a_container_image_pipeline
It's time for a challenge!

You’ve just joined a team that’s developing an API for mobile applications.  They’re using containers for the first time and need to set up a CI/CD pipeline to publish container images.

Their last project used an integration workflow for python and the plan is to use that workflow on the new project as well.

The team has asked you to help them set up a repo for the integration workflow so it can be used in the latest project and any new projects that come up in the future.

# REQUIREMENTS
1. Create two new repos.
1. In the first repo, add a workflow for a Python application.

    Use the supplied workflow,[python-app.yml](./python-app.yml), or create a workflow using the starter workflow for *Python applications*.

    If you use the supplied workflow, move it into the `.github/workflows` directory.

1. Update the workflow so it can be called from another workflow.
1. In the second repo, add a workflow that calls the integration workflow.

    Use the supplied workflow, [docker-publish.yml](./docker-publish.yml), or create a workflow using the starter workflow for *Publish Docker Container*.

    If you use the supplied workflow, move it into the `.github/workflows` directory.

    Additionally, add permissions for the integration workflow.  After the `uses` block, add:

        permissions:
          contents: read
          packages: write

1. Add a job to build and publish the code as a container image once the integration tests are complete.

    This requirement is partially completed in the previous step by adding the supplied workflow or using a starter workflow.

    However, the workflow still needs to be modified so the build and publish steps run _after_ the integration tests.

