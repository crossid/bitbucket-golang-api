package bb

// BitBucket API v2
//
// see: https://confluence.atlassian.com/bitbucket/version-1-423626337.html
const v1BaseUrl = "https://api.bitbucket.org/1.0/"

type permissionLevel string

const (
	AdminLevel permissionLevel = "admin"
	WriteLevel permissionLevel = "write"
	ReadLevel  permissionLevel = "read"
)

type ListGroupsOpts struct {
	// filter by group (e.g., "foo/bar" where foo = user or team and bar = group slug)
	Group string `url:"group,omitempty"`
}

type ListPrivilegesOfAccountOpts struct {
	// filter=read|write|admin
	// If you filter for the read permission, you also get the higher levels of permission such as write and
	// admin as they also include the ability to read.
	Filter permissionLevel `url:"group,omitempty"`
	// private=true query parameter to filter for private repositories
	Private bool `url:"private,omitempty"`
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

// Use the privileges endpoint to manage Bitbucket Cloud user privileges (permissions).
// It allows you to grant specific users access to read, write and or administer your repositories.
// Only the repository owner, a team account administrator, or an account with administrative rights on the repository can can query or modify
// repository privileges.
// To manage group access to your repositories, use the group-privileges Endpoint and to manage privilege settings for team accounts,
// use the privileges Resource
type v1Privilege interface {
	// GET a list of user privileges granted on all repos
	ListForAccount(teamOrUsername string, opts ListPrivilegesOfAccountOpts) ([]map[string]interface{}, error)

	// GET a list of user privileges granted of a specific repository
	ListOfAccountAndRepo(teamOrUsername, repoSlug string) ([]map[string]interface{}, error)
}

// Use the group-privileges resource to query and manipulate the group privileges (permissions) of a Bitbucket Cloudaccount's repositories.
// An account owner (or team account administrator) defines groups at the account level.
type v1GroupPrivilege interface {
}
