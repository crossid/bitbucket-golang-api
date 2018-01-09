package bb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV2RepositoriesImpl_ListPublic(t *testing.T) {
	// list
	pagelen := 2
	c := newTestV2Impl()
	lr, err := c.Repositories.ListPublic(ListReposOpts{Pagelen: pagelen})
	assert.Nil(t, err)
	assert.Equal(t, pagelen, len(lr.Values))
	assert.Nil(t, err)
	assert.Equal(t, pagelen, len(lr.Values))
}

func TestV2RepositoriesImpl_ListByOwner(t *testing.T) {
	pagelen := 1
	// list
	c := newTestV2Impl()

	teamUsername, err := getTeamName(c)
	assert.Nil(t, err)

	repos, err := c.Repositories.ListByOwner(teamUsername, ListReposByOwnerOpts{Pagelen: pagelen})
	assert.Nil(t, err)
	assert.Equal(t, pagelen, len(repos.Values))
}
