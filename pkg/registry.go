package pkg

import "fmt"

const (
	shellType         ComponentType = "shell"
	zshCustomizerType ComponentType = "zsh-customizer"
)

var registry = make(map[string]Component)

func Registry() map[string]Component {
	return registry
}

func registerComponent(c Component) Component {
	if c.Name == "" {
		panic("Component must have a name")
	}
	if _, found := registry[c.Name]; found {
		panic(fmt.Sprintf("Found duplicate component: %s", c.Name))
	}
	registry[c.Name] = c
	return c
}

var Tree = registerComponent(Component{
	Name:        "tree",
	IsInstalled: CommandExists("tree"),
	Install:     HomebrewInstall("tree"),
})

var Zsh = registerComponent(Component{
	Name:        "zsh",
	Type:        shellType,
	IsInstalled: CommandExists("zsh"),
	Install:     HomebrewInstall("zsh"),
})

var Zshrc = registerComponent(Component{
	Name:        "zshrc",
	IsInstalled: FileExists("$HOME/.zshrc"),
})

var OhMyZsh = registerComponent(Component{
	Name:        "oh-my-zsh",
	Type:        zshCustomizerType,
	DependsOn:   shellType, // TODO: actually depends on zsh, not on shell type
	IsInstalled: FileExists("$HOME/.oh-my-zsh"),
	Install:     CurlInstall("https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh"),
})

var Go = registerComponent(Component{
	Name:        "go",
	IsInstalled: CommandExists("go"),
	Install:     HomebrewInstall("go"),
})

var Goland = registerComponent(Component{
	Name:        "goland",
	IsInstalled: nil, // TODO: under some folder find *.app, then check vendor of app!
})

var Vim = registerComponent(Component{
	Name:        "vim",
	IsInstalled: CommandExists("vim"),
})

var Vimrc = registerComponent(Component{
	Name:        "vimrc",
	IsInstalled: FileExists("$HOME/.vimrc"),
	// TODO: specify configuration file itself
})

var Ag = registerComponent(Component{
	Name:        "ag",
	IsInstalled: CommandExists("ag"),
	Install:     HomebrewInstall("ag"),
})

var AltTab = registerComponent(Component{
	Name:        "alt-tab",
	Install:     HomebrewInstall("alt-tab"),
})

var Kind = registerComponent(Component{
	Name:        "kind",
	IsInstalled: CommandExists("kind"),
	Install:     HomebrewInstall("kind"),
})

var Iterm2 = registerComponent(Component{
	Name:        "iterm",
	IsInstalled: OsxAppExists("com.googlecode.iterm2"),
})

// TODO: needs postinstall zshrc-patch
var Autojump = registerComponent(Component{
	Name:        "autojump",
	IsInstalled: CommandExists("autojump"),
	Install: HomebrewInstall("autojump"),
})

var Telegram = registerComponent(Component{
	Name:        "telegram",
	IsInstalled: OsxAppExists("ru.keepcoder.Telegram"),
	Install: func() error {
		zipPath, err := Download("https://osx.telegram.org/updates/Telegram-8.3.223729.app.zip", ".zip")
		if err != nil {
			return err
		}
		unzipPath, err := Unzip(zipPath)
		if err != nil {
			return err
		}
		return OsxAppInstall(unzipPath)()
	},
})

var OnePassword = registerComponent(Component{
	Name:        "1password",
	IsInstalled: Or(
		OsxAppExists("com.agilebits.onepassword4"),
		OsxAppExists("com.agilebits.onepassword7"),
	),
	// TODO: debug why redirect is not working here for "https://app-updates.agilebits.com/product_history/OPM7"
	Install: PkgInstall("https://c.1password.com/dist/1P/mac7/1Password-7.9.2.pkg"),
})

var Tunnelbear = registerComponent(Component{
	Name:        "tunnelbear",
	IsInstalled: OsxAppExists("com.tunnelbear.mac.TunnelBear"),
})

var Docker = registerComponent(Component{
	Name:        "docker",
	IsInstalled: OsxAppExists("com.docker.docker"),
	Install:     HomebrewInstall("docker"),
})

var Slack = registerComponent(Component{
	Name:        "slack",
	IsInstalled: OsxAppExists("com.tinyspeck.slackmacgap"),
})
