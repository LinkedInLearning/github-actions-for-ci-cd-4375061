# 04_07_challenge_develop_a_deployment_pipeline

It's time for a challenge!

Your team is excited to use GitHub Actions for their next project.

The goal is to develop a model pipeline that provides continuous integration, continuous delivery, and continuous deployment to three different environments.

However, things aren’t lining up.

To start, the workflow jobs are being run in the wrong order.  And while some of the environments will use continuous deployment, others will need to be reviewed before they can be deployed.

They’ve asked you to help them straighten out the pipeline, set up more control for deployment to particular environments, and provide a summary once the pipeline completes.

# Requirements
1. Create a new repo and add the exercise files. Run and observe the provided workflow.
1. Edit the workflow to place the jobs in order.
1. Configure the following environments:

    - Development
    - Staging
    - Production

1. Use continuous deployment for the Development and Staging environments.
1. Protect deployments to the Production environment with a review.
1. Update the summary to indicate all jobs have completed successfully.

