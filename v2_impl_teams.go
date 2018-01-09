package bb

type v2TeamsImpl struct {
	v2Impl *V2Impl
}

func (t *v2TeamsImpl) Get(teamName string) (map[string]interface{}, error) {
	req, err := t.v2Impl.client.New().Path("teams/").Path(teamName).Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.Do(req)
}

func (t *v2TeamsImpl) List(opts ListTeamsOpts) (*ListResult, error) {
	req, err := t.v2Impl.client.New().Path("teams").QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.DoList(req)
}

func (t *v2TeamsImpl) Members(teamUserName string, opts ListTeamMembersOpts) (*ListResult, error) {
	req, err := t.v2Impl.client.New().Path("teams/").Path(teamUserName + "/").QueryStruct(opts).Path("members").Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.DoList(req)
}
