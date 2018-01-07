# Preface

Golang implementation for the BitBucket API.

Please see _v2 API_ section for API coverage, PR are very welcome.

# Other projects

- https://github.com/emicklei/go-bitbucket - not updated for long time, contains very minimal APIs
- https://github.com/ktrysmt/go-bitbucket - active but we found it not clean enough, insufficient unit tests, no OAUTH support, minimal pagination support.

# v2 API

# Teams

- [x] Get a team
- [x] List Teams (with support for pagination)
- [x] List Team's members

# Repositories

- [x] List Public Repos (with support for pagination, filtering and sorting)
- [x] List Repos by Owner (with support for pagination, filtering and sorting)

# Users

- [x] Get current user
- [x] Get public user

# Running tests

In order to run tests you should simply:

1. Clone the project
1. Set two env vars: `BITBUCKET_USER` & `BITBUCKET_PASSWORD` with your Bitbucket username and password respectively
1. go test

Note: Unit tests assume that your user have at least:

- 2 teams
- 1 member per team
- 2 repositories