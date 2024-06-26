package ghext

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/cli/go-gh/v2/pkg/config"
	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

type httpClientKey struct{}

type githubClientKey struct{}

var errInvalidContext = errors.New("invalid context")

// ContextHTTPClient returns a *http.Client from context.
func ContextHTTPClient(ctx context.Context) *http.Client {
	value := ctx.Value(httpClientKey{})
	if value != nil {
		if client, ok := value.(*http.Client); ok {
			return client
		}
	}

	cobra.CheckErr(errInvalidContext)
	return nil
}

// ContextGitHubClient returns a *github.Client from context.
func ContextGitHubClient(ctx context.Context) *github.Client {
	value := ctx.Value(githubClientKey{})
	if value != nil {
		if client, ok := value.(*github.Client); ok {
			return client
		}
	}

	cobra.CheckErr(errInvalidContext)
	return nil
}

// Context prepares GitHub CLI extension context with http.Client and github.Client values.
// You can use ContextHTTPClient and ContextGitHubClient later to get client instances from the context.
func Context(ctx context.Context) context.Context {
	htc, err := newHTTPClient()
	cobra.CheckErr(err)

	ctx = context.WithValue(ctx, httpClientKey{}, htc)

	return context.WithValue(ctx, githubClientKey{}, github.NewClient(htc))
}

func newHTTPClient() (*http.Client, error) {
	var opts api.ClientOptions

	opts.Host, _ = auth.DefaultHost()

	opts.AuthToken, _ = auth.TokenForHost(opts.Host)
	if opts.AuthToken == "" {
		return nil, fmt.Errorf("authentication token not found for host %s", opts.Host)
	}

	if cfg, _ := config.Read(nil); cfg != nil {
		opts.UnixDomainSocket, _ = cfg.Get([]string{"http_unix_socket"})
	}

	opts.EnableCache = true
	opts.CacheTTL = 2 * time.Hour

	return api.NewHTTPClient(opts)
}
