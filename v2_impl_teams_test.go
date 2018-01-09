package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV2Teams_List_Get_Members(t *testing.T) {
	// list
	pagelen := 1
	c := newTestV2Impl()
	lr, err := c.Teams.List(ListTeamsOpts{AdminRole, pagelen})
	assert.Nil(t, err)
	assert.True(t, lr.Size > 1, "expected more than 1 team")
	assert.Equal(t, pagelen, lr.Page, "should return pageLen size")
	assert.Equal(t, lr.Page, len(lr.Values))
	team1UserName := lr.Values[0]["username"].(string)
	team1DisplayName := lr.Values[0]["display_name"]
	assert.NotNil(t, team1DisplayName)
	next := lr.Next
	assert.NotNil(t, next, "1 team found only or next doesn't work")
	// paging
	lr, err = c.Next(next)
	assert.NotEqual(t, team1DisplayName, lr.Values[0]["display_name"])
	assert.NotEqual(t, next, lr.Next)

	// get
	team1, err := c.Teams.Get(team1UserName)
	assert.Nil(t, err)
	assert.Equal(t, team1DisplayName, team1["display_name"])

	// get 404
	team1, err = c.Teams.Get("foo")
	assert.Nil(t, team1)
	assert.Error(t, err)

	// get team's members
	lr, err = c.Teams.Members(team1UserName, ListTeamMembersOpts{1})
	assert.Equal(t, lr.Pagelen, len(lr.Values))
	assert.Equal(t, 1, len(lr.Values))
	assert.Nil(t, err)

	// get team's members if non existing team
	_, err = c.Teams.Members("foo", ListTeamMembersOpts{})
	assert.Error(t, err)
}
