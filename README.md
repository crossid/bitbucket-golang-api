# Preface

Golang implementation for the BitBucket API.

Please see [v1 API](#v1-api) and [v2 API](#v2-api) sections for API coverage, PR are very welcome.

# Add dependency

Use [Dep](https://github.com/golang/dep) to add this library to your project:

`dep ensure -add github.com/crossid/bitbucket-golang-api`

# Insantiate a v1 or v2 api

```go
v2 := bitbucket.NewV2BasicAuth(user, password)
v1 := bitbucket.NewV1BasicAuth(user, password)
```

Work with the API:

```go
v2.Teams.List(bitbucket.ListTeamsOpts{Role: bitbucket.AdminRole})
```

Please check unit tests as they cover all supported features.

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


# v1 API

## Groups

- [x] List groups matching one or more filters.
- [x] List of an account's (team / user) groups

## Privileges

- [x] List privileges of an account (team / user)
- [x] List privileges of an account (team / user) for a specific repo

## Group Privileges

- [x] List group privileges of an account (team / user)
- [x] List group privileges of an account (team / user) for a specific repo

# Running tests

In order to run tests you should simply:

1. Clone the project
1. Set two env vars: `BITBUCKET_USER` & `BITBUCKET_PASSWORD` with your Bitbucket username and password respectively
1. dep ensure
1. `export BITBUCKET_USER=<user> ; export BITBUCKET_PASSWORD="<password>"; go test`

Note: Unit tests assume that your user have at least:

- 2 teams
- 1 member per team
- 2 repositories

# Other projects

- https://github.com/emicklei/go-bitbucket - not updated for long time, contains very minimal APIs
- https://github.com/ktrysmt/go-bitbucket - active but we found it not clean enough, insufficient unit tests, no support for v1, no OAUTH support, minimal pagination support.
s
# License

APACHE 2