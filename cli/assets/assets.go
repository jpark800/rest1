// Code generated by go-bindata.
// sources:
// schema/gateway_schema.json
// DO NOT EDIT!

package assets

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

var _schemaGateway_schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xcd\x8e\xda\x30\x10\xbe\xf3\x14\x56\x76\x8f\x05\xf7\xd0\x13\xc7\xf6\xd4\xd3\x56\x6a\x6f\xd5\xaa\x32\xc9\x24\x98\x06\x3b\xb5\x87\x45\xa8\xda\x77\xaf\xe2\x84\x10\x83\xed\x84\x62\x10\x2b\xed\x25\x87\x99\xf1\xfc\xe5\x9b\xcf\x3f\x7f\x27\x84\x10\x92\x3c\xea\x74\x09\x6b\x96\xcc\x49\xb2\x44\xac\xe6\x94\xae\xb4\x14\xd3\x46\x3a\x93\xaa\xa0\x99\x62\x39\x4e\x3f\x7e\xa2\x8d\xec\x21\xf9\xd0\xac\xc4\x5d\x05\xf5\x32\xb9\x58\x41\x8a\x7b\x69\xa5\x64\x05\x0a\x39\xe8\x64\x4e\x9a\x18\x46\x5e\x30\x84\x2d\xdb\x59\x42\xbf\x9b\x4e\xab\xe0\xcf\x86\x2b\xc8\x92\x39\xf9\x69\x69\x8c\x56\xb0\x35\x1c\xad\x30\xf2\x17\x50\x9a\x4b\xe1\x52\xa5\x52\xe4\xbc\xd8\x28\x86\x5c\x0a\xed\xb2\x40\xc5\x8b\x02\x94\x53\x07\x2f\x20\xf0\xd7\x92\x89\xac\x0c\x5a\x94\x5c\xfc\xd6\x89\xa5\x7d\x3e\xaa\xcc\xd3\x28\xbb\x36\x97\x86\xf4\xbb\xa6\x51\x71\x51\x24\x27\x46\xaf\x81\xae\x44\x75\x9a\x81\x4e\x15\xaf\x30\xba\xe3\xa3\x1f\xe5\xf5\xcd\x11\xd6\x7e\xb5\x31\x79\x54\x90\xd7\xe1\x1f\x68\x06\x39\x17\xdc\x78\xa4\x56\x80\xd3\x9c\x3c\x79\x59\xf5\x30\xa5\xd8\x6e\x5c\x39\x47\xc8\x89\x5f\x8e\x15\xe0\x46\xe5\x34\x30\xbf\x56\x2d\xb5\xf7\xab\x17\xd2\x4d\x7b\xfc\x2a\x5a\xd7\x71\x4b\x98\x04\x9c\x24\x2c\xcb\x4c\x6c\x56\x7e\xeb\xb3\x4b\xce\x4a\x0d\x13\xdb\x45\xbb\x34\xe9\x25\x6c\x13\xb6\x3d\x1e\x37\xa1\x6d\xe3\xd3\x21\xd7\x80\xc8\x45\x71\x77\x84\xda\xae\x79\x0b\x6c\xda\xb5\x70\xd0\xab\xf3\x87\x76\x56\x1e\x84\x05\x66\x63\x28\x5b\x72\x82\xea\x53\xc9\xff\xe1\xfc\xb0\x68\x3f\xe5\xef\x30\x7e\x87\x31\x79\xc3\x30\xb6\xf7\xf8\xf8\x60\x0e\x02\x53\x0a\x78\xca\x9d\x8b\x3d\xbf\x22\x18\xb0\x67\x95\x83\x02\x91\x82\xbb\xa7\xcf\x63\x70\x71\x51\x02\x87\x1d\x70\x74\x06\x77\x36\xc0\x57\x1b\xb7\x8a\x29\x16\x38\xfb\xdc\xdb\xb0\x39\x2a\x38\xa0\x2b\x72\xc7\x3b\xcc\x5c\xd8\x9c\x91\x18\x65\xc8\x3c\x1e\x8c\x9e\x67\x21\x6d\x7d\x3a\x75\x43\xdb\x93\xd4\x00\x84\xed\xb4\x42\x16\xc4\xd1\x08\xaf\xb1\xe7\x48\x4c\xda\x02\x47\xc7\x09\xa1\x66\x28\x4e\x73\x90\x3f\x2f\x90\xdf\x1d\x69\xb1\xc2\x36\x25\xee\x5f\x58\xf4\x9c\xd2\x82\xe3\x72\xb3\x98\xa5\x72\x4d\x7f\x7c\xfd\xfc\xe5\xe9\xbb\xcc\x71\xcb\x14\xd0\xbc\x94\x85\x9c\xa6\x52\xa0\xe2\x0b\xba\x28\xe5\x82\xae\x99\x46\x50\x94\xa5\x35\xda\x6a\x83\x6d\xef\x92\xd1\xbe\xc9\xcc\x56\xda\xc7\x5e\xc4\x39\x2a\x81\x36\x8c\xdb\x95\xdc\xae\x23\x6d\x70\xe6\xe2\x17\x75\x77\x0b\xbd\xe9\x64\x5c\x57\x0c\xd3\x25\x5c\x78\x30\x1b\xbe\x49\xda\x77\x3b\x4f\xfb\x47\x5c\x37\xe3\xd1\x63\xaf\xf6\x9b\x65\x1d\xe4\x43\x32\x9a\x13\x8d\x65\xf0\xb9\x83\xf8\x19\x8e\x9c\xc1\x72\xc6\x96\x0f\xd3\x02\x39\x87\x83\x48\x98\x87\x48\xbf\xb8\xf8\x81\xcf\x65\x8a\xab\x8c\xfd\xa4\xf9\xbe\xfe\x0b\x00\x00\xff\xff\x93\x43\x4a\x84\x89\x16\x00\x00")

func schemaGateway_schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaGateway_schemaJson,
		"schema/gateway_schema.json",
	)
}

func schemaGateway_schemaJson() (*asset, error) {
	bytes, err := schemaGateway_schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/gateway_schema.json", size: 5769, mode: os.FileMode(420), modTime: time.Unix(1501277195, 0)}
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
	"schema/gateway_schema.json": schemaGateway_schemaJson,
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
	"schema": &bintree{nil, map[string]*bintree{
		"gateway_schema.json": &bintree{schemaGateway_schemaJson, map[string]*bintree{}},
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

