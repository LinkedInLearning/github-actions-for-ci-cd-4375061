# 04_03_ci_cd_for_github_pages
> **Warning**: GitHub Pages sites are publicly available on the internet, even if the repository for the site is private.
>
> If you have sensitive data in your site's repository, you may want to remove the data before publishing a site using GitHub Pages.
>
> Specifically, do not publish repository secrets that may include passwords or credentials.

[GitHub Pages](https://pages.github.com/) is a free service provided by GitHub that allows users to host static websites directly from their GitHub repositories.

Static sites can be stored in a repo as HTML, JavaScript, and CSS.  Or they can be stored as Markdown.

For Markdown sites, a static site generator is used to convert Markdown files into HTML.

Popular static site generators include:
- [Jekyll](https://jekyllrb.com/)
- [Hugo](https://gohugo.io/)
- [Gatsby](https://www.gatsbyjs.com/)

Project sites are available at URLs similar to the following:
- `http(s)://<username>.github.io/<repository>`
- `http(s)://<organization>.github.io/<repository>`.

GitHub Pages can also be configured to use a custom domain.

## Jekyll Front Matter
To help Jekyll with compiling and publishing Markdown files, files contain text near the top of the file called [front matter](https://jekyllrb.com/docs/front-matter/).

> Any file that contains a YAML front matter block will be processed by Jekyll as a special file. The front matter must be the first thing in the file and must take the form of valid YAML set between triple-dashed lines. Here is a basic example:

    ---
    layout: page
    title: Welcome to my site
    permalink: /
    ---

## Recommended Reading
- [GitHub Pages Documentation](https://docs.github.com/en/pages)

- [GitHub Pages Usage limits](https://docs.github.com/en/pages/getting-started-with-github-pages/about-github-pages#usage-limits)

## Using the Exercise Files
1. Create a new repo and add the exercise files for this lesson.

    Make sure all `*.png` files are moved into the `images` directory.

    _*Note that images will not display properly unless they are located in the `images` directory.*_

1. From the repo homepage, select **Settings**.
1. On the Settings page, under **Code and automation**, select **Pages**.
1. On the *GitHub Pages* page, under **Build and Deployment**, select the drop-down under **Source**.
1. Select **GitHub Actions**.
1. Note that repos that only contain HTML, JavaScript, and CSS can be deployed using the *Static HTML* workflow.
1. Under **GitHub Pages Jekyll**, select **Configure**.
1. Review the starter workflow, *Deploy Jekyll with GitHub Pages dependencies preinstalled*. Things to note include:
    - The workflow is triggered by pushes to the main branch or workflow dispatch
    - Permissions are requested for reading the repo, writing to the github pages service, and writing to
    the ID token.
    - The workflow establishes concurrency at the workflow level for a group named "pages" and will wait for in builds that are in progress to complete.
    - the **build** job checkouts out the repo and runs the `jekyll-build-pages` action to compile the site into a directory containing the HTML, JavaScript, and CSS.
    - The `upload-pages-artifact` action is used to create an artifact for the deploy step.
    - The **deploy** job waits for the build job to complete and then deploys the artifact using the `deploy-pages` action.
    - Note that the build job also uses an environment named github-pages and sets a URL for display on in the Actions UI.
1. Select **Commit changes...** and then **Commit changes**.
1. Go to the **Actions** tab.
1. Note the actively running workflow and select it.
1. Wait for the workflow to complete.  Optionally, click into the build and deploy jobs to observe them while they are running.
1. Note and select the URL displayed on the build job.
1. View the deployed site.
