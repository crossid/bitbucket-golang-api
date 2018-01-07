package bb

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type v1Impl struct {
	client    *sling.Sling
	basicAuth basicAuth
	groups    v1Groups
}

func newV1Of(client *http.Client) *v1Impl {
	if client == nil {
		client = http.DefaultClient
	}

	base := sling.New().Base(v1BaseUrl).Client(client)

	impl := &v1Impl{
		client: base,
	}

	impl.groups = &v1GroupsImpl{impl}

	return impl
}

// anonymous v1 api
func newV1() *v1Impl {
	return newV1Of(nil)
}

func newV1BasicAuth(user, pass string) *v1Impl {
	client := newV1()
	client.basicAuth = basicAuth{
		user, pass,
	}

	return client
}

func (impl *v1Impl) DoCustom(req *http.Request, successV interface{}) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")

	if impl.basicAuth.user != "" {
		req.SetBasicAuth(impl.basicAuth.user, impl.basicAuth.password)
	}

	resp, err := impl.client.Do(req, successV, nil)

	if resp.StatusCode > 299 {
		// TODO: Errors are string, handle it right
		return resp, fmt.Errorf("request has failed")

	}

	return resp, err
}

func (impl *v1Impl) Do(req *http.Request) (map[string]interface{}, error) {
	successV := map[string]interface{}{}

	_, err := impl.DoCustom(req, &successV)
	// this is a fatal error
	if err != nil {
		return nil, err
	}

	return successV, nil
}

func (impl *v1Impl) DoList(req *http.Request) ([]map[string]interface{}, error) {
	successV := []map[string]interface{}{}

	_, err := impl.DoCustom(req, &successV)
	// this is a fatal error
	if err != nil {
		return nil, err
	}

	return successV, nil
}
