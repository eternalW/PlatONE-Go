// Code generated by go-bindata.
// sources:
// ../../release/linux/conf/contracts/cnsManager.cpp.abi.json
// ../../release/linux/conf/contracts/fireWall.abi.json
// ../../release/linux/conf/contracts/groupManager.cpp.abi.json
// ../../release/linux/conf/contracts/nodeManager.cpp.abi.json
// ../../release/linux/conf/contracts/paramManager.cpp.abi.json
// ../../release/linux/conf/contracts/userManager.cpp.abi.json
// DO NOT EDIT!

package precompile

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

var _ReleaseLinuxConfContractsCnsmanagerCppAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x41\x6b\xb4\x30\x10\xbd\xfb\x2b\x42\xce\x7b\xfa\xbe\xd2\xc3\xde\xba\x85\x42\x2f\x16\xda\xe3\xb2\x87\xa0\xa3\x0c\xd4\x89\x24\x93\x05\x5b\xfa\xdf\x8b\x82\xae\xae\x60\x75\xb5\xed\x36\xe8\x41\xd1\x24\xf3\x98\xf7\x26\x2f\x83\xfb\x40\x08\x21\xde\xab\x7b\x79\x49\x52\x19\xc8\xad\x90\x11\xd9\x67\x48\xd1\x32\x98\x07\xa3\xb3\x47\x42\x96\x9b\xd3\x34\xa4\xdc\xb1\x95\x5b\xb1\x6f\xbe\x75\x03\xf5\x02\x56\xcf\x4d\x7f\x9c\x8b\xbc\x1a\xb7\x6c\x90\x52\xd9\x99\xf0\xb1\x19\x1b\xfd\x08\xc6\xa2\xa6\xc9\x00\xcd\xdb\xa1\x95\x9c\x76\x3c\x35\xbb\x21\x60\x24\xfe\xff\x6f\x0c\x6e\xa4\xc9\xb2\x22\x2e\x17\x25\xea\xd5\xb6\xf9\x6a\xa2\x25\x8e\x22\x2e\x53\x0d\x5a\x0c\x0d\x0b\xe8\xa3\x70\xa3\x01\x54\x1c\x1b\xb0\x76\xad\x8c\x6e\x65\xc4\x68\x20\x5a\xb7\xf4\x9f\x12\x2e\x05\xbe\xd7\xc4\x46\x45\x7c\xd7\xab\xea\x55\xbf\x33\xfc\x65\x80\xdb\x02\xb2\x71\x73\xf5\xab\x2d\x19\xe2\x5a\xc9\x99\x1a\xe6\x2a\x85\xd0\x65\x53\x8b\x75\xb4\x8a\x65\xfc\x17\x7c\x1b\xac\x93\xd1\xbb\xc1\x63\x15\x77\xc5\x93\xc1\x14\x69\x9e\x9a\xfa\x3c\xc6\x8c\x9c\xbd\x20\x1b\x93\x13\xd7\xbb\x22\xec\x1a\xd6\x4f\x59\xde\xd5\x9e\x28\xcb\xb2\xbb\xc8\x99\xe2\x5d\xbb\xf5\x5d\x76\xe1\x3d\xdb\xd7\xe5\xce\x1e\x5b\xc7\xef\x10\x1d\x6a\xc6\xa4\xb8\x8c\xd2\x1a\xcb\x21\xf1\xed\xcd\xd4\xd6\x64\x7a\xde\xf5\x0a\x38\x02\xf1\x97\xa6\x78\xf1\xaf\x8e\x79\x7d\xf3\x32\x69\x05\x87\xe0\x33\x00\x00\xff\xff\x6b\x0c\x17\xf6\xd4\x11\x00\x00")

func ReleaseLinuxConfContractsCnsmanagerCppAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsCnsmanagerCppAbiJson,
		"../../release/linux/conf/contracts/cnsManager.cpp.abi.json",
	)
}

