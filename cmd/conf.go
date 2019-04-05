package cmd

import (
	"github.com/mitchellh/go-homedir"
)

func ConfigPath() (string, error) {
	return homedir.Expand("~/.snippets")
}

func SnippetFilePath() (string, error) {
	return homedir.Expand("~/.snippets/snippets.toml")
}

func DefaultFilePath() (string, error) {
	return homedir.Expand("~/.snippets/default.toml")
}