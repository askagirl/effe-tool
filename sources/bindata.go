// Code generated by go-bindata.
// sources:
// effe/effe.go
// effe/logic/logic.go
// DO NOT EDIT!

package sources

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

var _effeEffeGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x53\x4f\x6f\xdb\x3e\x0c\x3d\xc7\x9f\x82\x3f\x03\xbf\x41\x0e\x0c\xa7\xc7\xa2\x43\x0e\x5b\x97\xfe\x01\x8a\x2d\x48\x30\xec\x58\xa8\x0e\xed\x08\x75\x64\x4f\xa2\x9b\x14\x59\xbe\xfb\x48\xc9\x0e\xda\xa1\x87\xc4\x92\xc8\x47\xbd\x47\x3d\x76\xba\x7c\xd6\x35\xc2\x4e\x1b\x9b\x24\x66\xd7\xb5\x8e\x40\x25\x93\xb4\x6a\x74\x9d\xca\x77\x47\xf2\xa9\x0d\x6d\xfb\xa7\xa2\x6c\x77\x33\x6f\x7c\x69\xf4\x0c\xab\x0a\x67\x4d\x5b\x9b\x52\xe2\xbc\x98\xf9\x57\xcf\x1f\xd9\x59\xa4\xd9\x96\xa8\x93\xb5\x7f\xb5\x9c\x91\x25\x09\xbd\x76\x08\x5c\xa0\x6b\xf0\x70\xdd\x5a\xc2\x03\x81\x27\xd7\x97\x04\xc7\x64\x52\xd2\x01\x42\xb5\x62\x88\x25\x13\x74\x0e\xf8\xd7\xba\xe4\x94\x24\x55\x6f\x4b\xa8\xd1\xa2\xd3\x84\x77\xda\x6e\x1a\x74\xaa\x6b\xdb\x06\xa6\x72\x43\xb1\xe4\x65\x2e\x15\x6a\x74\x72\x24\x54\x8a\x5f\xce\x10\xba\x0c\x04\xac\x84\x50\xb1\x42\xdf\xb5\xd6\x63\x8c\xe4\x30\x1d\x4e\x7f\xf7\xe8\x29\x13\x22\x0e\xa9\x77\x36\x42\xf6\xf0\x21\xc8\x7d\x00\x0b\x02\xae\xe6\x20\x94\x8a\x5b\x24\x95\x15\xea\xbd\xd8\x8c\x93\x36\x58\x31\xbd\x50\x3b\xa2\x26\xa6\xe2\x72\x8c\x73\x58\xb6\x2f\x2c\x29\xfb\xcc\xfb\xff\xe6\x60\x4d\x13\x13\x26\xfb\x28\xe3\x0e\xf5\x86\xe3\xe1\xe2\x35\x69\xea\xfd\x3d\x97\x75\x56\x37\x6b\x74\x8c\x5c\x48\xa7\xb2\x80\x88\x5d\x28\xae\x19\xa6\xd2\x07\x69\x2a\x2c\xb5\x35\xe5\x33\x6e\xd2\x90\x71\xe2\xbf\x93\x92\xa5\xf4\x98\x6f\x8f\x9d\x5f\xf5\x56\xb1\x8c\x82\x7f\x39\xc8\x82\xa3\x39\xec\x59\xb0\xa4\x32\x53\xc9\x7e\xcb\x6d\xb8\xe8\x1b\x3e\xf5\xb5\xe2\x60\x11\x48\xa8\x4c\xd2\x4f\x11\x32\x94\x81\xf9\x1b\x58\xe8\xd1\xb2\x27\xb9\x6c\x48\x3d\x9d\xdf\x58\x9c\x18\x9b\x13\xbc\xc8\xe4\xc4\x8a\x05\x8b\x55\xa9\x9c\xa4\x39\x5c\x5e\x5c\x5e\xe4\x90\x2e\x25\xbe\xdf\xa2\x43\xf0\xd2\x02\xa0\x2d\x82\xd8\xb2\x10\x95\xc6\x56\xed\x19\xfd\x95\x6f\x54\xa9\x1c\x31\xbc\xd2\x8d\x47\xc1\x3b\x63\xe9\x0c\x02\x89\xba\x9d\x26\xd3\xda\x5c\x4e\x2d\xe0\xc1\x50\xa8\x15\x6a\x2c\xb5\xf3\x28\x4d\x63\x55\xd3\x50\x5d\xc4\xf0\x7c\x14\xa1\x50\x63\x55\xec\xe2\x3d\x87\x44\x55\x34\x92\x48\x9b\xf4\xae\x09\x54\x38\x77\xdd\x49\x72\xa5\xd2\xab\xff\x37\xcc\x65\x2a\x92\x38\x7b\x84\xf2\x8b\xc5\x5d\x2d\x46\x7b\x14\xd4\xe0\xe5\xef\xb8\x57\xc3\xf2\xe1\xc7\xed\xe3\x62\xb5\xfa\xf3\x66\xfb\x73\xbd\x58\xb1\x24\x7e\x6c\x0f\x37\xae\xdd\xc1\x42\x24\x09\xf7\x17\xed\xe4\x11\x64\x3e\x60\x0e\x9f\xce\xc3\x72\xe4\x82\x57\xa3\x15\x8d\x58\xa9\xd2\x25\x1e\x4f\xa3\x99\x73\x78\xe7\x0d\xb6\x9c\x0b\xdc\xc6\x01\x79\x6f\xee\xe3\x88\x90\xb7\xe4\x5f\xb0\x69\x1c\xd3\x1b\xb9\x22\x9d\xb1\xd8\x7f\xc7\x77\xa0\x35\x8e\xad\xd8\x26\xe0\x1e\x8c\x27\xb4\x5f\xec\x26\x38\x5b\x71\xf7\x72\x31\x4f\xc6\x1e\xf9\x1b\x00\x00\xff\xff\x68\x0a\x0d\x0a\xb0\x04\x00\x00")

