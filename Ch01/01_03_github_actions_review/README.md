# 01_03_github_actions_review
GitHub actions is an automation framework that integrates seamlessly with code stored in a GitHub repository.  You can use Actions to automate many parts of the Software Development Lifecycle including continuous integration and deployment.

## Terminology

|Keyword|Definition|
|-------|----------|
| **Workflows**|Workflows define how all the automation in a process is tied together including how the process gets started.  Files are written using YAML and must be stored in the `.github/workflows` directory at the root of the repo.|
| **Events**|Events are the triggers that start a workflow. Common events are `push`, `workflow_dispatch`, and `pull_request`.|
| **Runners**|Runners define the compute layer where jobs are executed.  GitHub Actions provides runners with popular operating systems like Ubuntu, Windows, and macOS.|
| **Jobs**|Each workflow contains one or more jobs.  Jobs are high-level tasks that define one or more steps that need to be run for the job to complete.
| **Steps**|The actual work done by a job takes place in Steps.  Steps are simple commands, shell scripts, or they can be Actions.|
|**Actions**|Actions help speed up your workflow development by bundling programs, runtimes, and all of their dependencies into Docker containers or JavaScript modules.|

## Recommended Reading
- [Learning Github Actions](https://www.linkedin.com/learning/learning-github-actions-2/automating-with-github-actions-2)

- [Quickstart for GitHub Actions](https://docs.github.com/en/actions/quickstart)

- [Using workflows](https://docs.github.com/en/actions/using-workflows)

- [GitHub Actions Runner Images](https://github.com/actions/runner-images/tree/main)