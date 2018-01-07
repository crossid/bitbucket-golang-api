package bb

type v1GroupsImpl struct {
	impl *v1Impl
}

func (v1 *v1GroupsImpl) List(opts ListGroupsOpts) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("groups/").QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}

func (v1 *v1GroupsImpl) ListOfAccount(teamOrUsername string) ([]map[string]interface{}, error) {
	req, err := v1.impl.client.New().Path("groups/").Path(teamOrUsername).Request()
	if err != nil {
		return nil, err
	}

	return v1.impl.DoList(req)
}
