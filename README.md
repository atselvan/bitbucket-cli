# Bitbucket CLI

_This CLI can be used for managing an Atlassian Bitbucket Server_

## Usage

```console
Usage of bitbucket-cli:
  -bitbucketURL string
        Bitbucket server URL
  -getInactiveUsers
        Get a list of users who have not logged in since 3 months
  -getProjectsList
        Get a list of projects available in bitbucket
  -password string
        Password for authentication
  -projectExists
        Check if a project exists in bitbucket. Required Parameters: projectKey
  -projectKey string
        The Key of a bitbucket project
  -username string
        Username for authentication
  -verbose
        Set this flag for Debug logs
```

## TODO

* Add authorization validation on all functions
* List projects and inactive users into a file
* Add function to create a project with proper role mapping