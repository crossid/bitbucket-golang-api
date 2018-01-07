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
	teams, err := v2Impl.teams.List(ListTeamsOpts{Role: adminRole, Pagelen: 1})
	if err != nil {
		return "", err
	}

	return teams.Values[0]["username"].(string), nil
}
