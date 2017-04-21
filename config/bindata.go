// Code generated by go-bindata.
// sources:
// config/config.toml
// DO NOT EDIT!

package config

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x6f\xdb\x3c\x0c\xc6\xef\xfa\x14\x84\x73\x79\x5f\x60\x75\x1c\xbb\x5d\x3a\x03\x39\x14\x45\x0f\x1d\xd6\x0d\x68\x8f\x41\x31\xd0\x36\x63\x0b\x91\x44\x43\x92\xd3\x3f\x9f\x7e\xa0\x92\xb6\x0b\xd6\xc3\x06\xd4\x39\x28\xd2\x43\xf2\xf9\x91\xb2\x67\x70\xc9\xe3\x93\xd7\xfd\x10\xe1\xbf\xf6\x7f\x28\x8b\x45\x05\x27\xb2\x2c\xa1\x31\xd8\x6e\x23\x8f\xf0\x95\xc3\x30\x21\xdc\xa0\x76\xf4\x09\x2e\x8c\x81\x5b\x49\x08\x70\x4b\x81\xfc\x8e\xba\x5c\xcd\xe0\x8e\x08\xbe\x5d\x5f\x5e\x7d\xbf\xbb\x82\x0d\x7b\x30\xba\x25\x17\x08\xb4\xdb\xb0\xb7\x18\x35\xbb\x5c\xa9\xd9\xc7\x3c\x6a\x06\x37\x17\xe2\x06\x97\xec\x36\xba\x9f\x7c\x32\x80\x7f\xaf\xf3\x41\x3c\x2a\xea\x68\x08\x56\x90\xdd\xa0\x74\x0e\xb7\x93\x8b\xda\xd2\x31\x5f\xa6\x76\xe4\x83\x80\xae\x20\xdb\x15\x79\x95\x97\x99\x52\x6b\x9c\xe2\xc0\xfe\x5e\x01\x38\xb4\xa9\xc8\xcb\xe8\x33\x05\xc0\xbe\x47\xa7\x9f\xf7\x0d\xbe\x1a\x5c\xff\x90\xcc\x07\x6a\x24\x6d\xf2\x46\x94\x22\x4f\xbf\xfa\xbc\x90\x3c\xec\xac\x76\x3f\x0f\xd2\xa2\x5c\x26\x71\x51\x57\x55\x55\x49\x2a\x59\xd4\x46\x92\x07\x0e\x51\x42\x82\x8d\x63\x4e\x8f\x68\x47\x43\x79\xcb\x56\x6a\x8c\xec\x45\x2b\xcf\xc4\x24\x90\x97\x38\x59\x85\x33\xe9\x18\x82\x9c\xc9\xfa\xc0\xbe\x93\xc2\x1d\x46\x6c\x30\xd0\xef\xfd\xd8\xc4\x7c\x42\x06\x43\xd4\xad\x64\x6a\x8b\xfd\x51\xab\xf3\x83\x18\x08\x7d\x3b\xd4\x67\x79\x25\x61\xe9\xfd\x4a\xb6\x86\x5b\x34\xc2\xfa\xc2\x25\xc6\xeb\x2f\x65\x51\x88\x91\xcc\x9a\xa7\xc4\x5a\x28\x00\x72\xd8\x18\xea\x60\x05\xd1\x4f\xa4\xd4\x7a\xd2\xef\xe0\x6c\x75\x83\x0e\xdf\xa7\xd9\x6b\xf5\x3e\xf2\x6f\x48\x4e\x4f\xab\xfb\xf7\x9c\xc9\xed\xb4\x67\x67\xc9\x45\xd1\xfd\x94\x2e\xb1\xa3\x1d\x19\x1e\xe5\x34\xcd\x8c\xdb\x2d\xa5\x37\xc0\x62\x3b\x68\x47\x27\xc7\xa8\x59\xaa\xdc\x8d\xac\x5d\xba\xab\xd8\x8e\xf5\x7c\xfe\x0a\x52\x97\xd5\xf2\x73\x76\x34\x86\x45\x9a\x43\xa3\x5d\x17\xde\xca\xd4\x73\x8b\xe6\x01\x3d\xd5\x9e\x25\xdc\x68\xb7\x0d\x7f\xde\x4f\x7d\x74\x15\x12\xd8\x8e\x13\xac\xe0\xac\x38\x3c\xc2\x49\x96\xfd\x93\x1c\x96\xa7\xe5\xf9\xb9\x1c\xaa\xb5\xe1\xbe\xdf\xb7\xb1\xd1\x86\x8e\x5b\xc8\x0d\xf7\x59\x6a\xf0\x31\xe8\x67\x11\x16\xc5\x7e\xbb\x1f\x7d\x75\xd8\x35\xd8\x6e\xa7\x51\xa8\x96\x42\x28\x2d\xa6\x0f\x69\x05\x1b\x34\x41\x26\x3a\x7a\x7e\x7c\x7a\x9b\xf5\xab\x02\x30\xc4\x38\x8a\x63\x76\xf8\x1f\x5e\x36\xbf\x02\x00\x00\xff\xff\x00\x99\xc5\x5b\xe7\x04\x00\x00")

func configConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigToml,
		"config/config.toml",
	)
}

func configConfigToml() (*asset, error) {
	bytes, err := configConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.toml", size: 1255, mode: os.FileMode(420), modTime: time.Unix(1492819146, 0)}
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
	"config/config.toml": configConfigToml,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.toml": &bintree{configConfigToml, map[string]*bintree{}},
	}},
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

