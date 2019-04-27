// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (5.583kB)
// override/templates/singleton/psql_upsert.go.tpl (1.317kB)
// override/templates_test/singleton/psql_main_test.go.tpl (4.979kB)
// override/templates_test/singleton/psql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.743kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xdd\x72\xdb\xb6\x12\xbe\x26\x9f\x62\xa3\x39\x13\x93\x67\x64\xfa\x5c\xfb\x8c\x2e\x62\x3b\x49\x33\x69\x1c\x35\x8e\x9b\x99\x66\x32\x1e\x88\x5c\x4a\x18\x43\x00\x03\x82\x76\x54\x96\xef\xde\x59\x00\x14\x49\xfd\xc4\x4a\xd2\xb4\xe9\x95\x45\x60\xb1\xfb\xed\xee\xb7\xd8\x85\xeb\xfa\x18\xfe\xc3\x04\x67\x25\x9c\x4e\x20\x79\x42\xbf\xb0\x4c\xde\xb2\x99\x40\x70\x7f\x92\x4b\xb6\xc4\xa6\x09\xad\x68\x99\x2e\x70\xc9\xdc\x36\x1d\xe8\x24\xe0\x0f\x48\xae\xba\x5d\x7b\x80\xe7\x90\x3c\xc9\xb2\xe7\x42\xcd\x98\x80\xe3\xa6\x09\x4f\x4e\xe0\xba\x28\x51\x9b\xe7\xc0\x8c\xc1\x65\x61\x4a\x60\x12\xb8\xa4\xb5\x31\x30\x99\x41\xa6\xd0\xae\x55\x45\xc6\x0c\x82\xd2\xc0\xe7\x52\x69\x04\x25\x21\x55\x32\x17\x3c\x35\x49\x98\x57\x32\x85\x48\xc1\x7f\xeb\xda\xe1\x4f\xae\x8b\x2b\x2e\xe7\x95\x60\xba\x69\xe2\xd6\x4a\x64\x41\x48\x65\x20\xb9\x54\xe7\x4a\x1a\xfc\x64\x9a\x26\x35\x9f\x48\x15\x7d\x24\x7e\x71\x0c\x75\x8d\x32\x23\x90\xde\xf2\x6b\x79\xee\xad\xc1\x4c\x29\x31\x5e\x1b\x3f\x57\xa2\x5a\xca\x12\xde\x7f\x28\x8d\xe6\x72\x3e\xf6\x07\xfc\xfa\xd8\x7b\xd3\x8a\xcd\x14\x17\x89\xff\x88\x01\xb5\x56\x1a\xea\x30\xd0\x68\x2a\x2d\x41\x25\x0e\xa9\x03\xda\x07\x69\xcf\x3d\x47\x73\x71\x16\xc5\x75\x8d\xa2\x44\x0b\x7c\x0c\xed\x86\x97\xf4\xfb\x32\x6b\x9a\xf1\x16\xf4\x2d\xd4\x9f\x07\x1b\x87\x4d\x18\xae\x03\x11\xba\x14\x52\x52\x7a\x69\xa4\x9f\x53\x26\x79\xba\x91\xd0\xe9\xb7\x65\x14\xac\xce\x92\xd6\x6c\x8c\x0e\x4e\xf1\xf4\x87\xcb\x71\x1d\x06\x3c\x27\x2f\xa8\x44\x7e\xb0\x04\xff\xdf\xe2\x7a\x34\x01\xc9\x05\x01\x0d\x0a\x0a\x7b\x64\x4d\xbe\xd3\xac\x78\xaa\x75\x84\x5a\xc7\x71\x18\x34\xbb\xc8\xb0\x27\xfb\xbb\x92\x0f\x55\xc9\xe5\x9c\xbe\xf1\x13\xa6\x95\x51\xfa\x4b\x0a\xbc\xa7\xba\xf8\x3a\x66\x4c\xb7\x43\x4e\x40\x5c\x78\x9f\x7a\x48\xbd\xc0\x6f\xd3\xa5\x13\xf7\x4b\xbd\x53\xbb\xd3\xf1\x37\xd1\x68\x07\xd9\xfb\xe4\x26\xdc\xff\x28\x55\xd6\xc9\xfb\x1e\xb4\xb8\x42\x1c\x44\x0a\x32\x95\x56\x4b\x94\x86\x19\xae\x24\xe4\x4a\xc3\x42\xdd\x83\x51\x50\x68\x55\xa0\x16\x2b\xa8\x4a\x1c\xfa\x6a\x2d\x0e\xdc\x3d\x94\x55\xff\x72\x52\xad\xfb\x0f\xcf\x41\xc1\xa4\x4b\xae\xef\x47\x76\xbf\x4c\x2e\xf1\x3e\x1a\xd5\x75\x32\xbd\x9d\xbb\xf6\x7f\x0a\x52\x41\x5d\x0f\x46\x02\x8a\xef\x1d\xcf\x30\xb3\x31\xaf\x6c\x78\x46\x96\x0d\x61\x40\xd3\x02\x65\x5e\x50\x2e\x47\x86\x2f\xb1\x34\x6c\x59\xdc\x38\xa9\x9b\x05\x8a\x02\xf5\x08\x12\x68\x9c\x74\x47\xea\x9f\x94\xba\x2d\x2d\x8d\x06\xf4\xcf\xd4\x19\xe6\x4a\xa3\xcb\x82\x15\x3a\xb8\x16\xb6\xa9\xdc\x79\x4b\x70\x2d\x5a\x1b\xfc\x30\x0c\xe4\xef\x17\x98\xb3\x4a\x18\x3b\x12\x7d\xac\x50\x73\x2c\x93\x4b\x25\x7f\x43\xad\xfc\xd6\x15\x12\x0f\x3c\x4b\x2e\xd4\xbd\xec\x78\xe2\x23\xfd\x8e\x9b\x85\x17\x1e\x83\x8a\xc3\x30\x38\x39\x81\xb3\x8a\x8b\x0c\x52\x96\x2e\x10\x6e\x71\x05\x5c\x1e\x0b\x2e\x11\xaa\xb9\xe0\x62\x05\xc7\xb0\x5c\x95\x1f\x05\xdc\x95\x50\xd0\xdf\x42\xab\x99\xc0\x65\x19\x06\xb3\x2a\x27\x30\xa5\xd1\x4b\x26\xe7\x02\xa9\x3b\x9c\x55\x79\x8e\x3a\x8a\x6d\x98\xb6\x28\x43\x4e\xce\xaa\x3c\x79\xa7\xb9\xc1\xb3\x95\xc1\xe8\xc8\x1c\x51\x6e\x80\xa8\xb9\x6b\x3b\xb7\xdb\xe1\xe6\x72\x42\xcb\x94\xdf\x9b\x31\xa4\x04\x42\x33\x39\xc7\x2d\x32\x0e\x14\x5e\x59\x5e\x46\xe9\x7e\x85\x9b\xa2\xa5\xd1\xa9\x92\x77\xc9\x0b\xa3\x58\x34\xa0\x73\xf2\x92\xcb\x2c\xde\x89\x61\x28\x77\xae\xc4\x5f\x0b\x63\x78\x3d\xec\x87\x31\x94\xfb\x1a\x18\xdb\x3a\x7b\x24\xfc\x8c\x2e\xe2\xd0\xe9\x04\x68\xd7\x6f\xc4\x61\xd0\x91\x64\x5a\xb5\x24\x99\x55\x79\x6c\xcb\x6c\x27\x65\x5d\x49\x9d\x13\x2d\x5f\x55\x26\x79\xf3\xb3\x4a\x6f\x49\x93\x25\xea\xd8\xf1\x35\x23\x43\x0f\x9f\x7f\x7f\x8b\xab\x0f\x07\x1b\xba\x96\xc2\x99\x0a\x83\x3b\xa6\x6d\x8d\xda\xfb\x27\xb4\x9c\x7e\xe4\x0d\x53\x00\xda\x71\x52\xa3\x21\x20\xc3\x90\xbf\xe8\x7d\x51\x65\x86\x41\xb0\x0f\x41\x7b\x47\x3e\x2c\xd2\x2f\xe0\xc3\xa4\x55\x65\xfa\x07\xba\x14\xd2\x67\x1c\x06\x81\xef\x6c\xa7\x93\x0d\xe6\x5e\xf7\xbe\xbe\x1d\xff\x54\xf3\x25\xd3\xab\x97\xb8\xea\x09\x53\x88\x77\xde\x13\x8f\x1f\x83\x40\xe9\x4b\x2e\xa6\x86\xf0\x3f\x1b\xf0\x87\xfb\x41\x25\xed\x2b\xd0\x28\x7f\xf3\x6f\x76\x07\x6a\x58\x95\xc8\xec\xfd\x3c\xb3\x17\x9f\xf7\x3f\xb5\xb0\x40\xf0\xd2\x76\x0b\xdb\x2e\x82\xf6\x3e\xa1\xe8\x6c\xdc\x2d\x0e\x39\xa1\x6c\x37\xfa\x38\xd7\x07\x27\xb0\x64\xb7\x18\x75\x5d\x91\x4e\x1c\x1a\x23\xaa\x6c\xd2\x55\xac\xd6\x46\xc6\xfb\xe8\xbe\x7d\xd8\x3a\x11\xb8\x7a\x49\xa8\x63\xac\x60\xe2\x7c\x76\x8c\xff\x85\x96\xa6\xaa\x34\x73\x8d\x65\x94\x71\x26\x90\xf4\x8f\xea\xba\xff\xa0\x6e\x9a\xd1\xae\xa1\x4d\xa3\x69\x97\xbb\x19\xa0\x6d\xf2\x36\xaf\xce\xee\x1d\x13\x15\xbe\x62\x45\x61\x9d\xa7\x5a\xea\xba\xd7\x19\x97\x99\xdf\xda\x17\x92\xb7\xab\x02\xf7\xba\xbc\x56\xdb\x5a\x0d\xda\xde\xdc\xeb\xa9\x83\xa6\x6a\x03\xe2\xd3\xa6\xd1\xc4\x24\xd8\x66\xcc\xc2\xd5\x68\xbe\x37\x58\xb2\x4b\x06\x77\x40\x1d\x62\xb5\x60\x1b\x37\xb8\xd8\x30\xda\x8b\x18\x73\x4a\x53\xf2\x42\x66\x5c\x63\x6a\xa2\x76\xe1\x57\x92\x78\x9d\x47\x8a\x48\x73\xc7\xc4\x60\x4e\xb0\x9b\xe5\x33\xad\x96\xad\x0b\x56\xa1\xbf\x45\x07\x49\x8a\xdd\xad\xe7\x90\xd0\x38\xc7\xa5\x41\x9d\xb3\x14\x6b\x37\xfb\x58\xca\x6f\x04\xab\x17\xc8\xf6\x60\x67\x7c\x6a\xf4\x7e\xd3\x3d\x1d\xce\x53\x9e\xbb\xd9\xf0\x02\x67\xd5\xfc\x95\xca\xdc\x54\x90\x2f\x4d\xf2\xac\xd0\x5c\x1a\x21\xa3\x6e\xdf\x76\x1f\xdd\xea\xb2\x1c\x8f\x1f\x96\xa6\xe8\x74\xd6\x1e\xf0\x67\x63\xb0\x76\x23\x60\xe0\xb8\x41\x53\x5c\x62\xcb\xe8\x8d\xba\x8f\x7a\x20\x9c\x8d\x24\x49\xe2\xe4\x2a\x65\x96\x6b\x14\x14\x5a\xb0\x2a\xed\xb4\xb3\x57\x93\x37\x15\xd9\x99\xf1\x4b\xb4\xfa\x87\xce\x9a\x5b\x93\x09\x94\x1f\x45\xf2\x54\xeb\x4b\xf5\x46\xdd\xbb\xae\xed\x2d\x12\xe9\x4e\x4e\xa0\xad\x7f\xfb\xd0\x91\x47\xc6\x27\x1e\x98\x5c\x99\x05\xbd\x88\xee\x17\x28\xc1\x2c\x50\xe3\x51\x49\xd3\xb6\xab\x79\xcf\xcc\x6e\x6c\xdb\x1d\xa6\x9b\xb6\x8a\xac\x7f\xf4\xa4\xd8\x1d\xa5\xcd\xa0\x6c\x9f\x7b\x38\x26\xc3\x10\x74\x73\xfa\xce\xf9\x9a\xba\x07\xbd\x16\xe9\xa9\x68\xaf\xbc\x2f\xe9\x21\xa3\x8e\x3c\xfd\x69\xe0\xb0\xf1\xa2\x1d\x63\x0e\x10\xb7\x63\x0b\x4c\x9c\xbb\x07\x1b\x58\x8f\x2f\xc1\x67\xde\x30\xeb\xff\xf3\x65\xea\x49\x6e\x50\x7f\xd5\xfb\xc5\xbf\x50\xd6\x69\xf3\x4a\x25\x17\xfd\xb7\x4b\x13\xfe\x19\x00\x00\xff\xff\x1a\x87\x77\x76\xcf\x15\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb3, 0x8f, 0xae, 0x3b, 0x29, 0x6d, 0x86, 0xd3, 0x3f, 0x75, 0xcc, 0xd1, 0x99, 0x6a, 0x4e, 0xa8, 0x8e, 0xb0, 0xe6, 0x43, 0x27, 0xce, 0xf, 0x32, 0x60, 0xcc, 0x60, 0x9b, 0x99, 0xe4, 0xc7, 0x70}}
	return a, nil
}