func ReleaseLinuxConfContractsCnsmanagerCppAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsCnsmanagerCppAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/cnsManager.cpp.abi.json", size: 4564, mode: os.FileMode(436), modTime: time.Unix(1596005127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ReleaseLinuxConfContractsFirewallAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x4b\xc3\x40\x10\xbd\xe7\x57\x2c\x7b\xce\x49\xc5\x43\x6f\xe2\x07\x78\xd1\x83\xc7\x52\xc2\xd2\x4c\x64\x21\x9d\x0d\x3b\xb3\x6a\x90\xfe\x77\x69\x6b\xe2\x46\x8b\x66\xa5\x25\x65\xd9\x1c\x02\x49\x66\xdf\x63\xde\x07\x64\x9e\x09\x21\xc4\xfb\xf6\xbe\xb9\x24\xaa\x15\xc8\x99\x90\x45\x41\x2d\x15\x77\xaf\x8f\x0d\xa0\xcc\xbf\xbe\x6b\x6c\x1c\x93\x9c\x89\x79\xff\x6e\x88\xf0\x03\x49\x95\xa5\x05\x22\x0f\xa4\x1f\xe1\xb6\xd9\x8e\x10\x5b\x8d\xcf\x72\x30\xb0\xee\x9f\x16\x1e\xbd\x71\x1c\xca\xff\x1b\xb1\x46\x3e\x3f\x1b\xc3\xbb\x34\x48\xac\x90\x37\x87\x2a\x55\x13\xf8\x9a\x74\x68\x95\xc3\x25\x6b\x83\x3b\xc0\x75\xfe\x97\xb6\xd7\xb5\x19\x02\x25\x71\x0f\x27\xee\x55\x59\x4e\x23\x6d\x3e\x9a\x60\xb7\xcf\xd1\xf0\xad\xab\x21\x25\x63\x5f\xed\x40\xd9\x38\xb3\x11\xbf\x79\x37\x50\xc7\x69\x5d\xaa\xf5\x37\xb4\xd0\x64\x3c\x01\xa7\x64\xa4\x64\xec\x4b\x06\x2b\x76\x14\xef\x8f\xd6\x78\x62\x5f\x5e\xb6\xee\x20\xea\xde\xaf\x1a\x63\x4f\xbd\x7a\x27\xe9\xdb\xb4\xad\xb8\x7d\x9b\xce\xb7\x78\x5b\xf1\x60\x58\x57\xed\xff\x54\xed\xb8\x9c\x46\xbe\xbc\x08\x8d\x7a\xf8\xde\xdd\x09\x78\x01\xe4\xcf\xc5\xb2\x45\xf6\x11\x00\x00\xff\xff\x8c\x09\x5e\xaa\x14\x10\x00\x00")

func ReleaseLinuxConfContractsFirewallAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsFirewallAbiJson,
		"../../release/linux/conf/contracts/fireWall.abi.json",
	)
}

func ReleaseLinuxConfContractsFirewallAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsFirewallAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/fireWall.abi.json", size: 4116, mode: os.FileMode(436), modTime: time.Unix(1596005127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ReleaseLinuxConfContractsGroupmanagerCppAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\xb1\x4e\xc3\x30\x10\xdd\xfb\x15\x96\xe7\x4c\x80\x18\xba\x51\x55\x42\x65\x28\x48\x8c\x55\x07\x2b\xbe\x04\x4b\xc9\xd9\xf2\x9d\x87\x0a\xf5\xdf\x91\x8d\x1a\x52\x01\xa5\x60\x20\x15\x24\x43\xa4\x24\xe7\xe7\x7b\xef\xdd\x8b\x57\x13\x21\x84\x78\x4c\xf7\x78\x49\x54\x2d\xc8\xa9\x90\xa5\x07\xc5\x70\xed\x6d\x70\xb2\x78\xf9\x6c\xd0\x05\x26\x39\x15\xab\xee\xdd\x3e\xc0\x2b\xa0\x3a\x42\x2c\xb0\xb2\x3d\x98\xae\x88\x37\x2e\x15\x11\x7b\x83\xb5\xdc\x2b\xd8\x76\x4f\xeb\x5e\x03\x36\xf0\x67\x3b\x38\xb4\xb1\x41\x3e\x3f\x3b\x66\xdf\xd2\x22\xb1\x42\x8e\x8b\x2a\xd5\x10\xf4\x55\xd9\xa1\x55\x01\x4b\x36\x16\x9f\x01\xb7\xc5\x7b\xe2\xd6\xc0\x57\x4d\x93\xc4\xa5\xb7\xd5\xfd\x41\xca\xc7\x6b\xdd\xe7\xcc\x3e\x64\x51\x7e\x50\x94\xf8\xde\xba\x3b\xf0\xad\x21\x8a\x6b\x7e\x9b\xfa\x97\xdc\xce\x65\x8e\x56\x03\x2d\x43\x9b\x17\xa3\x88\x72\x43\x16\xef\xd9\xff\x9d\x20\xe5\x4a\x5b\x03\xa7\xa1\x9a\x6d\x16\xf3\xef\xf8\x4b\xcd\x0f\x31\x0c\x06\xf9\xf2\x62\x00\x69\x87\x09\x6c\x70\x5a\x31\xcc\xac\xe5\x65\x1c\xe0\x61\xe4\x2d\xfe\x75\x3c\xb2\xcf\x19\xa5\xf5\xce\xc0\xd1\xbf\x0f\x7b\x38\x3d\xff\x34\x34\xa3\x7f\xa7\xef\xdf\x44\xac\x9f\x02\x00\x00\xff\xff\x13\xb9\x10\x95\x48\x0b\x00\x00")