func effeEffeGoBytes() ([]byte, error) {
	return bindataRead(
		_effeEffeGo,
		"effe/effe.go",
	)
}

func effeEffeGo() (*asset, error) {
	bytes, err := effeEffeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "effe/effe.go", size: 1200, mode: os.FileMode(436), modTime: time.Unix(1461601460, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _effeLogicLogicGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\xcb\xae\xd3\x30\x10\x86\xd7\xf8\x29\x46\x96\x90\x6c\x88\x52\x0a\x88\x45\x25\x56\x15\x97\x6e\x2a\xd4\x82\xd8\x20\x51\x2b\x99\x34\x16\xc9\x38\x38\x93\xa6\x47\x55\xdf\xfd\x8c\xd3\xcb\xe9\x26\xb1\xc7\xff\xf8\xfb\x3c\x9d\x2b\xfe\xb9\x3d\x42\x13\xf6\xbe\x50\xca\xb7\x5d\x88\x0c\x46\xbd\xd2\x55\xcb\x5a\x7e\xad\xe3\x7a\x16\x1d\x95\x69\x43\xc8\xb3\x9a\xb9\x4b\x6b\xf6\x2d\x6a\x65\x95\x3a\xb8\x08\x2b\xaa\x02\xf4\x1c\x3d\xed\xe1\x33\xec\xd4\x29\x85\x9d\x04\x16\xa0\x6b\x6c\x9a\xf0\x17\xab\x0a\x75\x26\xe5\x03\xc6\xde\x07\x4a\x27\xef\xf2\xf9\x54\x2a\x43\xa1\x41\xf6\xdf\x90\x39\xdd\xd0\xb3\x13\x89\xd1\x73\x0d\x53\x9b\x3a\xab\x9d\x52\xfc\xd4\x21\x2c\x03\x31\x1e\x39\xb1\x86\x82\x41\x38\x07\xd7\x0c\x08\x9e\xf8\xd3\x47\xc9\xa9\x6a\xa0\x42\x74\x3c\x1b\x9b\x4e\x93\x79\xbe\x45\x2c\x4d\xf2\xcd\xd7\x61\x34\x36\xff\xf5\x73\x99\xbe\xe4\x8f\x6b\x47\xc1\x58\x7b\x6f\xdc\x26\xb2\x74\x9a\x2b\x27\x03\x8c\x31\xc4\xe9\x2a\x19\x48\xfe\x43\x5e\xc8\x0d\x19\x3d\x05\x81\x70\xbc\x19\x69\x2b\x30\xe4\x21\xd2\xad\x72\x9a\xc3\x5b\x98\xf8\x2b\x91\xfb\x40\xe6\xbd\x3d\x67\x40\xbe\xb9\xd3\x36\x03\x99\x82\x8f\xf0\x08\xbb\x00\x33\x18\x21\x0d\x3a\xdf\x60\xdf\x05\xea\xf1\x77\xf4\x8c\x52\x8e\xf0\xe6\x5a\xff\x3f\x60\xcf\xf6\x12\xbf\xe9\x7d\xed\x92\x5f\x65\xc6\x0c\xf4\xf7\x34\x76\xa8\x62\x68\xe1\x8b\x0c\x71\x01\xf0\xba\xfc\x43\x3a\x03\x21\xe6\xd3\xd0\x5e\x8c\x1f\xa5\xb6\x1c\xba\x47\x2b\x79\x3b\x5c\x63\x67\xf5\x1c\x00\x00\xff\xff\xad\x64\x5c\x87\x2f\x02\x00\x00")

func effeLogicLogicGoBytes() ([]byte, error) {
	return bindataRead(
		_effeLogicLogicGo,
		"effe/logic/logic.go",
	)
}

func effeLogicLogicGo() (*asset, error) {
	bytes, err := effeLogicLogicGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "effe/logic/logic.go", size: 559, mode: os.FileMode(436), modTime: time.Unix(1461600381, 0)}
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
	"effe/effe.go": effeEffeGo,
	"effe/logic/logic.go": effeLogicLogicGo,
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
	"effe": &bintree{nil, map[string]*bintree{
		"effe.go": &bintree{effeEffeGo, map[string]*bintree{}},
		"logic": &bintree{nil, map[string]*bintree{
			"logic.go": &bintree{effeLogicLogicGo, map[string]*bintree{}},
		}},
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
