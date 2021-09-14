package main

import (
	"context"
	"errors"
	"log"

	"github.com/cenkalti/backoff"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func NewClientForToken(token string) (*github.Client, error) {
	if token == "" {
		return nil, errors.New("Missing token parameter")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.TODO(), ts)

	return github.NewClient(tc), nil
}

// ListRepositories returns repos of specified organisation.
func ListRepositories(ctx context.Context, gh *github.Client, org string) ([]*github.Repository, error) {
	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())

	var repos []*github.Repository
	opt := &github.RepositoryListByOrgOptions{}
	var err error
	for range ticker.C {
		for {
			reposPage, resp, err := gh.Repositories.ListByOrg(ctx, org, opt)
			if err != nil {
				log.Println("failed to retrieve repos, retrying...")
				continue
			}
			repos = append(repos, reposPage...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
		ticker.Stop()
		break
	}

	return repos, err
}

// GetRepository is a helper which returns a single repository.
func GetRepository(ctx context.Context, gh *github.Client, org string, name string) (*github.Repository, error) {
	repo, _, error := gh.Repositories.Get(ctx, org, name)
	return repo, error
}

// AppendIfMissing appends a string to a slice if it is not already present.
func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}