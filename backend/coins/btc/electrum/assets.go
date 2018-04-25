// Code generated by go-bindata.
// sources:
// ../../../config/certificates/electrumx/dev/ca.cert.pem
// DO NOT EDIT!

package electrum

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

var _ConfigCertificatesElectrumxDevCaCertPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x96\xc9\xae\xab\x4c\xb2\x85\xe7\x3c\xc5\x9d\x5b\x57\x74\xc6\xc6\xc3\x4c\x12\x30\x3d\x49\x0f\x33\x7a\x30\x18\xb0\x0d\xa6\x79\xfa\xd2\x3e\x5b\xfa\x55\xaa\xa3\xaa\x1c\x2e\x2d\xc5\xb7\x94\xa1\x08\xc5\xff\xff\x3c\x28\xca\x8a\xf9\x7f\x82\xe8\x78\x8a\xa4\x08\xc0\x13\xff\xa8\x84\xa1\x28\xb2\xfc\x10\x04\x08\xb4\x1a\xac\x0a\x04\xb5\xa2\x02\x8b\x06\xe2\xcb\x39\x6d\xdf\x87\x63\x00\x4a\x16\xdc\x97\xec\x2a\x19\x8b\xb0\x08\x11\xf6\x81\xa1\xc8\x89\x81\x3f\xab\x80\x63\x44\x04\x18\xcb\xe2\xaa\x22\x17\x89\xf6\x8f\x19\xd0\xbe\x28\x80\x55\x0e\x87\x60\x4f\x43\x73\x34\x1c\x6a\x95\xd7\x18\x05\x18\x6b\x08\x3a\x5e\x2a\xf7\xcf\x42\x80\x28\x1f\xfa\xb5\x90\x6f\x07\x91\x84\x26\x9d\x3f\x7b\xaa\xf4\xc4\xc3\x40\xe2\x6f\x81\x75\x7d\xf9\x4c\x33\x25\x03\xae\x31\xab\x72\xf9\xdd\xf9\xe6\x4c\xf0\x28\x22\x75\x2a\xee\x5d\x8d\x99\x60\x2f\xe4\xfe\x99\x86\x66\x43\x14\xb2\x5f\xe3\x28\xa0\x52\xf9\xb6\xa7\x91\xc3\x19\xae\xbf\x2a\xbf\x44\x84\xe0\xf6\x17\xf1\x1f\xa0\x0b\xdd\x8c\xb9\x51\x84\x22\x9a\xd0\x80\xe7\x08\x79\xe2\xd9\x40\xc6\x6a\x1e\x22\x6b\x1c\xfe\x66\xf4\x63\x84\x3c\xe3\x8f\x66\x3c\xfe\xd1\xd6\x3a\xe9\x36\xe1\x00\x2a\xac\xcd\x00\x82\xd8\x23\x40\x67\x2a\x06\xe6\x57\x84\xff\x50\x15\x04\x92\xf4\x27\x6a\xcc\xd4\xdb\xdd\x03\xd9\xaf\x71\x34\x24\xc9\x1c\xd3\x30\xa1\x14\xd1\xdc\xcb\x08\x52\x19\x6b\xf6\x31\x1b\xec\xc4\x9f\xd8\x9e\xb1\x1a\xbf\x05\x74\x24\x4c\xff\x2b\x36\x4a\x22\x95\x4a\xc3\x64\x8a\x19\x89\x4a\x5c\x08\x89\x22\x72\xc6\x8c\xfd\xf3\x39\x9b\xea\x81\xc7\x2f\xd1\x30\xee\xff\x85\xf8\x03\x54\x24\xf5\x9b\xb1\xb8\xc6\x94\xb8\x12\x75\xad\xb4\xff\xd9\x6b\x80\x7d\x00\xce\x8a\x80\x56\xb0\xd6\xb5\xa2\x81\x51\x11\x00\x46\xfd\xc1\x32\x41\xd2\x91\x88\x7d\x57\xcf\x2b\x5e\xb7\x4b\xd8\x89\x84\x34\xdd\xf2\x02\x07\x8c\x64\xc6\x5e\xe9\x3d\xc2\x40\x7c\x07\xa5\x57\x7e\x29\x26\x9f\xef\xf7\x28\xa0\x2b\x87\x75\xbf\x1d\xdf\x2a\x6a\x98\x02\xdd\x55\xf7\xeb\xd4\x14\x22\x5a\x05\x92\xad\x91\x82\x89\xb8\x7f\xd2\x5c\x37\x0d\xb3\xd0\x66\x56\xd8\x62\x3b\x49\x36\xf9\x75\xf5\x8b\xf9\xc3\x50\xd5\x38\xe6\x79\x91\xbc\x66\x99\xb4\x35\xc9\x11\xec\x30\x61\x0c\xab\xbe\x03\x2b\x44\x9a\xdc\x5d\xb4\xec\x44\x7c\xda\x57\xf7\xe6\xb8\xa6\xdb\xd6\xb9\xbd\x2f\x6b\x27\xe4\x5e\x43\x62\x66\x7c\x8b\xca\x32\x97\xae\x93\x65\xf1\x5a\xfb\x4e\xb2\x16\x97\x02\x29\xf8\x75\x9e\xf7\xeb\xbb\x14\xd8\x07\xcb\x6c\x53\xd3\x10\x65\x91\x0d\x01\x1c\x91\x78\x41\x6e\x51\x66\xae\xcb\xb9\xaa\x4e\xd6\xd9\x45\xdf\xfc\xa2\x1d\x6f\x7c\x54\xa8\xf6\xda\xa5\x88\x67\x6e\xcc\x22\x5a\x68\xdb\x28\xa4\xad\xa4\xb5\x30\x25\x5d\x71\x98\x50\xbe\x74\x08\xfb\x93\x0e\x53\xf7\x9d\xb0\x1f\xd5\x95\xfc\xd7\xe8\x7e\x05\xbc\x42\xfc\x34\x80\x3d\xda\x2a\x6d\xd1\xec\x43\x18\xa4\x83\xde\x46\x73\xaf\x37\xbf\x3a\x98\xb2\xd5\x1c\x9d\x13\x99\x9e\x08\xbc\xa7\xd7\x5c\x0f\xa5\x3c\xce\xe3\xd0\x5a\x0d\x67\x58\x4f\x34\x68\x46\x50\x7b\xbe\x6c\xd3\x8a\xc2\xf8\x3e\xf7\xbe\xd8\xe0\xc5\xcc\x68\x5a\xcf\x7d\xb0\xee\xe5\xd1\xec\x7a\x0a\xcb\x7a\x35\x72\x62\xaa\x49\xbd\x1d\x32\x00\x37\x5f\x7d\xdb\x7c\xce\xcb\xcc\xfc\x2d\xa9\x7d\xf1\x40\xf3\x69\xdf\xd7\xf7\x49\x1b\xc7\xd3\x00\xc4\xab\xb4\xe6\x8b\xd2\x21\xca\xf7\x76\x9c\x8f\x69\xcd\x28\x86\xcb\x5b\x44\xa1\x25\x45\x6c\xc8\xd1\xa3\xf2\x6d\xd5\x81\x61\x7d\xa1\xf4\x2a\x52\xcb\x97\xb1\xfb\xf4\x78\x9f\xd0\xfb\xe3\x8c\x29\xd7\xa6\xb1\x3d\x5f\x13\x30\x6d\xf9\x99\xee\xf6\x7d\x78\x55\xcb\xb2\x28\x0e\x22\x8a\x85\x67\x99\x6b\x5d\xd0\x83\x7a\x11\x37\x63\x93\xef\xf1\xb5\x88\xcb\x5e\x3c\xcb\x66\x67\xd5\xac\x43\x9d\xb8\xfc\xd8\xbb\x27\xe9\x6c\x83\xbc\xa3\x05\xe5\x16\xe9\xe4\x31\x54\x3d\xa1\x31\x69\x9d\xb8\xc7\xaf\x53\xe1\xcd\x54\x11\xdb\xe0\x80\x73\x1b\x0d\xd5\xf2\x7d\xa0\xdd\xfa\x20\x8d\xab\x2e\xdc\x94\xbd\x51\x3d\x5a\xef\x0b\xc0\xe7\xfe\x8b\xd4\x34\xcb\x23\xe9\x13\xda\xef\xa5\x37\x6e\x68\x5f\x88\x89\x8a\xcf\xa7\x6a\x59\x23\x4b\x28\xf8\xfc\x4d\x27\x4f\xf6\x6c\xfc\x0c\x25\x06\x70\x64\x8c\x35\xf6\x40\xf1\x33\x1f\x77\x7c\x16\xa5\x1a\xfb\x67\xfe\x12\xf1\x17\xdd\xc8\x4c\xd3\x45\xeb\x55\x25\xcc\xdc\x63\x7c\x96\x32\xdf\x6d\xb7\xde\x7f\x16\x89\x43\x3d\x20\xac\x57\x69\x04\x7f\x99\xff\xdd\x8b\x7e\xbd\x1e\x01\xf0\x9d\x84\xc0\x5f\xc1\x2a\x42\xf2\x00\xd6\x2f\x8d\x87\xa0\xe2\x45\x08\x0c\x01\xc4\xf1\xcf\xae\x50\xb5\x31\x51\x9a\x6f\x6e\x02\x2c\x9a\x10\x03\x54\xd7\x0a\x04\x26\x65\x12\x8a\x1d\xc0\x2f\x9f\xa6\x1a\xba\x23\xed\x66\x7e\x16\xae\x5a\x5b\x79\xe2\xe5\xba\x03\x26\x05\xe9\x13\x62\xcf\x42\xb6\x79\xcb\x21\xa0\xe5\xed\x8d\x81\x71\x17\x6c\xb1\x18\x6f\x73\x47\x82\x43\x3c\x13\x20\xa5\x27\xb2\x33\xc2\x5b\xc4\x46\x36\xaf\xec\x82\x24\x3d\x4e\x70\x12\x83\x0e\x3a\xef\x5b\x15\x79\xc1\xd2\xb0\x91\x73\xcf\x76\xf3\x32\x47\x52\xd6\x69\x21\x26\x2f\xb8\xf4\xf3\xa1\xac\x5e\x4c\x47\x20\x61\x7a\x89\xf2\x43\xc5\xe1\x67\x79\x24\xe7\x59\xd5\x3a\xb5\x67\xee\xba\x02\x13\xfd\x22\x81\x94\xec\x52\xa4\x4b\xf7\x82\xd5\xcb\x80\x6e\x9f\x4f\xe1\x72\x11\x5a\x93\x6d\xcb\x7b\xf9\x10\x54\x9d\x26\x8e\x24\x6a\xc3\x96\x9f\xcc\xed\x2b\xa6\x83\xa7\xdb\x30\x85\x0f\x61\x25\x01\xe8\xc9\x71\xad\xc9\xdd\x5d\x98\x46\x4e\x0e\x9d\x5a\x3f\x52\x36\xda\x6f\x3f\x1b\xcf\xea\x49\xfb\xea\xf4\xf4\x5d\x5f\x3d\x61\x0b\x1e\x0f\xf6\x5e\xd5\x84\xe5\x94\x0f\xa4\x79\x3b\x12\x34\xfb\x9f\x3a\x51\x31\x54\x5e\xd7\x6c\x1e\xd3\x4e\x60\x9f\x82\x3b\x54\x81\xd7\xe7\xd9\x56\xdd\x83\xb2\xa2\x50\x56\x8d\x8d\xf4\x1a\x03\x22\x99\x8a\xa7\xb2\xec\x14\x39\x3c\xd6\xeb\x78\x62\x16\x9d\x04\x6f\x5b\xfd\xe4\xb6\xd5\x98\xfd\x85\x1a\x73\x94\x15\x92\xb2\xc7\xdf\x9c\xe7\xc6\x6d\x9f\x3f\xfc\x6e\x7c\x35\x54\x84\xc1\x0d\x3e\x49\x92\xe8\xfa\x6b\xff\x3d\x63\x1f\x7c\x1f\xaf\x47\x76\x4d\x6a\xbf\x89\xdb\x2c\xe8\x58\x71\xb9\x0c\xb4\x21\x27\x96\x7d\xa6\x7a\x9a\x7c\x3e\x6f\x0c\x43\x86\x79\xc1\x0c\xe4\x3d\x09\x3a\x52\xff\xa8\x9f\xf3\x3c\x13\xf0\x82\x74\x03\x4d\x15\x37\x1c\xa5\x42\xc7\x07\x89\x67\x24\x7f\xcd\x3d\x3b\xa7\xad\x3a\x3a\x01\x37\xe3\xec\xa6\x75\x52\xa5\x94\xee\x01\x5d\xd2\x72\x92\x79\x3c\xcf\xc1\x5d\x73\xd9\x6b\xbf\x05\x44\x41\xbf\xf9\x4e\xda\x85\x7a\xd2\x6f\x1a\x70\x8b\xb4\xa9\x76\x1d\x86\x82\x70\x4d\xd7\x3e\xb7\xb0\x4d\x63\x15\x70\x78\xb4\x6e\x0b\x27\x95\x2f\xb6\xf7\xa9\x60\x50\x25\x2a\x16\x92\x86\x97\xad\x9d\x98\x76\x76\xa0\x3d\xe7\xe2\x72\xb7\xd2\x3b\xdf\x38\xd8\x6a\xe8\x11\x0e\x4d\x11\xe4\x8d\x98\xf2\x87\x21\x2b\x21\xb9\x4a\xe1\x55\x8c\x48\x7d\x0f\x99\x23\x62\x1b\xfc\x29\xaa\xe1\x69\x84\xbe\x1d\xbc\x09\x8b\x1d\xac\xed\x51\xbb\x95\x03\x2a\x2d\xc4\x8c\xb9\xd3\xc0\xd6\xf2\xe1\xa2\xf0\xac\xcd\xd9\x92\xde\xcc\x16\xa7\xd0\x0c\xf1\xe7\x9c\x10\x4d\xf4\xf7\x89\xf1\xaf\x00\x00\x00\xff\xff\x3a\xd1\x0f\xa7\x7f\x08\x00\x00")

func ConfigCertificatesElectrumxDevCaCertPemBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCertificatesElectrumxDevCaCertPem,
		"../../../config/certificates/electrumx/dev/ca.cert.pem",
	)
}

func ConfigCertificatesElectrumxDevCaCertPem() (*asset, error) {
	bytes, err := ConfigCertificatesElectrumxDevCaCertPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../config/certificates/electrumx/dev/ca.cert.pem", size: 2175, mode: os.FileMode(420), modTime: time.Unix(1520957162, 0)}
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
	"../../../config/certificates/electrumx/dev/ca.cert.pem": ConfigCertificatesElectrumxDevCaCertPem,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"config": &bintree{nil, map[string]*bintree{
					"certificates": &bintree{nil, map[string]*bintree{
						"electrumx": &bintree{nil, map[string]*bintree{
							"dev": &bintree{nil, map[string]*bintree{
								"ca.cert.pem": &bintree{ConfigCertificatesElectrumxDevCaCertPem, map[string]*bintree{}},
							}},
						}},
					}},
				}},
			}},
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
