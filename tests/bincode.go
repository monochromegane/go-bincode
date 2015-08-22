// Code generated by go-bindata.
// sources:
// bindata.go
// cmd/go-bincode/main.go
// gen.go
// template.go
// writer.go
// DO NOT EDIT!

package bincode

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\x4d\x6f\xdb\x30\x0c\x3d\x5b\xbf\x82\xf5\xa1\x90\x01\x47\xd9\xae\x03\xb2\xcb\x3e\xb0\xee\xb0\x0d\xeb\x61\x87\x61\x18\x64\x5b\x76\x85\x38\x92\x20\xcb\x43\x82\x36\xff\x7d\xa4\xe4\x8f\x64\x68\xd7\x43\x63\x91\xd4\x7b\x8f\x8f\x94\x93\xf5\x5e\x76\x0a\x2a\x6d\x6a\xdb\x28\xc6\xf4\xc1\x59\x1f\x80\xb3\x2c\xb7\x43\x1e\xff\x6f\xd5\x51\xd5\xf4\xe9\x64\x78\xd8\xb6\xba\x57\xf4\x41\x81\x21\x78\x6d\x3a\x2c\x2b\x18\x0b\x27\x17\x61\x1a\x19\x24\x60\x62\xac\x03\x3c\xb2\xcc\x5b\x1b\x20\xfe\xa5\x62\x96\xb9\x7d\x07\xd7\x11\x3b\x06\x37\x86\xf7\xda\x2f\x11\xdd\x19\xeb\xd5\x80\x35\x3f\x7f\x4d\xb1\x33\x63\xed\x68\x6a\xe0\xd5\xcc\x52\x00\x09\xe3\xf8\xe3\xbd\xf5\x44\x46\xda\x86\x92\xce\xf0\x66\x07\x95\xe8\x6c\x8c\xf0\x02\x11\xdb\x18\xbe\xd9\x81\xd1\x3d\xd5\x66\x5e\x85\xd1\x1b\x8a\xb2\x0c\xc1\x33\xe9\xbb\x81\xae\xcd\x8c\x8f\xf9\xc6\xe6\x25\xcc\xfd\x8a\xcf\x56\x1b\x5e\x89\x45\x6c\x09\xf9\xe4\x1a\xf2\xe4\x05\x1e\x37\xd8\x1a\xde\xa8\x04\xfe\x12\x22\x72\xf6\x8a\xee\x4c\xed\x14\xf0\x16\x5e\x45\xee\x16\xf5\xfe\x2e\x41\x13\x9f\x97\x86\x06\x30\x17\xc5\x7c\x12\xb3\x03\xe9\x9c\x32\x0d\xa7\x13\xe1\xa7\x12\xa4\xd0\xd8\x11\x8a\x5e\x75\xff\x53\x1a\xdb\x16\x42\x60\xd9\xdc\x26\x5a\x25\xde\xd9\xc3\x41\x62\x51\xde\xd9\xcd\x64\x22\x82\xd1\x15\xaa\x15\xdf\x47\x83\x56\x3d\x67\xf4\x62\x24\xf0\xd9\x9e\x32\xd9\x5e\x90\xde\x3f\xd2\x27\xca\x75\x5c\xd9\x34\x84\xc5\xbe\x1f\xb2\xdf\xa3\x15\xb4\x0f\xa8\x0f\x09\x38\x85\x61\x06\xd3\xa6\xb5\x60\x07\xf1\x11\xeb\xef\xf0\x3b\x4d\x71\xa2\x58\x06\x4c\x96\x52\xa5\xb8\x1b\x70\x02\xbc\x48\x66\x19\x79\x50\xc4\x15\x33\x5f\xf0\x40\x03\xcf\xb2\xed\x16\xee\xf7\xda\xc1\x83\x6e\x1a\x65\xa0\xd1\x5e\xd5\xc1\xfa\x13\xe5\x10\x67\x5a\x5e\xf1\x49\x0e\xdf\xbc\x6a\xf5\x91\x13\x0e\xda\x2c\xf2\x02\x6e\x6f\xe3\xe8\x28\x42\x53\x7b\x9d\x88\x66\x33\x97\xa6\x08\x1f\x85\x50\xee\xcc\xd6\x3c\xae\x58\x9c\x0f\x8b\x22\xbe\x9a\xfe\x04\xb8\x23\xf1\x5a\xea\xe1\xe6\x82\xfc\x7e\x6c\x89\xfc\x42\x3c\x69\xa0\x95\x4a\xa4\xcf\x62\x7e\x38\xd6\xfd\xd8\x28\xe8\x94\x51\x5e\x06\xd5\x40\x7c\xbe\xab\x41\x09\x09\x76\xbb\xab\x2d\x85\xa7\xa7\x97\xd2\x3e\xe6\x5f\x60\x4c\xc3\x5d\x96\x6c\x7a\x67\x64\x41\xc1\xae\xeb\xcf\xff\x7b\x6b\x78\x2c\xd7\x07\x77\x61\x26\x82\xd1\xed\x33\xfb\x1b\x00\x00\xff\xff\xe6\xcd\x11\x5e\x90\x04\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 1168, mode: os.FileMode(420), modTime: time.Unix(1440241348, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _cmdGoBincodeMainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\x4d\x6b\xdc\x30\x10\x3d\xaf\x7e\xc5\xd4\xd0\x20\x87\x8d\xdd\x5e\x03\x39\x14\xfa\x01\xbd\x34\x34\x90\x4b\x5b\x82\x6c\x8f\x6d\x35\xb6\x64\xa4\x71\x3e\x28\xf9\xef\x9d\x91\xbd\xdb\xed\x2e\xb4\x17\x4b\x9a\xd9\x79\x33\xef\xcd\xdb\xc9\xd4\xf7\xa6\x43\x18\x8d\x75\x4a\xd9\x71\xf2\x81\x40\xab\x4d\xd6\x0e\xa6\xcb\xe4\x1c\x49\x0e\x1f\xe5\x1b\x29\x58\xd7\xf1\x95\xef\x9d\xa5\x7e\xae\x8a\xda\x8f\xe5\xe8\x9d\xaf\xfb\xe0\x47\xec\x8c\xc3\xb2\xf3\x17\x95\x75\xb5\x6f\x30\x53\xb9\x52\x0f\x26\xc0\x74\xdf\xc1\x52\x9c\x9e\x7e\xa6\x69\xa6\xf7\x36\x1c\x06\x6d\xe7\x7c\xc0\x08\xdf\x7e\xac\x41\xd5\xce\xae\x06\xeb\x2c\xe9\x1c\x7e\xa9\x8d\x8c\x54\xdc\xa4\xdc\xad\x09\xfa\x8c\x41\xb7\x90\xf1\x37\xe3\x43\x08\xc8\x79\xbd\x12\x72\x66\x44\x20\x0f\x73\x44\x86\x00\xea\x11\x3a\x74\x18\x0c\x61\x03\x32\x5b\x91\xe5\xa7\x90\xfb\xc1\x18\x69\xb9\x5f\x34\x36\x08\x6e\x21\x9f\x2f\x29\x04\x1c\xc2\x9a\x7c\x78\xde\x43\x48\xb1\x3e\x7f\x37\x4d\xe8\x9a\x9b\xc1\xd6\x78\x6b\x86\x19\x73\x7d\xb6\x92\xca\xb9\x78\xb9\x0a\xcc\x57\xec\xf0\x09\x26\x43\x84\xc1\xc9\x90\x6b\x6a\x87\x76\x6d\x42\x44\x9d\xab\x97\x45\x01\xa1\xb6\x28\x80\x21\xc0\xe5\x15\xac\xea\x16\x9f\x56\x42\x3a\x4d\x97\xe4\x38\x20\xb0\x6b\xad\x36\xb6\x05\xa9\x7c\x75\x05\xce\x0e\x82\xb3\xe1\xad\x16\x1f\x27\xa6\x4d\xad\xf6\x91\x15\x68\xf8\x07\x3c\xd9\xeb\x87\xef\xa2\x22\x3f\xb8\x6c\xc3\x99\x0f\x4f\xac\xfe\x5b\x7e\xbc\xf0\x34\xaa\x2c\xa1\xf2\x21\xf8\x47\xd6\xb0\xe5\x7d\x43\x4f\x34\xc5\xcb\xb2\x3c\x30\xc3\x4f\x42\x9c\x1f\xd1\xad\x36\x68\x0c\x99\xb2\x1a\x7c\x55\x8e\x26\x32\xdf\xc3\xf0\xb1\x5e\x45\xe7\xa5\x85\x0e\x78\xf1\xdf\x36\xbd\x89\xbd\xad\x7d\x98\xca\x88\xa1\xfd\xab\x03\xa7\x47\xe3\x9a\x92\x6d\xe0\xa8\x14\x45\xef\xa2\xf4\xb8\x7b\x58\x9b\xe4\x89\xc9\x71\x77\x60\xf7\x0f\x38\x72\x4d\x4c\x76\x59\x17\x9b\x32\x8e\x71\x5b\x53\x23\x30\x2e\x98\x61\xf0\x8f\x11\xc6\x79\x20\xcb\x15\x82\x55\x73\x2c\xca\x26\xa5\x30\x8a\xf5\xd8\xd1\xd6\x54\x43\xf2\xa0\x49\x9d\xc0\xc0\x60\x23\x15\x8a\x9e\x27\x3c\xed\x7e\x64\x7b\x1d\xe1\xd4\x4f\xb0\x58\x95\xcd\xb0\xfc\x56\x76\x19\x90\x66\x76\xd1\xfa\xd7\x2c\x3e\x7b\x36\xcb\x79\xe4\x5d\x6e\x33\x71\xd0\x3f\xd1\x90\x74\xd2\x64\xad\xce\x65\xf1\x3e\x08\x2a\x5b\xe6\x3c\xc2\xd5\x1f\xc7\xc8\x8b\x9d\x78\x8f\x7a\x37\xe8\x16\xde\x6c\x61\xb1\x86\x5a\xd2\x0b\xcf\xd4\x3d\xc1\xe6\xfb\xe9\x18\x85\x67\xf9\x1d\x00\x00\xff\xff\xf1\xaa\x7d\x39\x6d\x04\x00\x00")

func cmdGoBincodeMainGoBytes() ([]byte, error) {
	return bindataRead(
		_cmdGoBincodeMainGo,
		"cmd/go-bincode/main.go",
	)
}

func cmdGoBincodeMainGo() (*asset, error) {
	bytes, err := cmdGoBincodeMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/go-bincode/main.go", size: 1133, mode: os.FileMode(420), modTime: time.Unix(1440244452, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _genGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\x31\x6f\x83\x30\x10\x85\x67\xfc\x2b\xae\x4c\x20\x05\xd8\x91\x32\xb5\x6a\xa5\x4e\x95\x9a\xad\xea\xe0\xc2\x85\x58\x09\x3e\x74\x1c\x6d\xa5\x28\xff\xbd\x67\xe3\x24\x2a\x93\xf9\xde\xb3\xdf\xf3\xb9\x69\x06\x6a\x07\xf4\xc8\x56\x10\x06\x02\x5e\x3c\x74\x63\xaf\xb8\xfa\x72\xbe\xa3\x1e\x9b\xd1\x3a\x5f\xab\x54\x4d\xc7\x01\x12\x84\xca\x0d\x9e\x18\x41\x70\x96\x19\x2a\x5a\x64\x5a\xa4\xea\x1d\xaf\xc4\x4c\xb6\x3b\xda\x01\xaf\x7e\x63\xdc\x38\x11\x0b\xe4\x93\x95\x43\xb3\x77\x27\x0c\x8b\xdc\x98\xfd\xe2\x3b\x78\x49\x0d\x0a\x26\x92\x0d\x68\xd0\x06\xd6\x23\x9f\xf4\xc4\x59\xd8\x79\x25\x6b\xe4\x0c\x1f\x9f\x2b\x29\x01\x99\x89\xe1\x6c\xb2\xa6\x81\xdb\x2d\x62\x3f\xcb\xdd\xc1\x7d\xa3\xc9\xd4\x02\xed\x36\xf4\xe8\xad\x58\xb5\x66\x21\xa3\x85\xf8\xe5\x75\xbe\x51\xa2\x79\x09\xc4\x68\x25\xb7\xf0\xf6\xde\x23\xf0\x54\x21\xb8\xd3\x52\xe9\xa5\xc6\x5f\xec\x8a\xd2\x64\x6e\x1f\x2a\xc1\xc3\x16\xbc\x3b\x41\x0c\x43\x59\xd8\x07\xaa\x3e\xf3\xbf\x68\x9a\x8d\x2a\xc9\xf5\xc3\x4e\x70\x47\xcf\x3a\x9d\x42\xf7\x5e\xa7\x54\xbf\x92\xf3\xc5\xbd\x07\xe4\xd7\x9d\xfa\x2c\x79\x19\x7a\x85\xbf\x1d\x8e\xd3\x49\xcf\x9d\x6b\xa1\xf7\x38\xa0\x22\x6a\x6a\x7e\x54\xf9\xfc\xb6\x3e\x49\x1b\xee\x78\x51\xa1\x34\x17\xf3\x17\x00\x00\xff\xff\x50\x64\xdf\x8e\x00\x02\x00\x00")

func genGoBytes() ([]byte, error) {
	return bindataRead(
		_genGo,
		"gen.go",
	)
}

func genGo() (*asset, error) {
	bytes, err := genGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen.go", size: 512, mode: os.FileMode(420), modTime: time.Unix(1440244488, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templateGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xcd\x6e\xdc\x36\x10\x3e\x4b\x4f\x31\xd1\xc1\x10\xd3\x8d\xec\x14\x49\x0b\x6c\xe1\x83\x1d\xdb\x68\x0f\xb5\x0b\xd8\x68\x0f\x81\xd1\x68\x77\xa9\x5d\x22\x5a\x52\x20\xa9\x38\xc6\x62\xdf\xbd\x33\x43\xea\xc7\xb1\x36\xf5\x21\x17\x89\x9c\x9f\x6f\x66\xbe\x19\x92\x4d\xb9\xfc\x5c\xae\x25\x2c\x94\x5e\x9a\x95\x4c\x53\xff\xd8\xf0\xee\x03\xee\xc0\x79\xdb\x2e\x3d\xec\xd2\xe4\xaf\x68\x87\x12\xa5\xd7\xe9\x3e\x1a\x7a\xb9\x6d\xea\xd2\x4b\x07\x1f\xef\x5f\x77\x9b\x34\xad\x5a\xbd\x84\xdc\xbb\x41\x2f\xc0\x9b\x5b\xf6\xcd\x45\x04\x21\xd8\x2f\xa5\xed\x6d\x3a\xec\xa4\x32\x16\xfe\x9d\x81\x87\xf9\x29\xd8\x52\x63\x54\x44\x42\xe3\xa4\xb7\x3c\x1d\x9c\x7e\x02\x5f\x78\xf9\xd5\xa7\xc9\x3e\x4d\xac\xf4\xad\xd5\xbd\x92\xd2\xa4\x08\x54\xd9\x5d\x9f\xe9\xe0\xec\x10\x34\x12\x70\xd7\xd4\xb3\x34\x51\xdb\xc6\x58\x1f\xd7\x5a\xc5\x55\x64\xc3\xf2\xee\xdb\xca\x47\x1c\xe9\x72\x3b\x14\x41\x39\x8d\xd8\xa2\x34\x86\x50\x98\xc3\x51\x07\x10\x1d\xe7\x90\x45\x7d\x36\x0b\xde\x73\xf8\x74\x7c\x0c\xdc\x87\xb5\xd4\xd2\xa2\xed\x0a\x16\x8f\xb0\x36\x6f\x62\xb7\x8a\x14\x0d\x2e\x6e\xe0\xfa\xe6\x0e\x2e\x2f\xfe\xb8\x7b\x95\x76\xed\xdc\xed\xa0\xe8\x5a\xb6\xdf\xa7\x9f\x62\x06\x7d\x7d\xd3\x09\x04\xf5\x28\x7e\x1a\x24\x90\xa7\x49\xb6\x78\x44\xc2\x32\x5c\x54\x75\xb9\xe6\xff\xd6\xd3\x4f\x99\x63\x65\x5a\xaf\x6a\xda\x18\xb6\x68\x4a\xbf\x39\xae\x54\x2d\x69\x41\x82\xc0\x03\xea\xc4\x90\x4a\xa0\xf7\x40\x22\xa8\x1c\xa7\xc1\xf3\x44\x42\x1c\x9e\x38\x35\xcb\x38\x9f\x4c\x36\x09\x6a\xe5\x3c\x2c\x8c\xa9\xc3\xd6\x4a\xe7\x8d\x95\x41\x82\x33\x45\xc3\x44\x99\x17\xd7\xf2\xe1\x0a\xff\xb7\xd2\xe7\xc6\x15\x67\x76\xed\x3e\x9e\xdc\xcf\x82\xee\x83\xd1\x5e\xe9\x56\xde\xe8\x4b\x6b\x8d\x15\xe4\x58\x84\xb1\xfd\xbb\xb4\xf9\x11\x05\x9d\x41\xe6\x36\xe6\xe1\x0d\xad\x33\xdc\x64\x51\x00\xce\xb4\x76\x29\x39\xb1\x4c\x90\xe3\x39\x86\x66\x37\x4a\x0d\xad\xe8\xd7\xb9\x55\x65\xed\x64\x94\x7d\xc7\x33\x56\x81\x86\x71\xf5\xad\x7f\x57\xe6\x53\x08\x9c\xd9\x96\x2b\x3e\xe2\xae\x15\xe7\x6d\x55\x49\xbb\xdb\x73\x39\xd2\xdf\xb4\xbe\x69\x7d\x8e\x36\x64\x2a\xad\x65\x72\x70\x60\xac\x93\x3d\x29\x6f\xe7\xf7\x98\x8b\xaa\x80\xf4\xaf\x4e\x41\xab\x9a\x0f\x21\xcf\x64\xa0\x09\x50\x4b\xbc\x81\x36\x1e\x2a\xd3\xea\x15\xea\x51\x16\xdb\xcd\x74\x96\x4a\xbb\x1c\x21\x0a\x66\x34\x17\x98\x33\xbb\x34\xd6\x7c\x51\x2b\x9a\xe8\xd6\xb3\xff\x4a\x56\x4a\xcb\xd5\x3c\xe3\x16\x27\xf1\x2c\xe3\x6a\x7f\x20\xa8\x72\x90\x6d\x32\xc0\xcb\x22\xdb\xc8\xba\xc9\x42\x6c\xca\xf6\x34\xb6\x1a\x43\xfe\x8e\x9a\x80\x87\xe3\x5a\x5c\x35\x98\x98\xaf\x35\x15\x79\xeb\xf1\x48\xdb\x19\xc6\xef\x5a\x9c\x0b\x31\x15\xf8\xf2\xab\xf2\x7d\xd0\x4d\xe9\x28\x84\xb1\xe9\xcb\x11\x51\x45\x18\xf9\xcf\x82\x2e\x29\xe6\x94\xdb\x4e\x69\x69\xf9\x70\x1e\xaf\x97\x5c\x14\x24\xa6\x75\x3e\x76\x3b\x19\xdc\x78\xe8\xb1\x17\x59\xc6\xce\xb1\x73\x4f\x31\x68\x16\x19\x83\x8c\xc5\x40\xca\xa8\x85\x07\x33\xc7\x0f\x53\xd0\x85\x7e\x2b\x22\x0d\x13\xb9\x74\x93\xf7\xbc\x8a\xa8\xe9\x0b\xf9\xf1\x09\x0c\x37\xda\xe8\x6e\x9e\xbe\x4a\x3a\x83\xf1\x75\x32\x7e\xe3\xec\xe8\x02\x7f\xb0\xca\xa3\x40\x99\xe2\x1f\x5e\xd1\xc5\xcd\x57\xcf\x93\xfa\xe0\x75\xef\xba\xeb\x9f\x9c\xa3\x4e\xb6\x0b\x20\x73\x08\x55\xe1\xcd\xb8\xef\x61\xf2\x45\x1f\x55\xc0\xd0\x6b\x82\x89\x6f\x5e\x39\xbc\x79\x67\xce\x49\x7f\x8d\x45\xb8\x60\x91\x2c\x0a\x04\xcb\xcb\x48\xc0\x04\x62\xdf\xf9\xd1\x43\x24\xc2\xb4\x12\x40\x49\x80\x4c\x31\x05\x61\x78\xb6\x9c\x3c\xe5\xb1\x2c\x94\xf2\xc3\x1a\x62\x07\xc8\x9c\x81\x68\xb2\xa3\x11\xfa\x4c\x27\xf4\x64\x10\x86\x44\x62\xad\x9c\xe5\xe1\x72\xf1\xe4\x6d\x3f\xaf\x94\x1d\x06\x08\x6d\x91\xd4\x3f\x49\x78\x56\xd7\x79\xf7\xc0\x14\x17\xca\x86\x42\x66\x70\xf2\xeb\xfb\xf7\xe2\xb7\x67\xe3\x36\xae\xa6\x3b\xd6\xdc\x27\x20\x10\xdc\x7f\x87\x9b\xa9\xf1\x7d\x86\xc7\xb7\x0e\x84\x77\x30\x0c\xcf\x15\x02\x33\x04\x36\x35\x80\x9f\xfc\xf2\xee\xdd\x0b\xf1\xf6\xff\xcf\x2d\xf7\xa3\x6f\x32\xb1\x3a\x3e\x52\x8b\x22\x8c\xe1\x0c\x9c\x08\xa7\xe5\xbf\x00\x00\x00\xff\xff\xcf\x51\xb7\x6c\xe9\x09\x00\x00")

func templateGoBytes() ([]byte, error) {
	return bindataRead(
		_templateGo,
		"template.go",
	)
}

func templateGo() (*asset, error) {
	bytes, err := templateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.go", size: 2537, mode: os.FileMode(420), modTime: time.Unix(1440245822, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _writerGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xbf\x8e\xdb\x30\x0c\xc6\x67\xeb\x29\x54\x0f\x85\x0c\x18\xca\xde\xe2\x96\x0b\x9a\xad\xc5\x0d\x07\xdc\x50\x74\xb0\x1d\xca\x27\x54\xb1\x02\x8a\xae\x53\x14\xf7\xee\x25\x65\xc7\x4d\xff\xa4\x97\x25\xa1\x25\x8a\xfc\x7d\x1f\x79\x6c\xba\xaf\x4d\x0f\xba\xf5\x43\x17\xf7\xa0\x94\x3f\x1c\x23\x92\x36\xaa\x28\xdb\xd1\xf9\x58\x4a\xf0\x9d\x20\x49\x30\x7f\x46\x89\xe9\x70\x0c\xba\x24\x38\xd1\x86\x80\xe3\x86\xa0\x54\x7c\xd9\xc7\xd0\x0c\xbd\x8d\xd8\x6f\x4e\x1b\x8a\x31\xa4\xcd\x5c\x92\x1f\x55\x4a\xb9\x71\xe8\xf4\x84\x9e\xe0\x31\xee\x7c\x00\xe3\xf8\xa7\xd6\xe7\x12\x3a\x11\xfa\xa1\xaf\x85\x67\xcb\x3c\xe7\xff\x4a\x03\x62\x44\xfd\x43\x15\xae\x96\x58\xbf\xbb\xd3\x31\xd9\x2d\x02\xbf\xca\x45\x2a\x55\x78\x97\xaf\xde\xdc\xe9\xc1\x07\xc9\x2d\x10\x68\xc4\x41\x4e\x55\xf1\xa2\x8a\x3d\x38\x40\xed\xec\x36\xc4\x04\x86\x71\x8a\x49\x0a\x65\xa1\xf6\x13\x4c\x4f\x02\x86\xc6\x55\xe7\xd4\xc9\xee\xc2\x98\x9e\x73\x6a\xbb\x36\xce\xfc\x4f\x9e\x9e\x77\x11\x0f\x0d\xfd\xa1\x61\x85\xbf\x81\x88\x01\x6c\x6e\x6a\x5a\xce\x5e\xee\x38\x55\xbd\x5c\x5a\x75\xb5\xd5\x75\xbb\xcc\xe7\x2f\x32\xb5\x7a\xf6\xad\x92\xd6\xdf\x1a\xd4\xad\xce\xb3\xb4\xf7\xa3\x63\x79\x57\xe4\xbf\x6d\xb3\x33\x19\x6b\xfa\xb7\xae\x4b\x5f\x5c\x26\x23\xd8\xaf\xfe\x2c\x03\xb7\x0f\x18\x3b\x48\x69\x81\x6e\xed\xbd\xf4\x36\x55\x2d\x0a\xff\x67\x0e\x7f\xd6\xeb\xcc\x96\xb3\x8b\x2e\x7f\xf9\x63\x26\xcd\x02\x66\x7a\xe6\x3d\x86\x47\x5e\xcb\x1b\x56\x89\x33\x05\x57\x56\xd9\x7e\x1c\x13\x99\x1c\xb1\x11\xa6\xa4\xb2\xb2\x0f\x0d\xf2\x96\x2c\xe5\xaa\x5f\xc0\xf2\x84\xf3\x3e\x9c\xa0\x1b\x67\x8b\xce\xa5\xdf\xbf\x36\xed\xdf\x26\xfc\x33\x00\x00\xff\xff\xe5\xa0\xd4\x32\x7c\x03\x00\x00")

func writerGoBytes() ([]byte, error) {
	return bindataRead(
		_writerGo,
		"writer.go",
	)
}

func writerGo() (*asset, error) {
	bytes, err := writerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "writer.go", size: 892, mode: os.FileMode(420), modTime: time.Unix(1440236164, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"bindata.go": bindataGo,
	"cmd/go-bincode/main.go": cmdGoBincodeMainGo,
	"gen.go": genGo,
	"template.go": templateGo,
	"writer.go": writerGo,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"bindata.go": &bintree{bindataGo, map[string]*bintree{
	}},
	"cmd": &bintree{nil, map[string]*bintree{
		"go-bincode": &bintree{nil, map[string]*bintree{
			"main.go": &bintree{cmdGoBincodeMainGo, map[string]*bintree{
			}},
		}},
	}},
	"gen.go": &bintree{genGo, map[string]*bintree{
	}},
	"template.go": &bintree{templateGo, map[string]*bintree{
	}},
	"writer.go": &bintree{writerGo, map[string]*bintree{
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
