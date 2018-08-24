# GitHub PR Commits

When given a GitHub repo, PR, and valid token it will return, as JSON, commits
on the pull request. This is best used other projects as part of a chain of events.

## Usage

Prior to running the following environment variables are required to be set:

* `GITHUB_REPO`: The repo (e.g., `mattfarina/github-pr-commits`)
* `GITHUB_PR_NUMBER`: The pull request number to add it to (e.g., `1`)
* `GITHUB_TOKEN`: The GitHub API token to use for auth