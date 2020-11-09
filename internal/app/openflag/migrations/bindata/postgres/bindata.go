// Code generated by go-bindata.
// sources:
// 20200704133101_init.down.sql
// 20200704133101_init.up.sql
// DO NOT EDIT!

package postgres

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

var __20200704133101_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x48\xcb\x49\x4c\x2f\xb6\xe6\x02\x04\x00\x00\xff\xff\x51\x0c\x2c\x11\x1c\x00\x00\x00")

func _20200704133101_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20200704133101_initDownSql,
		"20200704133101_init.down.sql",
	)
}

func _20200704133101_initDownSql() (*asset, error) {
	bytes, err := _20200704133101_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20200704133101_init.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1604830254, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20200704133101_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x50\x4b\x6e\x84\x30\x0c\xdd\xe7\x14\x5e\x12\xa9\xab\x4a\xb3\xe2\x30\xc8\x10\x43\xdd\x86\x04\xc5\x9e\x96\xb9\x7d\x95\x09\x2d\x64\x94\x2c\xbc\x78\x1f\xbf\xf8\x4d\x89\x50\x09\x14\x47\x4f\xc0\x33\x84\xa8\x40\x3b\x8b\x0a\xcc\x1e\x17\x31\x9d\x01\x00\x60\x07\xd5\x1b\x79\x11\x4a\x8c\xfe\xed\x49\x2b\x2e\x72\xa5\x3f\x25\x86\xb1\x50\x8e\x64\x4a\xbc\x29\xc7\xf0\xa4\x94\x76\xfd\xd7\xe5\xb4\x70\xf7\xc7\x96\x1c\x78\xdd\xf2\x8d\x69\xfa\xc0\xd4\xbd\xdf\x6e\xf6\x45\x2a\xb4\xac\x14\x54\xaa\xc0\xf6\xd6\x72\xa1\x1b\xf0\x88\x55\x5e\x49\x14\xd7\xed\x2a\x05\x47\x33\xde\xbd\x42\x88\x3f\x9d\xfd\xfb\xb9\xa7\xa6\xb1\xd0\x5b\xe2\x15\xd3\x03\xbe\xe8\x01\x1d\x3b\x6b\x6c\x6f\xcc\x51\x27\x07\x47\x7b\x29\x70\xc8\xdd\x0c\xec\x76\x88\xa1\x20\x5d\x46\x6c\xdf\xd2\xe6\x59\x6b\xf3\x6c\x6b\xcf\xbb\x6a\xc7\x89\xb7\x7d\xe7\x59\xb5\xef\xc4\x6d\x6f\x7e\x03\x00\x00\xff\xff\x5f\x38\xa5\x90\x17\x02\x00\x00")

func _20200704133101_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20200704133101_initUpSql,
		"20200704133101_init.up.sql",
	)
}

func _20200704133101_initUpSql() (*asset, error) {
	bytes, err := _20200704133101_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20200704133101_init.up.sql", size: 535, mode: os.FileMode(420), modTime: time.Unix(1604911518, 0)}
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
	"20200704133101_init.down.sql": _20200704133101_initDownSql,
	"20200704133101_init.up.sql":   _20200704133101_initUpSql,
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
	"20200704133101_init.down.sql": {_20200704133101_initDownSql, map[string]*bintree{}},
	"20200704133101_init.up.sql":   {_20200704133101_initUpSql, map[string]*bintree{}},
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
