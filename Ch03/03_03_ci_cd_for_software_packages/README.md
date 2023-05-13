# 03_03_ci_cd_for_software_packages
Continuous delivery workflows for software packages follow a pattern with these steps:
- Configure the project to work with the package registry
- Authenticate with the target registry
- Build the package
- Publish the package to the registry

## Registry configuration and authentication
Each language has a specific configuration that identifies the target registry and how to authenticate with it.

|Language |Config File          |
|----------|---------------------|
|JavaScript|package.json         |
|Ruby      |.gemspec             |
|Java      |settings.xml, pom.xml|
|.Net      |.csproj              |

## Build and publish
Each language will also use its own, native tooling to build and publish a package.

|Language |Build, publish Commands       |
|----------|------------------------------|
|JavaScript|npm ci; npm publish           |
|Ruby      |gem build; gem push           |
|Java      |mvn package; maven deploy     |
|.Net      |dotnet pack; dotnet nuget push|

## Package Versions
The configuration files for a package also define a version number for the package being published.

Version numbers can't be reused.

Update code to reference a new version number with each new release.

## Recommended Reading
- [Working with a GitHub Packages Registry - GitHub Docs](https://docs.github.com/en/packages/working-with-a-github-packages-registry)

## Using the Exercise Files
1. Create a new repo and upload the files for this lesson.  Note that the Java files need to be located in subdirectories.  Follow the steps in this document to move the files after uploading them: [Moving a file to a new location
](https://docs.github.com/en/repositories/working-with-files/managing-files/moving-a-file-to-a-new-location).

    Specifically, the file [HelloActions.java](./src/main/java/com/example/HelloActions.java) needs to have the following path:

        ./src/main/java/com/example/HelloActions.java

    And the file [HelloActionsTest.java](./src/test/java/com/example/HelloActionsTest.java) needs to have the following path:

        ./src/test/java/com/example/HelloActionsTest.java

1. Edit [pom.xml](./pom.xml).

        Completing this step is key to having the workflow run properly!

    Replace all occurrences of `GITHUB_USERNAME` and `GITHUB_REPONAME` with your GitHub username and the name for the GitHub repo you are using for this exercise.

    Save the [pom.xml](./pom.xml) file and commit the changes to the repo.

1. In the repo select the *Actions* tab.
1. Select and configure the workflow *Publish Java Package with Maven*.
1. Select *Start Commit* and then select *Commit New File*.
1. Select the *Code* tab.
1. Select *Create a new release*.
1. Select *Choose a tag*.
1. For the tag name, enter `v1.0.0`.
1. Select *Create a new tag on publish*.
1. For the release title, enter `v1.0.0`.  Enter text for the body of the release as well.
1. Select *Publish*.
1. Select the *Actions* tab and wait for the workflow to finish successfully.  If there are errors in the workflow run, review them and make corrections where needed.  If you get stuck, ask for help in the Q&A for the course.
1. Once the workflow completes, select the *Code* tab.
1. Refresh the page as needed until the package is listed under *Packages*.
1. Select the package and review the details on the package page.
