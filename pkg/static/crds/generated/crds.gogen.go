// Code generated by vfsgen; DO NOT EDIT.

package generated

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// CRDs statically implements the virtual filesystem provided to vfsgen.
var CRDs = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 12, 30, 3, 47, 46, 163018444, time.UTC),
		},
		"/devops.k8s.io_clustercredentials.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_clustercredentials.yaml",
			modTime:          time.Date(2020, 12, 30, 3, 52, 45, 668053745, time.UTC),
			uncompressedSize: 3261,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x56\x4d\x6f\x1b\x37\x10\xbd\xef\xaf\x18\xa4\x87\x5c\xba\x2b\x07\x6d\x81\x76\x6f\xae\x9c\x02\x86\xdb\xc0\x88\x0d\xa3\x40\xd1\x03\x45\x8e\x24\xc6\xbb\x43\x76\x66\x28\x44\xfd\xf5\x05\xb9\x2b\x6b\xe5\xaf\x28\x75\x92\xbd\x2d\x3f\xde\x1b\x3e\xce\xcc\x63\x55\xd7\x75\x65\xa2\xbf\x41\x16\x1f\xa8\x05\x13\x3d\x7e\x54\xa4\xfc\x27\xcd\xed\xcf\xd2\xf8\x30\xdb\xbc\xa9\x6e\x3d\xb9\x16\xe6\x49\x34\xf4\xef\x51\x42\x62\x8b\x67\xb8\xf4\xe4\xd5\x07\xaa\x7a\x54\xe3\x8c\x9a\xb6\x02\x30\x44\x41\x4d\x1e\x96\xfc\x0b\x60\x03\x29\x87\xae\x43\xae\x57\x48\xcd\x6d\x5a\xe0\x22\xf9\xce\x21\x17\xf0\x1d\xf5\xe6\xa4\xf9\xb1\x39\xa9\x00\x2c\x63\xd9\x7e\xed\x7b\x14\x35\x7d\x6c\x81\x52\xd7\x55\x00\x64\x7a\x6c\xc1\x76\x49\x14\xd9\x32\x3a\x24\xf5\xa6\x93\xc6\xe1\x26\xc4\x5d\xbc\x95\x44\xb4\x99\x7a\xc5\x21\xc5\x16\x0e\x27\x07\x94\x31\xb4\xf1\x58\x03\xe0\xfc\x0e\xb0\xcc\x75\x5e\xf4\xe2\xf1\xf9\xdf\xbd\x68\x59\x13\xbb\xc4\xa6\x7b\x2c\xa4\x32\x2d\x9e\x56\xa9\x33\xfc\xc8\x82\x0a\x40\x6c\x88\xd8\xc2\xbb\x1c\x4e\x34\x16\x5d\x05\x30\xaa\x51\xc2\xab\xc7\xf3\x6e\xde\x0c\x60\x76\x8d\xbd\x19\xe2\x06\x08\x11\xe9\xf4\xf2\xfc\xe6\x87\xab\x83\x61\x00\x87\x62\xd9\x47\x2d\x9a\x3e\x88\x1c\x18\x6d\x60\x27\xa0\x6b\x84\x7d\x34\xe0\x69\x19\xb8\x2f\xb2\x03\x21\x3a\x74\xa0\xe1\x0e\x13\xc0\x58\x8b\x32\xee\x1a\x30\x9b\xbb\xd9\xc8\x21\x22\xab\xdf\x89\x3a\xee\xd8\x67\xd5\x64\xf4\x5e\x7c\xaf\xf3\x11\x86\x55\xe0\x72\x3a\xe1\xc0\x31\xca\x80\x6e\x3c\x35\x84\x25\xe8\xda\x0b\x30\x46\x46\x41\x1a\x12\xec\x00\x18\xf2\x22\x43\x10\x16\x1f\xd0\x6a\x03\x57\xc8\x19\x06\x64\x1d\x52\xe7\x72\x16\x6e\x90\xb5\x08\xb0\x22\xff\xef\x1d\xb6\x80\x86\x42\xda\x19\xc5\xf1\x5e\xf7\x9f\x27\x45\x26\xd3\xc1\xc6\x74\x09\xbf\x07\x43\x0e\x7a\xb3\x05\xc6\xcc\x02\x89\x26\x78\x65\x89\x34\xf0\x47\x60\x2c\x8a\xb6\xb0\x56\x8d\xd2\xce\x66\x2b\xaf\xbb\x6a\xb2\xa1\xef\x13\x79\xdd\xce\x4a\x61\xf8\x45\xd2\xc0\x32\x73\xb8\xc1\x6e\x26\x7e\x55\x1b\xb6\x6b\xaf\x68\x35\x31\xce\x4c\xf4\x75\x09\x9d\x4a\x45\x35\xbd\xfb\x8e\xc7\xfa\x93\xd7\x07\xb1\xea\x36\x67\x93\x28\x7b\x5a\x4d\x26\x16\x21\xa8\x28\x9b\x78\x1d\x6e\xf1\xb9\xbb\xf8\x2d\x30\xe4\xda\x34\xae\x87\x5c\xd7\x10\x18\x3e\x04\x4f\xc7\x90\x58\x33\x47\xd6\x4f\x80\xdb\x40\x94\x35\x9b\x24\xd1\xc1\x86\x21\x03\x5b\x58\x6c\x15\x8f\x23\xbd\xc0\x6d\xfb\x32\x88\x9c\xb7\x4b\x6f\x8d\xe2\x03\xac\x2f\x27\x0e\xb2\xca\xaf\x9e\x0c\x6f\xcf\xc6\x1e\x39\x29\x13\xe7\x4a\x0b\x35\xdd\xe5\xa3\x65\xf4\xec\xa9\x9e\xa4\xdc\x4d\x0c\xb5\x30\x8d\xa5\xf3\x48\x7a\xc4\x65\xe5\xc3\xd6\x26\x7a\x29\x55\x04\x7f\xfe\x74\xf2\x0b\x98\xa4\xeb\x97\xc9\x5d\xd8\x8f\x53\xfa\x2b\x90\x97\x94\xcb\xed\xb6\x3d\x66\x3d\xaa\x75\xa7\x97\xe7\xf3\x27\x14\xfb\x5c\xfa\x03\xb8\x17\x27\x6e\x46\x9b\x9f\x1e\x71\x8f\xd7\x17\x6f\xc1\x13\xac\xba\xb0\x28\xdd\x3f\x09\x7e\x01\xe2\x97\xc7\xff\x51\xff\x5f\x2d\x7c\x6e\xc2\x17\x87\x7f\xc6\x80\xb2\xc3\x83\x17\x30\x23\xe6\xd0\xc4\xf7\x3e\x93\x87\x72\xc3\x7a\xff\xf6\xea\x1a\x76\x9d\xb7\x78\xd1\x7d\xf3\x29\xcc\xfb\x8d\xb2\x77\xa0\xec\x17\x9e\x96\xc8\x83\x87\x2d\x39\xf4\x05\x13\xc9\xc5\xe0\x69\xd7\x11\x73\x62\xdc\x03\x95\xb4\xe8\xbd\x66\xdb\xfb\x27\xa1\x68\xb6\xaa\x06\xe6\xe5\x85\x05\x0b\x84\x14\x9d\x51\x74\x0d\x9c\x13\xcc\x4d\x8f\xdd\xdc\x08\x7e\x75\xff\xc9\x4a\x4b\x9d\x85\x3d\xce\x81\x72\x35\x7f\x9b\xcb\xee\x0d\xf9\x65\xd6\xe9\x1b\xd1\x4d\x5e\xbd\x9f\x5c\xac\x48\x86\xf4\xfc\xec\xa8\xde\xa3\x47\x7a\xf5\xa4\x49\x96\x2d\x0f\xbb\xe4\xa3\x04\x39\x9d\x3c\xe3\xa4\x30\xea\x69\x7b\x9c\x8c\xee\xa2\xae\x9e\x3c\x5d\xa1\x77\x2d\x28\xa7\x61\xa3\x68\x60\xb3\xc2\x71\x44\xd4\x68\x2a\x4a\xe7\x27\x64\x54\x74\xef\xee\x3f\xbf\x5f\xbd\x3a\x78\x4b\x97\x5f\x1b\x68\xb8\x2f\x69\xe1\xaf\xbf\xab\x01\x15\xdd\xcd\xee\x79\x9c\x07\xff\x0b\x00\x00\xff\xff\xc1\x87\xc4\x67\xbd\x0c\x00\x00"),
		},
		"/devops.k8s.io_clusters.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_clusters.yaml",
			modTime:          time.Date(2020, 12, 30, 3, 52, 45, 722049159, time.UTC),
			uncompressedSize: 20489,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x3c\x5b\x73\xdb\x36\xd6\xef\xfa\x15\xe7\xcb\xf7\x75\x6c\x7f\xb5\xe8\x74\xdb\xd9\xe9\xea\xa5\xe3\xb5\x9d\xd6\x93\xc4\xd5\x58\x4e\x5e\xd2\xec\x0c\x44\x1c\x49\xa8\x40\x80\x01\x40\xd9\xec\x66\xff\xfb\x0e\x2e\xa4\x44\x8b\xa4\x48\xf9\xd6\x87\xfa\xc9\x22\x0e\x0e\x70\xee\x17\x80\x1c\x0c\x87\xc3\x01\x49\xd9\x47\x54\x9a\x49\x31\x02\x92\x32\xbc\x33\x28\xec\x2f\x1d\x2d\x7f\xd4\x11\x93\x27\xab\xef\x06\x4b\x26\xe8\x08\xce\x32\x6d\x64\x72\x8d\x5a\x66\x2a\xc6\x73\x9c\x31\xc1\x0c\x93\x62\x90\xa0\x21\x94\x18\x32\x1a\x00\x10\x21\xa4\x21\xf6\xb1\xb6\x3f\x01\x62\x29\x8c\x92\x9c\xa3\x1a\xce\x51\x44\xcb\x6c\x8a\xd3\x8c\x71\x8a\xca\x21\x2f\x96\x5e\xbd\x8e\x7e\x88\x5e\x0f\x00\x62\x85\x6e\xfa\x0d\x4b\x50\x1b\x92\xa4\x23\x10\x19\xe7\x03\x00\x41\x12\x1c\x41\xcc\x33\x6d\x50\xe9\x88\xe2\x4a\xa6\xc5\x2e\x07\x3a\xc5\xd8\x2e\x38\x57\x32\x4b\x47\x50\x1d\xf4\x73\xc3\x86\x02\x31\x1e\x8d\x7b\xc2\x99\x36\x6f\x37\x9f\xbe\x63\xda\xb8\x91\x94\x67\x8a\xf0\xf5\xa2\xee\xa1\x5e\x48\x65\xae\xd6\x08\x87\xb0\x8a\xfd\x00\x13\xf3\x8c\x13\x55\xc2\x0f\x00\x74\x2c\x53\x1c\x81\x03\x4f\x49\x8c\x74\x00\x10\x88\x76\xd3\x87\x40\x28\x75\x6c\x24\x7c\xac\x98\x30\xa8\xce\x24\xcf\x12\x51\x22\xa7\xa8\x63\xc5\x52\xe3\xd8\x74\xb3\xc0\x02\x39\x50\xa1\x2f\xc7\x91\x83\x02\xf8\x5d\x4b\x31\x26\x66\x31\x82\x48\x1b\x62\x32\x1d\xb9\xe1\x30\xea\x59\x77\x7e\x35\x29\x9f\x98\xdc\x6e\x4b\x1b\xc5\xc4\xbc\x69\xa1\xb0\x4f\x90\x33\xb0\x62\x53\x02\x0d\xea\xe6\x05\xa3\x00\x5f\x59\xf3\xe3\xc5\xf5\xe4\xf2\xd7\xab\x1e\xab\xc6\x3c\xb3\xd4\xa5\x0b\xa2\xb1\x79\x31\x37\x5c\x59\x69\xfc\xcb\xe9\xe4\xa2\xe3\x3a\x07\x67\xf7\xb5\x0c\x98\x06\x02\xa6\xfc\xa9\x30\x55\xa8\x51\x18\x26\xe6\x60\x16\x08\x1a\xd5\x0a\x95\x83\x08\x8b\x00\xdc\x2e\x50\x80\x59\x30\x0d\x72\xfa\x3b\xc6\x06\x6e\x89\xf6\x0a\x8c\x34\x82\x83\xed\xcd\x17\x96\x12\x6d\x69\x79\x85\x94\xd3\x9f\xab\x84\x50\x62\xfc\xa2\x7e\x78\xf5\x9d\x57\xb7\x78\x81\x09\x19\x05\x48\x99\xa2\x38\x1d\x5f\x7e\xfc\x7e\x52\x79\x0c\x55\xc2\x83\x82\x5b\x6a\x2d\x51\x1e\x16\x66\x52\xb9\x9f\xc5\xe8\xe9\xf8\xb2\x9c\x9e\x2a\x99\xa2\x32\xac\xd0\x76\xff\xb7\xe1\x35\x36\x9e\xde\x5b\xec\xc0\xee\x27\xe8\x10\xb5\xee\x02\xfd\xaa\x41\x4f\x90\x06\x12\xac\x82\x39\x2e\x96\x4c\x77\xbc\xa9\x20\x06\x0b\x44\x44\x60\x74\x04\x13\x27\x0e\x6d\x8d\x31\xe3\xd4\x7a\x99\x15\x2a\x03\x0a\x63\x39\x17\xec\x8f\x12\xb7\x06\x23\xdd\xa2\x9c\x18\x0c\x56\xbd\xfe\x73\xf6\x26\x08\x87\x15\xe1\x19\x1e\x03\x11\x14\x12\x92\x83\x42\x27\xce\x4c\x6c\xe0\x73\x20\x3a\x82\xf7\x52\x21\x30\x31\x93\x23\x58\x18\x93\xea\xd1\xc9\xc9\x9c\x99\xc2\x5b\xc6\x32\x49\x32\xc1\x4c\x7e\xe2\x1c\x1f\x9b\x66\x46\x2a\x7d\x42\x71\x85\xfc\x44\xb3\xf9\x90\xa8\x78\xc1\x0c\xc6\x26\x53\x78\x42\x52\x36\x74\x5b\x17\xce\x63\x46\x09\xfd\x5f\x15\xfc\xab\x3e\xa8\xec\x75\x4b\xa3\xfd\x9f\x73\x66\x2d\x12\xb0\x6e\xcd\xab\xb6\x9f\xea\xa9\xd8\xd6\xee\xeb\x8b\xc9\x0d\x14\x4b\x3b\x61\xdc\xe7\xbe\x57\xf0\x72\xa2\x5e\x8b\xc0\x32\x8c\x89\x99\x35\x0e\x2b\xc4\x99\x92\x89\xc3\x89\x82\xa6\x92\x09\xe3\x7e\xc4\x9c\xa1\xb8\xcf\x7e\x9d\x4d\x13\x66\xac\xdc\xbf\x64\xa8\x8d\x95\x55\x04\x67\x2e\x84\xc0\x14\x21\x4b\xa9\xb7\xa4\x4b\x01\x67\x24\x41\x7e\x66\x5d\xc2\x53\x0b\xc0\x72\x5a\x0f\x2d\x63\xbb\x89\x60\x33\xfa\xdd\x07\xf6\x5c\xdb\x18\x28\xc2\x54\x83\xbc\x82\x01\x4e\x52\x8c\x2b\x16\x43\x51\x33\x65\x75\xda\x10\x83\xd6\x12\x36\xc3\x57\xbb\xa5\x06\x6b\xf5\xc2\xba\xb8\x33\x8a\x9c\xaa\xf9\x16\x04\x54\xc2\x50\x13\x9e\x16\x2e\xb4\x52\xed\xf2\x00\xbf\xe3\xb3\xcb\xf3\xeb\xd1\xa0\x07\xd2\x30\xef\xc6\x42\xf4\x9a\x57\xe6\x1d\xef\x89\x20\xf3\x97\xa5\x5d\xb1\xfa\xfd\x57\x84\x7f\x9d\x09\x1b\x5d\x2c\x64\x45\xf8\xca\x3f\x77\x62\x97\xc2\x10\x26\x50\x45\x7d\x58\x41\x99\x4e\x39\xc9\x6d\x0e\xd2\x8b\x85\x54\xe8\x73\x99\x10\x26\x76\x6c\xfc\xfc\x6a\xe2\xe1\x8a\xb0\x42\x85\x06\xea\x9f\x64\x1a\x29\x4c\x73\x58\xfe\xa8\x5d\x08\x65\xb1\xf5\xa1\xe7\x38\x23\x19\x37\xba\x8e\xc7\x12\x5e\x05\x99\x47\x5c\xc6\x84\xbf\xea\x47\xab\x8c\x97\x2f\x2a\x6b\x34\x31\xdd\xc1\xaf\x0b\x13\x53\x58\x48\x4e\xb5\xd5\xd2\x19\x9b\x67\xca\x05\x3c\x17\x87\xed\xfc\x6d\x8a\xd3\xd6\xbd\xda\xa4\xdd\x86\xb1\xba\xb1\xfb\x6b\x07\xd0\xf0\x74\x8a\x1a\x16\xf2\xd6\x72\x3d\x96\x42\x58\x17\x6f\xa4\x8d\xb3\x05\xca\x5a\x8c\x9e\xca\x32\x11\x7d\x67\xc5\xe4\x62\x67\x89\x9d\x28\x84\x24\x33\x19\xe1\x3c\x07\xbc\xb3\x90\x6c\x85\xb5\xc8\xda\x49\x73\xe6\x43\xde\x30\x8e\x4d\xa3\xf7\x5d\xe8\xa9\x05\x76\x31\x4f\xc0\x64\xf2\x0e\xce\x2c\xf2\x19\x8b\xad\xe7\x3c\xcd\xcc\x42\x2a\x66\x72\x98\x59\x20\xab\x9c\x8d\x58\x9d\x2a\x6a\x8c\x33\x85\x81\x5c\x1f\x59\x62\x27\xab\x08\xae\xf1\x4b\xe6\x9c\x32\x9b\x41\x66\x53\x7f\x20\x70\xf3\x6e\x52\xf0\xd1\xc2\x34\xe2\x6e\x55\xae\x40\x34\x2a\xd3\x87\xec\x00\xbe\x41\x78\x5c\x12\xee\x74\xab\x20\x18\x8c\x6c\xa1\xf9\xe5\x08\x2e\xd2\x05\xdd\x91\xe2\x8b\x02\xde\x3a\x46\xb7\xdf\x04\x93\xa9\xad\x0c\xd7\x3b\xb5\x06\x55\xe8\xe4\x45\xad\x61\x95\x99\xa0\xc1\xa4\x65\xe5\x4e\x14\x14\x40\x44\x29\x92\x37\xc0\x2c\x31\xef\x21\xd5\xb7\x1e\x7a\x43\xa8\x4b\xcc\x2b\xa2\xdc\x14\x58\xcb\xee\x9f\x53\x94\x2a\x20\xaf\xa7\x71\x18\xcc\xb9\x69\x30\xe8\x71\xc3\x70\xa9\x24\x0d\xe3\x81\xbd\x83\xe6\x8d\xd7\x3a\x6d\xd7\x05\xb0\x5e\xac\x83\x07\xf5\xde\x2e\x55\x72\xc5\x28\xde\xf7\xe0\x4b\x21\xa7\xda\xa9\x5d\xf1\xbc\x59\x5d\x5c\x55\xe2\x90\x39\xed\x65\x42\x1b\x22\x62\x7c\x72\x77\x6a\x93\xd5\x73\xa6\x3a\xaa\xe0\xb9\x87\x2e\x43\x3b\x53\x18\x1b\xa9\x72\xbf\xe9\x5b\xc6\x39\xa4\x9c\xc4\x08\xac\x41\x28\xeb\x45\xd7\x71\xdf\x45\xf9\x93\x15\x51\x27\x9c\x4d\x4f\x2c\xa6\x57\x0f\xf3\x1d\xcd\xf1\xbe\x6f\xdc\xef\x65\xef\xf7\x43\xab\xdf\x84\x13\x97\xdb\x12\x10\x35\xcf\x12\x5b\x2e\x15\x0a\x43\x43\x3d\xda\xb2\xb2\x63\xec\x94\x09\xa2\x72\xdf\x61\x50\x99\xb0\xda\xc1\x28\xba\x3a\x8e\x18\x16\x43\x2a\xe9\x2e\x8e\x35\x6a\xba\x53\x13\x44\x65\x63\xc6\xe4\xf4\xaa\xab\xc3\x1d\x6f\x4c\x01\x8d\x46\x07\x1a\x27\x99\xaf\x0d\x4f\xb9\xd3\x56\xc3\x56\xe8\xfb\x5d\x2d\x34\x16\x1d\x07\x47\xab\xdd\x0b\x68\x36\x17\xd6\x11\x59\x07\xf0\xf2\x6e\xda\x77\x7b\x7a\x32\x68\x52\x99\xb4\x83\x45\x2d\x34\x38\xe6\x55\x59\x14\xba\x4f\x7f\x26\x26\xed\x72\xf3\xc1\xcd\xf4\x77\xc5\x2d\x83\x33\x24\xb6\x6c\xd7\x3b\x12\xec\x50\x1d\xbf\xf1\xd0\xae\x29\xa4\xa8\xf7\x5f\x05\x06\x30\x0b\x62\xbc\xa1\x0a\x32\xe5\xb5\x79\xe0\x34\x0f\xad\x0b\x5f\x8d\xf4\x4d\xca\x1d\xde\xf7\xc4\x15\xf4\xf1\x02\x69\xd6\x14\xf6\x3d\xc1\x53\x29\x39\x12\x51\x03\x61\xe3\x7d\x83\x3c\x5b\x45\x9d\x76\x70\x74\x54\x9b\x07\x6b\x8a\x56\xf1\x03\x71\xb4\xeb\x92\xd3\x26\x6d\x1a\xc7\xb4\x8a\x07\x7b\x3a\xc2\x76\x25\x9f\xa7\x59\x7d\xdd\xbe\xa5\x71\x3f\x8f\x3f\x6c\xd5\xed\xf3\x34\x73\xf8\x6d\x7e\xda\xa8\x43\x1d\x18\xb4\x20\xa3\x7d\x23\xfd\xb2\x25\xd1\x4c\x3b\x85\xc1\x15\x4b\xdb\x86\x3b\x6a\xc8\x2e\xf9\xba\x43\x14\x96\x3e\x24\xa0\x99\x05\x53\x74\x4c\x94\xc9\xff\x1c\x24\x03\xac\x52\xa9\x4c\x3b\xa6\x99\x54\x09\x31\x23\x60\xc2\x7c\xff\xb7\x0e\x6b\x32\x61\x70\x8e\xea\xc9\xf8\x3c\xf4\x9b\xde\x5f\x0e\x3b\x00\x16\x52\x2e\x1b\x78\xdf\x27\x3f\xdb\x29\x80\x1d\xdb\x28\xda\xfe\xef\xfe\xb9\x9f\x43\x66\xe9\x4a\xef\x37\x33\xcd\xa6\x9c\xc5\xfb\xae\xab\x97\x2c\x3d\x93\xc2\x33\x6a\x9f\x88\xd0\x91\x71\xf5\xfe\xb0\x2d\x2e\x33\x41\x38\xfb\x03\xd5\xae\xc8\xfc\xa6\x04\x0c\x35\xad\x4c\xc9\x97\x0c\xdd\x11\xac\xf5\x93\xfe\x90\xc5\x07\xe7\x24\xd3\xee\x18\x00\x93\xd4\xe4\xf5\x9d\xc2\x14\x55\x42\x04\x0a\xc3\x73\x50\x98\xc8\x15\x16\x07\x15\xee\x14\x42\x1b\xa9\xc8\x1c\xb7\xbd\x6e\x23\x93\xea\x37\x6b\x13\xb2\xa2\x00\x12\xee\x7f\x8a\xc2\xb0\x59\xee\x2b\xe7\x92\x7a\xa0\xcd\xf5\x5e\xd1\x2c\xe3\x6c\x86\x71\x1e\x73\x8c\xf6\x6b\x3a\xd6\xc9\x66\x99\x4d\x91\xa3\x79\xc1\xae\x67\x42\xe2\x85\x0d\x7c\xa3\x3d\x59\x1d\x32\xb6\xf7\x1e\x4d\xc1\xeb\xc4\x25\x4f\x05\x72\x1f\x5c\x1d\x19\x20\x67\x6d\x4c\xae\x63\xed\x2e\xdf\xdf\xec\xf4\x77\x1a\x0d\x27\x53\xe4\x8d\x36\xd7\xaf\xf0\xec\x10\x5b\x76\x3a\xe1\x94\x68\x3d\x5e\x28\xa2\x1b\xa3\x7f\x11\x77\xa6\xb9\xc1\x7d\xa9\xb6\xab\xdc\x4a\x45\xf7\x66\x5b\x5b\x78\xec\x12\x18\x77\x87\xc4\x54\xb1\x15\x31\xf8\x16\xf3\xa7\x64\x84\x21\x6d\x5d\xcb\x8a\x9a\x5f\xce\xdc\x09\x20\x9b\x31\xa4\xc7\xde\x9d\x48\x8a\x07\x3a\xe0\x68\x2a\xe9\x76\x14\x74\x5b\x57\x28\x2c\x52\x7f\x18\x7b\x63\xf1\x3a\x57\x6b\x0c\xb1\x45\x88\xf5\x9a\x0b\xe2\xcd\xeb\x15\xce\x66\x18\x9b\x57\x2d\x89\x87\x14\x40\x44\x0e\xa9\xa4\xde\x27\x53\x89\x1a\x84\x34\x60\x24\x47\x45\x0c\x3a\x44\x6e\x95\xe8\x81\x69\x97\xdf\x4c\x7b\xbe\x54\x3d\x2f\x0b\xc9\x4e\xe4\x68\xf6\xd3\xfd\x55\x02\xf4\xfc\xb4\xbb\x4f\x25\xd5\xad\x28\xa1\x20\x6c\x9b\x2c\x87\x24\x82\x8f\x84\x33\x1a\xf0\xfb\xde\xce\x95\x2c\x0a\xba\xe3\x1d\xb8\xc7\x0a\x67\xa8\xd6\xf0\xae\xad\x77\x25\x2f\xee\x30\xce\x0c\x46\x8f\x91\x66\x2e\x31\xdf\x9b\x69\x9e\x4d\x4b\xcc\xad\x5a\x4c\x11\x48\x9a\x72\xb6\xeb\x88\xc0\x39\x35\xa7\x63\x8f\xb2\x7f\xc3\x12\x3c\xa5\xb4\x2d\x69\xdd\x56\xf2\x62\xce\xc6\x85\x04\x2f\x32\x96\x20\x10\x03\xb7\x0b\x16\x2f\x76\xf4\xd9\x4a\xdb\x75\xd7\x76\x88\x45\x17\xc1\xa5\xb3\x16\x29\x78\x0e\xb7\x8a\x19\x83\xfe\x54\xae\x14\xd9\x0e\x4b\xad\x7a\x15\x4a\x0c\x0e\x2b\x37\x86\x1e\x52\x4e\xd8\xd4\xa8\x0f\x8f\x4a\xf9\xfa\x3b\x1f\xb1\x54\x0a\x75\x6a\xd3\x47\x31\x2f\x6e\xc5\x38\x80\x1d\x3c\x5a\x62\x1e\x3d\x4f\x09\xe8\x6d\xac\x05\x60\x89\xf9\x83\x6a\xc4\x1d\x2d\xad\x4c\xdb\xb2\x20\xc1\x3d\x43\x5a\x1b\x89\x43\xa8\x2d\xbb\x86\xd0\x50\x6f\x0d\xcb\xcd\x0c\xf6\x6a\x9f\xd5\x11\x29\xd0\xdc\x4a\xb5\x7c\x99\x0c\x31\x2c\x7e\x8e\x2b\x16\xf7\xbb\x86\x10\x66\x76\xb8\x41\x71\xb5\x86\xac\x74\x62\x02\x86\xdd\xdd\x98\x96\x5d\x48\xdd\x61\x03\xbf\xa6\x36\x7a\x30\x31\x9f\xe4\xda\x60\x52\xd9\x84\x2c\xc6\xec\x0e\xb4\x1b\xef\xb5\x81\x94\x64\x7a\xd7\xfa\x63\x0b\x33\xe8\x53\x56\xb6\xc5\xe6\xba\x0c\x3d\x28\x47\x5e\x69\xaa\x12\xe3\x2f\x5a\xf9\x6b\x59\xd6\xf7\x35\xe6\xe8\x0f\x68\xa8\x26\xe4\xae\xb8\x1f\xe5\x2f\x92\x5c\x65\x49\xbd\xa5\xee\xce\x1e\x77\xe5\x8e\x09\xb9\xbb\x92\x14\xc7\x92\x3e\xe1\x22\x72\x85\x4a\x4b\x4e\xaf\x2d\xbf\x5e\xba\x2b\xd2\x32\xe8\x9b\x16\x1b\x67\x17\x1b\x97\x9f\x3b\xa5\xa9\x7b\x96\xb5\x3a\x24\x4d\x2f\x79\x9d\x27\xdc\x59\xaa\xbf\xb6\xb6\x75\xfa\x13\x20\x6d\x0a\xb1\x3e\x9f\x37\x40\x40\x63\x4a\x6c\x5e\x49\xc1\x8d\xdb\x94\x62\xe3\x46\x54\x5d\x16\xc9\xcc\x81\x5e\x1f\xf5\xc2\x2d\x33\x0b\x78\x5f\x63\x01\xbd\x7c\x88\x41\x41\x84\xb9\x3c\xef\xe5\x7f\xb3\x74\xae\x08\xdd\xe5\x7a\x3e\x78\xa8\xe2\xde\x5d\x31\xcb\xda\x74\x8c\x5a\xf7\x36\x76\x49\xbb\xf4\xdd\x8b\x55\x2d\xf8\xb1\x75\xb7\x24\xe3\x45\xc2\xc3\x34\x9c\x66\x46\xee\xd5\x73\xd7\xc6\x4a\x6b\x9e\xf7\xd8\x42\x31\x25\x9c\xfc\x47\x7b\xf6\xec\x13\x72\xf7\x41\x28\x24\xb4\x25\xab\x27\x22\xff\x75\xd6\x96\x48\x75\xeb\x16\x0f\xf7\x38\xe4\xbe\x71\x6d\x99\x3b\x96\x64\x09\x88\x2c\x99\xa2\xb2\xe1\xcc\x56\x59\xbe\x8e\x8a\x89\x70\x57\x77\x3d\x09\x6d\x17\x02\x5c\xab\xcc\xc5\x83\xa0\x2a\x11\xbc\xfe\x06\x12\x24\x42\x03\xe1\xdc\xe3\x14\xe8\xad\x68\x8a\xe0\x10\x02\x99\x99\x16\x8a\x00\x70\xc5\xfc\x15\x97\xef\x5e\x97\xd8\xd8\x5c\x48\x85\x45\x21\xab\x8b\xcd\x85\x22\x21\x21\x39\x4c\xdb\x32\x74\x67\xc9\x4c\x80\x14\x08\x36\x8d\x40\xe5\x8a\x9f\x63\x3b\xe0\xab\xec\x98\x28\x9c\x65\x9c\xe7\xff\x53\x55\xc1\x16\xa4\x4c\xc3\xeb\x6f\x9a\xd3\xea\xbb\xe1\xfa\x8d\x8b\x21\x13\x66\x28\xd5\xd0\x8b\x69\x04\x46\x65\xf8\xa8\xfe\x7d\x55\x77\x93\xbf\xd5\x48\xea\x33\xdd\x61\xe9\x5f\x3a\x5d\x85\x76\x2f\x72\x74\xb8\x0c\xed\xe0\x36\x0b\xbd\xcd\x24\x83\x4c\x65\xe6\x6f\x98\x7b\x7c\xdb\x59\x07\xa9\xcf\x37\x5a\x6e\x4b\x53\xaa\x50\xeb\x9d\x19\xd1\xbb\xd0\xb3\x2e\xe1\xad\x8a\xc6\x0b\x32\xe5\x58\x94\x57\x8d\x99\x4e\xaf\xae\xe8\xa9\x5f\xc0\xb9\x56\xc2\x44\x95\x01\xc5\x35\x81\xb0\xd4\x81\x6e\xca\x23\x54\xad\x17\xde\xed\x92\x16\xb2\xf9\x88\xb8\xf1\xdd\xa4\x96\xf5\xfe\x24\xbd\x41\xd3\x78\xb0\xdb\xf0\xca\x4c\x20\xc9\x4d\x3c\x76\xce\x40\xce\x60\xec\x52\xa3\xe3\xf2\xa6\x56\xf9\x7e\x55\x4d\xb6\xa7\xe0\x52\x14\x50\xd1\x53\x54\x98\x56\x52\xfd\x6a\x4c\xbb\xdc\xe3\xd5\x97\x7b\x5f\xf1\x97\x49\x2a\x05\xd6\x36\x51\x7b\x19\xca\x59\x81\xa8\x52\x9d\xac\xc3\x54\x2c\x53\x86\xfe\xbe\x28\x89\x17\xf5\x27\x08\x25\x8a\xd0\x49\x2c\x34\xda\x9f\x45\xec\x63\x40\x0a\x53\xce\x62\xa2\xfb\x68\x5b\x49\xc9\x75\x98\x5c\x4f\x51\xa3\xb2\x55\x29\x5d\xbf\x37\xe2\x7e\xed\x41\x63\xd7\x36\x2e\x59\x11\xc6\xad\x07\x1c\x0d\x1e\x7e\xea\xdd\x2d\x8b\x89\x33\xa5\x50\x98\xe7\x5b\x30\xbc\x8c\xf3\x7c\x0b\x86\xf7\xa0\x9e\x6b\xc1\xdd\x97\x71\x4a\x29\x37\x42\x04\xa1\x34\x5f\xe7\xf1\x3c\x6c\x1c\x0f\x24\x3f\xe8\x5a\xcf\x53\xb8\xd8\xc2\x96\x9f\xc3\x9b\xb6\x1c\xf4\xf7\xf4\x8a\x01\xd1\x3a\x81\xa0\x68\x08\xe3\xeb\x3b\x86\x41\x5c\xeb\x35\x6b\x79\x57\xbc\xa6\xf9\x80\x53\x56\x4e\xb4\x19\x2b\x39\xc5\x1b\x96\x74\x0b\xbf\xef\x88\x36\xbe\xad\x7f\xeb\x8a\xc9\xa9\x2d\x08\x16\xb8\xde\x6a\x34\x78\x58\x13\xbe\xc3\xf1\xae\x36\x37\x8a\x08\xcd\x8a\x37\x76\x7b\x6e\xbc\xb2\x5d\x30\x25\x2a\xa4\xfe\x9a\x82\xcd\x25\x7c\xf6\xda\xdc\xdc\x96\x40\x84\x34\x8b\x66\xef\xfc\x68\xe4\x26\xa8\x35\x99\x77\xa3\xf1\x97\x2c\x21\x62\x68\xab\x29\x97\xf5\x86\xa9\xc0\x04\x75\xaf\x1a\x88\x79\xa9\x69\x2e\x4f\x6f\x24\x8f\x3b\x5e\x95\x8c\xd9\x3b\x6d\x54\x48\x74\x5d\x21\x53\x57\xb6\x0b\xf6\x25\xf3\x89\xdc\xf0\x56\x2a\x7a\xbc\x7e\xc5\x34\xa0\x59\x5b\x47\x21\xbb\x03\xfd\xe4\x14\xd4\x55\x45\x4d\x0d\x27\x5f\xf0\x84\xbb\x12\x65\xf9\x73\xcf\x3a\xe0\xcc\x57\xe4\x37\x2a\x6b\x39\xad\x7c\x43\xb8\xc6\x63\xf8\x20\x96\x42\xde\xee\xbf\xfb\xce\x49\xb5\x6b\xd1\x87\x9d\x17\x5d\xf9\x4e\x56\xfd\x20\xef\xdd\x68\x64\x8f\xed\xbb\xdd\xb7\x17\x7a\xe5\xc0\x5c\xc6\xcb\xba\x6d\xb7\xf5\xee\x1b\x2d\xb5\xc2\xea\x53\x58\x58\x2b\x85\xce\x56\x0a\xb7\x8b\xbc\xbd\x73\x6f\x25\xc7\xc2\x97\x0f\x5a\x24\xd6\x76\xa6\x23\xa9\xeb\x94\xbe\x27\x7a\x39\x61\x7f\xd4\xd0\xd0\x9e\xd1\xb4\xe5\x31\xf7\x71\x5f\x8e\x57\x3f\x3c\x31\xfe\xbf\x3f\x26\x7e\xf7\x69\x8b\x8e\xa7\x31\x16\xb4\x72\xca\xe4\x26\x6f\x9c\x71\x59\xf9\x68\xa3\xb2\xd8\xc8\x7e\xe7\x5d\x4d\xae\xf4\x9e\x6e\x4d\x15\xc3\xd9\x86\xeb\x7c\x5c\xe5\x72\x75\x4b\xcf\x6d\xcf\x99\x36\x2a\xbf\x1c\x3f\xc3\x01\x45\xf1\xc1\x84\x6e\xc2\x2a\xbe\x9c\x53\x29\xe5\x8a\xcc\xab\x4c\xa9\xc3\x37\x28\x5c\x87\xb5\xd6\xc1\x05\x24\x5f\x32\x69\x48\x5b\x27\xa8\x6f\xc7\x9d\x70\x2e\x63\x62\x9a\x8b\xb7\x3e\x67\x51\xad\x1d\xea\x6e\xfd\xe9\x4e\xdd\xe9\x94\x18\x83\x4a\x8c\xe0\x5f\x87\xbf\x7d\xfb\x75\x78\xf4\xd3\xe1\xe1\xa7\xd7\xc3\x7f\x7c\xfe\xf6\xf0\xb7\xc8\xfd\xf3\xff\x47\x3f\x1d\x7d\x2d\x7e\x7c\x7b\x74\x74\x78\xf8\xe9\xed\xfb\x9f\x6f\xc6\x17\x9f\xd9\xd1\xd7\x4f\x22\x4b\x96\xfe\xd7\xd7\xc3\x4f\x78\xf1\xb9\x23\x92\xa3\xa3\x9f\xfe\x6f\xf0\x88\x8d\xdb\xaa\x4d\xad\xe5\x70\xff\x76\x4b\xf9\x6d\x10\xd7\x76\xdc\xf8\xca\x50\xe3\xb5\x26\xa2\x70\x43\xb5\xac\x86\x84\x43\x35\x26\xe6\xd5\xd7\xf1\xce\x48\x4a\x62\x66\xf2\x68\x9f\x0b\xd5\x41\x77\x9a\x6a\xc6\xbf\x34\xe7\x59\x34\xa7\x70\x30\xae\x31\xed\x3f\xfa\x82\xae\xc7\x73\x58\x7a\x0d\x41\x12\x3c\x86\x2f\x19\x11\x86\x99\xfc\xa8\x91\x37\x4c\xe9\xbd\x14\x21\x0e\x5a\xf4\x97\x1e\xbc\xa0\x1e\x14\xa6\xbc\x75\x39\x4e\x1a\xc2\x1b\x9c\x48\xf4\xa8\x27\x4a\x1a\x6d\x5a\x48\x54\x7e\xb6\x67\x4b\xb8\x44\x30\x69\x3b\x7f\x6f\x45\xb0\xdf\xbc\xfe\x67\x61\xb5\x6c\xd8\x7a\xe8\x5e\xe0\xa4\x1b\x12\x0c\xef\x26\x6c\x3e\xc9\xa6\xa5\x6c\x8a\xf5\x43\x39\x08\xff\xfe\xcf\x60\x5d\x19\x92\x38\xc6\xd4\x20\xbd\xba\xff\x51\xbb\x57\xaf\x2a\x5f\xad\x73\x3f\x37\xda\x48\xf0\xe9\xf3\xc0\x2f\x8c\xf4\x63\xf1\x0d\x3a\xfb\xf0\xbf\x01\x00\x00\xff\xff\x62\x4d\x08\x6d\x09\x50\x00\x00"),
		},
		"/devops.k8s.io_machines.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_machines.yaml",
			modTime:          time.Date(2020, 12, 30, 3, 52, 45, 726000874, time.UTC),
			uncompressedSize: 10540,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x3a\x5b\x6f\x1b\xbb\xd1\xef\xfa\x15\x83\x7c\x0f\xfe\x0a\x58\xab\x93\x9e\x06\x2d\xf4\xe6\x3a\x49\x63\x24\xce\x11\x2c\x27\x2f\x45\x11\x8c\x96\x23\x2d\x8f\xb8\xe4\x86\x1c\xda\xd1\x29\xfa\xdf\x0b\x92\xbb\xab\xdb\xae\x2e\x8e\x0d\x54\x4f\x5a\x5e\xe6\x3e\xc3\x99\x21\x07\xc3\xe1\x70\x80\x95\xfc\x4a\xd6\x49\xa3\xc7\x80\x95\xa4\x1f\x4c\x3a\x7c\xb9\x6c\xf9\x37\x97\x49\x33\x7a\x78\x3d\x58\x4a\x2d\xc6\x70\xed\x1d\x9b\xf2\x8e\x9c\xf1\x36\xa7\xb7\x34\x97\x5a\xb2\x34\x7a\x50\x12\xa3\x40\xc6\xf1\x00\x00\xb5\x36\x8c\x61\xd8\x85\x4f\x80\xdc\x68\xb6\x46\x29\xb2\xc3\x05\xe9\x6c\xe9\x67\x34\xf3\x52\x09\xb2\x11\x78\x83\xfa\xe1\x97\xec\x2f\xd9\x2f\x03\x80\xdc\x52\xdc\x7e\x2f\x4b\x72\x8c\x65\x35\x06\xed\x95\x1a\x00\x68\x2c\x69\x0c\x25\xe6\x85\xd4\xe4\x32\x41\x0f\xa6\x6a\xa8\x1c\xb8\x8a\xf2\x80\x70\x61\x8d\xaf\xc6\xb0\x3d\x99\xf6\xd6\x04\x25\x66\x6e\x13\x98\x38\xa2\xa4\xe3\x8f\x9b\xa3\x9f\xa4\xe3\x38\x53\x29\x6f\x51\xad\x91\xc6\x41\x57\x18\xcb\x9f\xd7\x00\x87\x50\xe6\x69\x42\xea\x85\x57\x68\xdb\xf5\x03\x00\x97\x9b\x8a\xc6\x10\x97\x57\x98\x93\x18\x00\xd4\x4c\xc7\xed\x43\x40\x21\xa2\x18\x51\x4d\xac\xd4\x4c\xf6\xda\x28\x5f\xea\x16\xb8\x20\x97\x5b\x59\x71\x14\xd3\x7d\x41\x90\x2b\xcf\x64\xa1\x2a\xd0\x51\x16\x17\x01\xfc\xee\x8c\x9e\x20\x17\x63\xc8\x1c\x23\x7b\x97\xc5\xe9\x7a\x36\x49\x6e\xf2\xe1\x6a\xfa\xae\x1e\xe1\x55\xa0\xca\xb1\x95\x7a\xd1\x85\xe7\xe2\x7a\x57\x0d\x20\x1d\x20\x70\xfb\x69\xa9\xb2\xe4\x48\xb3\xd4\x0b\xe0\x82\xc0\x91\x7d\x20\x1b\x57\xd4\x48\x00\x1e\x0b\xd2\xc0\x85\x74\x60\x66\xbf\x53\xce\xf0\x88\x2e\x69\x98\x44\x06\x17\xfb\xc4\x37\xa6\x94\xed\x99\xc1\x16\x2b\x57\xff\xd8\x66\x44\x20\x27\xa4\x69\xfa\xe1\x75\xd2\x47\x5e\x50\x89\xe3\x7a\xa5\xa9\x48\x5f\x4d\x6e\xbe\xfe\x3a\xdd\x1a\x86\x6d\xc6\x6b\x0b\x08\xdc\x06\xa6\xd2\x5a\x98\x1b\x1b\x3f\x9b\xd9\xab\xc9\x4d\xbb\xbd\xb2\xa6\x22\xcb\xb2\x31\x87\xf4\xdb\x70\xab\x8d\xd1\x1d\x64\x17\x81\x9e\xb4\x0a\x44\xf0\x27\x4a\x58\x6b\x03\x21\x51\xb3\x00\x66\x9e\xa4\xd8\x0a\x3d\xca\x66\x0b\x30\x84\x45\xa8\x6b\x41\x67\x30\x8d\xea\x70\xc1\x5a\xbd\x12\xc1\x0d\x1f\xc8\x32\x58\xca\xcd\x42\xcb\x3f\x5a\xd8\x0e\xd8\x44\xa4\x0a\x99\x6a\xb3\x5f\xff\xa2\x41\x6a\x54\xf0\x80\xca\xd3\x25\xa0\x16\x50\xe2\x0a\x2c\x45\x75\x7a\xbd\x01\x2f\x2e\x71\x19\xdc\x1a\x4b\x20\xf5\xdc\x8c\xa1\x60\xae\xdc\x78\x34\x5a\x48\x6e\xc2\x49\x6e\xca\xd2\x6b\xc9\xab\x51\x8c\x0c\x72\xe6\xd9\x58\x37\x12\xf4\x40\x6a\xe4\xe4\x62\x88\x36\x2f\x24\x53\xce\xde\xd2\x08\x2b\x39\x8c\xa4\xeb\x18\x52\xb2\x52\xfc\x9f\xad\x03\x90\xbb\xd8\xa2\x75\xcf\xa2\xd3\x2f\x7a\xfb\x01\x0d\x04\xbf\x4f\xa6\x9d\xb6\x26\x2e\xf6\xad\xfb\xee\xdd\xf4\x1e\x1a\xd4\x51\x19\xbb\xd2\x4f\x06\xde\x6e\x74\x6b\x15\x04\x81\x49\x3d\x0f\xce\x11\x94\x38\xb7\xa6\x8c\x30\x49\x8b\xca\x48\xcd\xf1\x23\x57\x92\xf4\xae\xf8\x9d\x9f\x95\x92\x83\xde\xbf\x7b\x72\x1c\x74\x95\xc1\x75\x8c\xb1\x30\x23\xf0\x95\x48\x9e\x74\xa3\xe1\x1a\x4b\x52\xd7\x21\x24\xbc\xb4\x02\x82\xa4\xdd\x30\x08\xf6\x34\x15\x6c\x1e\x0f\xbb\x8b\x93\xd4\x36\x26\x9a\x38\xde\xa3\xaf\xda\x01\xa7\x15\xe5\x49\x6b\x1b\xb3\xc1\x01\xea\xc0\x9b\x6d\x41\xe8\xf6\xd0\x78\x38\x29\xef\x98\x6c\x88\xce\xbb\x53\xbd\xec\x84\xdf\x9c\x30\x48\x67\x7f\x4f\x3f\xaa\xb8\x4d\xaa\xee\x09\x00\xc9\x54\xf6\x4c\x1d\x83\x5a\x8b\xc9\x71\xff\xe4\x41\x66\x36\x84\x6f\xf3\x9f\x84\x11\x0c\x55\x5a\x12\x7d\x60\x86\x81\xce\xde\x39\x67\xf3\xc1\x21\xd4\x7b\xd6\xb2\xbb\x00\xad\xc5\x55\xc7\x7c\x61\xcc\xb2\x47\x76\x9b\xc7\xef\x31\x29\x1f\x15\xc0\x11\x32\xdd\x52\x56\xd7\x46\x27\x84\x4f\x31\x84\x13\x09\xe8\x16\xc3\x01\xe2\xe6\x52\xa3\x92\x7f\x90\xed\xc0\xbc\xe5\x7f\xef\xdb\x85\xd1\xfd\x34\x98\x0a\xbf\x7b\x8a\x29\x54\xf0\xbf\x74\x06\x00\x17\xc8\x50\x7a\x17\xa3\x14\x95\x15\x77\x29\x85\x0d\x54\x64\x4b\xd4\xa4\x59\x85\x23\xa5\x34\x0f\xd4\xc4\xd1\x18\x24\x1d\x1b\x8b\x8b\x1d\x6f\x3e\x28\xa4\x6e\x62\x83\x7f\x37\x27\xba\x8e\xff\x45\x88\x67\xf3\x55\x88\xee\xb8\xe6\x1e\x84\xef\x95\x6c\x1d\x2a\x40\xc9\x39\xe5\xab\x5c\x75\x50\x75\x44\x3f\xfd\xba\xa9\xa3\xd6\x11\xd9\x5f\x27\x0a\x76\x32\x94\x12\x23\x59\x35\x88\x94\x46\xc8\x26\x1c\xd6\x44\x67\x67\xc6\x29\x59\x8d\x07\x4f\x30\x3f\x85\x33\x52\xff\x03\x6e\x56\xa1\x73\x93\xc2\xa2\xa3\x6e\x0c\x73\x63\x4b\xe4\x31\xcc\x56\x4c\x4f\xe1\x33\xc0\x7f\x34\x56\x3c\x49\x48\x95\xb1\x7c\x98\x2c\xa9\xf9\xd7\x3f\x1f\x00\x1d\x72\xb2\x05\xd9\x2e\xd8\x56\x3e\x20\xd3\x47\x5a\xbd\x0c\xe3\x8c\x52\x73\x8f\xda\xb6\x4c\xf5\x66\x1e\x0f\x72\x39\x97\x24\x2e\x93\xdb\x19\x41\x17\xae\x86\x90\x9d\x1f\xf9\xf6\xaa\xa0\x00\x30\xe5\x53\xf7\x01\x66\x0c\x47\xcc\x98\x17\x24\x42\x64\x29\x30\xb9\xc7\x2b\x9a\xcf\x29\xe7\x57\xbd\xc7\x9a\xd1\x80\x7a\x05\x95\x11\x29\x6a\x09\x43\x0e\x42\x7e\xc5\x46\x91\x45\xa6\x08\x26\xe2\xc8\x7e\xe2\x78\x4e\x64\x1c\x3a\x5d\xb7\x38\xbc\xab\xcf\xd1\x2c\xf2\x9a\x36\xa7\x2a\x80\x92\x0c\x03\xdd\x95\x11\x29\xd4\x1e\x80\x0a\x20\xcc\x3e\x3b\x11\x44\x06\x5f\x51\x49\x51\x43\x77\x80\x96\xe0\xb3\x09\x15\x8f\xf0\x8a\x2e\x0f\x02\x9d\x58\x9a\x93\x5d\xaf\x8e\x85\xc1\x67\xf3\xee\x07\xe5\x9e\x29\xfb\xd9\x44\x64\xd9\x67\xc1\x47\x45\x95\x84\xb3\xa4\x55\x30\x82\x19\x01\x56\x95\x92\xc9\x24\xf0\x20\x47\xc1\x9e\x7e\x9a\xee\x50\xfc\x5e\x09\xd1\x9f\xff\xec\x9b\x72\xb3\x63\xa3\x72\x48\x2a\x92\x25\x01\x32\x3c\x16\x32\x2f\xc2\xc8\x41\xea\x13\xdb\xa1\xba\xc6\x00\x2c\x83\x9b\xe8\x11\x46\xab\x15\x3c\x5a\xc9\x4c\x3a\x16\xb1\xad\x8a\x0e\x7a\xe2\x76\xb4\x08\x35\xc6\x70\xab\xac\x7f\xa2\x74\x62\x72\x70\xba\x64\x5a\x6d\xa6\x92\x2c\x37\xd6\x92\xab\x42\xfa\x14\x6a\x32\xb3\x36\xe4\x83\x92\x59\xd2\x2a\x7b\xe9\x9c\x36\x79\x50\xef\xf4\x92\x56\x2f\x93\xd6\x7a\x17\x8a\xf3\x92\x9e\x70\x10\xf5\x33\x35\x04\x59\x75\x0c\x86\x73\xab\x63\xb8\x21\xe1\x9c\x6c\xb3\x42\xef\x7a\xeb\xad\x99\x31\x8a\x70\xb7\xb7\xc1\xa4\x51\xf3\xcd\xdb\xb3\xaa\xb4\x38\x75\xfa\x86\x6e\x91\x0c\x37\x8b\xc4\x9d\x99\x00\xea\xa4\xa2\x36\xb6\xe4\x4e\x28\x6b\xe3\xba\xcd\x48\x10\xaa\xf8\xe0\x85\x21\x9f\xc3\x99\xf1\xa9\x57\x90\xe0\x81\x99\xef\x30\x87\xfa\xdc\x02\x18\x85\xb0\xe4\x1c\x1d\xcb\xfb\x3f\xd5\xf9\x7d\xbb\x1e\x2c\x61\x5e\xe0\x4c\x51\xe3\x8a\x9d\x98\x4f\x4f\xd6\x6b\x11\x5c\x25\x04\xb1\x6d\x8c\x52\x6f\x4b\xa0\x69\xc3\xd5\xa8\x2e\x5c\x5f\xaa\x19\x40\x64\x83\xf3\x4f\xea\x7a\xeb\xc9\x49\x48\x93\x75\x1f\x40\x79\x7a\x42\x7b\x0a\xd2\xdb\x6d\x84\x71\xe3\x25\x18\x4d\x41\x39\x13\x3f\x53\x32\xbf\x84\x77\x3f\x52\xd3\xee\x66\xd2\x9f\xf5\x58\xb8\xd1\xcd\xaa\x27\x92\x7d\x28\x2e\x0e\x1b\x0a\x3b\xe7\xf6\xfc\xe6\x84\x68\xd8\x1f\x09\xf3\x03\x15\xf5\x59\xb6\xd7\x96\xe6\x6b\xeb\x13\xc4\x28\x95\x6b\x2d\x2f\xf7\xd6\x92\xe6\x35\xce\x4e\xd1\x35\xed\xda\xdb\x3e\x97\x38\x6e\x89\x0a\x1d\x4f\xac\x99\x51\x48\x10\x4e\x32\x8d\x4f\xe8\x38\x65\x0d\x8f\x14\xc0\xcf\x42\xd6\x13\x48\x6e\x48\xed\x53\xf3\xa9\xe7\xfc\x51\x2b\x0e\x34\xdf\x5b\xd4\x4e\x36\x9d\xfb\x33\x09\xdf\x22\x17\xb8\x05\x45\x22\xf5\x03\x82\x9d\xa7\xd8\xd7\x9f\x81\x19\x40\x6d\xb8\xe8\x2a\x7a\x9f\x99\xdd\x92\x9c\xc3\xc5\x69\x3c\x7e\xf0\x25\xea\xa1\x25\x14\x31\x64\xd6\x5b\x41\x6a\x21\x73\x8c\x4d\xe6\xc6\xd2\x62\x94\xef\x65\x4f\x45\x59\xb5\x82\x79\x72\xc0\xb1\x84\x6e\xf7\x6a\xa2\x87\xf4\x2f\x5a\x7e\xf7\x29\xc8\x0c\x43\xd5\x7b\xb9\x6e\x35\xd7\x60\xd6\xde\xd1\xe8\xee\xc2\xbd\x38\x07\x5d\x67\x6a\x0f\x07\xf5\xb1\x5a\x37\x4c\xda\xc3\x73\xc7\x3b\xe0\x1a\x75\xa8\x18\xee\xad\x3f\x50\xfc\xbc\x47\xe5\xe8\x12\xbe\xe8\xa5\x36\x8f\xfa\xe5\x03\xfe\xfd\xaa\x6a\x5b\x3d\x61\xd3\x3e\xdd\x2f\x11\xbc\x7b\x9d\xec\xb9\x63\xb7\x32\xf9\xb2\x8b\x88\x43\xb9\x60\x7d\xe8\xde\xe8\xb9\x39\x92\xb5\x4c\x29\x26\x2d\x52\xb8\x91\xf7\x52\xc4\xab\x2e\x1f\xcd\x59\xad\xda\x1e\x60\xdb\x9e\x38\xb7\x4b\xb6\x79\x4f\x72\x42\x4f\x24\xe4\x0b\x57\x1b\x5b\x42\x9a\x67\x2c\x93\x80\xd9\x9a\x86\xa7\x74\x65\x66\xc6\x74\x66\xc6\x7b\x14\xfc\xdd\x18\x86\x9b\xb7\x9d\x88\xb3\xa7\x60\xae\x8f\x49\xb2\x77\x5e\x87\x48\xda\x79\xe3\xd9\xdd\xcb\xdc\xd9\x09\xcd\x35\xe8\xb3\xd1\xb6\x24\xab\x49\x9d\x4e\xd1\xc7\xb8\xfe\x05\xe8\xf0\x33\x9a\x58\xf3\x63\x75\x06\x29\xcd\x96\x97\xa1\x46\x11\x9f\x47\x8b\x22\x7e\x7e\x4a\x1a\x2f\x3e\xc5\x70\x2f\x6e\x9b\xc5\xdd\xf8\xe1\xbd\xb1\xb5\x63\x6f\x3c\xbd\xe8\xec\x31\x26\xa7\x8f\x87\xae\xd1\x20\x75\x7d\xf7\x9a\x7a\xfb\xe9\x7a\x56\x92\x8a\x57\xc2\x55\xec\x71\xc5\xce\xd2\x27\x42\xab\x7b\x40\x96\xc6\x52\x4a\x4f\x4a\xd4\xff\xff\xe6\x4f\x0d\x05\x43\x29\xd2\xfd\xeb\x78\x34\x2a\x51\xff\x35\x33\x76\x31\x52\x52\xfb\x1f\xe1\x73\x58\xe1\x82\x5c\xf8\xf7\x66\xb4\xde\x90\xbd\xc9\x0a\x2e\xd5\xc5\x53\x04\x1a\x42\x55\x4c\x25\xa6\x2b\xc7\x54\x9e\x18\x91\x7e\x6b\x76\x41\xda\xf6\x6c\x51\xc9\xb8\x9b\xb2\x37\x3b\xda\x22\xe3\xb7\x29\xc4\xa5\xcf\x67\x5b\x2e\xb2\xf2\xe5\xcb\x49\xc6\x35\x6d\x17\x3f\xb3\x71\xad\x8d\x76\xdb\x98\xee\xb7\xac\xac\xee\x93\xf7\x5e\x7c\x1a\xb8\x23\x01\x1f\x90\xa1\x30\x8e\x5d\x7b\xa3\x8f\x79\x1e\x2a\x4e\x4b\xa2\x40\xce\x72\x53\x8e\x84\xc9\x7d\xd9\xbc\x0d\x19\x91\x1e\x7e\x99\x8e\xee\x48\x7c\xfb\x80\xfc\x6d\xea\x67\x2d\xcb\xdf\x6e\x51\xe3\x82\xc2\xd2\xd1\xeb\x51\xb0\xb7\xd1\xdd\x87\xe9\xed\x68\x41\x1c\x0c\x61\x98\xa4\x37\x0c\x27\x66\xb4\xc6\xf3\x35\x70\x20\x19\xe8\x4d\x9a\xb7\x74\x72\x05\x45\x48\x98\xe1\xe4\x84\x19\x1e\x8b\xce\x2b\xc6\x8d\x1a\x5d\xba\xe4\xee\xd2\x1d\x4a\x9e\x0e\xf0\x15\x5f\x54\x1d\x21\xbc\xd6\xf9\x24\x2c\xdd\x7a\xd2\x13\x37\x6f\xbc\x50\x08\x34\x38\xb6\x3e\x67\x63\xcf\x21\xa2\x2f\x71\xdf\x11\xdf\xcc\x4a\x9a\x6f\x24\xea\xcf\x2b\xbf\x90\x1e\xd2\x19\xb2\xeb\xb4\x87\xbd\xc1\xf8\x80\x4c\x8c\x81\xad\x4f\x1e\x56\x5f\xff\x6e\x8e\xf8\x59\xfb\xfc\xa7\x91\x41\x5d\x08\xc0\xbf\xff\x33\x58\xd7\x04\xc1\x3b\x2a\x26\xf1\x79\xf7\xdd\xdf\xab\x57\x5b\x0f\xfb\xe2\xe7\x46\x03\x01\xfe\xf9\xaf\x41\x42\x4c\xe2\x6b\xf3\x4c\x2f\x0c\xfe\x37\x00\x00\xff\xff\x53\xc6\x77\x07\x2c\x29\x00\x00"),
		},
		"/workload.k8s.io_addons.yaml": &vfsgen۰CompressedFileInfo{
			name:             "workload.k8s.io_addons.yaml",
			modTime:          time.Date(2020, 12, 30, 3, 52, 45, 670775182, time.UTC),
			uncompressedSize: 2276,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x55\xcd\x8e\x23\x35\x10\xbe\xf7\x53\x94\x96\xc3\x5c\x48\x67\x57\x70\x40\x7d\x8b\xc2\x0a\x46\xb0\xab\x68\x33\x9a\x0b\xe2\x50\xb1\x2b\x69\xef\xb8\x6d\xe3\x2a\x67\x18\x10\xef\x8e\x6c\x77\x67\x3a\xc9\x0e\x20\xa1\xf1\xcd\x55\x65\x57\xd5\xf7\xd5\x4f\xb3\x58\x2c\x1a\x0c\xe6\x9e\x22\x1b\xef\x3a\xc0\x60\xe8\x77\x21\x97\x6f\xdc\x3e\x7c\xc7\xad\xf1\xcb\xe3\xbb\xe6\xc1\x38\xdd\xc1\x3a\xb1\xf8\xe1\x13\xb1\x4f\x51\xd1\xf7\xb4\x37\xce\x88\xf1\xae\x19\x48\x50\xa3\x60\xd7\x00\xa0\x73\x5e\x30\x8b\x39\x5f\x01\x94\x77\x12\xbd\xb5\x14\x17\x07\x72\xed\x43\xda\xd1\x2e\x19\xab\x29\x96\xcf\x27\xd7\xc7\xb7\xed\xb7\xed\xdb\x06\x40\x45\x2a\xcf\xef\xcc\x40\x2c\x38\x84\x0e\x5c\xb2\xb6\x01\x70\x38\x50\x07\xa8\x75\x0e\xee\xd1\xc7\x07\xeb\x51\x8f\x51\x36\x1c\x48\x65\x87\x87\xe8\x53\xe8\xe0\x52\x5d\x5f\x8f\x21\xd5\x74\x56\xe5\xa3\x22\xb0\x86\xe5\xa7\x99\xf0\x67\xc3\x52\x14\xc1\xa6\x88\x76\x72\x5a\x44\x6c\xdc\x21\x59\x8c\x33\x21\x2b\x1f\xa8\x83\x8f\xd9\x43\x40\x45\xba\x01\x18\x13\x2b\x1e\x17\xd9\xb4\x40\x85\x76\x13\x8d\x13\x8a\x6b\x6f\xd3\x30\x41\xb4\x00\x4d\xac\xa2\x09\x52\xa0\xb8\xeb\x69\x8c\x03\x42\x8f\x4c\x6d\x31\x02\xf8\xcc\xde\x6d\x50\xfa\x0e\x5a\x16\x94\xc4\x6d\x51\x8f\xda\x8a\xce\xe6\xc7\xd5\xf6\xfd\x28\x91\xa7\x1c\x15\x4b\x34\xee\xf0\x25\x3f\x37\xeb\x4b\xa8\xc1\x30\x20\xc8\xe9\x1a\x29\x44\x62\x72\x62\xdc\x01\xa4\x27\x60\x8a\x47\x8a\xc5\x62\x74\x02\xf0\xd8\x93\x03\xe9\x0d\x83\xdf\x7d\x26\x25\xf0\x88\x5c\x59\x24\xdd\xc2\xcd\x75\xf0\x53\xb9\xb4\x57\x54\x9f\xa5\xb2\xfa\xe1\x3c\x11\x8d\x52\x9d\x56\xf5\xf1\x5d\xa5\x43\xf5\x34\x60\x37\x5a\xfa\x40\x6e\xb5\xb9\xbd\xff\x66\x7b\x26\x86\xf3\xc4\x3f\xa0\xea\x8d\xa3\x9c\x6d\x4e\xaa\xda\xc2\xde\xc7\x72\x9d\xb4\xab\xcd\xed\xe9\x79\x88\x3e\x50\x14\x33\x55\x50\x3d\xb3\xd6\x99\x49\x2f\x9c\xdd\xe4\x78\xaa\x15\xe8\xdc\x33\x54\xbd\x8e\x05\x42\x7a\x4c\x01\xfc\xbe\xa2\x78\x02\xbd\x60\x73\xf6\x31\x64\x23\x74\x23\xd0\x2d\x6c\x0b\x1d\x0c\xdc\xfb\x64\x75\x6e\xb5\x23\x45\x81\x48\xca\x1f\x9c\xf9\xe3\xf4\x37\x83\xf8\xe2\xd4\xa2\xd0\x58\xda\xcf\xa7\x14\xa4\x43\x0b\x47\xb4\x89\xbe\x06\x74\x1a\x06\x7c\x82\x48\x85\xce\xe4\x66\xff\x15\x13\x6e\xe1\x83\x8f\x04\xc6\xed\x7d\x07\xbd\x48\xe0\x6e\xb9\x3c\x18\x99\x46\x86\xf2\xc3\x90\x9c\x91\xa7\x65\xe9\x7e\xb3\x4b\xe2\x23\x2f\x35\x1d\xc9\x2e\xd9\x1c\x16\x18\x55\x6f\x84\x94\xa4\x48\x4b\x0c\x66\x51\x42\x77\x65\x6c\xb4\x83\xfe\x2a\x8e\x43\x86\x6f\xce\x62\xbd\xaa\xe8\x7a\x4a\x3f\xff\x03\x03\xb9\xb5\x6b\x69\xd7\xa7\x35\x8b\xeb\xea\xfe\xf4\x7e\x7b\x07\x93\xeb\x42\xc6\x25\xfa\xb5\xc0\x4f\x0f\xf9\x99\x82\x0c\x98\x71\xfb\xdc\x1c\x99\xc4\x7d\xf4\x43\xf9\x93\x9c\x0e\xde\x38\x29\x17\x65\x0d\xb9\x4b\xf8\x39\xed\x06\x23\x99\xf7\xdf\x12\xb1\x64\xae\x5a\x58\x97\x39\x0a\x3b\x82\x14\x74\xed\xa4\x5b\x07\x6b\x1c\xc8\xae\xf3\x48\x78\x6d\x02\x32\xd2\xbc\xc8\xc0\xfe\x37\x0a\xe6\x2b\xe0\xd2\xb8\xa2\x36\x53\x4c\xb3\xfa\x05\xbe\xea\xec\xdb\x06\x52\x95\xb4\x99\x32\xd7\x7f\x55\xb7\x67\xef\xbf\xdc\x9f\xf9\xec\xbd\xbf\x14\xbd\x98\xc4\xcb\x01\x97\x71\xfb\xef\x21\x17\xb3\x59\x79\x14\x82\xe2\x50\x1a\x19\x70\xe7\x53\x2d\x83\xfa\x5d\xed\xe5\x8b\xd8\x5e\x33\x39\xa8\xeb\xe4\xff\xc2\x71\x25\x2c\x4b\x41\x77\x20\x31\xd5\x01\xcd\xe2\x23\x1e\x68\x2e\x49\xbb\x53\x4b\x4f\xfe\x47\x50\xe1\xcf\xbf\x9a\x67\x7c\x51\x29\x0a\x42\xfa\xe3\xe5\xb6\x7e\xf3\xe6\x6c\x21\x97\xab\xf2\xae\xae\x55\xee\xe0\x97\x5f\x9b\xea\x98\xf4\xfd\xb4\x7a\xb3\xf0\xef\x00\x00\x00\xff\xff\xe1\xa6\x02\xca\xe4\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/devops.k8s.io_clustercredentials.yaml"].(os.FileInfo),
		fs["/devops.k8s.io_clusters.yaml"].(os.FileInfo),
		fs["/devops.k8s.io_machines.yaml"].(os.FileInfo),
		fs["/workload.k8s.io_addons.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
