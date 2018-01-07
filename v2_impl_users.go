package bb

type v2UsersImpl struct {
	v2Impl *v2Impl
}

func (t *v2UsersImpl) GetCurrent() (map[string]interface{}, error) {
	req, err := t.v2Impl.client.New().Path("user").Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.Do(req)
}

func (t *v2UsersImpl) Get(userName string) (map[string]interface{}, error) {
	req, err := t.v2Impl.client.New().Path("users/").Path(userName).Request()
	if err != nil {
		return nil, err
	}
	return t.v2Impl.Do(req)
}
