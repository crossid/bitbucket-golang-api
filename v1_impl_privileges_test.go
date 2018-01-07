package bb

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestV1PrivilegesImpl_ListOfAccount(t *testing.T) {
	c := newTestV1Impl()
	teamUsername, err := getTeamName(newTestV2Impl())
	assert.Nil(t, err)
	privileges, err := c.privileges.ListForAccount(teamUsername, ListPrivilegesOfAccountOpts{Private: true, Filter: AdminLevel})
	assert.Nil(t, err)
	assert.True(t, len(privileges) > 0)
}

func TestV1PrivilegesImpl_ListOfAccountAndRepo(t *testing.T) {
	// get team and repo
	v2 := newTestV2Impl()
	team, repo, err := getRepoAndTeam(v2)
	assert.Nil(t, err)

	v1 := newTestV1Impl()
	privileges, err := v1.privileges.ListOfAccountAndRepo(team, repo)
	assert.Nil(t, err)
	assert.True(t, len(privileges) > 0)
	for _, priv := range privileges {
		currRepoName := strings.Split(priv["repo"].(string), "/")
		assert.Len(t, currRepoName, 2, "repo name must contain team + repo slug")
		assert.Equal(t, currRepoName[0], team)
		assert.Equal(t, currRepoName[1], repo)
	}
}
