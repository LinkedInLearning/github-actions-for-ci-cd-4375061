# 04_04_create_a_service_account
Service accounts are used to manage credentials and permission for accessing a service.

They’re called service accounts to indicate that they aren’t associated with a specific user.  Instead it's an account that a service will use to authenticate with the deployment platform.

Permissions given to a service account should be limited to the specific tasks a service needs to complete.

This prevents service accounts from using elevated permissions or potentially accessing services that aren’t required for the task at hand.

In summary, service accounts provide a secure way for our deployment workflows to interact with services outside of Github.

## Recommended Reading
- [Creating encrypted secrets for a repository](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-a-repository)

- [Creating an AWS account](https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-creating.html)

- [Creating an IAM user in your AWS account](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console)

## Using the Exercise Files
### Create a service account
1. Create an AWS as needed.
1. Login and open the [IAM dashboard](https://us-east-1.console.aws.amazon.com/iam/home)
1. Under `Access Management`, select `Users`.
1. Select `Add users`.
1. Enter a name for he service account.  Perhaps `github-actions-service-account`.
1. Select `Next`.
1. Select `Attach policies directly`.
1. Search for and select the following policies:
    - **AmazonEC2FullAccess**
    - **AmazonEC2ContainerRegistryFullAccess**
    - **AWSAppRunnerFullAccess**
    - **AmazonS3FullAccess**

    Once all policies are in place, select `Next`.
1. Select `Create user`.
1. After the user is created, select the user.
1. Select `Scurity credentials1`.
1. Under access keys, select `Create access key`.
1. On the best practices screen, select `Application running outside of AWS` and select `next`.
1. Select `Create access key`.

    This displays two values: an Access key and a Secret Access key.  Select `Show` to reveal the secret access key.  Select `Download csv file` to save these values to a file.

    **This is the only time you’ll be able to retrieve these values**

    **Don’t reveal these values or store them inside code in your GitHub repo**

    Download the credentials, copy them or leave the screen open for the next steps.

    When you’re all finished, select `Done`.

1. Follow these steps to find your Account Number:

    1. In the navigation bar on the upper right, choose your account name or number and then choose `Security credentials`.
    1. Under the Account details section, the account number appears next to the label `AWS account ID`.

### Save the service account credentials as repository secrets
1. Create a new repo.
1. Select `Settings` -> `Secrets and variables` -> `Actions`.
1. Repeat the following steps for `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_ACCOUNT_NUMBER`:



    1. Select `New repository secret`.
    1. Add the name for the current secret.
    1. Paste in the value.
    1. Select `Add secret`.

1. Select the `Variables` tab.
1. Select `New repository variable`.
1. Add a variable called `AWS_REGION`.

    This value should reflect the AWS Region where you plan to deploy resources.

    For more information, see [Regions, Availability Zones, and Local Zones](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).

1. Paste in the value.
1. Select `Add variable`.
