package app

import (
	"context"

	"github.com/Nansky/custom-fetch/pkg/client"
)

type FetchCli struct {
	httpCli *client.Client
}

// FetchImplementor contain 2 functions related to Fetch process
type FetchImplementor interface {
	FetchPages(ctx context.Context, urls ...string)
	FetchPageMetadata(ctx context.Context, url string)
}

// NewFetchImplementor will return FetchCli object instance
func NewFetchImplementor(cli *client.Client) FetchCli {
	return FetchCli{
		httpCli: cli,
	}
}