func ReleaseLinuxConfContractsGroupmanagerCppAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsGroupmanagerCppAbiJson,
		"../../release/linux/conf/contracts/groupManager.cpp.abi.json",
	)
}

func ReleaseLinuxConfContractsGroupmanagerCppAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsGroupmanagerCppAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/groupManager.cpp.abi.json", size: 2888, mode: os.FileMode(436), modTime: time.Unix(1596005127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ReleaseLinuxConfContractsNodemanagerCppAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x4f\x4b\xc3\x30\x18\xc6\xef\xfb\x14\x21\xe7\x9e\x54\x3c\xec\x26\xe8\x65\x42\x2f\x1e\xc7\x0e\xb1\x79\x5b\x02\xe9\x9b\xd2\x3c\x19\x14\xd9\x77\x97\x54\x5b\x3b\xd4\xb1\x5a\x65\xd0\xda\x43\x4b\xc3\xfb\x27\xcf\x2f\x79\x92\xed\x4a\x08\x21\x5e\xda\x77\x7c\x24\xab\x92\xe4\x5a\x48\xa5\xb5\x4c\x3e\x86\x0d\x57\x01\x5e\xae\xc5\xb6\x1f\x3b\x4e\xfc\x54\x80\x9d\xa6\x8d\x77\xfc\x84\x7a\x50\xa8\x0f\x43\x53\xb5\x61\x1e\xb5\xe1\x42\x1e\x05\x1c\xfa\xbf\xdd\x60\x0a\x2e\x60\xec\x1c\x4e\x35\x36\x8c\xeb\xab\x73\xfa\x66\x8e\x3d\x14\x23\x26\xe5\xca\x7a\x1a\x72\xe9\xaa\xe5\x81\x33\x18\xc7\x6f\x05\x0f\xc9\x77\x58\x0b\xc2\x9d\xb5\xa9\xd3\xe4\xbf\xc6\xfb\x87\x8a\xcf\x47\x3d\x94\x8c\x3a\x4c\x52\xbc\x57\xd6\xe8\x8d\x33\x1c\x45\x4f\xdb\x52\x55\x78\xb6\x26\x7b\xa4\x66\x3e\x1b\x6a\x2a\xdd\xe8\x32\x9f\x86\xf2\xdf\xab\xbf\x8e\xb6\x20\x9c\xf0\xe9\x5c\xd0\x5e\xe6\x50\x08\x95\x56\x98\x78\x1a\xb4\xdf\xb1\xca\x92\xb9\xac\xdb\xc5\xae\xaf\xd4\xd5\xa5\xb2\x0f\x91\xcf\x92\xae\xb1\x82\x70\x4f\x96\x40\x7a\x71\xd2\x53\x07\x93\x37\x3f\x33\x6b\xd7\x2b\x18\xc6\xed\xcd\x58\x33\x8e\xd7\xdd\x65\xd0\x9e\x18\xef\xc2\x56\xbb\xd7\x00\x00\x00\xff\xff\x67\x87\x5c\x95\xe6\x0a\x00\x00")

func ReleaseLinuxConfContractsNodemanagerCppAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsNodemanagerCppAbiJson,
		"../../release/linux/conf/contracts/nodeManager.cpp.abi.json",
	)
}