var _templatesSingletonPsql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templatesSingletonPsql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonPsql_upsertGoTpl,
		"templates/singleton/psql_upsert.go.tpl",
	)
}

func templatesSingletonPsql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonPsql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/psql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0x25, 0xb2, 0xa1, 0x78, 0x77, 0xde, 0x3, 0x46, 0x95, 0xac, 0xcf, 0xb0, 0xa1, 0xf5, 0xf1, 0x14, 0x37, 0x11, 0x4d, 0x5f, 0x91, 0x90, 0x62, 0xc7, 0x92, 0xef, 0xee, 0x53, 0xcb, 0xb2, 0x5f}}
	return a, nil
}

var _templates_testSingletonPsql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\x39\x03\x39\x48\x5b\x87\x3e\xf4\xe5\x4b\x0e\xc6\x22\x71\x9c\x74\x71\xd9\x24\x6b\xbb\x3d\x14\xdd\xf6\x8e\x96\x46\x0e\x11\x89\x64\x48\x2a\x59\xf7\x90\xff\x5e\x0c\x29\xd9\xb2\x63\x25\xdb\x5e\x0b\xdc\xa7\x44\xe4\x33\xef\x0f\x87\x43\x3f\x72\x03\x66\xf5\xe5\xf6\xf2\xe2\x1e\xd7\x30\x06\x83\x2b\xfc\xa2\xd9\xc7\xda\xba\x89\xaa\xb4\x28\x31\xf9\x39\x79\x5f\xa5\xff\x3c\xbd\x5a\x4c\x67\xb0\x38\x3d\xbb\x9a\xc2\xcd\xf5\xd5\xdf\x80\xbd\xfb\x2c\x3f\xdb\xdf\x9d\x9e\x9f\xc3\xe4\xe6\x7a\xbe\x98\x9d\x7e\xb8\x5e\x00\x7b\xf7\x1e\x2e\x6e\x66\xd3\x0f\x97\xd7\xf0\xc3\x94\x50\xef\xbf\xff\x2c\x7f\x4e\xe3\xd8\xad\x35\x82\x5e\x2d\xd0\x3a\x34\x60\x9d\xa9\x33\x07\xbf\xc4\x51\xbe\x9c\x28\x29\xe1\x9d\x7d\x28\xd9\xf9\x59\x4c\x0b\xd7\xbc\x42\x20\x88\x90\xab\x38\xba\x53\xd6\x01\x6c\xbf\x6b\x8b\xa6\xfb\xad\xb9\xb5\xdd\x6f\x6b\xcb\x4a\xe5\xb8\xdd\x57\xc6\xcb\x0b\xe9\xe2\x38\xd2\xab\x5b\x6e\xed\x85\x28\x37\x80\x38\x72\x68\xdd\xf9\x99\xb7\xba\x51\x72\x2f\xf4\xfc\xd3\xd5\xa4\xca\x61\xa9\x54\x19\x3f\xc7\x71\x51\xcb\x0c\x84\x14\x2e\x49\x83\xdf\x1f\xb9\x90\x30\x86\x6f\xdb\xa0\x7e\x79\x26\xd8\x68\x04\x16\x5d\xad\x21\xaf\x2b\x6d\xc1\xdd\x21\xe4\xdc\xf1\x25\xb7\x08\x36\xbb\xc3\x8a\x03\x97\x39\x88\x8a\xfc\xb2\x20\x1c\x39\xa6\x80\x83\x43\x5a\xe2\x66\x0d\x86\xcb\x5c\x55\xe5\x9a\x74\xad\x50\xa2\xe1\x0e\x73\x20\x2f\x3b\xaa\x14\xb8\x3b\xee\xfc\xaa\x85\x8c\x4b\x58\x22\x98\x5a\x02\x5f\x71\x21\xad\x23\xc5\xb5\x15\x72\x45\x1e\xec\x2a\xb2\x0f\xe5\x52\x89\x12\x0d\xdc\xcc\x3e\x82\xe6\xd9\x3d\x5f\x21\x0b\xf1\x25\x1a\xde\xb5\xf1\xa4\x21\x90\x24\x05\x34\x46\x19\x0a\x9a\xe8\x82\xc6\x84\x85\x38\x8e\x1e\x85\x46\xc3\xe6\xe8\xce\xb1\xe0\x75\xe9\x92\x81\xa6\x3a\x86\x38\x07\x43\x18\xe8\x7a\x59\x8a\x6c\x90\xf6\x42\x29\x0b\x83\x21\xfc\xe9\x8f\x7f\xf8\x7d\x3f\xa8\x29\x29\x29\x34\xf8\x50\x0b\x83\x83\x94\x6a\xc9\x1a\xae\x8c\x21\x08\x5e\xa2\x9b\xfb\x02\x36\x72\xf9\x52\xf2\x8a\xb0\x91\x66\x9e\x46\x7d\x40\xda\x0c\x30\xcf\xae\x3e\x18\x6d\x06\x98\x27\x5d\x1f\x8c\x36\x1b\x18\x71\xaf\x03\xfb\x20\x77\xe2\xf6\x98\x96\xaf\x7d\xda\xda\xe0\x3d\xb8\x43\xd5\x3e\x3c\x41\xba\x81\x77\xa8\xdc\x11\x39\x53\xaa\x6c\x0d\xdc\x0b\xfa\x9b\x55\xb9\xcf\x2a\xd5\x77\x0c\x8f\xbc\xe4\xec\x0c\x57\x42\xfe\x95\x97\x22\xe7\x4e\x28\x99\xa4\xac\xf9\xc0\x24\x8e\x22\x0f\x09\xa6\xaf\x95\x9b\x56\xda\xad\x93\x90\x40\x2a\xfc\x36\x5f\xc3\x5e\x2c\xa5\xbd\xc5\x86\x12\x6c\xb0\xd7\xca\x25\xfe\x9f\xe9\x43\xcd\x4b\x9b\x84\x5c\x0e\xe1\xbb\x16\x1f\x12\xf8\x8a\xf2\xc0\x8d\x16\xde\x66\xa4\x1f\xdf\xe4\xb9\x15\xd8\xa4\x7d\x18\x47\x29\x9b\xdc\x61\x76\x9f\x50\x7a\x44\xe1\x4f\xc0\x37\x63\x90\xa2\xa4\x33\x11\x19\x74\xb5\x91\xb4\x1a\x47\xcf\x71\x1c\x8d\x46\x20\x0a\x90\xca\x9f\x4d\x3a\x81\xe7\x67\x40\x94\xc0\xdc\x4b\x97\x28\x93\x6e\x21\x53\x18\x8f\xe1\x3b\xaf\x69\x34\x82\x89\x41\xee\x10\x78\xd3\x04\xc4\xbf\x30\x87\x7c\x09\xe4\x3c\x8b\xa3\x7d\x06\x6c\x40\x6c\xee\xf8\xb2\xc4\xb0\xb1\x09\x3e\x0d\x0e\x35\x2e\x8f\x41\xb3\x8a\xdf\xe3\xed\x65\xdb\x02\x93\xf4\xfb\xb7\x82\x11\x05\x7c\xb3\xc3\x21\x02\x75\x14\xe6\x46\xe9\x85\x77\xe9\x80\xb2\x1d\x6d\xd1\xf3\xae\x64\xe6\x23\xfd\x6a\xd9\x38\x8a\xa8\xa3\x92\x0b\x27\x63\xc0\x2f\x98\xb1\x89\xaa\x2a\x2e\xf3\x64\xa0\x57\x3f\xd1\x1e\xf5\x87\xe3\xe3\xd0\x7c\x8e\x95\x2c\xd7\x83\x21\x74\x52\xd1\xca\xb3\xa9\x7c\x84\x31\x70\xad\x51\xe6\x89\xb2\xf4\x2d\x0c\xd1\x9b\xe0\x7a\x35\x95\x8f\x49\xca\x18\x23\x91\xe0\xe4\x61\xa3\xf6\xa1\xf4\x06\x3a\xa5\xec\x4a\x7c\xbd\x19\x4a\xfb\x10\x9e\xc8\x84\x50\xec\x56\x68\x4c\x3a\xee\xce\x5d\x4e\xa9\x39\x19\xc3\xb7\xcb\xb5\x43\xcb\xce\xea\xa2\xf0\xb7\x4d\xc7\x58\x3f\xa8\x13\xf7\xdc\xe5\xaa\xa6\x7e\xf4\xb4\xbb\x18\x2a\xb2\x63\x2e\xde\x89\x64\xee\x72\x7f\xd5\x49\x7c\xba\xf8\x01\xd7\xe7\x68\x9d\x51\x6b\x34\xc9\x66\x74\x18\x82\x49\xf7\x45\x82\xda\x3d\x17\xe3\x2e\x09\xb6\x3e\x70\xe3\x5e\xe7\x80\x32\x96\xfd\x68\xb8\x4e\xd0\x50\x7b\x29\xb8\x28\xe9\x4e\x54\x60\x49\x16\x1a\x06\x40\x16\xaa\x43\x9d\x6f\x97\x6f\x5d\xcf\x7e\xb5\x31\xfb\x50\xee\x59\x3a\x14\xd5\x8f\x5c\x1c\xb4\x53\x54\x8e\xdd\x1a\x21\x5d\x29\xc9\x40\xba\xbf\xb6\x53\x88\xa6\x4f\x25\x69\xfa\x95\x2e\x3e\x71\xe1\xa0\x50\xa6\x27\x25\x71\x14\xfd\x44\x0c\x60\x93\x52\x59\x4c\x52\x18\x8d\xe0\xb4\xa0\x91\xac\x3d\x5d\xc2\x42\xae\x24\x0e\x21\x23\x84\x1f\x60\x9e\x8c\x70\x08\x28\x73\x50\x85\x5f\xd0\x42\x63\x7c\x38\xbd\xff\x6d\xd4\x7b\x3c\xf9\x15\x71\xbf\xac\x8e\x8f\xbb\xd1\x21\xc5\x76\x9a\xdb\x9d\x76\x4c\x2d\x27\x55\x9e\x58\x22\xfb\xb0\xd5\xd0\x4c\x84\x43\xe0\x66\x65\x81\x31\x16\xbe\x3b\x33\x51\x76\xa0\x39\x34\xc2\x41\x2a\xb4\x92\xec\x3f\xeb\x08\xcd\x45\xe1\x9d\x49\x29\x91\xe1\x86\xc8\x3a\xa7\x31\x78\x62\xd9\x35\x3e\xcd\x90\xe7\x68\x1a\x74\x08\xd7\x86\xc3\x7e\xa8\x6d\xd8\xfe\x8e\x92\x75\xdb\x44\x50\xb1\x59\x0c\x95\x0e\xc2\x9b\x4b\xe5\x64\x0c\xb4\x3d\xab\xe5\x81\xa2\x77\xeb\xdb\x96\xca\xd4\x52\x0a\xb9\x3a\x19\x6c\x52\x1c\xb2\x94\xee\xe1\x83\xf1\x1d\x1a\xec\x6d\xef\xb3\x64\xff\xea\x7a\xb3\xe0\x4d\xc6\xe1\xef\xff\x08\xa9\x24\x9f\x1b\xa1\x76\xa9\x8d\x62\xae\xc9\x6e\x91\x0c\x6e\x2f\xff\x7c\x33\x5f\x8c\x8f\xac\x6f\xfd\x34\xb4\xf8\x91\x62\x0f\x73\x7b\x33\x5b\x8c\x8f\x72\x8f\xa1\x41\xe5\x10\xe6\x2f\xf3\xe9\xac\xd5\x43\x83\xd2\x41\x3d\xa7\xf3\xf9\xc5\x87\xab\x69\x8b\xdb\xbe\x5e\x08\xfd\xdc\x13\xd7\xfe\x25\xbf\xe5\xaa\xab\xf4\xb0\x2d\x9b\x50\xb5\x13\x25\x5b\x60\xa5\x3d\x6c\xe0\xe7\xf5\x55\x3b\xbc\xbe\x36\xe7\xf4\x1e\xc2\x70\x88\x41\x69\x1a\x17\xa1\x10\xa5\x9f\x41\xa9\x18\x14\xd8\x45\x13\x98\xf7\x62\x70\x64\x4f\x8e\xf2\x13\xad\xac\x5b\x19\xb4\x27\x9d\x8c\xb6\x59\xdb\x64\xa6\x33\x37\x91\x7b\x9d\xf3\xf0\x52\x6d\xab\xc8\x03\xc9\x76\x07\x53\x4a\x02\xa5\xaf\xb8\x73\xd4\xeb\x48\x3b\x4e\xfe\x86\x5c\xda\x0e\x1e\xff\x47\xb7\xba\xa4\x83\x31\xb8\x4a\x33\x3f\x63\xa6\x9b\xb3\x42\x4b\xcd\x6d\xd2\x43\xc8\xdd\x51\x6f\x4b\xc7\x46\x81\x66\x4d\xeb\xf5\x14\x0c\xe0\x7c\xf9\x62\xb6\x3a\xac\xbb\x3b\x80\xbe\xa1\x99\xa0\x5e\xef\xe0\xf8\x58\x14\xc7\xf8\x45\x58\x67\x0f\x99\x19\x8d\xc0\x21\x37\xb9\x7a\x92\xbe\xaf\xd7\x0e\x2d\x64\x25\x72\x59\x6b\x70\xdc\xde\x5b\x78\xba\x43\xe9\xaf\xc2\xf0\x00\x2f\x84\x14\xf6\xae\x6d\x6e\x87\xfc\x6c\x15\xf6\x3f\xa7\x77\xc6\x6a\xff\xab\x48\x9b\xd6\x37\xa6\xf4\xa8\xc5\x83\x47\xfc\xcf\xa7\xf6\x4e\x33\x55\x96\xcd\xb0\x52\x8f\xf4\xc6\xe8\x34\xa3\xbe\xba\x2b\x49\xf1\x26\xcd\x8f\x3b\xc3\x10\xa8\xff\xf9\x44\x14\x9b\x28\x0f\x04\xd6\x6e\x0d\x7d\x3c\xde\x81\xbd\x5c\x6d\x11\xcd\xb5\xf4\x50\xb2\x1b\x8d\x32\x19\xb4\x1d\x65\x30\x84\xdc\x88\x47\x34\xec\x76\xfe\xe9\xea\xac\x16\x65\xfe\xa9\x46\xb3\x6e\xae\x8c\xf6\xa5\x1a\xf8\xff\xf2\x38\xed\x1f\xb6\xe6\x3d\x98\xbe\xd6\x1a\xa5\x28\x87\x2f\xee\x9f\xdd\x58\x9e\xe3\x7f\x07\x00\x00\xff\xff\x7d\xe6\x96\x96\x73\x13\x00\x00")

