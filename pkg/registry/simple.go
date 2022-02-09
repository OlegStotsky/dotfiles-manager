package registry

import dotfiles "github.com/asverdlov/dotfiles/pkg/dotfiles"

var homebrewPackages = []string{
	// useful CLI tools
	"fzf",
	"tree",
	"ag",
	"autojump",
	"htop",
	"jq",
	"mkcert", // can create locally trusted self-signed certificates
	"cfssl",  // certificate authority
	// essentials
	"zsh",
	"node",
	"go",
	"helm",
	"kind",
	"nvm", // node.js version management
	"golangci-lint",
}

func init() {
	for _, pkgName := range homebrewPackages {
		register(dotfiles.Component{
			Name:        pkgName,
			IsInstalled: dotfiles.CommandExists(pkgName),
			Install:     dotfiles.HomebrewInstall(pkgName),
		})
	}
}
