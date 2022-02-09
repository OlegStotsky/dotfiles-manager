package dotfiles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/asverdlov/dotfiles/pkg/resources"
)

// TODO: allow cancealable async install
var HomebrewInstall = func(pkg string, options ...string) Installer {
	opts := append([]string{"install"}, options...)
	opts = append(opts, pkg)
	return func() error {
		return Execute("brew", opts...)
	}
}

var PkgInstall = func(url string) Installer {
	return func() error {
		pkgPath, err := Download(url, ".pkg")
		if err != nil {
			return err
		}
		return Execute("/bin/sh", "-c", fmt.Sprintf("sudo installer -pkg %s -target /", pkgPath))
	}
}

var DmgInstall = func(url, volumeName string) Installer {
	return func() error {
		dmgPath, err := Download(url, ".dmg")
		if err != nil {
			return err
		}
		return Execute("/bin/sh", "-c", fmt.Sprintf(`
sudo hdiutil detach '/Volumes/%s';
sudo hdiutil attach %s;
pushd '/Volumes/%s';
sudo cp -pPR '/Volumes/%s/%s.app' /Applications
popd;
sudo hdiutil detach '/Volumes/%s';
		`, volumeName, dmgPath, volumeName, volumeName, volumeName, volumeName))
	}
}

var CurlInstall = func(url string) Installer {
	return func() error {
		bashScriptPath, err := Download(url, ".sh")
		if err != nil {
			return err
		}
		return Execute("/bin/sh", bashScriptPath)
	}
}

var ExecuteInstall = func(executable string, options ...string) Installer {
	return func() error {
		return Execute(executable, options...)
	}
}

var Execute = func(executable string, options ...string) error {
	cmd := exec.Command(executable, options...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stderr, stderr)

	return cmd.Run()
}

// TODO: make it canceleable!
var Download = func(url, extension string) (string, error) {
	savePath := "*.download" + extension

	out, err := os.CreateTemp("", savePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

var Unzip = func(path string) (string, error) {
	panic("Unzip not implemented")
}

var OsxAppInstall = func(path string) Installer {
	return func() error {
		panic("osx app install not implemented")
	}
}

type PlistKvModification struct {
	Key   string
	Type  string
	Value string
}

func PlistSetBool(key string, flag bool) PlistKvModification {
	value := "true"
	if !flag {
		value = "false"
	}
	return PlistKvModification{
		Key:   key,
		Type:  "-bool",
		Value: value,
	}
}

var Sequentially = func(installers ...Installer) Installer {
	return func() error {
		for _, install := range installers {
			err := install()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

var PlistModify = func(plistPath string, modifications ...PlistKvModification) Installer {
	return func() error {
		var installers []Installer
		for _, mod := range modifications {
			installers = append(installers, ExecuteInstall("/bin/sh", "-c", fmt.Sprintf(`
defaults write %s %s %s %s
`, os.ExpandEnv(plistPath), mod.Key, mod.Type, mod.Value)))
		}
		return Sequentially(installers...)()
	}
}

func ResourceInstall(resourcePath, installPath string) Installer {
	return func() error {
		data, err := resources.Asset(resourcePath)
		if err != nil {
			return err
		}

		f, err := os.Create(os.ExpandEnv(installPath))
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(data)
		return err
	}
}

func ConfigPartInstall(path, partId string) Installer {
	return nil
}

func GitInstall(repoUrl, installPath string) Installer {
	return ExecuteInstall("git", "clone", repoUrl, os.ExpandEnv(installPath))
}

func GoInstall(url string) Installer {
	return ExecuteInstall("go", "get", "-u", url)
}
