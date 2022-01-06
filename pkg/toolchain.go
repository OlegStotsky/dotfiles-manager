package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// TODO: allow cancealable async install
var HomebrewInstall = func(pkg string) Installer {
	return func() error {
		return Execute("brew", "install", pkg)
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

var CurlInstall = func(url string) Installer {
	return func() error {
		bashScriptPath, err := Download(url, ".sh")
		if err != nil {
			return err
		}
		return Execute("/bin/sh", bashScriptPath)
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
