package bitbucket

// BitBucket API v2
//
// see: https://developer.atlassian.com/bitbucket/api/2/reference/
const v2BaseUrl = "https://bitbucket.org/api/2.0/"

// List teams options
type ListTeamsOpts struct {
	// Role filters the teams based on the authenticated user's role on each team.
	//  - member: returns a list of all the teams which the caller is a member of at least one team group or repository owned by the team
	//  - contributor: returns a list of teams which the caller has write access to at least one repository owned by the team
	//  - admin: returns a list teams which the caller has team administrator access
	Role role `url:"role,omitempty"`
	// The amount of entries to return per page, default to 10.
	Pagelen int `url:"pagelen,omitempty"`
}

// List team members options
type ListTeamMembersOpts struct {
	// The amount of entries to return per page, default to 10.
	Pagelen int `url:"pagelen,omitempty"`
}

// List repositories options
type ListReposOpts struct {
	// Filter the results to include only repositories create on or after this ISO-8601 timestamp. Example: YYYY-MM-DDTHH:mm:ss.sssZ
	After string `url:"after,omitempty"`
	// The amount of entries to return per page, default to 10.
	Pagelen int `url:"pagelen,omitempty"`
	// query according to:https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
	Query string `url:"q,omitempty"`
	// sorting according to:https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
	Sort string `url:"sort,omitempty"`
}

type ListReposByOwnerOpts struct {
	// Filters the result based on the authenticated user's role on each repository.
	// - member: returns repositories to which the user has explicit read access
	// - contributor: returns repositories to which the user has explicit write access
	// - admin: returns repositories to which the user has explicit administrator access
	// - owner: returns all repositories owned by the current user
	Role role `url:"role,omitempty"`
	// The amount of entries to return per page, default to 10.
	Pagelen int `url:"pagelen,omitempty"`
	// query according to:https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
	Query string `url:"q,omitempty"`
	// sorting according to:https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
	Sort string `url:"sort,omitempty"`
}

type v2Teams interface {
	// Returns all the teams that the authenticated user is associated with
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
	// required scopes: [account:read]
	//
	// Note that members using the "private profile" feature are not included.
	// see: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/members
	Members(teamUserName string, opts ListTeamMembersOpts) (*ListResult, error)
}

type v2Repositories interface {
	// Returns a paginated list of all public repositories
	//
	// required scopes: [repository:read]
	ListPublic(opts ListReposOpts) (*ListResult, error)

	// Returns a paginated list of all repositories owned by the specified account or UUID.
	// The result can be narrowed down based on the authenticated user's role.
	// E.g. with ?role=contributor, only those repositories that the authenticated user has write access to are returned
	// (this includes any repo the user is an admin on, as that implies write access).
	ListByOwner(teamOrUser string, opts ListReposByOwnerOpts) (*ListResult, error)
}

type v2Users interface {
	// Returns the currently logged in user.
	// scopes: [account:read]
	GetCurrent() (map[string]interface{}, error)
	// Gets the public information associated with a user account.
	// If the user's profile is private, location, website and created_on elements are omitted.
	Get(userName string) (map[string]interface{}, error)
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