func templates_testSingletonPsql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testGoTpl,
		"templates_test/singleton/psql_main_test.go.tpl",
	)
}

func templates_testSingletonPsql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6d, 0xc1, 0x99, 0x24, 0x98, 0xcb, 0x39, 0x6d, 0x51, 0xc0, 0x1d, 0x6c, 0xd8, 0x44, 0xfe, 0x54, 0xcc, 0x3, 0x28, 0x8d, 0xc8, 0x86, 0xcd, 0x6b, 0xf6, 0xf3, 0x33, 0x91, 0x7f, 0x4c, 0x2b, 0x94}}
	return a, nil
}

var _templates_testSingletonPsql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonPsql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testGoTpl,
		"templates_test/singleton/psql_suites_test.go.tpl",
	)
}

func templates_testSingletonPsql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4f\x6f\xdb\x3e\x0c\x3d\x5b\x9f\x82\xbf\xe0\xb7\x41\x1e\x5c\x15\xbb\x76\xc8\xa1\xff\x0e\xc5\xb0\x20\x68\x9c\xf3\xa0\xda\x74\x2a\x44\x91\x0c\x89\x5e\x92\x19\xfa\xee\x83\xe4\xb4\x4d\xdb\x74\x08\x86\x0d\xc3\x0e\x89\x2d\xe1\xf1\x3d\x92\x8f\x74\xdf\x9f\xc0\xff\x52\x2b\xe9\xe1\x6c\x0c\xe2\x3c\xbe\xa1\x17\xa5\xbc\xd3\x08\xc3\x43\x4c\xe4\x0a\x43\x60\x4d\x67\x2a\x20\xf4\xd4\xf7\x43\x84\x98\xb7\x53\xdd\x39\xa9\x43\x98\xb7\x1e\x1d\x71\x82\x0f\x11\xa0\xcc\x42\x94\x39\xf4\x2c\x23\x31\x95\x4e\x6a\x8d\x9a\xe7\x8c\x65\xaa\x01\x8d\x86\x3f\x12\x5c\xd9\xb5\x99\x29\xb3\xe8\xb4\x74\x21\x5c\x5a\xdd\xad\x8c\xcf\x61\x3c\xfe\x19\x6c\xea\xd4\x4a\xba\xed\x67\xdc\x3e\x06\xf4\x2c\xcb\x48\xcc\x96\xaa\xe5\xa3\xf8\xdf\x2a\xb3\x00\x4a\x35\xac\x15\xdd\x83\x35\x7a\x0b\xed\x10\x07\x4b\xdc\x42\x35\x44\x8e\x72\x96\x05\xc6\x32\x8f\x58\xc7\xfa\x9d\x34\xb5\x5d\xa9\xef\x28\x26\xb8\x9e\x21\xd6\x3c\x67\xd9\x37\xe9\x00\x5d\xfa\x59\xc7\xb2\xd3\x53\x38\x27\xc2\x55\x4b\x40\xf7\x08\x37\x93\xd9\xf5\x6d\x09\x5e\xd5\x08\xb6\x01\x69\x60\x3e\x8d\x37\x2c\xb3\x91\x71\xaf\x57\x4f\x15\xf4\x21\xb5\x22\x92\xee\x6b\xce\xc8\x75\x15\xf1\x98\x4c\x01\xef\x6d\x01\x6f\x34\xe0\xea\xa2\xdc\xb6\xe8\x0b\x20\xd7\x61\xfe\x29\xf1\xfc\x37\x06\xa3\xf4\xae\x11\xd7\x31\xd3\x86\x8f\xe6\x26\xb5\x80\xec\x93\xc8\xe1\x84\xc0\x27\xe9\x33\x78\xe7\x47\x45\xe4\xdb\xf5\xa5\xef\x55\x03\xc6\x12\x88\x89\xbd\xb4\x86\x70\x43\x21\x54\xb4\x89\x95\x55\xc3\x59\x5c\xc8\x6a\xb9\x70\xb6\x33\x35\xcf\xfb\x1e\x4d\x1d\x02\xcb\x06\xc8\x97\xce\x53\xb9\xe1\x89\x65\x9f\xe1\xd5\xc5\x9d\x55\x5a\x5c\xe0\x42\x99\xc4\xa1\x3d\xee\xdf\x95\x1b\x5e\xd1\xa6\x88\x05\x3e\x28\x1c\x05\xca\x59\x56\x63\x83\x0e\xe2\xe4\xf2\x1c\x7a\xf8\x0a\x63\xa0\x8d\xb8\xb5\x5a\xdf\xc9\x6a\xc9\x73\x08\xd1\xe1\x47\x2f\xac\xd8\x0d\xf2\x5b\x85\x47\x4f\xd0\xd4\x70\x12\x02\xc4\x53\x23\xb5\xc7\x24\x5a\x40\xca\xe5\xc6\x34\xe8\x78\xfe\xfc\x74\x9c\x47\x5d\x92\x3e\x6c\xd0\x2b\x67\x2a\xdb\x19\x4a\x17\x2f\xa6\xec\x61\x23\x79\x2e\x2e\x23\xe6\xc8\x52\x9e\xba\xf0\x3a\x4b\xfe\x20\x1b\x21\x49\x38\x82\x3e\x3e\x83\x8c\xd6\xd2\x10\x58\x83\xe0\xb0\xb2\xae\x2e\x60\x61\xe9\x6c\x54\x0c\xf8\x5d\xd2\x2f\x56\x67\x3e\xbd\x3a\x2f\xaf\x0f\xad\xce\xef\x58\x8e\x9d\x35\xc7\x7e\x44\x84\x10\x7f\x74\x95\x7e\x7d\xc6\xe2\x96\xff\xe5\x11\xfb\x47\x26\x2c\xb0\x1f\x01\x00\x00\xff\xff\xcc\x1c\xe7\xf3\xcf\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x51, 0x9b, 0x35, 0xb1, 0xd2, 0xc5, 0xc, 0x1b, 0xe, 0x34, 0x1f, 0xf9, 0x0, 0xb9, 0xad, 0x23, 0xa8, 0xb4, 0x8c, 0x49, 0x7b, 0x29, 0x5f, 0xf1, 0x89, 0x83, 0x8c, 0xd4, 0x90, 0xbd, 0x17, 0x1}}
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/psql_upsert.go.tpl": templatesSingletonPsql_upsertGoTpl,

	"templates_test/singleton/psql_main_test.go.tpl": templates_testSingletonPsql_main_testGoTpl,

	"templates_test/singleton/psql_suites_test.go.tpl": templates_testSingletonPsql_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_upsert.go.tpl": &bintree{templatesSingletonPsql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.go.tpl":   &bintree{templates_testSingletonPsql_main_testGoTpl, map[string]*bintree{}},
			"psql_suites_test.go.tpl": &bintree{templates_testSingletonPsql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
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
