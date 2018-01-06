package bb

import "os"

func newTestV2Impl() *v2Impl {
	user := os.Getenv("BITBUCKET_USER")
	pass := os.Getenv("BITBUCKET_PASSWORD")
	return newV2BasicAuth(user, pass)
}
