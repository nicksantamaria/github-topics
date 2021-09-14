package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v38/github"
	"gopkg.in/alecthomas/kingpin.v2"
)

// cmdAdd removes specified topics from repo.
type cmdRemove struct {
	Owner string
	Repository string
	Topics []string
	Client *github.Client
}
func (cmd *cmdRemove) run(c *kingpin.ParseContext) error {
	r, err := GetRepository(context.TODO(), cmd.Client, cmd.Owner, cmd.Repository)
	if err != nil {
		return err
	}

	persistTopics := make([]string, 0)
	for _, topic := range r.Topics {
		topicKeep := true
		for _, topicRm := range cmd.Topics {
			if topicRm == topic {
				topicKeep = false
			}
		}
		if topicKeep {
			persistTopics = append(persistTopics, topic)
		}
	}

	newTopics, _, err := cmd.Client.Repositories.ReplaceAllTopics(context.TODO(), cmd.Owner, cmd.Repository, persistTopics)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to update topics - %s", err.Error()))
	}

	fmt.Println(fmt.Sprintf("Updated %s with topics %s", r.GetName(), newTopics))
	return nil
}