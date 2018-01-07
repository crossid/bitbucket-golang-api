package bb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV1GroupsImpl_ListOfAccount_List(t *testing.T) {
	c := newTestV1Impl()
	teamUsername, err := getTeamName(newTestV2Impl())
	assert.Nil(t, err)
	groups, err := c.groups.ListOfAccount(teamUsername)
	assert.Nil(t, err)
	assert.NotNil(t, groups)
	assert.True(t, len(groups) > 0)

	fooGrp, err := c.groups.List(ListGroupsOpts{Group: fmt.Sprintf("%s/foo", teamUsername)})
	assert.Empty(t, fooGrp, "expect empty groups for foo slug")

	// get some group (must not be empty as the .List only works for groups attached to some repo(s)!)
	group1 := groups[1]
	group1Slug := group1["slug"]
	grp1Res, err := c.groups.List(ListGroupsOpts{Group: fmt.Sprintf("%s/%s", teamUsername, group1Slug)})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(grp1Res), "expecting 1 group to be returned for a single team/slug")
}
