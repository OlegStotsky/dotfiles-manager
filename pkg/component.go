package pkg

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"howett.net/plist"
)

type ComponentSearcher func() (bool, error)
type Installer func() error
type ComponentType = string

type Component struct {
	Name        string
	Type        ComponentType
	DependsOn   ComponentType
	IsInstalled ComponentSearcher
	Install     Installer
}

var FileExists = func(paths ...string) ComponentSearcher {
	return func() (bool, error) {
		for _, path := range paths {
			if _, err := os.Stat(os.ExpandEnv(path)); err == nil {
				return true, nil
			} else if !errors.Is(err, os.ErrNotExist) {
				return false, err
			}
		}
		return false, nil
	}
}

var CommandExists = func(commands ...string) ComponentSearcher {
	return func() (bool, error) {
		for _, cmd := range commands {
			_, err := exec.LookPath(cmd)
			if err == nil {
				return true, nil
			}
		}
		return false, nil
	}
}

type osxApplication struct {
	CFBundleIdentifier         string `plist:"CFBundleIdentifier"`
	CFBundleShortVersionString string `plist:"CFBundleShortVersionString"`
}

var installedApplications map[string]*osxApplication

var osxAppRoots = []string{"/", "$HOME/"}

var OsxAppExists = func(cfBundleIdentifier string) ComponentSearcher {
	return func() (bool, error) {
		if installedApplications == nil {
			apps := make(map[string]*osxApplication)
			for _, root := range osxAppRoots {
				found := false
				err := filepath.WalkDir(os.ExpandEnv(root)+"Applications", func(path string, d os.DirEntry, err error) error {
					if found {
						return filepath.SkipDir
					}
					if err == nil && d.Name() == "Info.plist" {
						app, err := readOsxApplication(path)
						if err != nil {
							return err
						}
						apps[app.CFBundleIdentifier] = app
					}
					return nil
				})
				if err != nil {
					return false, err
				}
			}
			installedApplications = apps
		}
		_, found := installedApplications[cfBundleIdentifier]
		return found, nil
	}
}

func readOsxApplication(path string) (*osxApplication, error) {
	infoPlist, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	app := &osxApplication{}
	_, err = plist.Unmarshal(infoPlist, app)
	return app, err
}

var Or = func(searchers ...ComponentSearcher) ComponentSearcher {
	return func() (bool, error) {
		for _, searcher := range searchers {
			found, err := searcher()
			if err != nil {
				return false, err
			} else if found {
				return true, nil
			}
		}
		return false, nil
	}
}
