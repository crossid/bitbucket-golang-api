package bb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV2UsersImpl_GetCurrent(t *testing.T) {
	c := newTestV2Impl()
	curr, err := c.users.GetCurrent()
	assert.Nil(t, err)
	assert.Equal(t, getUser(), curr["username"])
}

func TestV2UsersImpl_Get(t *testing.T) {
	username := "asaf000"
	c := newTestV2Impl()
	curr, err := c.users.Get(username)
	assert.Nil(t, err)
	assert.Equal(t, username, curr["username"])
}
