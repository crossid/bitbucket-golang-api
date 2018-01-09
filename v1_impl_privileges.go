package bb

type v1PrivilegesImpl struct {
	impl *V1Impl
}

func (v1 *v1PrivilegesImpl) ListForAccount(teamOrUsername string, opts ListPrivilegesOfAccountOpts) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("privileges/").Path(teamOrUsername).QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}

func (v1 *v1PrivilegesImpl) ListOfAccountAndRepo(teamOrUsername, repoSlug string) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("privileges/").Path(teamOrUsername + "/").Path(repoSlug).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}
