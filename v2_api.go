package bb

const v2BaseUrl = "https://bitbucket.org/api/2.0/"

// BitBucket API v2
//
// see: https://developer.atlassian.com/bitbucket/api/2/reference/
type v2UsersApi interface {
}

type ListTeamsListOpts struct {
	Role    role `url:"role,omitempty"`
	Pagelen int  `url:"pagelen,omitempty"`
}

type ListTeamMembersOpts struct {
	Pagelen int `url:"pagelen,omitempty"`
}

type v2Teams interface {
	// Returns all the teams that the authenticated user is associated with
	//
	// options:
	// opts.RoleFilters the teams based on the authenticated user's role on each team.
	//  - member: returns a list of all the teams which the caller is a member of at least one team group or repository owned by the team
	//  - contributor: returns a list of teams which the caller has write access to at least one repository owned by the team
	//  - admin: returns a list teams which the caller has team administrator access
	// opts.Pagelen amount of entries to return per page, default to 10
	//
	// required scopes: [team:read]
	//
	// see: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams
	List(opts ListTeamsOpts) (*ListResult, error)

	// Gets the public information associated with a team
	// If the team's profile is private, location, website and created_on elements are omitted.
	//
	// params:
	// username -The team's username or UUID.
	//
	// see: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D
	Get(username string) (map[string]interface{}, error)

	// List all members of a team
	// Returns all members of the specified team. Any member of any of the team's groups is considered a member of the team.
	// This includes users in groups that may not actually have access to any of the team's repositories.
	//
	// params:
	// teamUserName - the team user name to get the members for
	//
	// options:
	// opts.Pagelen amount of entries to return per page, default to 10
	//
	// required scopes: [account:read]
	//
	// Note that members using the "private profile" feature are not included.
	// see: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/members
	Members(teamUserName string, opts ListTeamMembersOpts) (*ListResult, error)
}

type ListResult struct {
	Pagelen int    `json:"pagelen"`
	Size    int    `json:"size"`
	Page    int    `json:"page"`
	Next    string `json:"next"`
	Values  []map[string]interface{}
}

type Error struct {
	Message string `json:"message"`
}

type BitbucketError struct {
	Type string `json:"type"`
	Err  Error  `json:"error"`
}

func (err BitbucketError) Error() string {
	return err.Err.Message
}
