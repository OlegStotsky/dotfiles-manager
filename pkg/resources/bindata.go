// Code generated for package resources by go-bindata DO NOT EDIT. (@generated)
// sources:
// resources/vimrc
package resources

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _vimrc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x55\x41\x8f\xdc\x36\x0f\xbd\xeb\x57\x10\xfe\x2e\xbb\xf8\x26\x33\x49\xd0\x00\xed\x14\x73\x29\x5a\xa0\xb9\x14\x01\x12\xf4\x4e\x4b\xb4\xc5\x8e\x2d\xaa\x12\x35\x5e\xff\xfb\x42\xb2\x33\xbb\xd9\xb4\x17\x63\x44\x4a\x24\xdf\x23\xf9\x26\xaf\x41\xf1\x09\x24\x40\x07\xfb\x6f\xcf\xa3\x9f\x78\xf4\xca\x61\x34\x99\x14\x42\x01\x80\xea\xf7\xb2\xc0\xc4\x81\x20\x94\xb9\xa7\x94\xcd\xe6\x16\x2b\x73\x44\xe5\x7e\xa2\xdd\xd0\xa3\xbd\x96\xb8\x1f\x96\xc4\x4a\xdf\x58\xf2\x82\x71\xe0\x89\x8c\xa9\x5f\x5d\x23\x41\x9c\xca\xc8\x01\x24\x6c\x31\x67\x29\x99\x2e\x68\x4c\x07\x5f\x04\x46\x52\xf0\x72\xa3\x04\x8b\xa4\x2b\x87\x11\x38\x80\x7a\x02\xa5\x34\x73\xc0\x09\x16\x82\x40\xe4\x40\x05\xea\x73\xd5\xb5\x45\x38\xc2\x67\x22\xd3\x99\x0e\xce\x9e\xa6\x78\xb7\x37\xd3\x20\xa9\x05\xc1\x18\x93\xc4\xc4\xa8\x54\x1f\x57\xd8\xcd\xb7\x4a\x49\xf7\x0c\x47\xf8\x43\x94\x40\x3d\x2a\x38\xca\x91\xdb\x81\x4c\x07\x58\x54\x66\xd4\x9a\x9c\xb2\x66\x28\xb9\x06\x78\xaa\x0f\x01\xf3\x37\x65\x1e\x00\xef\x19\x64\xb8\x57\x73\x69\x97\x4d\x07\x4e\x28\x43\x10\x6d\x30\xc1\x4a\x4a\x64\x75\x5a\xa1\xa7\x55\x82\x03\x04\x4b\x49\x91\x03\x58\x99\xca\x1c\xf6\x36\xc0\x83\x65\x45\x65\x09\x8d\x03\x72\x8f\xa6\x03\x4f\xc1\x52\xa5\xa5\x64\x7a\x4e\x94\xc7\x64\x5e\xf2\xd3\x0c\xa6\x83\xcf\x65\x1c\x29\xd7\x10\x67\xf8\x65\x05\x47\x03\x96\x49\x0f\x30\xca\x8d\x67\x88\x12\xcb\x84\x4a\x1b\x98\xbf\x0b\xdb\xeb\xc0\x4f\xb0\x70\x70\xb2\xc0\xc2\xea\xc1\x31\x8e\x41\xb2\xb2\xcd\xa6\x83\x44\x51\x52\x65\xa4\x5f\x61\x94\x38\x65\xc0\x41\x29\x01\x42\xa4\xc4\xe2\x2a\x78\x0e\x68\x95\x6f\xac\xeb\x61\xe3\x88\x67\xfa\xea\xee\xa9\x8e\x5e\x57\xeb\xe0\xb0\x85\x29\xd1\xa1\x52\xbb\xf4\xd0\x5a\xf9\x6c\x78\x3c\xc2\xef\x94\x1a\xd8\xbc\xe1\xa8\x34\x7b\x49\xfa\xe2\x92\xe9\xb6\x0c\x75\xc6\x92\xa3\xb4\x75\xb2\xe1\x3b\xfd\xc9\x33\x60\x22\x98\x25\x11\x24\xca\x51\x42\xe6\x1b\x9d\x3e\xfe\xfa\xdb\x9b\x89\xaf\xdb\x50\x3f\x87\xba\x7c\x78\xfb\xf6\x35\x69\x5f\x04\x66\xbc\xd2\x8b\x80\xff\x15\xec\xf0\x2f\x75\x9a\x0e\x7a\x9c\x26\x91\xe0\x68\xc2\xb5\xe5\x7b\x69\xb8\xbc\xff\xf0\x7d\xc6\x92\xea\xba\x34\xee\x32\x8f\xf7\x99\xc8\x52\x07\x17\x2c\x06\xc8\x44\x40\x29\x49\x82\x19\xd3\x35\xd7\xdb\x75\x7b\x6b\x87\x16\x5f\x19\xd3\xf6\xad\xd0\xef\x4d\x6d\xf7\xf3\x11\x3e\xcb\xdc\x66\x27\x65\x58\xbc\x00\x4e\x89\xd0\xad\xdf\x49\x80\xe9\x60\xae\x52\x01\x31\xd1\x50\x59\x15\xe0\x90\x95\xd0\x81\xc7\x1b\xdd\xab\xcb\xed\x65\xf8\xba\xb7\xfb\xe0\x6e\x25\xff\x5c\xad\x8b\x67\xeb\x4d\x07\x16\x73\xdb\xc2\xf6\x6a\xf3\x5f\xf6\x54\xaf\xac\x2b\xe5\xd7\x9c\xe8\xce\x49\x5d\xc9\x37\x1c\x1c\x85\xba\x6a\x47\xf8\x38\x34\x4e\x16\x0c\x0a\x76\x92\xb6\xa0\x11\x13\x05\xf5\x94\x29\x1f\xa0\x4f\x68\x1b\x2f\xa4\xb6\x62\xe8\x09\xd0\x39\x72\x07\xf0\xaa\x31\x9f\x4f\xa7\x91\xd5\x97\xfe\x68\x65\x3e\xfd\xc5\x18\xc6\x99\x51\x4e\x2d\x4f\x44\xae\x84\x7d\x0c\x30\x14\x2d\xdb\x1c\x36\x4e\x4c\x07\x1c\xec\x54\x5c\xa5\x81\x73\x9d\xe2\x7d\xb1\x2a\xe0\x36\x29\xc7\x06\xaa\x86\xd9\xaa\xdd\x30\xce\x98\x74\x3f\xdf\xf5\x71\x3b\x37\x7d\xfc\x16\xf3\xb6\x23\x90\x29\xe4\x2a\xc0\x50\x85\x36\x47\xb4\x04\x3d\x79\xbc\xb1\x94\xd4\x54\x70\x57\xc0\x67\xf7\x20\xad\x7b\x75\x4c\x1d\x29\xf2\x94\xf7\xb9\xdb\x2f\x5c\xde\xbf\x4e\xd5\x9a\xcf\x61\x90\x26\x8f\x55\xf4\x27\x6a\xb2\x63\x31\x38\x76\x4d\x21\x38\xd4\x2d\x97\x58\x22\xcc\x14\x8a\xe1\x01\x3c\xe6\x87\x2e\xa2\x5a\xff\xe6\xc7\xe3\xbb\xe3\xbb\x9f\xde\xfe\xd0\x3d\x1a\x68\x5d\xde\x83\x90\x44\xfd\xff\xa5\x3d\x7b\xe5\x68\xb6\x0b\x4e\x3c\x86\x73\x0d\x78\xe8\xdb\xfa\x9e\x65\x18\x0e\xf7\x3f\xaa\xf3\xa7\x96\x8b\x82\xe3\xa1\x16\xfd\x09\xd5\xcb\x48\xc1\xd0\x13\xd9\xa2\x04\x71\x37\xfc\x8f\xc3\x40\x56\x1f\x1e\xcd\x3f\x01\x00\x00\xff\xff\x48\xe3\x34\x60\xf8\x06\x00\x00")

func vimrcBytes() ([]byte, error) {
	return bindataRead(
		_vimrc,
		"vimrc",
	)
}

func vimrc() (*asset, error) {
	bytes, err := vimrcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vimrc", size: 1784, mode: os.FileMode(420), modTime: time.Unix(1641904602, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"vimrc": vimrc,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"vimrc": &bintree{vimrc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