func ReleaseLinuxConfContractsNodemanagerCppAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsNodemanagerCppAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/nodeManager.cpp.abi.json", size: 2790, mode: os.FileMode(436), modTime: time.Unix(1596005127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ReleaseLinuxConfContractsParammanagerCppAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\xc1\x6e\xea\x30\x10\xbc\xe7\x2b\xac\x9c\x39\xbd\xf7\xf4\x0e\xdc\x52\x5a\x21\xa4\xaa\xe2\x40\x4f\x88\x43\x6a\x0c\xb5\x48\xec\x28\xbb\xa1\x44\x15\xff\x5e\x85\x10\x08\xc4\x21\x2e\x0e\xd4\x0d\x07\x24\x8c\x19\xcf\xcc\x0e\xd9\xf5\xd4\x21\x84\x90\xcf\xdd\x7b\xf6\x72\x85\x1f\x32\xb7\x4f\x5c\x60\x38\xf4\x61\x20\x05\xc6\x3e\xc5\x97\x6c\xb5\x77\xdc\xc5\x45\x94\x20\xb8\x7d\x32\x3d\xac\x9d\xe2\x54\xf0\xa8\x1a\xe9\xb0\x0f\xd3\x28\x3f\x17\x63\x2e\x96\xee\xc9\x86\xed\xe1\xd3\xac\xc4\x41\x26\x58\x90\x28\x2f\x53\x29\x00\x7d\x81\x19\xd8\xc2\x0f\xe0\x84\x77\x71\xca\x22\x11\x14\xb9\x14\xf9\x39\xdb\x5e\x9d\x0b\x4b\x4d\x17\x6a\x78\xe9\x9a\xd3\x8e\x21\x65\xe5\x18\x27\x46\xc2\x81\xe1\x08\xc6\xb1\x9c\x27\x94\x3d\x85\x11\xa6\x0f\x81\xa4\x2b\xb3\x08\xf0\x8b\x80\x15\xaa\x09\x17\xf8\xf7\x8f\x45\x49\xd0\x36\xe4\x86\x69\xd0\x37\xa5\xe5\x34\x4c\x36\x43\x1f\x9e\x79\xc8\xd1\x2c\x04\xa8\xc2\x51\xca\xfc\xff\xcf\xa2\xda\x37\xc9\xbf\x71\xc9\xf5\xbc\x68\xb9\xe4\xbb\x84\xb7\x53\xf5\xb7\x1a\x28\x03\xb1\x77\x2a\xbc\x86\x09\x5d\xac\xbd\x17\x04\xf2\xc3\x13\xa9\x47\xa9\x4c\x04\x3e\xb2\x28\x90\x69\xd1\x07\x4d\xdb\x80\x2e\xb8\xd2\x0e\x5b\xb2\x01\x0c\x07\xef\x8c\xae\x0a\xde\xb9\x8a\x31\x8b\x43\x0e\x90\xfd\xde\x6c\x5c\xca\xa0\x95\x60\x4a\x53\xac\xea\x93\x57\xd8\xd2\xb9\x8e\xb9\xbc\xea\x2f\xd4\x39\x1b\x76\x63\xa4\x17\x45\xb1\x5c\xb3\xdc\x00\x36\x6f\xed\x29\xd2\x08\x6b\x60\xc2\xdd\x46\xca\x6f\x99\xd3\xcd\x7c\x4c\x36\xaf\xc0\x86\x3e\x98\x06\x42\x81\xa3\x94\x69\xd5\x7c\xd1\x24\xff\x97\x0f\x17\x84\x90\x73\xd9\x6e\xe5\x46\x4d\xaa\xc2\xcf\x04\xee\x4f\x29\x52\xba\xed\xa9\xbf\xde\xdf\x95\xd5\x22\xf6\x7b\xd8\x9a\x15\x43\xc4\x25\x92\xca\xcb\x9e\x8d\x44\xcb\x37\x13\x1b\xf9\x9d\x0d\xd0\x36\x52\x1c\x41\xc3\xc8\x62\x27\xe9\xda\xd6\x61\x27\xdd\xe3\x73\xee\xc7\xf8\x11\x67\xe6\x38\x5f\x01\x00\x00\xff\xff\x6c\x94\xa7\x29\xf3\x14\x00\x00")

func ReleaseLinuxConfContractsParammanagerCppAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsParammanagerCppAbiJson,
		"../../release/linux/conf/contracts/paramManager.cpp.abi.json",
	)
}

