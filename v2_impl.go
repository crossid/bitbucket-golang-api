package bb

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type basicAuth struct {
	user, password string
}

type v2Impl struct {
	client       *sling.Sling
	teams        v2Teams
	repositories v2Repositories
	users        v2Users
	basicAuth    basicAuth
}

func newV2Of(client *http.Client) *v2Impl {
	if client == nil {
		client = http.DefaultClient
	}

	base := sling.New().Base(v2BaseUrl).Client(client)

	impl := &v2Impl{
		client: base,
	}

	impl.teams = &v2TeamsImpl{impl}
	impl.repositories = &v2RepositoriesImpl{impl}
	impl.users = &v2UsersImpl{impl}

	return impl
}

// anonymous v2 api
func newV2() *v2Impl {
	return newV2Of(nil)
}

func newV2BasicAuth(user, pass string) *v2Impl {
	client := newV2()
	client.basicAuth = basicAuth{
		user, pass,
	}

	return client
}

func (impl *v2Impl) DoList(req *http.Request) (*ListResult, error) {
	successV := &ListResult{}
	failureV := &BitbucketError{}
	resp, err := impl.DoCustom(req, successV, failureV)

	if err != nil {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("unauthorized")
		}
		return nil, err
	}

	// this is a status code > 299 and there's formatted error returned from bitbucket
	if failureV.Err.Message != "" {
		return nil, failureV
	}

	return successV, nil
}

func (impl *v2Impl) Next(next string) (*ListResult, error) {
	req, err := impl.client.Path(next).Request()
	if err != nil {

	}
	return impl.DoList(req)
}

func (impl *v2Impl) DoCustom(req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")

	if impl.basicAuth.user != "" {
		req.SetBasicAuth(impl.basicAuth.user, impl.basicAuth.password)
	}

	resp, err := impl.client.Do(req, successV, failureV)
	return resp, err
}

func (impl *v2Impl) Do(req *http.Request) (map[string]interface{}, error) {
	successV := map[string]interface{}{}
	failureV := &BitbucketError{}

	_, err := impl.DoCustom(req, &successV, failureV)
	// this is a fatal error
	if err != nil {
		return nil, err
	}

	// this is a status code > 299 and there's formatted error returned from bitbucket
	if failureV.Err.Message != "" {
		return nil, failureV
	}

	return successV, nil
}
