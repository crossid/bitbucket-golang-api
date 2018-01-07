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

func newTestV2Impl() *v2Impl {
	return newV2BasicAuth(getUser(), getPassword())
}

func newTestV1Impl() *v1Impl {
	return newV1BasicAuth(getUser(), getPassword())
}

func getTeamName(v2Impl *v2Impl) (string, error) {
	teams, err := v2Impl.teams.List(ListTeamsOpts{Role: AdminRole, Pagelen: 1})
	if err != nil {
		return "", err
	}

	return teams.Values[0]["username"].(string), nil
}

func getRepoAndTeam(v2Impl *v2Impl) (team string, repo string, err error) {
	team, err = getTeamName(v2Impl)
	if err != nil {
		return
	}

	repos, err := v2Impl.repositories.ListByOwner(team, ListReposByOwnerOpts{})
	if err != nil {
		return "", "", err
	}

	return team, repos.Values[0]["slug"].(string), nil
}