func ReleaseLinuxConfContractsParammanagerCppAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsParammanagerCppAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/paramManager.cpp.abi.json", size: 5363, mode: os.FileMode(436), modTime: time.Unix(1596006921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ReleaseLinuxConfContractsUsermanagerCppAbiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\x4d\x8b\xdb\x30\x10\x86\xef\xfb\x2b\x84\xcf\x3e\xb5\xb7\xbd\xa5\xbb\x50\x0a\x25\x85\x96\x9e\x96\x3d\x88\x68\x92\x35\x78\x25\x23\x8d\x0b\xa1\xe4\xbf\x17\xe7\xc3\xf1\x87\x64\x3b\x95\x9d\xc8\xd9\xd9\xc3\x1e\x64\x21\xbf\xf3\x68\xfc\x4a\x91\xe6\xe5\x81\x31\xc6\xfe\xee\xff\x17\x7f\x91\xe4\xef\x10\x3d\xb2\xc8\x00\xfe\xca\x33\xd0\x0b\xf1\x9e\xc8\x28\x3e\x77\x48\x64\x96\xa3\x89\x1e\xd9\xcb\x6b\xa5\x55\xe5\x68\x6b\x5e\x29\x69\x90\x4b\x2c\x46\x5c\xf3\xd4\x40\x75\x24\xdc\x66\xfb\x57\xad\x73\xb9\xc2\x44\xc9\x68\xff\x68\x17\xbb\x24\xa1\xe6\xd2\xac\x41\x9f\x75\x7d\xd9\x2e\x84\xd0\x60\x8c\x5d\x60\xd9\x56\x1f\xb0\x35\x30\x17\x42\x57\x46\x68\x09\x34\xa8\x13\xb9\x89\xca\x07\xb5\x9e\xbb\x6a\xc3\xb5\x91\x6c\x00\x0b\x02\xdf\x13\x83\x3f\xd6\x3f\x55\x0a\x7e\x20\x90\xeb\x0d\x60\x63\x1c\x17\x8e\x3a\x85\x3e\x04\x43\x25\x8c\xf3\xe2\x2a\x64\xd4\xb9\x2f\xe3\x82\x88\xb9\x6e\xb2\xdd\x1b\xdd\xe3\x47\xd2\xc9\x77\x59\xb4\x78\xc1\x95\xf5\x11\x3e\x0c\x5c\x17\x5b\x2e\xc4\xd3\x1b\x4f\xe4\x0d\xdc\x72\x2e\x8c\x07\x7b\xf0\x50\xc8\xfe\x59\x4c\x84\x1b\x9a\x04\xa4\x94\xc6\x63\x41\x1e\x4a\x99\xdc\x78\x0a\xa7\xf8\xaa\x55\x9e\x11\xe1\x2b\x11\x26\xa7\xf0\x81\xdc\xe5\x14\x44\x79\x24\xca\x03\x19\x93\x57\x8c\x0d\x98\x0b\xb1\x54\x02\x88\xef\x84\x5e\x5c\x01\x4c\x26\xe1\xc3\xb8\xcb\x8a\x09\xf2\x38\x90\x87\x21\x26\xa3\x98\xc0\x88\x9f\x94\x44\xcd\x57\x48\x69\x3c\xa5\x57\xb4\x41\x53\x32\x7b\x52\xee\xf0\x0c\x42\x3d\x1a\xea\x4b\x38\x93\x77\x4c\x84\xba\xe2\x1e\xcf\x90\xa5\x6a\x0b\x9a\x68\xfb\xd1\xbe\x08\x35\x9d\x2b\x4f\x73\xe4\x49\xa0\xc7\x03\x7d\x11\x67\xf2\x0e\x0f\xd4\x1d\xd6\xf1\xdb\x80\xf6\x83\x9a\x1b\xd0\xdf\xe4\x5a\x11\xd8\x0a\x93\x4c\x70\x84\x82\xed\x33\x98\x55\x83\xce\x7f\x26\x6e\x3d\xfb\x87\x45\x1a\x0f\x7d\x81\x68\xcb\xf4\x60\x79\x1f\x93\xb8\x01\x5c\xa4\x69\x31\x89\x0e\xdb\x09\x2e\xe4\x11\xaa\x57\x8a\x70\x7b\xdc\xd6\x1d\x5e\x4f\xbe\x3a\xa3\xdb\xdd\x2f\x49\xf7\xfe\x80\x30\xf6\x63\x7c\xe3\xc6\x5d\xa8\x36\x3e\x40\x76\x81\x67\x46\x5a\xa5\xb0\xec\xfc\x95\x7d\x3b\xc7\x4c\x24\x7e\xfe\x74\xdd\xc9\x6a\x16\x81\x36\xa4\xb5\x27\xb0\x11\xe8\x89\x59\x7e\xd0\xde\x9c\x86\x06\x52\x7b\x34\xc7\x3e\xf0\x07\x24\xf6\xa8\xed\xaa\x0f\x65\x2c\x9e\x85\xde\x7d\xf6\x15\x9d\xe2\xf0\xf8\x3a\x8a\xc9\x58\x1c\x68\x3e\xd8\xea\xb2\x18\x3b\x67\x71\x80\x7a\x6d\xb7\xea\xa5\xe4\xc0\xf5\x9e\x52\x37\x0e\x98\xaf\xe5\xaa\x6c\x2e\x7a\x4b\x67\x08\xfa\x7b\xb3\x9f\xc5\x1e\x34\x87\xaf\xf7\xc0\x78\x06\x7c\xdb\x87\x28\x01\xae\x6f\xee\x53\x4c\x16\x87\xc8\xd7\x51\x65\x5a\x7e\x71\x81\xeb\xad\x6e\x1d\x42\xd5\x6b\x5f\xdf\xc2\xe5\x6b\x5b\xdf\x42\xe6\x3b\xab\xf5\xcd\x52\x57\xc1\x02\xe7\xeb\x5c\xdf\x42\xdc\xaf\x3b\xee\xa0\x4b\xc4\x01\xeb\x9d\xc5\xfa\xe6\xbe\x3c\x0a\x33\x1f\x4e\x57\x03\x2c\xa0\x4d\x7a\x4d\xc1\x51\xa7\xed\xa4\xfd\x66\x3a\x1f\x5e\xff\x05\x00\x00\xff\xff\x9a\x9c\x7e\xa1\x23\x3b\x00\x00")

