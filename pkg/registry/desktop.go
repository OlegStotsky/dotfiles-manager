package registry

import "github.com/asverdlov/dotfiles/pkg/dotfiles"

var Alfred = register(dotfiles.Component{
	Name:        "alfred",
	IsInstalled: dotfiles.OsxAppExists("alfred"),
	Install:     dotfiles.DmgInstall("https://cachefly.alfredapp.com/Alfred_5.0.5_2096.dmg", "Alfred", "Alfred 5"),
})
