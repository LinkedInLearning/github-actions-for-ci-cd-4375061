# 04_05_ci_cd_for_container_images

## Recommended Reading
- [Creating a reusable workflow](https://docs.github.com/en/actions/using-workflows/reusing-workflows#creating-a-reusable-workflow)

- [Deployment protection rules](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#deployment-protection-rules)

## Using the Exercise Files
### 1. Create a repo; add service account credentials
1. Create a new repo.
1. Review the steps in [Create a Service Account](../04_04_create_a_service_account/README.md) to add service account credentials to the repo

### 2. Create a repo in the Elastic Container Registry (ECR); Push an image
1. In your AWS account, go the [ECR homepage](https://console.aws.amazon.com/ecr/home)
1. Select `Get started`.
1. Under settings, confirm `Private` is selected
1. Give the repo a name.

    **For our deployment workflow to operate correctly, the container repo needs to have the same name as the GitHub repo.**

1. Keep the defaults for everything else, scroll to the bottom, and select `Create repository`.
1. Back in the GitHub repo, use the `workflow dispatch trigger` to start a new run of the Container Deployment pipeline.
1. Wait for the workflow run to complete. It will likely fail.
1. On the ECR page, select the repo you created earlier.
1. Confirm that an image has been pushed to the repo.

### 3. Create two App Runner services
1. Browse to the [App Runner homepage](https://console.aws.amazon.com/apprunner/home)
1. Create an App Runner service for the `Staging` environment.

    1. Select `Create an App Runner service` or `Create service`.
    1. Under Source, confirm `Container registry` and `Amazon ECR` are selected.
    1. Under `Container image URI`, select `Browse`.
    1. Select the image repo you created and confirm the `latest` appears under `Image tag`.
    1. Select `Continue`.
    1. Under `Deployment settings`, select `Automatic`.

        This will configure App Runner to watch our container repo in ECR.  When a new image is pushed, the service will be updated with the new image right away.

    1. Under `ECR access role`, select `Create new service role`.  If you already have a service role in place, select `Use existing service role` and select the role.
    1. Select `Next`.
    1. Enter a name for the service. Add `staging` as a prefix or suffix so you can identify the service.  You may also consider using the repo name here as well.
    1. Under `Virtual CPU`, select `0.25 VCPU`.
    1. Leave everything else as the defaults and go to the next screen by selecting `Next`.
    1. Review your settings, scroll to the bottom of the screen, and select `Create & Deploy`.

1. Create an App Runner service for the `Production` environment.

    Use the same steps and settings as `Staging` with one exception:

    **For the `Production` service, make sure that the Manual trigger is selected.**

### 4. Create environments for the App Runner services

1. Select each App Runner service and save the values for:
    - `Default domain`
    - `Service ARN`

1. Back in the repo, create an environment for `Production`. (The `Staging` environment should already be in place, created by a workflow run.)
    1. Select `Settings` -> `Environments` -> `New environment`.
    1. For the name, `Production`.
    1. Select `Configure environment`.
    1. Select `Required reviewers`.
    1. In the search field, enter your GitHub user name and select it.
    1. Select `Save protection rules`.
    1. Under `Environment variables`, select `Add variable`.
    1. Create variables for `SERVICE_ARN` and `URL` using the values from the App Runner configuration for Production.
    1. Select `Save protection rules`.

1. Edit the `Staging` environment.
1. Under `Environment variables`, select `Add variable`.
1. Create variables for `SERVICE_ARN` and `URL` using the values from the App Runner configuration for Staging.

### 5. Run the workflow
1. Select the `Actions` tab.
1. Select `0-Container Deployment Pipeline`.
1. Select `Run workflow` -> `Run workflow`.
1. Observe the pipeline's progress.
1. When prompted, select `Review deployments`.
1. Select `Production` -> `Approve and deploy`.
1. Observe the pipeline's progress.
1. Follow the links on the tiles for each job to view the deployed application.

### 6. Remove the resources
**To avoid costs associated with running resources in AWS, please remove them.**

1. Remove App Runner services:
    1. Browse to the [App Runner homepage](https://console.aws.amazon.com/apprunner/home)
    1. For each service:
    1. Select the service.
    1. Select `Actions` -> `Delete`.
    1. Enter the word `delete` and select `Delete`.
1. Remove ECR repos.
    1. Browse to the [ECR homepage](https://console.aws.amazon.com/ecr/home)
    1. Select the checkbox next to the name of the repo.
    1. Select `Delete`.
    1. Enter the word `delete` and select `Delete`.

