// Code generated by go-bindata. DO NOT EDIT.
// sources:
// embd/dojoConfig.yml (12.316kB)

package main

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

var _embdDojoconfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x7a\x5b\x7b\xe2\xb8\xb2\xf6\x3d\xbf\xa2\xbe\xe4\xa2\x67\xbe\xd5\x84\x24\x3d\x9d\xce\x70\xb5\x1c\x70\x12\xaf\xe6\x34\xb6\x99\x3e\xcc\x33\xcb\x8f\xb0\x0b\xd0\xc4\x48\x8c\x24\x48\x33\xbf\x7e\x3f\x25\xd9\xc6\x06\x92\xee\xbd\xaf\x42\x64\xd5\xfb\x96\xaa\x4a\xa5\xd2\xe1\x1c\xfa\x38\x67\x9b\xdc\x40\x2a\xc5\x9c\x2f\x60\x2e\x15\x2c\x64\x26\xff\x92\xc0\x85\x36\x2c\xcf\x99\xe1\x52\xe8\xd6\x79\xeb\x1c\x1e\xe5\x33\x18\x09\x0a\x59\x06\x66\x89\x90\x73\x81\x1a\x66\x98\xcb\xe7\x6e\xeb\x1c\x7c\x96\x2e\x6d\x1b\x28\x5c\x2b\xd4\x28\x8c\x06\x06\x5b\x96\x6f\x10\x36\x1a\x33\x98\xed\xac\x5c\x81\x8c\x0a\xb8\x00\xb3\xe4\x9a\x68\x57\xcc\x10\xc8\x1f\x82\xad\xf0\xcf\x2e\xfc\x91\x39\xc5\xfe\x84\x73\xf8\xc3\x1f\xfd\xfe\x27\xb4\xe1\x8f\x3e\xea\x54\xf1\x35\x69\xf4\x67\xeb\x1c\x9e\x97\xa8\xb0\x92\x01\xae\x2d\x3a\xfd\x03\x72\x6e\x7f\xbb\x51\x6d\x94\x1d\x04\x70\x83\x2b\xea\x5e\x41\x17\x12\xc5\xff\x85\xa6\x64\x82\x97\x65\xad\x2a\x85\x1c\x8a\x2d\x57\x52\xac\x50\x18\x96\xc3\x96\x29\xce\x66\x79\x31\x54\x23\x41\x6e\x51\x29\x9e\x61\x0d\xcd\xc2\x00\x33\xa0\x36\x02\x0c\x5f\x59\xed\xeb\xa3\x22\x68\x06\xd9\xbe\xc5\x8d\x84\x99\x43\x90\x37\x1a\xd6\x1b\xb5\x96\x1a\xad\x6b\x02\x37\xde\xb9\xcc\x73\xf9\xcc\xc5\xc2\x8d\x45\x03\x53\x08\x39\xce\x0d\xe0\x6a\x6d\x76\xc0\x9e\x18\x9c\x9d\xbd\xa5\xbe\x3b\x78\xe6\x79\x0e\x33\x04\xc5\x44\x26\x57\xf9\x0e\x16\x28\x50\x31\x83\x19\x69\x58\x38\xa9\xd4\xb2\xa7\x30\x43\x61\x38\xcb\x3d\x3f\xba\x7e\x7f\xf3\x11\x77\xad\x73\x88\x30\x55\x68\xe8\x77\x2b\x70\xfd\xbb\x2d\x80\xdf\x51\x69\x2e\x45\x17\xce\xae\x2e\x3e\x5c\x5c\x9d\xc1\x39\xf4\xfb\x49\xd1\x0a\x6d\x08\x31\x47\xa6\x11\xb6\x45\x8b\x9c\x53\x18\x62\x6a\xfa\x14\x77\x73\x25\x57\xf0\xc0\xcd\x72\x33\x2b\x7b\xea\x16\x40\x24\x37\x2a\xc5\x92\x05\xe6\x2c\xd7\xe8\x80\x1b\x5f\xa0\x0d\x77\x52\xe6\xc8\x04\xf0\x39\x30\xd0\xf6\x63\x35\x1a\xae\xc9\xb8\x5c\x61\x06\x3f\x6d\xc9\xd4\xca\x31\xfc\xdc\x02\xb0\x56\xfc\xaf\x75\xae\xda\xe0\x5b\x27\x9c\x92\xd5\xb3\x3d\x80\x35\x9a\x4c\xd3\x8d\x2a\xfd\x4b\xd6\x26\xdb\xab\x83\x51\xad\x95\xdc\xf2\x0c\xb3\x4a\xf7\x3b\xc5\x44\xba\xec\xc2\x59\x86\xdb\xb3\xba\xea\xee\x03\xb4\x21\x5e\x22\xcc\xec\x3f\x6f\x34\x3c\xfa\x5e\x9f\xc2\x68\x86\x90\x2e\x31\x7d\xc2\x0c\xe4\xc6\xd0\xa8\x9a\x03\x2e\xf4\xad\x68\x7a\x72\xb5\xe2\xa6\x0b\x70\x7d\x7d\xfd\xeb\x2f\x6c\x76\x93\xde\xfc\xfa\xcb\xcd\xed\xe5\xfb\x0f\xb3\x14\x3f\xfc\xfa\xee\xf6\xe6\xc3\xcd\xed\xed\xcd\xaf\xb3\xeb\x0f\xb7\xb7\x19\xbe\xcf\xea\xaa\x38\x61\x68\x17\x01\xa5\xd0\x05\xa4\x9b\x19\xf4\xbf\x8d\x1d\xd0\x6b\x4c\xf9\x9c\xa7\x90\xba\xfe\x65\x28\xd9\xd0\x27\xbb\xd8\x5e\x6e\x28\xf0\xdf\x16\xc0\x6f\x1b\x8e\xa6\xe1\x35\xdb\x02\x6d\x88\x36\x6b\x4a\x16\x1a\x04\xe5\x80\x9c\xc6\xb8\xde\xd0\x07\x29\xf2\x1d\xa0\x52\x52\xe9\x0a\x5f\x2f\xe5\xb3\x68\x01\xc4\x8a\xa5\xd8\xb5\xe3\x76\x68\xb6\xa1\xe6\x7b\x23\x01\x85\x9d\x8d\xa4\xc8\x4a\x6a\x43\x7e\x99\x49\x8d\x90\xcb\xc5\x82\x7c\x96\x6d\x14\xfd\x29\x1c\xdb\x02\x08\x31\x63\xa9\xa9\xa3\xba\x96\x26\xac\x72\x6d\x1a\x85\xe6\x86\x6f\x29\x32\xe6\x45\xd0\xda\x9c\x28\x17\x14\xad\x7d\xdc\x9e\x0a\xd5\x3e\x6e\x93\xe3\x40\xa5\x8c\x93\xe1\x16\x73\xb9\xa6\x5c\x52\xaa\xa4\xdf\x92\x3d\x35\xcc\xf9\x37\xcc\x8a\x09\xdd\x02\x98\x28\xb9\x5a\x37\x6d\xe9\x9a\x9a\x7a\xae\x5d\x1b\x61\x37\x33\x59\x91\x19\xda\x30\x1a\xc7\x10\x0c\x27\x03\x7f\xe8\x8f\x62\xbf\x0f\x5f\xfc\xb8\x05\x30\x64\x69\x03\x7b\xc8\xd2\x26\xb0\x46\x53\xcf\xe2\x60\x98\x5a\xa0\x01\x46\x71\xe2\x3a\x9f\x06\x0e\xa5\x34\x5d\x38\xeb\xc8\xb5\xe9\xd0\x02\x53\x4c\x01\x6a\xa6\x9c\x40\x7f\x32\xae\x30\x35\x52\xed\xaa\x24\x5c\xcb\x0b\x6c\xbd\x86\x91\x34\xd8\x85\x91\x04\xa3\x18\x27\xef\x75\xaa\xb0\xa7\x79\xf5\x17\x13\x0b\xd9\xde\xcb\x34\x66\x19\xb4\xa1\x5f\xe1\x73\x51\x71\x13\x55\x8d\xa6\x48\x18\x34\xe9\x5b\x00\xf7\x3c\x47\xdd\x85\xb3\x5c\xa6\x2c\x2f\xe0\x6c\xdb\x6b\x68\xb6\x33\xcc\x6d\xb7\x9f\xb4\x61\x86\xa7\xc0\xb4\x46\x43\x1e\x5d\xe7\x92\x65\xfa\x2d\xa0\x49\x29\xe9\x0c\x31\xe3\xac\x0b\x67\x2b\xfa\x5b\x10\xd8\xb6\x13\x04\x8e\x98\x18\x1c\x0a\x66\x15\x49\xaa\x10\x85\x5e\x4a\xa2\x30\xa8\x0d\x30\x65\xf8\x9c\xa5\x66\xcf\x14\x59\x45\xba\x70\xe6\x34\x2a\x6d\xe3\xd4\x7b\x8d\xac\x3e\x82\x92\xf0\x3f\xd1\x5b\xe0\x2b\xb6\xc0\x3d\xbe\xb7\x5e\x93\x0f\xf6\x66\xf7\xd6\xeb\x13\xb8\x85\x33\xec\xda\x7d\xe8\xe2\xbe\x75\xa0\xf5\x34\xd7\xd6\x8c\xc6\xe5\x4f\xb6\x5a\xe7\x98\x31\xc3\x9a\x89\xbf\x6a\x3e\x98\x4c\x45\x68\x52\x80\x68\xdb\x07\x6c\xa7\x62\xc6\xd7\x82\x97\xa6\xd3\x26\xcf\xcb\x08\xda\xcf\xfc\x7d\xeb\x69\x68\xca\x70\x12\x32\xf9\x2c\xc8\x11\x65\xd0\x1c\x84\x92\x8b\xd6\xa9\xde\xb0\x3c\xdf\x15\xb3\x7c\xb6\x71\xd9\xe7\xa0\x0e\x32\x1a\xf3\x39\xe5\x8d\x3b\x5a\x40\x01\x7c\xb1\xe0\x82\x62\x7a\xb8\x8b\x7e\x1b\x14\x06\xed\xdf\x25\xae\x9d\xcc\xca\x0c\x9b\xd1\x72\x83\xae\xc5\x48\x4a\x16\xf0\x53\xf4\xdb\x80\x1b\x7c\x0b\x56\xee\x2d\x4c\xa4\x36\x0b\x85\xf6\xf7\x90\x0a\x94\xfe\xdd\xcf\x85\x5e\x3d\x2f\xf2\x41\xfb\x22\xe2\x31\xff\xdd\xff\x7f\x96\x76\x40\x91\x5b\xb7\x43\xff\x2e\xb1\x6d\x07\x56\x78\x5e\xa2\x80\xfe\x1d\x79\x49\x0a\xb7\x22\x50\xd5\xb5\x94\xda\x74\x34\xaa\x2d\xaa\xce\x76\x05\x3f\xd9\x89\xf0\xb3\x1b\xd0\x37\xae\x8d\x6e\x26\xc3\xbb\xc4\xb5\xbe\x00\x7e\x60\x4d\x96\x53\xe1\xb9\x03\x74\x32\x5a\x82\xa8\x8a\x54\x10\x88\x6e\xa9\x05\x08\x37\x1a\x55\x17\xce\x94\x94\x66\x6f\x38\xdb\xea\x96\x1b\x54\xf6\x77\x99\x62\xb2\xd2\x92\xce\x2c\xb6\x1e\x65\x22\x83\x70\xcd\x74\x51\xdd\x42\xe8\xff\x36\x0d\x42\xbf\x6f\x85\x14\xae\xa4\x41\xdb\xc7\xea\x42\xee\xec\xdf\x69\x47\x4e\x42\x5d\x38\xdb\x22\x5e\xc6\x4b\xc9\x04\xc3\xab\x8c\xe1\x44\xca\x7f\x2e\x39\x3e\xb1\x9a\x42\x16\xbe\x0d\x13\xa6\xf5\xb3\x54\xd9\xb1\x3e\xba\x52\xd5\x69\x66\x13\xaf\x84\xeb\x5f\x8a\x8a\x0d\xd2\x25\x53\x2c\x35\xa8\x1c\xf7\x88\xad\xb0\x98\x81\xd9\x6c\x4f\x44\xcd\x94\x94\x6b\x65\x71\x45\x61\x2b\xcb\x9a\x8d\xed\x3a\xbb\xd1\x68\xf1\xa6\xce\x8e\x0e\x6f\xa3\xd5\x1e\x72\xea\x6c\x49\x7f\xc4\x29\x58\xab\xf4\xab\xd8\x93\x1f\x35\xd3\xe4\x7b\x56\xb2\x5c\x27\x68\x7e\xd0\x68\x8f\x52\x9b\x32\xbf\x53\xf8\xee\x99\xe9\x4b\x7d\xa2\xd1\x57\x1a\xad\xd3\x5f\x2a\xd3\x85\x77\xef\x2e\x6f\xf6\x9a\x4a\x45\xfd\xed\x9f\x86\x86\x94\xca\xb8\x36\x28\x28\x50\xa4\xb0\xf2\x7d\x25\xd7\x87\x93\x81\xda\x9a\x4b\xad\x41\x5a\x5f\x1b\x99\x82\xb2\x0e\xf5\x63\xa2\x1e\x7c\x54\x15\xce\xe5\x46\xd0\x0c\x18\x47\xdd\x43\xf7\xb5\xb5\x2a\x2b\xce\x71\x54\x7a\x6f\x1c\x15\x7e\x92\x20\x9f\xc5\x61\x2a\x26\xc6\x3c\xb7\x41\x6e\x53\x7e\xdd\x6b\xcf\x6c\x99\x73\x9c\x49\xf9\x97\xfc\xc8\x6e\x19\x37\x4b\xe4\x33\xc6\xdf\xed\x29\x5e\xf2\x5a\xc9\x79\x32\x51\x7e\xcf\x55\x0f\x4a\x6e\xd6\x27\x47\x64\xbf\xb8\x21\xb9\x9f\x2f\x8f\x89\x1d\x0d\x6a\x1a\xf4\xbb\x70\xf5\xee\xdd\x87\xbd\x81\x82\x7e\x11\xdd\x10\xf4\x4f\x95\x22\xc5\x30\x9c\x56\x47\xe2\x0f\x56\xdc\x29\xf2\xa2\xfc\x82\x3e\xd3\xd2\x86\x86\x7c\xa8\x9d\xcf\xfa\xdc\x46\xa3\xad\x8f\x3a\xba\xf8\x54\xfd\xb8\xc8\xb8\x36\x17\xeb\x5d\xb9\x74\xfb\x71\x42\x02\xd6\xd0\x66\x59\xcd\x42\xae\x8d\xe2\xb3\x0d\xed\xe3\x4a\x49\x3b\x5c\xda\xa1\x30\x5b\xb2\x1a\xb9\x5f\x89\x2d\x2f\x2d\xf6\xaf\xf0\x36\x28\xa9\xef\x01\x65\xad\xe3\x2b\x4c\xe5\xba\xa3\xf0\xe4\x1e\xd3\x2d\x7b\xdb\x63\x35\x2e\x50\x6c\x2f\xd6\x4a\x66\x35\x25\x7c\xb1\x2d\x75\x20\x8a\xca\xb4\x6f\xf4\x0b\xbb\x72\x67\x02\x2a\x53\xb2\x15\x17\x8d\x19\xc2\xa8\xa5\xac\x59\xfa\xc3\x60\x54\xce\x10\xdb\x15\x1a\x8b\x45\x3d\xd1\xe0\x8c\x2a\x95\xfa\xc4\x38\x46\x7a\x69\x22\x1c\xe3\x00\xdb\xb3\xfd\xd8\x7c\xf0\x57\x8c\xe7\x25\xeb\xbf\x0f\x33\x98\xe3\xb7\x7d\xa0\xed\xfa\x02\xcb\x32\xbb\xd5\x2a\xb5\x38\xa6\x2e\xa2\x41\xd9\x30\x2c\xce\x86\x1a\x80\xf6\x1b\xb4\xe1\xd3\x91\xa8\x7e\x43\x36\xa6\xaf\x55\x96\x1c\x30\x8b\xe3\x1d\x99\x85\x3e\xbc\x08\x92\xb3\x3a\xc6\x98\xb6\xa1\x64\xdc\x06\x80\x6b\x85\x36\x0c\x68\x02\xc8\x39\x0d\x8d\xd3\xf6\x86\xe5\x2f\xda\x56\x03\x5e\x2c\x2e\x2c\x0d\xa1\x5f\x75\x69\x05\xbe\x7a\x5b\xfe\x7f\x6d\xff\xbf\x6e\xb5\xea\x53\xd2\xcb\x73\xf9\x8c\x19\xad\x05\xa4\xc3\xff\x2f\x95\x18\x0c\xc6\x9f\xfc\x7e\xf2\x38\x8e\xe2\xa8\xa6\x45\xb9\x46\xe8\x4e\x30\xd1\xc0\x9c\x30\x39\x31\x95\x42\x60\x6a\x9a\xc1\xda\x02\xe8\x61\x8e\x6a\x77\x87\xcc\x44\xe9\x12\xb3\x4d\x8e\x34\xbd\x84\x5b\xc9\x0b\xb2\x9e\x3f\xf0\xc3\x2f\xc9\x9d\xef\xc5\x49\xd4\x7b\xf4\xfb\xd3\x81\x9f\xdc\x07\x03\x7f\xe4\x0d\x7d\x5a\xa1\xca\x62\xfd\x5f\x90\xc9\xbf\x2e\x52\x0b\x79\x31\x43\x66\x2e\xb2\xd9\x9e\x43\xc9\x27\x54\xc5\x72\x77\x80\x1c\x8e\x3f\xfa\xa1\x1d\x0c\xb4\xed\x8a\x68\x57\xf5\x32\x4a\x9c\x3c\xcc\x2c\xc0\x01\x5e\x19\xd8\x2f\x61\x4e\xbc\x28\xfa\x34\x0e\xfb\xa7\xe6\xc0\xeb\xb8\x66\x59\x66\x82\x62\x44\xd9\xec\x42\xff\x9d\x73\x83\x2f\x10\xc5\x8f\x4d\x5b\x9c\x94\xa5\x85\xb2\xfc\xa5\x8b\x43\x8a\x53\xf4\x76\x79\x6f\x5f\x9d\x64\x1a\x87\x71\xb9\xce\x97\xa7\x9a\x33\x96\x3e\xa1\xc8\xe0\xa7\x90\xcd\x66\xdc\x0c\x7f\x03\xa9\x20\xc4\x8c\xeb\x9f\x0f\x90\xc9\xcd\xd6\xb9\xfa\xef\x9c\x25\xaf\x0c\xc8\x75\x84\x36\xb0\xd5\xdf\x6b\x57\x13\x57\xe0\xaf\xa9\x3e\x0d\x07\x2f\x39\x63\x1a\x0e\x68\x5d\x0b\x07\x07\x11\xf9\x3d\x67\x14\x79\xf2\x34\x66\xe4\x87\xe5\x62\xf9\xa3\xa8\x03\xb9\x18\xe0\x16\xf3\x23\xcc\xc1\xf8\x21\x19\xf8\xbf\xfb\xa4\x66\x3c\xee\x8f\x2b\x89\x10\xf5\x26\x37\x77\xce\xca\xfb\x9d\x7e\x55\xe6\x16\xf2\xa1\x1f\x4d\x07\x71\x72\xe7\xf5\x3e\xfa\xa3\xfe\x69\x10\xff\xdb\x9a\x2b\xda\xd3\xdf\xde\xfc\x72\x79\x79\x4a\xdc\xff\x3c\x09\x42\x3f\x3a\x14\x8f\x99\x7e\x0a\x16\x42\x2a\x74\x40\xf5\x6d\x52\x01\x10\x7b\xd1\xc7\x24\x78\x18\x8d\x43\xbf\x00\x3b\x05\x12\xa1\xe2\x2c\xe7\xff\x58\xa3\xae\x79\xfa\x94\x1f\x44\x80\x85\x89\x22\x3f\x0c\xbc\x41\xf0\xd5\xda\xb7\x04\x39\x3e\x85\xad\x19\x31\xf4\xfb\xfe\x28\x0e\xbc\x41\xe2\xf9\x51\x72\xfd\xfe\x26\xf9\xe8\x7f\x81\x36\x7c\xc4\x5d\x75\x28\xed\xf9\x11\xa0\x48\xd5\x6e\x6d\x20\xad\xd0\x34\x68\x23\x15\x66\xee\x40\xbe\xb1\x3c\x95\x65\x2c\xb1\x47\xe1\x7d\x4f\xca\x27\x8e\x8f\x71\x3c\x19\x8b\x7c\xd7\xb0\x41\x14\xde\x27\xbd\xf1\xf8\x63\xe0\x27\xf6\xfb\x68\xf0\xa5\x59\xce\x52\x55\x5e\x4a\x96\xfb\xc6\x62\xcb\x4f\xc2\x90\x5a\xec\x06\x51\x84\xe9\x46\x61\xa3\x52\xae\xf3\x44\x7e\x6f\x1a\xfa\xc7\x2c\xae\x28\x21\x51\x98\xe7\x6c\xf1\x2a\x59\x59\xe0\x57\x5b\xee\x72\x0b\xe0\xc5\xde\x9d\x17\xf9\x89\x3f\x7a\x08\x46\x44\x12\xa1\x71\x07\x7d\xf5\xbd\x78\x0d\xa1\x99\x5c\x2b\xf9\x22\xaf\x1e\x48\x53\xe7\x9a\xec\xa8\x91\xf2\x2b\xd9\x22\xc5\x1f\xc8\x8e\xdc\xe2\x58\xca\x1e\x27\xe1\x4a\xbe\x96\x7f\x0f\x30\x48\xa8\x8e\x61\x33\xde\xb1\xbc\x4b\x76\x87\xb2\x52\x91\xee\xe7\xa5\x70\xbc\x5b\x57\xca\xfb\x93\xd0\xef\x79\xb1\xdf\xaf\x81\x37\x72\x52\x85\xed\xd2\x51\x05\x6d\x1d\x64\xe1\xf7\x45\x5b\x1d\xa3\x91\x84\xf6\x20\x2e\xff\x1c\x28\x38\x75\xa5\x0c\xc9\x4e\xed\x51\xd9\x90\x7d\x1b\xe2\x4a\xaa\x5d\xc4\xff\xc1\x2e\xdc\xbe\xbb\xbd\xbd\xb9\xbc\xdd\x63\x25\xd3\xc9\x60\xec\xf5\x93\xa1\xf7\x39\x19\xfa\xc3\x71\xf8\x25\x89\x82\xaf\x64\x7a\xf6\xc4\xe0\x16\x56\xb3\x72\x0a\xf6\x71\xb6\x59\x34\x8e\x48\xfc\xbb\xe9\xc3\xf1\x11\x69\x11\x6b\xfb\x13\x1f\xbb\xf3\x73\xad\xb6\x20\xf2\xed\xf9\x74\xd6\xdc\x05\xfe\xc7\x1b\x3d\x8c\xcb\xda\x6d\xe4\xdd\x0d\xfc\x5a\x1a\xb3\x45\x9c\xb3\xa5\x5e\x99\x75\xb7\xd3\xa1\xda\xe5\xdf\xb6\x6e\x21\xff\xef\x6b\xc0\xee\xf5\xfb\xc2\x50\xfe\xd0\x0b\x06\x85\xa9\x4b\x18\x5b\x64\x97\xdf\x47\xbf\x37\xbe\x14\x0b\xef\xfe\x6b\xb9\xb2\x16\x5d\xee\xa5\x4a\x71\x20\x9f\x51\xa5\xe4\x78\xb6\xd0\x75\x5b\xdc\x8f\xc3\x9e\x9f\x50\x6d\x14\xf6\xc8\x3b\xb1\xf7\x10\x35\x4d\x93\x97\xa2\xc0\xc4\x0e\x0c\x5b\x50\xc1\x6e\x50\xd9\xc3\x9b\xf3\xda\x0c\x2a\xf8\xce\x03\xc1\x8d\x4b\x97\xcd\x0f\x03\x26\x16\x5d\x38\x43\xd1\xde\xe8\x42\xdb\x81\x37\x22\x47\xd0\x97\x0d\x5b\x1c\x1e\xc7\x15\x32\xf4\xa5\x27\x33\x3c\x21\x3b\xf5\x1e\xfc\xa4\x37\xee\xfb\x75\x10\x7b\xfb\x73\x8c\x24\x17\x5c\xd0\x02\xaf\x30\x35\xce\x25\x9d\x12\x69\xfc\x10\x8c\x92\xd0\xef\x07\xa1\xdf\x8b\x0b\xd3\x7f\x72\xc7\x9d\xf6\xfa\xc0\xca\x00\x9b\x1b\x54\xc0\x40\x6f\xd2\x14\xb5\x9e\x6f\x72\xc8\x09\xd4\x1e\xc1\x7f\x8b\xd9\x62\x80\x62\x41\xbe\xb8\x7e\x5f\x9c\x0a\x7b\x9f\xc9\x9e\xc9\xc0\x1f\x3d\x58\x8f\x0c\xd9\xb7\xda\x96\xc0\x95\x55\x64\x4f\x2e\x9a\xba\xda\xe3\xe4\xe2\xf4\xbd\x3c\x61\xf6\xfb\x81\x97\x84\xe3\x71\xec\x6a\x26\x7b\x74\xfd\xaf\x7a\xf5\x54\x1e\x43\x97\xf2\xc5\x18\xed\x31\x75\xa7\x81\xe2\x06\x48\xb3\x77\x1a\x06\xb0\x2e\xb6\x64\xb6\xa3\xbd\xa9\xb4\x77\x16\x8d\xec\x52\x24\x95\xc2\x91\xf4\x2d\x4a\x99\xe8\x49\x61\x58\x6a\xca\x5d\x0d\xd2\xdf\xa3\x5d\x0d\x89\x26\x51\xcf\x1b\x25\xbd\xf1\x28\xf6\x7a\xb1\x0b\xf0\xb2\x1e\xd3\x29\x13\x54\x83\x10\x10\x60\x7d\xcf\x53\x23\xf2\xbf\x19\xaa\xf8\xf3\xa9\xe0\x8e\x6d\xe0\x76\xdd\x7f\xbc\x39\xe0\x7c\xf3\xe7\x11\xab\xff\x39\xf6\xc3\x91\x37\x48\xa6\xa3\xa0\xe0\x4e\x06\x81\xcd\xef\xe5\x1e\xa0\x41\x8b\x1a\x72\xfe\x84\x50\x62\xe3\xb7\x8b\x54\xae\xde\xfc\x49\xf5\xe1\x41\xdb\x5b\x78\xc3\x84\xa4\x7d\x4d\xd5\xab\xa6\x74\x51\xbc\x90\xc0\xbd\x92\xab\x1f\x30\x50\x59\xcd\x58\x1d\xef\xc3\xf1\xf0\xd8\xe4\xc5\x45\xe2\xa4\x0b\x67\x57\xd7\x1f\x2e\x2e\x2f\x2e\xab\xdb\xda\x3d\x4e\x34\x9e\xd2\xa4\x0e\x26\x7b\xf9\xea\xda\x77\xef\xd3\xc8\xef\x85\x7e\x5c\x14\x1c\xd3\xa2\x2c\x2e\x72\x20\x45\x66\x51\x73\x70\x29\x9c\x41\x52\x85\xcc\x1e\x5e\xd9\xc5\xd8\xc8\x27\x14\xba\xd8\xf7\x06\x73\x77\x5b\xfd\xb6\x76\x21\x7d\x75\x7d\xbb\x8f\xf5\x72\x3f\xec\x6e\x25\xb9\xb6\xc5\x8d\x53\x6b\xa3\xf0\x4e\xc9\x67\x8d\xea\x73\x14\xdd\xf3\xdc\xd0\xa2\x51\xbf\x32\xb0\x05\x03\x15\xaf\x9f\x22\x3f\x4c\x3e\x47\x11\xed\xa4\x62\xbb\x86\x9c\xb8\xe3\xfa\xdc\xfe\x1c\x45\xed\x89\x92\x06\x53\xab\xfb\x12\x59\x66\x97\x16\xc7\x45\x31\x8b\xc2\xd0\xfa\x37\x92\x5a\xf0\xf9\xbc\x61\x11\xa2\xa2\x38\xf5\x47\x71\x12\x7f\x99\xf8\xc9\x68\x1c\x8d\x82\xfb\xfb\xd3\x64\x35\x30\x28\xd0\x0e\xf9\x1e\xa3\x38\x0a\x44\x9a\x6f\x32\x8c\x36\xb3\x4c\xae\x18\x17\xfa\xd4\x00\xa9\x63\x12\x8c\x7a\x83\x69\xdf\x4f\xa2\xe9\x5d\x7f\x3c\xf4\x82\x51\x3d\x17\x4b\x01\xdc\x22\xd9\x5b\x92\x0a\x8c\xb2\x07\x09\x37\x28\x23\x4c\xa5\xc8\x74\x17\xde\x5d\xbd\x7f\x77\x73\x59\x96\xcd\x75\xae\xc8\xef\x8d\x47\x7d\x22\xb8\x82\x1d\x32\xfb\xb4\x43\x3b\x31\x68\xc3\x52\x3e\x43\x2e\xc5\xc2\x42\x93\xc7\xb6\x2c\xe7\x7b\x97\x4d\x94\xfc\xb6\x8b\xa2\xc1\xa3\x1d\xed\xa9\xf1\x4c\xc2\xf1\xe7\x2f\x49\x14\x0d\x92\x47\xdf\xeb\xbf\xe4\x2e\x8b\x03\x51\x34\x38\xb4\x5b\x14\x0d\xca\x64\x7d\x0a\x9d\x70\xcb\x84\xdd\x34\x51\xe3\xfe\xe1\x99\xed\x74\x95\xc0\xed\x25\x8e\xb4\xf5\xac\x33\x96\xd6\x5c\x8a\x97\xcb\xe3\xc8\x8f\xa2\x60\x3c\x7a\xa5\x42\x3e\xa0\xcb\xac\x67\xaa\x7a\xd9\x0e\xd3\x72\xec\x4b\xd7\x06\xe9\x89\x52\xf9\x80\xf3\xa8\x5a\x3e\xc9\x58\xaf\x9b\x4f\x92\x72\x83\xf6\xcc\xb3\xe0\x08\x62\x3f\x09\x6a\x55\x4a\x24\x53\xda\x9f\x6c\xcc\xd2\xfb\x67\xa3\x90\x65\x31\x0a\x26\xcc\x98\x6d\xcc\xf2\xba\x2a\x7a\xaa\x59\x32\xee\xd9\x0d\xcb\x34\x7e\x4c\xbc\xaf\xd3\xd0\xf7\xfa\x49\xec\x8f\xbc\x51\x9c\x8c\xa9\xf1\xba\x56\x06\x1d\x5f\xee\x5b\x06\xf0\xfa\xd5\xd2\xf9\x2a\x79\x33\x61\x7d\x97\xd8\xe5\xb2\x8a\x62\x4c\xa8\xd7\xf0\x84\xbb\xef\xf1\x84\xa8\xcb\x2b\xe9\x1f\x26\x0b\x7d\x97\x67\x4f\x30\xaa\x02\xee\x7b\xb4\x2e\x2b\xff\x6f\x48\x5d\xd2\x3e\x41\xa9\x2d\xd4\xf7\x08\xdd\x6f\x8a\x85\x1f\xa7\x2c\xfe\xb3\x01\x73\xc8\xea\xf0\x20\xe8\x37\x88\x1f\xa4\x5c\xe4\x58\x0f\x9f\xd3\x7c\x0f\xe3\xf1\xc3\xc0\x6f\x46\xcd\xc9\xa0\x71\x80\x27\x42\xa6\xce\xf4\x62\xac\x34\x69\x5c\x8c\x14\x88\x2f\x44\x48\x1d\xf6\x35\x1f\x35\x91\x2b\xdf\x34\xc1\x4f\x78\x66\xfc\x64\x98\x43\xf7\x26\x41\x63\xab\x55\x47\x1f\x7f\x8c\xbd\x12\xdb\x9b\x04\xd0\x06\x92\x03\xfa\x49\xeb\xb3\x83\x7f\x01\xf7\xd5\x59\x5b\x07\x7e\x6d\xae\x5a\xba\x63\xa3\xef\x59\x5e\x34\x79\x9d\xc1\x19\xdc\x62\xbd\x60\xee\x3d\xe0\x6b\xc6\xae\x63\x56\xa6\xae\xc3\xee\x0d\x6d\x9f\x3f\x34\xeb\xe7\x28\xf6\xe2\xa0\xf7\xbd\x02\xda\x49\x56\x10\x45\x09\xed\x9e\x53\x74\x9a\x48\xe5\xd1\x5c\x51\x40\xd7\x9e\x5d\xcc\xab\x77\x18\x8d\xa2\x3e\x46\xb6\x2a\xce\x0c\x6c\xe2\xe7\x66\x67\xdb\x0a\xd8\xd8\xf7\x86\xe5\xe1\xc1\xa8\x7e\x98\xab\xcb\xce\x06\xd9\xca\xe2\x68\x73\x70\x06\x61\x50\x9b\x24\xb3\x64\xb5\xd7\x1c\xb1\x1f\xc5\x47\xe7\x12\x45\xca\x27\x90\x70\x23\x84\xbb\x70\xb4\x45\xdf\x05\xa1\x5c\x28\xdb\x78\xd1\xe7\x3a\x95\x5b\x54\xae\x4f\x1d\x30\x9c\x8e\x46\xf5\x03\xae\x98\xaf\xf0\xab\xb4\xc7\x30\xd3\xb8\x57\xf6\x0c\x86\x7e\xf2\x75\x6c\xcf\x60\xa8\x03\xfc\x23\xc5\xd1\x63\x0b\x77\x15\x13\x2b\x96\x3e\x0d\xf9\xc2\xbd\x75\x6a\x96\x46\x71\xe8\xf5\x3e\x26\xc3\xe0\x21\xf4\xe2\x60\x3c\xaa\x1d\xf0\x4d\x55\x3e\x51\x38\xe7\xdf\xf6\x1e\x9e\x86\x83\x64\x12\xfa\xf7\xc1\x67\x72\xb0\xbb\x08\xd1\x36\x9a\xed\x43\x4a\x6d\x14\x17\x8b\xb7\xd5\xb1\xda\x7a\x63\x9a\x37\x8d\x40\xcb\xeb\x7c\xde\x56\x14\x18\xd6\xa7\x3f\xd9\x03\xe9\x9f\x5b\xf6\xde\x29\xb8\xba\x15\xf5\xea\x60\x1a\xf9\x49\x70\x75\x3b\x3a\x7a\x6d\x62\xb7\x2a\xee\x62\x83\xff\x63\x7f\x38\x80\xc1\xd5\xe5\x11\xc0\xe0\xea\xf2\xc7\x01\xe2\xaf\x87\xe2\xf1\xd7\x9a\x45\xa6\xf5\xb4\x3e\x9d\xd6\x17\xf8\xe9\xb3\x5e\x70\x5f\x64\x6b\xc9\x45\x6d\x4e\x4c\x3f\x45\x0f\x41\xe2\x8f\xfa\x93\x71\x30\xa2\x69\x51\x76\x71\x2f\x96\x48\xa8\x14\x6e\x1e\x94\x39\xc1\x53\xb7\x0f\x0d\xa1\xa1\xdb\xb4\x37\x84\x86\x6e\xb3\x3e\x2c\xf7\xe8\x0d\x81\xe2\x52\xae\x21\x30\xf1\xa2\xe8\xf0\x2e\xa2\x29\xd4\xd8\xa9\x16\x42\xb5\x13\xff\xba\xc0\xa7\x25\x37\x28\x24\xd7\xcd\xaa\xeb\xd3\x63\x10\xfb\xa3\x71\x10\x1d\x2c\x3f\x4b\xb6\x45\x78\x5e\x1a\xee\x64\xc0\x3e\x94\xf9\xce\x34\xff\xf4\xb4\x34\xab\xdc\xc8\x75\x46\x3b\x8b\xce\x46\xab\x8e\xdd\xf1\x75\x66\x5c\x74\x9e\xf7\x1f\x0b\x7d\x3f\x7d\x7c\x8c\x87\x83\x78\x3c\xe9\xd3\x1e\xa3\x76\x16\xd7\x3a\x87\x78\x89\x1a\xed\x83\x60\xb9\x2e\x6e\xca\x8a\xf7\xc4\xc5\x8b\x40\xfb\x60\xc3\x6d\xb9\xf2\x7c\x07\x02\xb7\xa8\xec\xc3\x9a\xea\x15\x2a\x13\x0b\xcc\xec\xf5\x39\xc1\xd8\xd8\x7f\x5e\xf2\x1c\xed\x93\x33\xaa\x1d\xdd\xd3\xf1\xd6\x39\x28\x46\x1b\x59\xca\x38\x02\x58\x6a\x36\x2c\xaf\x5e\x35\xb6\xc6\x6b\x37\x37\x5b\x00\x8f\x98\xaf\x5d\x5a\x5c\x1a\xb3\xd6\xdd\x4e\x67\x61\x1f\xfe\xd2\xf6\xb7\xb3\x32\xa8\xd9\x46\xc9\x8e\x43\x2d\x86\x68\x1f\x6a\xda\x97\xb6\x2f\x0a\xee\x2d\xd8\x39\x7a\x12\xd8\x61\x2a\x5d\xf2\x2d\x76\xf6\x70\xbd\x5c\x8a\xff\x1b\xd8\xc5\x82\x9b\x3d\xce\x17\xa6\xc4\xc3\xe4\xa1\x06\x93\xe5\x17\x3b\xa6\xc4\xfa\x69\x61\xa1\x32\x9c\x71\x26\x3a\xeb\xcd\xec\x09\x77\x17\x8b\xf5\xa2\x29\x1b\xe2\x5a\xda\xd7\xc1\x33\x78\x1d\x80\x62\x86\x56\x54\xda\xb0\xed\x21\x46\x32\x3b\x18\x45\x86\xb3\x0b\x21\xb3\xa2\x7c\xb4\x08\x1a\xcd\x66\x9d\x5c\x5d\x5f\x7c\xdb\x0b\xfa\xab\x59\x79\x0a\xd9\x02\x28\x56\xe1\x56\xeb\x7f\x02\x00\x00\xff\xff\xab\x9b\xfc\xea\x1c\x30\x00\x00")

func embdDojoconfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_embdDojoconfigYml,
		"embd/dojoConfig.yml",
	)
}

func embdDojoconfigYml() (*asset, error) {
	bytes, err := embdDojoconfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "embd/dojoConfig.yml", size: 12316, mode: os.FileMode(0664), modTime: time.Unix(1602635062, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0xf1, 0x2c, 0x20, 0x48, 0x85, 0x8c, 0xcf, 0x33, 0xf6, 0xeb, 0x15, 0x9c, 0xdc, 0x10, 0xf7, 0x33, 0x9b, 0x23, 0xee, 0x4b, 0x91, 0xb5, 0x50, 0x7a, 0x59, 0x33, 0x6c, 0xab, 0xdf, 0xae, 0x1a}}
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
	"embd/dojoConfig.yml": embdDojoconfigYml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"embd": &bintree{nil, map[string]*bintree{
		"dojoConfig.yml": &bintree{embdDojoconfigYml, map[string]*bintree{}},
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