func ReleaseLinuxConfContractsUsermanagerCppAbiJsonBytes() ([]byte, error) {
	return bindataRead(
		_ReleaseLinuxConfContractsUsermanagerCppAbiJson,
		"../../release/linux/conf/contracts/userManager.cpp.abi.json",
	)
}

func ReleaseLinuxConfContractsUsermanagerCppAbiJson() (*asset, error) {
	bytes, err := ReleaseLinuxConfContractsUsermanagerCppAbiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../release/linux/conf/contracts/userManager.cpp.abi.json", size: 15139, mode: os.FileMode(436), modTime: time.Unix(1596005127, 0)}
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
	"../../release/linux/conf/contracts/cnsManager.cpp.abi.json": ReleaseLinuxConfContractsCnsmanagerCppAbiJson,
	"../../release/linux/conf/contracts/fireWall.abi.json": ReleaseLinuxConfContractsFirewallAbiJson,
	"../../release/linux/conf/contracts/groupManager.cpp.abi.json": ReleaseLinuxConfContractsGroupmanagerCppAbiJson,
	"../../release/linux/conf/contracts/nodeManager.cpp.abi.json": ReleaseLinuxConfContractsNodemanagerCppAbiJson,
	"../../release/linux/conf/contracts/paramManager.cpp.abi.json": ReleaseLinuxConfContractsParammanagerCppAbiJson,
	"../../release/linux/conf/contracts/userManager.cpp.abi.json": ReleaseLinuxConfContractsUsermanagerCppAbiJson,
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
			"release": &bintree{nil, map[string]*bintree{
				"linux": &bintree{nil, map[string]*bintree{
					"conf": &bintree{nil, map[string]*bintree{
						"contracts": &bintree{nil, map[string]*bintree{
							"cnsManager.cpp.abi.json": &bintree{ReleaseLinuxConfContractsCnsmanagerCppAbiJson, map[string]*bintree{}},
							"fireWall.abi.json": &bintree{ReleaseLinuxConfContractsFirewallAbiJson, map[string]*bintree{}},
							"groupManager.cpp.abi.json": &bintree{ReleaseLinuxConfContractsGroupmanagerCppAbiJson, map[string]*bintree{}},
							"nodeManager.cpp.abi.json": &bintree{ReleaseLinuxConfContractsNodemanagerCppAbiJson, map[string]*bintree{}},
							"paramManager.cpp.abi.json": &bintree{ReleaseLinuxConfContractsParammanagerCppAbiJson, map[string]*bintree{}},
							"userManager.cpp.abi.json": &bintree{ReleaseLinuxConfContractsUsermanagerCppAbiJson, map[string]*bintree{}},
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

