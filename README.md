# Github Topics CLI

A simple cli tool to manage topis on github repositories.

## Usage

### List repos with a topic
```
export GITHUB_TOKEN=xxx

# List on any of provided topics. 
github-topics list --owner=my-org topic-to-filter,topic-2  
```

### Add topic to a repo
```
export GITHUB_TOKEN=xxx
github-topics add --owner=my-org repo-name new-topic 
```

### Remove topic from a repo
```
export GITHUB_TOKEN=xxx
github-topics remove --owner=my-org repo-name old-topic 
```
