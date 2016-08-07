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

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x7f\x4b\xe3\x4c\x10\xfe\xbf\x9f\x62\x68\xe1\xe5\x7d\xc1\xc6\x4d\x8a\xd6\x06\x44\xc4\xd7\xfb\x85\xbd\x03\x15\x0e\xce\x13\xd9\x64\xa7\xc9\xe2\x66\x37\xec\x6e\xb4\xfa\xe9\x6f\x66\x53\x15\xff\x3c\x30\x2d\xdd\xce\x3c\x33\xf3\x3c\x33\x3b\x99\xc1\x99\xeb\x9f\xbc\x6e\xda\x08\xff\xd6\xff\x41\x21\xf2\x05\xcc\xf9\x38\x84\xca\xc8\xfa\x3e\xba\x1e\xbe\xb9\xd0\x0e\x12\xd6\x52\x5b\xdc\x83\x53\x63\xe0\x92\x13\x02\x5c\x62\x40\xff\x80\x2a\x9b\xcc\xe0\x0a\x11\x2e\xbe\x9e\x9d\x7f\xbf\x3a\x87\x8d\xf3\x60\x74\x8d\x36\x20\x68\x4b\x56\x27\xa3\x76\x36\x9b\x4c\x66\x1f\xf3\x10\xdf\xfa\x94\xd9\x48\xbe\xdd\xe8\x66\xf0\x89\x00\xfe\xbe\xce\x07\xe9\x99\x44\x1d\x0d\xc2\x31\x4c\xd7\x92\x3b\x87\xcb\xc1\x46\xdd\xe1\x7b\x7d\xd3\xc9\xe4\x46\x0e\xb1\x75\xfe\x76\x02\x60\x65\x97\x32\x5e\xe6\x3c\x25\xdf\x0c\x9c\x6f\xa4\xd5\xcf\x63\x3f\x84\x7e\xd6\xf1\xcb\x50\x8d\x58\xa5\xdd\x9b\x8b\x4a\x6f\xdc\x60\x15\x7a\xf8\x07\xce\xce\x7f\xfc\xb6\x17\xfa\x1e\x03\x44\x19\xc9\x15\x1d\xdd\x8f\xb4\x0a\x2a\x44\x9f\x8d\xe9\xca\x55\x94\x9e\xaf\x96\xab\xb9\x38\x98\x17\xcb\x6b\xb1\x2c\x17\x45\x29\xc4\x2f\x02\x3f\x69\x1f\x22\xd4\x46\x86\x00\x8a\x4a\x84\x13\xf8\xd9\x3e\x81\x75\xf1\x84\x54\x3f\x62\xc5\x92\x07\x6f\x58\x80\xc8\xd2\xa7\x3c\x12\x5c\x58\xaa\x4e\xdb\xbb\x1d\x94\x17\xcb\x04\xe6\xe5\x82\x1e\x6e\x18\x3b\xa9\x0d\x27\xb7\x8e\x08\x28\x24\x74\xb1\xcf\x70\x2b\xbb\xde\x60\x56\xbb\x8e\x6b\xf4\xce\x33\x56\x1c\x30\x09\x2d\x15\xc7\xf1\xc9\x33\x4a\x38\xcb\x22\x1f\x9f\x8f\xce\x2b\x2e\x4c\x2a\x65\x25\x03\xde\x26\x3c\xb6\x8c\x77\x69\xfa\x99\x4a\x03\x4b\xcb\x99\x4a\x19\x57\x4b\xc3\xfc\x53\xea\xf4\xba\xd5\x01\xe8\xcb\x5b\xea\x07\x6b\xb5\x6d\x80\x66\xfd\xbf\xab\xef\x29\x9a\xbd\x6b\x59\xd3\xde\x86\x28\x8d\x49\xf7\x10\xb2\x9d\x44\xd6\x70\x03\x47\x42\xe4\x7b\x6f\xbf\x05\xb0\x82\xda\x59\x8b\x35\x47\xdf\x75\x72\x4b\x71\x07\x42\x08\xf2\xa3\x95\x95\x41\x45\x8e\xe8\x07\xe4\x79\xd8\x07\xed\x9d\xed\xd0\x46\xce\x23\x05\xac\x50\xe1\x03\x1a\xd7\xb3\x77\xbc\xac\x9d\xbf\xf7\x4e\x0d\xf5\xcb\xf2\xa8\xa4\x91\xd3\x3a\x59\xb7\xf4\x3a\xce\x5f\x96\x68\x6c\x7c\x9a\x08\x55\xef\xb4\x4d\xa3\x8e\x75\x5f\xee\xef\xe7\xab\x22\xcb\x0f\x8f\xb2\xd5\x2a\xcb\x85\x28\x8b\xc5\xf2\x90\x03\x2b\x6d\x55\x78\xcb\x2d\xf7\xe9\x7c\x94\x1e\x4b\xef\x18\xe6\xf5\x75\x03\x97\xc9\x0b\x41\xdc\xc6\x35\xcd\xc8\xbd\xd1\x06\xdf\xf3\x66\x04\x4e\x93\xaa\x6d\xd0\xcf\x0c\xe4\x62\x34\x65\xc3\xd6\x62\x67\x55\xb4\xe9\x43\xcf\xa4\x4b\x72\xa4\x5b\x49\x6f\xc9\x31\x6c\xa4\x09\x3c\x1d\xea\x77\xfb\x74\xfb\x3a\xb7\x57\x84\xd6\x27\xc6\x9e\x19\xa7\xbb\xff\x61\x34\xfe\x04\x00\x00\xff\xff\x83\xf4\x45\x4e\xc2\x04\x00\x00")

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

	info := bindataFileInfo{name: "config/config.toml", size: 1218, mode: os.FileMode(420), modTime: time.Unix(1470525399, 0)}
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
