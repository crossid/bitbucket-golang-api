package bb

// BitBucket API v2
//
// see: https://confluence.atlassian.com/bitbucket/version-1-423626337.html
const v1BaseUrl = "https://api.bitbucket.org/1.0/"

type ListGroupsOpts struct {
	// filter by group (e.g., "foo/bar" where foo = user or team and bar = group slug)
	Group string `url:"group,omitempty"`
}

// The groups endpoint provides functionality for querying information about Bitbucket Cloud user groups,
// creating new ones, updating memberships, and deleting them. Both individual and team accounts can define groups.
// To manage group information on an individual account, the caller must authenticate with administrative rights on the account.
// To manage groups for a team account, the caller must authenticate as a team member with administrative rights on the team
//
// see:https://confluence.atlassian.com/bitbucket/groups-endpoint-296093143.html
type v1Groups interface {
	// Get a list groups matching one or more filters.
	// The caller must authenticate with administrative rights or as a group member to view a group.
	List(opts ListGroupsOpts) ([]map[string]interface{}, error)

	// Get a list of an account's groups.  The caller must authenticate with administrative rights on the account or as a group member to view a group.
	//
	// params:
	// teamOrUsername -The team or individual account name. You can supply a user name or a valid email address.
	ListOfAccount(teamOrUsername string) ([]map[string]interface{}, error)
}
