package bitbucket

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type V1Impl struct {
	client          *sling.Sling
	basicAuth       basicAuth
	Groups          v1Groups
	Privileges      v1Privilege
	GroupPrivileges v1GroupPrivilege
}

func NewV1Of(client *http.Client) *V1Impl {
	if client == nil {
		client = http.DefaultClient
	}

	base := sling.New().Base(v1BaseUrl).Client(client)

	impl := &V1Impl{
		client: base,
	}

	impl.Groups = &v1GroupsImpl{impl}
	impl.Privileges = &v1PrivilegesImpl{impl}
	impl.GroupPrivileges = &v1GroupPrivilegesImpl{impl}

	return impl
}

// anonymous v1 api
func NewV1() *V1Impl {
	return NewV1Of(nil)
}

func NewV1BasicAuth(user, pass string) *V1Impl {
	client := NewV1()
	client.basicAuth = basicAuth{
		user, pass,
	}

	return client
}

func (impl *V1Impl) DoCustom(req *http.Request, successV interface{}) (*http.Response, error) {
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

func (impl *V1Impl) Do(req *http.Request) (map[string]interface{}, error) {
	successV := map[string]interface{}{}

	_, err := impl.DoCustom(req, &successV)
	// this is a fatal error
	if err != nil {
		return nil, err
	}

	return successV, nil
}

func (impl *V1Impl) DoList(req *http.Request) ([]map[string]interface{}, error) {
	successV := []map[string]interface{}{}

	_, err := impl.DoCustom(req, &successV)
	// this is a fatal error
	if err != nil {
		return nil, err
	}

	return successV, nil
}
