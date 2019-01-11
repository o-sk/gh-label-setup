# gh-label-setup

`gh-label-setup` is a CLI to create labels in Github.

## Background
When creating a new repository in Github, creating same labels every time. It is troublesome task. Using `gh-label-setup` make it easy. 


## Getting Started
### Clone
Clone this repository.
```
git clone git@github.com:o-sk/gh-label-setup.git
cd gh-label-setup
```

### Edit config.toml
Copy config file, and edit it.
```
cp config.toml.sample config.toml
```
Genrate a personal access token, if you haven't yet.

https://github.com/settings/tokens

### Run
```
go run main.go owner/repository
```

#### Example
Run follow command, then `wip` and `ready` are added.
```
go run main.go o-sk/gh-label-setup
```
https://github.com/o-sk/gh-label-setup/labels