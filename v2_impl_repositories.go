package bb

type v2RepositoriesImpl struct {
	v2Impl *v2Impl
}

func (t *v2RepositoriesImpl) ListPublic(opts ListReposOpts) (*ListResult, error) {
	req, err := t.v2Impl.client.New().Path("repositories").QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.DoList(req)
}

func (t *v2RepositoriesImpl) ListByOwner(teamOrUser string, opts ListReposByOwnerOpts) (*ListResult, error) {
	req, err := t.v2Impl.client.New().Path("repositories/").Path(teamOrUser).QueryStruct(opts).Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.DoList(req)
}
