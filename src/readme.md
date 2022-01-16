# Microservice to create a repository

## Before run

Please follow below topics to make API work correctly.

### PAT(Personal access tokens)

API internally uses PAT to interract with GitHub. To get a PAT you need to authenticate to your GitHub account and go to `Settings -> Developer settings -> Personal access tokens` and generate a new PAT.

> Copy PAT and store it in a secure place because yo won't be able to see it again!

### Environment variables

As previously mentione API needs PAT to interract with Github. To do so, add a `SECRET_GITHUB_ACCESS_TOKEN` variable to your environment.

``` 
    export SECRET_GITHUB_ACCESS_TOKEN=your_PAT_goes_here
```


