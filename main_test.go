package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestSite(t *testing.T) {
	q := url.Values{}
	q.Set("actor", os.Getenv("GITHUB_ACTOR"))
	q.Set("refname", os.Getenv("GITHUB_REF_NAME"))
	q.Set("SHA", os.Getenv("GITHUB_SHA"))
	u, err := url.Parse("https://infinite-waters-17773.herokuapp.com?" + q.Encode())
	require.NoError(t, err)
	fmt.Println(u.String())
	req, err := http.Get(u.String())
	require.NoError(t, err)
	assert.Equal(t, 200, req.StatusCode, "%s: %s", req.Status, readBody(req.Body))
}

func readBody(r io.Reader) string {
	b, _ := ioutil.ReadAll(io.LimitReader(r, 1024))
	return string(b)
}
