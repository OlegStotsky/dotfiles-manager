package registry

import dotfiles "github.com/asverdlov/dotfiles/pkg/dotfiles"

var AltTab = register(dotfiles.Component{
	Name:        "alt-tab",
	IsInstalled: dotfiles.OsxAppExists("com.lwouis.alt-tab-macos"),
	Install:     dotfiles.HomebrewInstall("alt-tab"),
})

var Iterm2 = register(dotfiles.Component{
	Name:        "iterm",
	IsInstalled: dotfiles.OsxAppExists("com.googlecode.iterm2"),
})

var Telegram = register(dotfiles.Component{
	Name:        "telegram",
	IsInstalled: dotfiles.OsxAppExists("ru.keepcoder.Telegram"),
	Install: func() error {
		zipPath, err := dotfiles.Download("https://osx.telegram.org/updates/Telegram-8.3.223729.app.zip", ".zip")
		if err != nil {
			return err
		}
		unzipPath, err := dotfiles.Unzip(zipPath)
		if err != nil {
			return err
		}
		return dotfiles.OsxAppInstall(unzipPath)()
	},
})

var OnePassword = register(dotfiles.Component{
	Name: "1password",
	IsInstalled: dotfiles.Or(
		dotfiles.OsxAppExists("com.agilebits.onepassword4"),
		dotfiles.OsxAppExists("com.agilebits.onepassword7"),
	),
	// TODO: debug why redirect is not working here for "https://app-updates.agilebits.com/product_history/OPM7"
	Install: dotfiles.PkgInstall("https://c.1password.com/dist/1P/mac7/1Password-7.9.2.pkg"),
})

var OnePasswordCli = register(dotfiles.Component{
	Name:        "1password-cli",
	IsInstalled: dotfiles.CommandExists("op"),
	Install:     dotfiles.PkgInstall("https://cache.agilebits.com/dist/1P/op/pkg/v1.12.4/op_apple_universal_v1.12.4.pkg"),
})

var Tunnelbear = register(dotfiles.Component{
	Name:        "tunnelbear",
	IsInstalled: dotfiles.OsxAppExists("com.tunnelbear.mac.TunnelBear"),
})

var Slack = register(dotfiles.Component{
	Name:        "slack",
	IsInstalled: dotfiles.OsxAppExists("com.tinyspeck.slackmacgap"),
})

var Docker = register(dotfiles.Component{
	Name:        "docker",
	IsInstalled: dotfiles.CommandExists("docker"),
	Install:     dotfiles.HomebrewInstall("docker", "--cask"),
})

var JetbrainsToolbox = register(dotfiles.Component{
	Name:        "jetbrains-toolbox",
	IsInstalled: dotfiles.OsxAppExists("com.jetbrains.toolbox"),
	Install:     dotfiles.DmgInstall("https://download.jetbrains.com/toolbox/jetbrains-toolbox-1.22.10970-arm64.dmg", "JetBrains Toolbox"),
})
