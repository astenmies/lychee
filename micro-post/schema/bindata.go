// Code generated by go-bindata.
// sources:
// schema.graphql
// DO NOT EDIT!

package schema

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xbb\x6e\xeb\x30\x0c\x86\x77\x3d\x05\x9d\xe9\xe4\x15\x34\x9e\xb4\x43\x96\x20\x4d\x51\x74\x28\x3a\x08\x16\xe3\x08\xb0\x49\x57\xa4\x9b\x06\x45\xde\xbd\x90\x7c\x69\x9a\x00\xf5\xd0\xc9\x32\x2f\x1f\xff\x5f\x94\x94\x07\x6c\x1c\x7c\x1a\x80\xb7\x0e\xe3\xc9\xc2\x43\xfa\x98\xb3\x31\x81\x14\xe3\xde\x95\x08\x1b\xf6\x98\x4b\x82\xb7\xb0\xbe\x2b\x52\x76\x71\x3c\x38\x85\x20\x40\x88\x1e\x3d\xec\x39\x82\x83\x96\x45\x17\x46\x4f\x2d\xc2\x96\x45\x21\x34\x6d\x8d\x0d\x92\xca\x2d\x04\x40\x83\xd6\x68\xe1\x51\x63\xa0\xca\x00\x44\x7c\x0f\x78\x14\x0b\xbb\x7c\x58\x31\x11\x96\x1a\x98\xd2\xc0\x0c\xed\x13\x33\x58\x51\x17\xc5\xc2\x9a\xd4\x40\x56\x64\xb3\x98\x2b\xc8\x37\x3d\xb7\xa3\xaf\x50\x2c\xbc\xf4\xc9\x7b\x5f\xe1\x6b\xea\x76\x15\xae\x69\xcf\x16\xb6\xc3\xa9\xb8\xc2\xa4\xca\x0c\x28\xbb\x28\x1c\x47\x0d\xc4\x1e\x47\x1f\x53\xc7\xc8\xc8\xf5\x07\x27\x1b\xfc\xd0\x14\xb3\xf0\x9f\xb9\x46\x47\x45\x1f\xdf\xa6\x8b\xe0\x4e\x6e\x72\xc9\x99\xae\xa6\x41\x49\x36\xf9\x8b\xff\x69\x10\x8b\xfe\xd9\x5f\x82\xfc\xea\xee\xc7\xa5\x3e\x09\xc6\x99\xbd\x94\x8e\x9e\x63\x50\xdc\x8d\x6b\xbe\x70\x36\xbf\xfa\xfc\x2e\x33\x31\xad\xf4\x5f\x8f\x5d\x0e\x2a\x7a\x49\x43\xb0\x58\xda\x3c\xde\x9c\xcd\x57\x00\x00\x00\xff\xff\x35\x5c\xb4\x08\xdf\x02\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 735, mode: os.FileMode(420), modTime: time.Unix(1578162893, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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

