package registry

import dotfiles "github.com/asverdlov/dotfiles/pkg/dotfiles"

var OsxSystemPreferences = register(dotfiles.Component{
	Name:        "osx-system-preferences",
	IsInstalled: nil, // TODO
	Install: dotfiles.Sequentially(
		dotfiles.PlistModify(
			"$HOME/Library/Preferences/com.apple.driver.AppleBluetoothMultitouch.trackpad.plist",
			dotfiles.PlistSetBool("TrackpadThreeFingerDrag", false),
		),
		// TODO: add more
	),
})

var SshKeys = register(dotfiles.Component{
	Name:        "ssh-keys",
	IsInstalled: dotfiles.FileExists("$HOME/.ssh"),
	Install: dotfiles.ExecuteInstall("/bin/sh", "-c", `
		mkdir ~/.ssh/
		ssh-keygen -t rsa -b 2048 -C "booking gitlab" -f ~/.ssh/id_rsa_booking -N ""
		ssh-keygen -t rsa -b 2048 -C "public github" -f ~/.ssh/id_rsa_github -N ""
`),
})

var GitConfig = register(dotfiles.Component{
	Name:        "git-config",
	IsInstalled: dotfiles.FileExists("$HOME/.gitconfig"),
})

// TODO: needs post-install     source "/Users/asverdlov/.sdkman/bin/sdkman-init.sh"
var SdkMan = register(dotfiles.Component{
	Name:        "sdkman",
	IsInstalled: dotfiles.FileExists("$HOME/.sdkman"),
	Install:     dotfiles.CurlInstall("https://get.sdkman.io"),
})

var Zshrc = register(dotfiles.Component{
	Name:        "zshrc",
	IsInstalled: dotfiles.FileExists("$HOME/.zshrc"),
})

var OhMyZsh = register(dotfiles.Component{
	Name:        "oh-my-zsh",
	Type:        zshCustomizerType,
	DependsOn:   shellType, // TODO: actually depends on zsh, not on shell type
	IsInstalled: dotfiles.FileExists("$HOME/.oh-my-zsh"),
	Install:     dotfiles.CurlInstall("https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh"),
})

var Goland = register(dotfiles.Component{
	Name:        "goland",
	IsInstalled: nil, // TODO: under some folder find *.app, then check vendor of app!
})

var Vim = register(dotfiles.Component{
	Name:        "vim",
	IsInstalled: dotfiles.CommandExists("vim"),
})

var Vimrc = register(dotfiles.Component{
	Name:        "vimrc",
	IsInstalled: dotfiles.FileExists("$HOME/.vimrc"),
	Install:     dotfiles.ResourceInstall("vimrc", "$HOME/.vimrc"),
})

var GoVim = register(dotfiles.Component{
	Name:        "govim",
	IsInstalled: dotfiles.FileExists("$HOME/.vim/pack/plugins/start/govim"),
	Install:     dotfiles.GitInstall("https://github.com/govim/govim.git", "$HOME/.vim/pack/plugins/start/govim"),
})

var Pathogen = register(dotfiles.Component{
	Name:        "pathogen",
	IsInstalled: dotfiles.FileExists("$HOME/.vim/autoload/pathogen.vim"),
	Install: dotfiles.ExecuteInstall("/bin/sh", "-c", `
mkdir -p ~/.vim/autoload ~/.vim/bundle && \
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
`),
})

// TODO: requires => pathogen
var VimPuppet = register(dotfiles.Component{
	Name:        "vim-puppet",
	IsInstalled: dotfiles.FileExists("$HOME/.vim/bundle/puppet"),
	Install:     dotfiles.GitInstall("https://github.com/rodjek/vim-puppet.git", "$HOME/.vim/bundle/puppet"),
})

// TODO: needs postinstall in PATH
var Postgres = register(dotfiles.Component{
	Name:        "postgres",
	IsInstalled: dotfiles.CommandExists("psql"),
	Install:     dotfiles.HomebrewInstall("postgresql@10"),
})

var Cobra = register(dotfiles.Component{
	Name:        "cobra",
	IsInstalled: dotfiles.CommandExists("cobra"),
	Install:     dotfiles.GoInstall("github.com/spf13/cobra/cobra"),
})
