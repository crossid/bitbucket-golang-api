package bb

import "os"

func newTestV2Impl() *v2Impl {
	user := os.Getenv("BITBUCKET_USER")
	pass := os.Getenv("BITBUCKET_PASSWORD")
	return newV2BasicAuth(user, pass)
}

func getTeamName(v2Impl *v2Impl) (string, error) {
	teams, err := v2Impl.teams.List(ListTeamsOpts{Role: adminRole, Pagelen: 1})
	if err != nil {
		return "", err
	}

	return teams.Values[0]["username"].(string), nil
}
