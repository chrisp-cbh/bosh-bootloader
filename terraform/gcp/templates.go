// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x95\xd1\x6e\xa3\x30\x10\x45\x9f\xc3\x57\x8c\xac\x7d\x6c\x48\xca\x42\x12\x55\xea\x97\xac\x22\x64\xc0\xa5\x6c\x1d\x6c\x19\x93\x54\xaa\xf8\xf7\x95\xb1\x0d\x24\x10\xe4\x64\x1b\x55\x7d\x08\x85\x3b\x97\x99\xc3\x1d\x60\xb5\xe4\xb5\x04\x54\x12\x79\x62\xe2\x23\x2e\xf1\x81\x20\xf8\xf2\x00\x00\x8e\x98\xd6\x04\x5e\x01\xfd\xfa\xca\x19\xcb\x29\x89\x53\x76\xe0\xb5\x24\xb1\x51\xfb\x49\x42\x97\xf6\x58\x55\x36\xc8\x6b\x3c\xcf\x7a\x56\x75\x72\x9b\x6d\x5f\xd0\x3a\xeb\x7f\x27\x8c\x13\x56\xbd\xc7\x8c\x93\x32\x96\x38\x77\xf4\x7e\x2b\x04\x39\x61\x4a\x7d\x55\xbc\x54\xc5\xd7\x8c\xb3\x42\x90\x54\x32\x71\x66\xbe\x70\x75\xb6\xd5\x13\xee\x7f\xeb\x03\x4f\xd8\xe7\x55\xdf\x23\x16\x3e\x29\x8f\x71\x91\x35\x4b\xa3\x3d\xab\x2f\x4a\x49\x44\x89\xe9\x3d\x53\xdb\xda\x89\xb6\xba\x79\x71\x96\x09\x52\x55\xe7\x6d\xbd\x4b\xc9\xab\x97\xd5\x6a\x64\x6d\xd4\x7a\x6a\xf2\xa9\xed\x97\x05\xf7\xcd\x85\xe6\x25\x88\xa2\x28\x6a\x6f\x25\x48\xc5\x6a\x91\x12\x40\xd3\x41\x42\x80\x06\x51\xd2\x63\xa9\x46\x17\x8b\x31\x19\x2b\x9a\xf5\xed\x93\x64\xac\xf5\x89\xa1\xf3\x62\xe4\x6c\x34\x1e\x40\xc1\xe3\xb4\xc8\x44\x2c\x70\x99\xb7\x14\x9e\xd7\x7e\xfb\xb7\x7a\xde\xa8\xeb\xc6\xdb\x78\x38\x2c\x47\x45\xe8\x5b\x4c\x8b\xf2\xa3\x99\xef\xbb\x7b\x04\xe8\x12\x6b\xdf\xfa\x18\xc9\x48\x3b\x7b\x0f\x9b\x09\x04\xc8\xd6\x0c\xcc\x01\xc6\xfe\x9d\xac\x1f\xfd\xd6\xd7\x82\x07\xa0\xdb\xd1\x4c\x2b\x78\x85\x3f\xc8\x42\x5d\xa3\xbd\x12\x60\x4a\xd9\xc9\x44\x9a\x33\x21\xb5\x28\x08\xd0\x13\xa0\xcd\x6e\xb3\x53\xbf\x3a\x53\x7b\xad\x11\x4c\xb2\x94\x51\xd5\x8b\x4c\xb9\xea\xae\x51\x3e\x12\x8b\x9c\x48\xb5\x25\xda\x61\x02\x96\xda\x7f\xb4\x77\xc5\xd4\x97\xcc\x73\xea\x75\xdf\x01\xca\xa1\x7f\x47\x68\xbb\x30\xfc\xdd\xfe\xee\xc2\xf0\x1b\x21\xda\x17\xc7\x8d\x20\xbb\x32\x07\x98\x9d\xf6\xd1\x40\x07\xb3\x5c\x42\xbd\x0b\x90\x7d\xdf\xba\xb3\xb1\x15\x4b\xc9\x5c\x11\x4d\x96\x3c\x90\xd4\x60\xa8\xab\xc9\x0b\x03\x9d\xbd\x20\x0a\xa2\xb5\x3e\xd8\x6e\xb7\x3f\x11\x36\xf3\x09\x55\x70\xda\x13\xb3\x28\x2f\xc4\x0f\x84\x68\xbf\xec\xf3\xdb\xfb\x3f\xbc\xba\xc7\xf4\x34\xbf\x53\x37\x47\xd3\x31\x8e\x3f\x14\xc1\x01\xab\x22\x3d\xf4\xb0\x5c\x96\xf9\x9a\xa6\xce\xee\x5a\xf8\x7f\x01\x00\x00\xff\xff\x02\xc3\x4b\xc8\x56\x0b\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 2902, mode: os.FileMode(420), modTime: time.Unix(1506965001, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcf\x8a\xdb\x30\x10\x87\xef\x7e\x8a\x41\xf4\x54\xb0\x59\xe8\x39\x87\x42\xcf\xbd\xf4\x58\x16\xa3\x58\x13\x47\x20\x6b\xc4\xcc\xd8\xd9\x74\xf1\xbb\x17\x39\x71\x36\x69\x6b\xba\x3e\x04\x72\xd8\x5c\xf2\x47\x3f\xcd\x7c\xf3\x29\x18\x0d\x96\xbd\xdd\x06\x04\x23\x47\x51\xec\x6a\x47\x9d\xf5\xd1\xc0\x6b\x01\xa0\xc7\x84\xb0\x01\x23\xca\x3e\xb6\xa6\x18\x8b\x82\x51\xa8\xe7\x06\xc1\xb4\x44\x6d\xc0\xda\x45\xa9\x3b\x1b\x6d\x8b\xae\xfe\x45\x11\x0d\x18\x8c\xc3\xf4\xf3\xe9\x6b\x2e\x14\x6d\x87\x70\x7e\x6d\xc0\x7c\x7a\x1d\x2c\x57\x39\xe6\xdd\x58\x4e\xb1\x02\x20\x6f\x99\x83\x97\xd0\x0d\xd5\x58\x4d\x39\x94\x86\x7d\x52\x4f\x31\xe7\xbe\x7d\xff\x01\xb9\x04\xec\x88\x41\xf7\x08\x37\xd5\x01\xe3\xe0\x99\x62\x87\x51\xa7\x01\xa8\xd7\xd4\xeb\x1f\xe3\x4e\xb8\x82\x3c\x20\xcb\x89\x78\xb0\xa1\xc7\x13\xc6\xc2\xa0\xd5\xf5\x98\x55\x06\x9f\x2b\x8c\xcb\xa6\x18\x1b\x62\x57\x0b\xaa\x01\x73\xf0\xc1\x35\x96\x5d\xe9\xa2\xfc\xe5\x69\x03\xe6\x73\xf5\xce\xe6\xb3\xb9\xf1\xa4\x27\x61\x74\x52\x4f\x76\x7e\xce\xcd\x1b\xea\x52\xaf\x58\xb7\x81\xb6\x36\xd4\xd6\x39\x46\x91\xaa\xd9\x95\xe7\x8f\xe6\x79\x3e\xf0\x4b\xff\xaf\xb9\x9c\x6a\x78\x3b\xb9\x2f\x4f\x4f\x45\x01\x70\x4d\xb2\xd2\xd1\x68\x72\x01\x66\x67\xd5\xca\x04\x78\xd9\xfc\x5f\xc4\xea\xfc\x3e\x9a\xe7\xf7\x09\xde\x92\xec\x97\xe4\xe6\xb5\x3b\xf8\x9d\x51\xa7\xd6\xf8\xa2\xc8\xd1\x86\xd2\xa7\xc7\xd1\xbb\x44\xb8\xda\x6e\xb3\x2b\x45\xf6\x65\x62\x7a\x39\xfe\xcb\xb0\xdc\x55\xf0\x4d\xf7\x87\x93\x7b\x4d\xb7\x5a\xac\x36\x69\xe9\x5f\xab\x4d\xba\xaf\xd3\xdc\x9b\xa9\x57\xe4\x87\x94\xfa\x86\xb7\xda\xaa\xa3\x94\x02\xf2\x92\xd9\xf3\xf2\x7d\xed\x1e\x1e\xe8\x31\x7b\x83\xb5\xda\x66\xa0\xb6\x65\x6c\xad\xd2\xa2\xd1\xab\xc8\x87\xd5\x95\x37\x82\x83\x2c\x5f\x0a\x0e\xf2\xa1\xb3\x18\x8b\xdf\x01\x00\x00\xff\xff\xc6\x89\x38\x3c\xba\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2746, mode: os.FileMode(420), modTime: time.Unix(1506702731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4f\x6f\xe3\xb6\x13\x3d\xff\xfc\x29\x08\xe2\x77\x68\x81\x95\x63\x3b\xe9\xd6\x3d\xec\xa9\xe8\x75\xdb\x43\x6f\x8b\x80\xa0\xa8\x91\x4d\x98\x11\x59\x92\xb2\xd7\x08\xf2\xdd\x0b\x92\x52\x2c\x59\x12\x25\x39\x59\x14\xd9\x3d\x58\x11\x67\xde\x50\x6f\xfe\xe8\x89\x47\xaa\x39\x4d\x05\x20\x6c\x8c\x20\x0c\xb4\xe5\x39\x67\xd4\x02\x46\xcf\x0b\x84\xec\x59\x01\xfa\x82\xb0\xb1\x9a\x17\x3b\xbc\x78\x59\x2c\x06\x3d\x88\xd2\xfc\xe8\x7e\x0f\x70\x1e\xf4\x96\xa5\x55\xa5\x45\x58\xcb\xd2\x82\x26\x29\x65\x07\x28\x32\x62\x40\x1f\x39\xab\x82\x1e\xa9\x28\xbd\xdf\xff\x9f\x77\x52\xee\x04\x10\x26\x9f\x54\x69\xe1\xda\x7c\x19\x50\x12\x91\x26\xd5\x4a\x52\xaf\x14\xf4\x09\x5e\xfa\x22\x8a\x94\x70\x15\xe2\xc4\x22\xed\x84\x4c\xa9\x20\x34\xcb\x34\x18\xb3\x64\x79\x52\x5f\x56\xbf\x6d\x70\x63\xf6\x44\x69\xf9\xfd\x3c\x15\xbf\x01\x6c\xcc\x3e\xf1\xbe\xfd\xd0\x96\x29\x32\x6f\xef\x0d\x6c\xcb\x54\x12\x9c\xfb\xc1\x4f\xe6\x06\xd0\xd3\x15\x09\x1a\x8c\x2c\x35\x03\x84\xaf\x7c\x72\xae\xe1\x44\x85\xc0\x08\xd7\x97\x09\xcb\x43\x2c\x97\x20\x14\xfe\xf9\x70\x47\xaa\x97\x50\x1c\x09\xcf\x5e\x12\x96\x27\x52\x41\x81\x17\x08\x65\xa0\xa0\xc8\x0c\x91\x05\xfa\x82\xbe\x5d\x07\x28\xc0\x9e\xa4\x3e\x2c\xd3\x54\x24\xd5\x35\x7e\x74\xe0\xe1\xfa\x15\x7c\xdc\xad\x2e\x98\x05\x42\x54\x08\x79\xaa\xf8\x50\x5a\x5a\xc9\xa4\x70\x30\x96\x29\x1c\x6e\x4a\x6d\x4d\xc0\xfe\x86\xb7\x2b\xfc\x09\xe1\x87\x87\x7b\x1f\xf8\xc5\x01\x04\x36\x88\xa6\xc5\x0e\x8c\x37\x5a\x2d\xfd\xff\xbb\x15\x7e\x74\x06\x96\xea\x1d\x58\x62\xe9\x2e\x2c\xbf\xb9\xce\x1f\xa3\x69\x68\xd7\x32\x46\xf8\x52\xcd\x8d\x5c\xf4\x64\x21\x9e\xdd\x0a\x36\x97\xfa\x44\x75\xc6\x8b\x1d\xd1\xa5\x80\x00\xbf\xb7\x56\x25\x97\x95\x24\xac\x4c\xc8\xbb\x73\x74\x2c\x73\x55\xef\xf7\xe6\xf6\xac\x79\x46\x43\x65\x50\xa5\xc1\x85\x0c\xcd\xbb\xac\x77\x2e\xd2\xaa\x23\x0d\x88\x9c\x08\x5e\x1c\x3c\x9e\x4b\x7c\x48\xab\xc3\xdb\xae\xde\xc6\x8f\xb9\x99\x20\xf3\x1f\x30\x64\xda\x14\x99\x69\x1c\xb9\xbe\x88\x92\xd4\xc9\x41\xa3\x7e\xea\x08\x1d\x5e\xba\xc4\x78\xfb\x60\xec\x87\x86\x61\x9a\x2b\xcb\xfd\xd4\xc0\x1a\xa8\x10\x67\x44\x91\x90\x34\x43\x29\x15\xb4\x60\xa0\x51\x5a\x5a\x24\xb8\xb1\x90\x21\x6a\x10\x2d\x90\x03\x41\xaf\x20\xa5\x16\xe4\x89\xaa\x41\x6e\xaa\xf5\x16\x21\xa5\x16\x89\xbb\xd7\xa4\x64\xe2\xd3\x9b\xeb\xc7\x37\x91\xe7\x1f\x26\xc1\xf4\xb3\x50\x3b\xcc\xa1\xc2\xf4\x73\xf1\x66\x42\x10\xba\x12\x0e\x03\x43\xf0\xca\xca\xe1\xba\x3f\x9b\x58\xf1\xb9\xd7\x51\x34\xb8\x82\xb8\x10\x4a\x94\x86\x9c\x7f\xef\x70\xd9\x53\x45\xa5\x01\xed\x18\x39\xf2\x0c\x32\xf7\x08\xa8\xd2\x3b\xe8\x00\x67\x74\xe7\xef\x34\xa2\x21\x45\xb9\xf6\x0d\x71\x51\x45\x21\x4c\xce\x05\xfc\xe4\x62\x45\xf4\xd3\xcf\x7e\x07\x4d\xb8\xa8\x6b\x30\x17\x3c\x07\x76\x66\x02\xd0\xf3\xe2\x7f\x4c\x83\xc3\x4a\x21\x97\x1a\x48\x06\xc6\x6a\xe9\x36\x60\x75\x09\xfe\x45\x15\x63\xae\x4a\xe5\x55\x31\x56\xc9\x8c\xbf\x33\xaa\x09\xee\xf9\xcb\x69\x29\x6c\xfd\x12\x7b\xa3\xa0\x9b\xda\x52\x7b\xa0\xc2\xee\x09\xdb\x03\x3b\x84\xfd\xab\x32\x15\x9c\x25\x61\x21\xa9\x16\x7a\x3b\x6a\x60\xe4\x06\x00\xff\x4c\x7e\x4e\x35\x43\x38\xaa\xc3\xd0\xeb\x22\x6d\x57\xdb\x95\x5b\xd5\xf0\x4f\x09\xc6\x12\x45\xed\xbe\x11\xe7\x2e\xe0\xe0\xd1\x6c\x74\x82\xbe\xcf\x73\xd5\xd3\x7a\x60\xe3\xe3\xfb\x9e\x28\xfd\x5c\x4d\xc4\xf6\xd8\x5b\x44\x4d\x87\x8f\x21\x03\x83\x10\xdc\xae\x62\x3a\x70\x7d\xbf\x5a\x6e\xd6\x6b\xaf\x05\x37\x1b\x67\x7f\xff\xcb\x72\xfd\x5b\xb8\xb1\xfe\xec\x5d\x9b\xe2\x10\xbd\xa3\x3c\xec\x7e\xaa\x54\x91\x94\x94\x62\xec\xc3\xab\x61\xda\xfe\x64\xb9\x7c\x67\x0d\x96\x42\x4b\x77\xbe\x7a\x8e\x4c\x91\x8b\xdd\x8c\x32\xeb\x03\x1f\xae\xb1\x57\xeb\x8f\xf3\xb1\xb1\xd9\x6c\x36\x97\xfa\x1a\xfd\x8c\x18\xc9\x5a\xfc\xed\xd9\xaa\x8e\x1b\x53\xe7\x9a\x00\x8c\xe1\xb2\x20\x34\xcf\x79\xc1\xad\x7f\x07\x7e\xfd\xf3\xeb\x1f\x23\x79\xed\x13\xcd\xc3\xe9\x1d\xdb\x47\x4b\xe8\xce\x2b\xf0\x41\x75\xeb\x60\x7c\x3e\x82\x16\x6f\x26\xef\xef\xdf\xff\xba\x52\xe8\xef\x77\x06\x70\x7b\xd3\x36\xce\x02\x26\x74\x6d\xbb\xb3\x2e\xbe\x93\x5a\xab\x61\xfe\x11\xda\x6a\xbd\xda\x3c\x24\xf7\x9b\x5f\x3f\x6f\x6f\x6f\xae\x0e\xbb\xf1\xee\x6a\x0d\xc5\x5e\x76\xc7\x78\xbd\x41\x31\x44\xb2\x18\x15\x40\xed\x74\x0e\x69\x86\x37\x2a\x86\xce\xbc\xb9\x89\x95\xe8\xc4\x71\x02\xae\x41\x8a\x4f\xac\xaf\x86\x6e\x76\x3b\x0c\xf6\xe6\xf8\xd3\x02\xa1\x78\x9e\x7b\x07\x59\x2c\x0f\xe3\xfc\xcf\x1c\x65\x8d\x4d\x47\x67\x59\xa3\x09\xde\x63\xa2\x4d\x38\x79\xbc\x7d\x94\x9d\xcc\x6c\xe1\x71\x1a\x39\xeb\x72\x06\xf3\xea\x73\x12\xe2\xec\x7a\x9c\x58\x8a\x3d\xe2\x7f\xd2\xdc\xe9\xad\xc7\x93\xa9\x8e\x95\x26\x55\xe3\xab\xf5\xfc\x5a\x3c\x99\x78\x0d\xfa\xe3\xa2\x77\x28\xbe\xe9\x27\xd4\x11\x3a\x66\xb1\xf1\x03\xc8\xd8\xae\x7e\x04\x17\xff\x06\x00\x00\xff\xff\xcb\x4b\x2e\xce\xf1\x19\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 6641, mode: os.FileMode(420), modTime: time.Unix(1506702731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x3f\x8f\xdc\x20\x10\xc5\xeb\xf0\x29\x46\xa3\x94\xc1\x85\x73\x6d\xaa\x28\xed\x25\x45\xba\xe8\x84\x58\x7b\xec\x45\xc7\x31\x08\xf0\x5a\xd1\xc9\xdf\x3d\xc2\x66\xbd\xbe\xdc\x9f\x5d\x69\xb5\x95\x47\xe3\xc7\x63\xf8\x3d\xe0\x21\xf9\x21\x01\x36\xec\x1a\x1e\x42\x24\x95\x74\xe8\x29\x29\xcf\x6c\x11\x9e\xc5\xa7\x83\xb6\x03\xc1\x37\xc0\xcf\xcf\x3d\x73\x6f\x49\x35\xfc\xe4\x87\xf4\x42\x59\x2d\xb5\x9c\x6b\xa7\x9f\x68\x42\x31\x09\xf1\xda\xdd\xee\x94\xf1\xd9\x17\x00\xe0\x7d\x6b\xdd\xb6\x81\x62\xac\xd6\x85\xf2\xd8\x29\xdf\xc5\x3f\x50\xe4\x21\x34\x04\xf8\xdf\xfa\xce\x04\x1a\xb5\xb5\x08\x78\x2c\xe5\xea\xb5\x6c\x9f\xa7\xcc\x43\xcc\xdb\x1f\x74\xa8\xc8\x1d\x94\x69\xa7\x93\x4e\xb2\x27\x87\x59\x4a\x69\xe4\xf0\xf8\xe6\xa4\xe5\x5f\xb5\xdb\x59\x79\xac\x0b\x00\x01\xa0\xad\xe5\xb1\x9c\xd6\x07\x4e\xdc\xb0\xcd\x36\xa9\xf1\xb8\x34\x39\xa4\xb8\x8c\xf1\x07\xef\xee\xbe\xe2\x17\xc0\xba\xae\x6b\x7c\x10\x00\x53\xb6\x28\x94\x93\xee\xe3\x2c\x3a\x1d\xe3\xe1\x43\x04\x05\x14\x6e\xe8\xcb\xb5\xb7\x02\x78\xff\xf4\x1f\x03\x7e\x71\x4b\x70\x93\xfe\x85\xde\x02\x20\x52\x8c\x86\x9d\xd2\x5d\x67\x9c\x49\x7f\xb3\xfe\xfe\xe7\xfd\x8f\x33\xc9\x72\x18\x75\x68\x8d\xeb\x55\x18\x2c\x21\x60\x8c\x7b\x79\xea\xca\xa5\xbb\x4d\xf8\x4c\xca\x31\xee\x71\xe5\xbc\x51\x5f\x78\xdb\x23\xd9\x4e\x59\xe3\x1e\xa7\xec\x92\xf3\x54\x41\xbb\x9e\x66\x97\x39\x4a\x01\x60\xbc\xda\xc6\xff\xfb\xfb\xaf\xd2\x2d\x89\xbc\xbd\xe5\xd5\xaf\xe0\x15\xab\x7d\x4a\x3e\x5e\x45\x6b\x76\xb8\x19\xaf\xfc\x02\x6e\x8c\xeb\x5f\x00\x00\x00\xff\xff\x5c\xc2\x73\xd4\xf2\x04\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1266, mode: os.FileMode(420), modTime: time.Unix(1506702731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xce\xcd\x0a\xc3\x20\x0c\xc0\xf1\xbb\x4f\x11\x64\xd7\x7a\xe8\x71\xb0\x67\x11\x57\x43\xe9\x50\x23\xd1\x48\xa1\xf8\xee\x63\xac\xfb\x82\x9d\x7a\x0d\xf9\xff\x12\xc6\x42\xc2\x13\x82\x9e\x89\xe6\x80\x76\xa2\x98\xa5\xa2\x75\xde\x33\x96\xa2\x41\xdf\x24\xe6\x2b\xad\xc3\x92\x35\x6c\x0a\x20\xb9\x88\x70\x01\x7d\xda\x9a\x63\x83\xa9\xd9\xc5\xf7\xe1\x6b\x4b\x75\xa5\x48\x6a\x96\xfa\x8e\xad\x70\x78\xd6\x00\xcd\x05\xd9\x81\xff\x37\xcd\xc7\x32\xfb\xa8\x9f\xc7\xf1\xc7\xc5\xb5\x22\x27\x17\xec\xeb\xab\x83\xee\x03\xbd\x07\x00\x00\xff\xff\xa3\x5f\xd0\xdf\x04\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 260, mode: os.FileMode(420), modTime: time.Unix(1506702731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\x0a\xc3\x20\x10\x05\xd0\x75\x3c\x85\x0c\x59\xb4\x9b\xde\xa0\x67\x29\x36\x99\xca\x14\x99\x91\x89\x08\xad\x78\xf7\x22\x16\xcc\xa6\xa4\x4b\xf9\x4f\x3e\xf3\xb3\x53\x72\xf7\x80\x16\xa2\xca\x13\x97\x74\xa3\x15\x6c\x31\x53\x7a\x45\xb4\x57\x0b\x5b\x52\x62\x0f\xa6\x1a\x33\xac\xa2\x27\xe1\x63\xf7\x16\xc6\x63\x85\x9c\xff\x6a\x5d\x14\x57\xe4\x44\x2e\x6c\xbf\x70\x54\xc9\xb4\xa2\x5a\xf0\x22\x3e\xf4\xf2\xdd\xbf\xc6\xe7\xf2\xa0\x80\x27\x98\x4b\x76\x7a\xd9\x85\x15\xce\x15\xcc\xf4\x1d\xa2\xd3\x46\xc6\x32\x2d\xee\xb7\x8f\xb4\xbf\x6b\xab\xff\x04\x00\x00\xff\xff\x39\xda\x2a\x22\x4d\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 333, mode: os.FileMode(420), modTime: time.Unix(1506702731, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

