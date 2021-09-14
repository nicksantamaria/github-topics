package main

import (
	"fmt"
	"context"

	"github.com/google/go-github/v38/github"
	"gopkg.in/alecthomas/kingpin.v2"
)

// cmdList lists all repos with a topic.
type cmdList struct {
	Owner string
	Topics []string
	Client *github.Client
}
func (cmd *cmdList) run(c *kingpin.ParseContext) error {
	repos, err := ListRepositories(context.TODO(), cmd.Client, cmd.Owner)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		filterKeep := false
		for _, topic := range repo.Topics {
			for _, topicFilter := range cmd.Topics {
				if topic == topicFilter {
					filterKeep = true
				}
			}
		}
		if filterKeep {
			fmt.Println(repo.GetName())
		}
	}

	return nil
}