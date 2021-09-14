package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)


func main() {
	app := kingpin.New("github-topics", "cli interface to list repos by topic, and add / remove topics")

	client, err := NewClientForToken(os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		panic(err)
	}

	list := new(cmdList)
	list.Client = client
	cl := app.Command("list", "List all repos matching a label")
	cl.Action(list.run)
	cl.Flag("owner", "owner organisation or account").Required().Envar("GITHUB_OWNER").StringVar(&list.Owner)
	cl.Arg("topics", "comma-separated list of topics to filter on").Required().StringsVar(&list.Topics)

	add := new(cmdAdd)
	add.Client = client
	ca := app.Command("add", "Add topic(s) to a repo")
	ca.Action(add.run)
	ca.Flag("owner", "owner organisation or account").Required().Envar("GITHUB_OWNER").StringVar(&add.Owner)
	ca.Arg("repo", "Name of repo to add labels on").Required().StringVar(&add.Repository)
	ca.Arg("topics", "comma-separated list of topics to filter on").Required().StringsVar(&add.Topics)

	remove := new(cmdRemove)
	remove.Client = client
	cr := app.Command("remove", "Remove topic(s) from a repo")
	cr.Action(remove.run)
	cr.Flag("owner", "owner organisation or account").Required().Envar("GITHUB_OWNER").StringVar(&remove.Owner)
	cr.Arg("repo", "Name of repo to add labels on").Required().StringVar(&remove.Repository)
	cr.Arg("topics", "comma-separated list of topics to filter on").Required().StringsVar(&remove.Topics)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}