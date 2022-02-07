package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestSite(t *testing.T) {
	u, err := url.Parse("https://infinite-waters-17773.herokuapp.com/")
	require.NoError(t, err)
	u.Query().Set("actor", os.Getenv("GITHUB_ACTOR"))
	u.Query().Set("refname", os.Getenv("GITHUB_REF_NAME"))
	req, err := http.Get(u.String())
	require.NoError(t, err)
	assert.Equal(t, 200, req.StatusCode, req.Status)
}
