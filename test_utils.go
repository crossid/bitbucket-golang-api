package bb

import "os"

func getUser() string {
	user := os.Getenv("BITBUCKET_USER")
	return user
}

func getPassword() string {
	pass := os.Getenv("BITBUCKET_PASSWORD")
	return pass
}

func newTestV2Impl() *V2Impl {
	return NewV2BasicAuth(getUser(), getPassword())
}

func newTestV1Impl() *V1Impl {
	return NewV1BasicAuth(getUser(), getPassword())
}

func getTeamName(v2Impl *V2Impl) (string, error) {
	teams, err := v2Impl.Teams.List(ListTeamsOpts{Role: AdminRole, Pagelen: 1})
	if err != nil {
		return "", err
	}

	return teams.Values[0]["username"].(string), nil
}

func getRepoAndTeam(v2Impl *V2Impl) (team string, repo string, err error) {
	team, err = getTeamName(v2Impl)
	if err != nil {
		return
	}

	repos, err := v2Impl.Repositories.ListByOwner(team, ListReposByOwnerOpts{})
	if err != nil {
		return "", "", err
	}

	return team, repos.Values[0]["slug"].(string), nil
}
