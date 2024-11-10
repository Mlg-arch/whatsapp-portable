// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/WhatsApp.lnk (1.836kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _whatsappLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x4f\x68\x13\x69\x18\xc6\x7f\xd9\x76\x69\xda\x5e\xb2\x6c\x59\xb6\xa5\xa5\x39\x6c\x96\x85\xdd\x64\x93\x6c\xd2\x66\x02\xcb\xb6\x9b\x3f\x9b\x6e\x87\x26\x4d\x36\xb4\x2c\x23\x3a\x4d\xa7\xa4\xd0\xd8\x69\x92\x42\xdb\x83\xe0\xa9\x58\xf4\x52\x0f\xd6\x83\xb4\x88\x60\x45\x0d\x28\x1e\x84\x8a\x78\xee\x41\x7a\xa8\x54\x14\xd4\x4a\x3d\x28\x68\x41\x4f\x42\x41\x99\x2f\x13\xb1\x4d\xc5\x8b\x17\xc5\x07\xbe\x79\xe7\x7b\xff\x3c\xef\xf3\xce\xf0\xcd\xc8\x80\xa5\xe9\x1b\x0c\xdc\x12\x57\xa2\x1b\x5d\x60\x07\xae\xd9\x0e\xbe\x68\xdb\x5a\xb7\xac\x07\xca\x76\x7e\x43\xd9\x32\xec\x75\xa9\x41\x24\x5a\xd8\x8d\x6d\x9a\x68\x4f\x3c\x8a\xaf\xd9\x9f\x05\x47\x6d\x67\xef\x5a\xf9\xd5\xed\x3e\xd3\xcc\xef\xa1\xa0\x42\x35\xfe\xc7\x23\xec\x4c\x5f\xa2\x64\x63\x20\xab\x16\x0b\xdd\xba\x0e\x7f\x53\x4f\x2d\xdb\x37\x0d\xbf\xb1\x5c\x40\xb0\xdf\x90\xd8\xb8\x87\xe1\x4d\x70\x99\x01\xb2\xa8\x14\x29\xd0\x8d\x8e\x0e\xfc\xc8\x08\x5e\x0c\x91\x33\x7d\xff\x94\xec\xef\x98\x5d\xda\x94\x06\xff\xee\xc3\x7e\x31\x65\xb0\x5b\xf7\xb0\x1f\xde\xf8\xad\x8a\xdd\x85\xc6\x14\x1a\xd0\x0a\xc4\x4d\x6b\x31\xad\xdf\xac\xec\x03\x9a\x81\x1a\x20\xb4\xd2\xaa\xd8\x80\xf0\xe8\x4f\x13\x93\x9a\x3d\x44\x28\xa8\x54\x14\x29\xbb\xa5\x59\xab\xda\x35\x12\x22\x88\x52\xe5\xf7\x9a\xfe\x34\x05\x34\xf2\x14\x50\xc8\x90\x47\x65\x86\x69\x14\x33\x2b\x2c\x6a\x54\x14\x64\xc6\xc9\xa0\x32\xb6\x0f\x57\xb5\xa7\x32\x64\x53\x0d\xd4\xc1\x92\x23\x9d\x8a\x24\x13\xc9\x78\xb4\x47\x8e\x38\x94\x6e\x5d\x0f\xab\x45\x55\x91\xc7\x33\xea\xd8\x87\x46\xf9\xe2\xe0\x20\x4d\x8a\x08\x49\x12\x24\x89\x13\xa5\x07\x99\x08\x8e\x4f\xf6\xb4\xbf\xe2\x73\xc0\xa1\xf2\xb1\x5e\x1a\x34\xf7\x7a\xc6\x99\x55\x87\xb5\x42\x25\x3e\x5c\x2a\x15\xa7\x5f\x5e\x8a\xcc\x1f\x7d\x72\xe1\xde\xf2\x6c\xed\xce\x5f\xc3\xf3\xd6\xc4\xd3\xef\x4e\x17\xff\xbb\x5a\xd3\x30\x27\x7f\x2c\xbe\x60\x81\x7a\x58\x3a\x06\x78\x52\x89\xd4\xe3\xb9\xc1\xe8\x0d\x39\x10\x5a\xd9\xf9\xfe\xe4\xcf\x0b\xb9\xdb\x39\xa0\xd6\x68\xd4\x0e\x18\xdf\xae\x14\x4e\x3c\x38\xf1\xe3\xc4\x2b\xee\xfe\xa0\x13\x09\x1f\x3e\x24\x24\x3a\x08\xd0\x21\x62\x12\x7e\xdc\xf8\xc4\xde\x4b\x40\x64\xfa\xf1\xe1\xc6\x8d\x87\x4e\x51\xd1\x21\xd8\xca\x9e\x32\x56\x4d\x21\xe9\x5f\xe4\xc5\xe9\x45\xa9\x77\x79\x6d\x73\xdd\xb9\x79\xe7\xd5\x01\xa0\xa5\x22\xc4\x01\xa2\x68\xc8\x6c\xab\x31\x82\x13\x1f\x7e\x34\x34\x21\xce\x43\x00\x2f\x4e\x54\xd1\xb0\x2c\x48\x13\x32\x0c\x89\x6e\x24\xb1\x57\x19\x12\x4d\xfb\x81\x6f\x2b\xe4\xc6\xca\x30\x4e\x0e\x17\x05\x26\x98\x64\x94\x3c\x79\x34\xc6\x70\xed\x73\x9c\xaa\xff\x08\x65\x48\xe6\x24\x57\x7e\xc8\x85\x2f\x9f\xd0\x63\xe7\x63\x5d\xae\x73\x7f\x4e\x1d\x6f\x03\xb2\x46\x42\x0c\x78\xf8\xa0\xe5\xf9\xec\xeb\xed\xd8\xa9\xd2\xfd\xde\xba\xd5\x23\xee\xf7\xdf\xfd\xdb\x00\x00\x00\xff\xff\x5b\xbe\xb4\x2a\x2c\x07\x00\x00")

func whatsappLnkBytes() ([]byte, error) {
	return bindataRead(
		_whatsappLnk,
		"WhatsApp.lnk",
	)
}

func whatsappLnk() (*asset, error) {
	bytes, err := whatsappLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "WhatsApp.lnk", size: 1836, mode: os.FileMode(0666), modTime: time.Unix(1553637056, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x32, 0xa8, 0xad, 0xa, 0x1e, 0xe6, 0x59, 0xef, 0xbc, 0x8e, 0xe, 0x70, 0x80, 0x7a, 0xe5, 0x96, 0x86, 0x8c, 0x31, 0xac, 0xf7, 0xad, 0x5e, 0x6f, 0xd8, 0xa3, 0x7e, 0x12, 0xd2, 0xbd, 0x9d, 0x6}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"WhatsApp.lnk": whatsappLnk,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"WhatsApp.lnk": {whatsappLnk, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
