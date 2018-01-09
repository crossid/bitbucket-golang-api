package bitbucket

type v1GroupPrivilegesImpl struct {
	impl *V1Impl
}

func (v1 *v1GroupPrivilegesImpl) ListForAccount(teamOrUsername string, opts ListPrivilegesOfAccountOpts) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("group-privileges/").Path(teamOrUsername).QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}

func (v1 *v1GroupPrivilegesImpl) ListOfAccountAndRepo(teamOrUsername, repoSlug string) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("group-privileges/").Path(teamOrUsername + "/").Path(repoSlug).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}
