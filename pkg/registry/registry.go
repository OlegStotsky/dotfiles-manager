package registry

import (
	"fmt"

	dotfiles "github.com/asverdlov/dotfiles/pkg/dotfiles"
)

const (
	shellType         dotfiles.ComponentType = "shell"
	zshCustomizerType dotfiles.ComponentType = "zsh-customizer"
)

var registry = make(map[string]dotfiles.Component)

func Registry() map[string]dotfiles.Component {
	return registry
}

func register(c dotfiles.Component) dotfiles.Component {
	if c.Name == "" {
		panic("Component must have a name")
	}
	if _, found := registry[c.Name]; found {
		panic(fmt.Sprintf("Found duplicate component: %s", c.Name))
	}
	registry[c.Name] = c
	return c
}
