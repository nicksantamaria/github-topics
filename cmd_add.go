package main

import (
	"errors"
	"fmt"
	"context"

	"github.com/google/go-github/v38/github"
	"gopkg.in/alecthomas/kingpin.v2"
)

// cmdAdd adds specified topics to repo.
type cmdAdd struct {
	Owner string
	Repository string
	Topics []string
	Client *github.Client
}
func (cmd *cmdAdd) run(c *kingpin.ParseContext) error {
	r, err := GetRepository(context.TODO(), cmd.Client, cmd.Owner, cmd.Repository)
	if err != nil {
		return err
	}

	for _, topic := range cmd.Topics {
		r.Topics = AppendIfMissing(r.Topics, topic)
	}

	newTopics, _, err := cmd.Client.Repositories.ReplaceAllTopics(context.TODO(), cmd.Owner, cmd.Repository, r.Topics)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to update topics - %s", err.Error()))
	}

	fmt.Println(fmt.Sprintf("Updated %s with topics %s", r.GetName(), newTopics))
	return nil
}