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
			modTime: time.Date(2020, 12, 29, 6, 51, 58, 658636542, time.UTC),
		},
		"/devops.k8s.io_clustercredentials.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_clustercredentials.yaml",
			modTime:          time.Date(2020, 12, 29, 6, 52, 41, 967435914, time.UTC),
			uncompressedSize: 3261,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x56\x4d\x6f\x1b\x37\x10\xbd\xef\xaf\x18\xa4\x87\x5c\xba\x2b\x07\x6d\x81\x76\x6f\xae\x9c\x02\x86\xdb\xc0\x88\x0d\xa3\x40\xd1\x03\x45\x8e\x24\xc6\xbb\x43\x76\x66\x28\x44\xfd\xf5\x05\xb9\x2b\x6b\xe5\xaf\x28\x75\x92\xbd\x2d\x3f\xde\x1b\x3e\xce\xcc\x63\x55\xd7\x75\x65\xa2\xbf\x41\x16\x1f\xa8\x05\x13\x3d\x7e\x54\xa4\xfc\x27\xcd\xed\xcf\xd2\xf8\x30\xdb\xbc\xa9\x6e\x3d\xb9\x16\xe6\x49\x34\xf4\xef\x51\x42\x62\x8b\x67\xb8\xf4\xe4\xd5\x07\xaa\x7a\x54\xe3\x8c\x9a\xb6\x02\x30\x44\x41\x4d\x1e\x96\xfc\x0b\x60\x03\x29\x87\xae\x43\xae\x57\x48\xcd\x6d\x5a\xe0\x22\xf9\xce\x21\x17\xf0\x1d\xf5\xe6\xa4\xf9\xb1\x39\xa9\x00\x2c\x63\xd9\x7e\xed\x7b\x14\x35\x7d\x6c\x81\x52\xd7\x55\x00\x64\x7a\x6c\xc1\x76\x49\x14\xd9\x32\x3a\x24\xf5\xa6\x93\xc6\xe1\x26\xc4\x5d\xbc\x95\x44\xb4\x99\x7a\xc5\x21\xc5\x16\x0e\x27\x07\x94\x31\xb4\xf1\x58\x03\xe0\xfc\x0e\xb0\xcc\x75\x5e\xf4\xe2\xf1\xf9\xdf\xbd\x68\x59\x13\xbb\xc4\xa6\x7b\x2c\xa4\x32\x2d\x9e\x56\xa9\x33\xfc\xc8\x82\x0a\x40\x6c\x88\xd8\xc2\xbb\x1c\x4e\x34\x16\x5d\x05\x30\xaa\x51\xc2\xab\xc7\xf3\x6e\xde\x0c\x60\x76\x8d\xbd\x19\xe2\x06\x08\x11\xe9\xf4\xf2\xfc\xe6\x87\xab\x83\x61\x00\x87\x62\xd9\x47\x2d\x9a\x3e\x88\x1c\x18\x6d\x60\x27\xa0\x6b\x84\x7d\x34\xe0\x69\x19\xb8\x2f\xb2\x03\x21\x3a\x74\xa0\xe1\x0e\x13\xc0\x58\x8b\x32\xee\x1a\x30\x9b\xbb\xd9\xc8\x21\x22\xab\xdf\x89\x3a\xee\xd8\x67\xd5\x64\xf4\x5e\x7c\xaf\xf3\x11\x86\x55\xe0\x72\x3a\xe1\xc0\x31\xca\x80\x6e\x3c\x35\x84\x25\xe8\xda\x0b\x30\x46\x46\x41\x1a\x12\xec\x00\x18\xf2\x22\x43\x10\x16\x1f\xd0\x6a\x03\x57\xc8\x19\x06\x64\x1d\x52\xe7\x72\x16\x6e\x90\xb5\x08\xb0\x22\xff\xef\x1d\xb6\x80\x86\x42\xda\x19\xc5\xf1\x5e\xf7\x9f\x27\x45\x26\xd3\xc1\xc6\x74\x09\xbf\x07\x43\x0e\x7a\xb3\x05\xc6\xcc\x02\x89\x26\x78\x65\x89\x34\xf0\x47\x60\x2c\x8a\xb6\xb0\x56\x8d\xd2\xce\x66\x2b\xaf\xbb\x6a\xb2\xa1\xef\x13\x79\xdd\xce\x4a\x61\xf8\x45\xd2\xc0\x32\x73\xb8\xc1\x6e\x26\x7e\x55\x1b\xb6\x6b\xaf\x68\x35\x31\xce\x4c\xf4\x75\x09\x9d\x4a\x45\x35\xbd\xfb\x8e\xc7\xfa\x93\xd7\x07\xb1\xea\x36\x67\x93\x28\x7b\x5a\x4d\x26\x16\x21\xa8\x28\x9b\x78\x1d\x6e\xf1\xb9\xbb\xf8\x2d\x30\xe4\xda\x34\xae\x87\x5c\xd7\x10\x18\x3e\x04\x4f\xc7\x90\x58\x33\x47\xd6\x4f\x80\xdb\x40\x94\x35\x9b\x24\xd1\xc1\x86\x21\x03\x5b\x58\x6c\x15\x8f\x23\xbd\xc0\x6d\xfb\x32\x88\x9c\xb7\x4b\x6f\x8d\xe2\x03\xac\x2f\x27\x0e\xb2\xca\xaf\x9e\x0c\x6f\xcf\xc6\x1e\x39\x29\x13\xe7\x4a\x0b\x35\xdd\xe5\xa3\x65\xf4\xec\xa9\x9e\xa4\xdc\x4d\x0c\xb5\x30\x8d\xa5\xf3\x48\x7a\xc4\x65\xe5\xc3\xd6\x26\x7a\x29\x55\x04\x7f\xfe\x74\xf2\x0b\x98\xa4\xeb\x97\xc9\x5d\xd8\x8f\x53\xfa\x2b\x90\x97\x94\xcb\xed\xb6\x3d\x66\x3d\xaa\x75\xa7\x97\xe7\xf3\x27\x14\xfb\x5c\xfa\x03\xb8\x17\x27\x6e\x46\x9b\x9f\x1e\x71\x8f\xd7\x17\x6f\xc1\x13\xac\xba\xb0\x28\xdd\x3f\x09\x7e\x01\xe2\x97\xc7\xff\x51\xff\x5f\x2d\x7c\x6e\xc2\x17\x87\x7f\xc6\x80\xb2\xc3\x83\x17\x30\x23\xe6\xd0\xc4\xf7\x3e\x93\x87\x72\xc3\x7a\xff\xf6\xea\x1a\x76\x9d\xb7\x78\xd1\x7d\xf3\x29\xcc\xfb\x8d\xb2\x77\xa0\xec\x17\x9e\x96\xc8\x83\x87\x2d\x39\xf4\x05\x13\xc9\xc5\xe0\x69\xd7\x11\x73\x62\xdc\x03\x95\xb4\xe8\xbd\x66\xdb\xfb\x27\xa1\x68\xb6\xaa\x06\xe6\xe5\x85\x05\x0b\x84\x14\x9d\x51\x74\x0d\x9c\x13\xcc\x4d\x8f\xdd\xdc\x08\x7e\x75\xff\xc9\x4a\x4b\x9d\x85\x3d\xce\x81\x72\x35\x7f\x9b\xcb\xee\x0d\xf9\x65\xd6\xe9\x1b\xd1\x4d\x5e\xbd\x9f\x5c\xac\x48\x86\xf4\xfc\xec\xa8\xde\xa3\x47\x7a\xf5\xa4\x49\x96\x2d\x0f\xbb\xe4\xa3\x04\x39\x9d\x3c\xe3\xa4\x30\xea\x69\x7b\x9c\x8c\xee\xa2\xae\x9e\x3c\x5d\xa1\x77\x2d\x28\xa7\x61\xa3\x68\x60\xb3\xc2\x71\x44\xd4\x68\x2a\x4a\xe7\x27\x64\x54\x74\xef\xee\x3f\xbf\x5f\xbd\x3a\x78\x4b\x97\x5f\x1b\x68\xb8\x2f\x69\xe1\xaf\xbf\xab\x01\x15\xdd\xcd\xee\x79\x9c\x07\xff\x0b\x00\x00\xff\xff\xc1\x87\xc4\x67\xbd\x0c\x00\x00"),
		},
		"/devops.k8s.io_clusters.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_clusters.yaml",
			modTime:          time.Date(2020, 12, 29, 6, 52, 42, 24155000, time.UTC),
			uncompressedSize: 20536,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x3c\x5b\x73\xdb\x36\xd6\xef\xfa\x15\xe7\xcb\xf7\x75\x6c\x7f\xb5\xe8\x74\xdb\xd9\xe9\xea\xa5\xe3\xb5\x9d\xd6\x93\xc4\xd5\x58\x4e\x5e\xd2\xec\x0c\x44\x1c\x49\xa8\x40\x80\x01\x40\xd9\xec\x66\xff\xfb\x0e\x2e\xa4\x44\x8b\xa4\x48\xf9\xd6\x87\xfa\xc9\x22\x0e\x0e\x70\xee\x17\x80\x1c\x0c\x87\xc3\x01\x49\xd9\x47\x54\x9a\x49\x31\x02\x92\x32\xbc\x33\x28\xec\x2f\x1d\x2d\x7f\xd4\x11\x93\x27\xab\xef\x06\x4b\x26\xe8\x08\xce\x32\x6d\x64\x72\x8d\x5a\x66\x2a\xc6\x73\x9c\x31\xc1\x0c\x93\x62\x90\xa0\x21\x94\x18\x32\x1a\x00\x10\x21\xa4\x21\xf6\xb1\xb6\x3f\x01\x62\x29\x8c\x92\x9c\xa3\x1a\xce\x51\x44\xcb\x6c\x8a\xd3\x8c\x71\x8a\xca\x21\x2f\x96\x5e\xbd\x8e\x7e\x88\x5e\x0f\x00\x62\x85\x6e\xfa\x0d\x4b\x50\x1b\x92\xa4\x23\x10\x19\xe7\x03\x00\x41\x12\x1c\x41\xcc\x33\x6d\x50\xe9\x88\xe2\x4a\xa6\xc5\x2e\x07\x3a\xc5\xd8\x2e\x38\x57\x32\x4b\x47\x50\x1d\xf4\x73\xc3\x86\x02\x31\x1e\x8d\x7b\xc2\x99\x36\x6f\x37\x9f\xbe\x63\xda\xb8\x91\x94\x67\x8a\xf0\xf5\xa2\xee\xa1\x5e\x48\x65\xae\xd6\x08\x87\xb0\x8a\xfd\x00\x13\xf3\x8c\x13\x55\xc2\x0f\x00\x74\x2c\x53\x1c\x81\x03\x4f\x49\x8c\x74\x00\x10\x88\x76\xd3\x87\x40\x28\x75\x6c\x24\x7c\xac\x98\x30\xa8\xce\x24\xcf\x12\x51\x22\xa7\xa8\x63\xc5\x52\xe3\xd8\x74\xb3\xc0\x02\x39\x50\xa1\x2f\xc7\x91\x83\x02\xf8\x5d\x4b\x31\x26\x66\x31\x82\x48\x1b\x62\x32\x1d\xb9\xe1\x30\xea\x59\x77\x7e\x35\x29\x9f\x98\xdc\x6e\x4b\x1b\xc5\xc4\xbc\x69\xa1\xb0\x4f\x90\x33\xb0\x62\x53\x02\x0d\xea\xe6\x05\xa3\x00\x5f\x59\xf3\xe3\xc5\xf5\xe4\xf2\xd7\xab\x1e\xab\xc6\x3c\xb3\xd4\xa5\x0b\xa2\xb1\x79\x31\x37\x5c\x59\x69\xfc\xcb\xe9\xe4\xa2\xe3\x3a\x07\x67\xf7\xb5\x0c\x98\x06\x02\xa6\xfc\xa9\x30\x55\xa8\x51\x18\x26\xe6\x60\x16\x08\x1a\xd5\x0a\x95\x83\x08\x8b\x00\xdc\x2e\x50\x80\x59\x30\x0d\x72\xfa\x3b\xc6\x06\x6e\x89\xf6\x0a\x8c\x34\x82\x83\xed\xcd\x17\x96\x12\x6d\x69\x79\x85\x94\xd3\x9f\xab\x84\x50\x62\xfc\xa2\x7e\x78\xf5\x9d\x57\xb7\x78\x81\x09\x19\x05\x48\x99\xa2\x38\x1d\x5f\x7e\xfc\x7e\x52\x79\x0c\x55\xc2\x83\x82\x5b\x6a\x2d\x51\x1e\x16\x66\x52\xb9\x9f\xc5\xe8\xe9\xf8\xb2\x9c\x9e\x2a\x99\xa2\x32\xac\xd0\x76\xff\xb7\xe1\x35\x36\x9e\xde\x5b\xec\xc0\xee\x27\xe8\x10\xb5\xee\x02\xfd\xaa\x41\x4f\x90\x06\x12\xac\x82\x39\x2e\x96\x4c\x77\xbc\xa9\x20\x06\x0b\x44\x44\x60\x74\x04\x13\x27\x0e\x6d\x8d\x31\xe3\xd4\x7a\x99\x15\x2a\x03\x0a\x63\x39\x17\xec\x8f\x12\xb7\x06\x23\xdd\xa2\x9c\x18\x0c\x56\xbd\xfe\x73\xf6\x26\x08\x87\x15\xe1\x19\x1e\x03\x11\x14\x12\x92\x83\x42\x27\xce\x4c\x6c\xe0\x73\x20\x3a\x82\xf7\x52\x21\x30\x31\x93\x23\x58\x18\x93\xea\xd1\xc9\xc9\x9c\x99\xc2\x5b\xc6\x32\x49\x32\xc1\x4c\x7e\xe2\x1c\x1f\x9b\x66\x46\x2a\x7d\x42\x71\x85\xfc\x44\xb3\xf9\x90\xa8\x78\xc1\x0c\xc6\x26\x53\x78\x42\x52\x36\x74\x5b\x17\xce\x63\x46\x09\xfd\x5f\x15\xfc\xab\x3e\xa8\xec\x75\x4b\xa3\xfd\x9f\x73\x66\x2d\x12\xb0\x6e\xcd\xab\xb6\x9f\xea\xa9\xd8\xd6\xee\xeb\x8b\xc9\x0d\x14\x4b\x3b\x61\xdc\xe7\xbe\x57\xf0\x72\xa2\x5e\x8b\xc0\x32\x8c\x89\x99\x35\x0e\x2b\xc4\x99\x92\x89\xc3\x89\x82\xa6\x92\x09\xe3\x7e\xc4\x9c\xa1\xb8\xcf\x7e\x9d\x4d\x13\x66\xac\xdc\xbf\x64\xa8\x8d\x95\x55\x04\x67\x2e\x84\xc0\x14\x21\x4b\xa9\xb7\xa4\x4b\x01\x67\x24\x41\x7e\x66\x5d\xc2\x53\x0b\xc0\x72\x5a\x0f\x2d\x63\xbb\x89\x60\x33\xfa\xdd\x07\xf6\x5c\xdb\x18\x28\xc2\x54\x83\xbc\x82\x01\x4e\x52\x8c\x2b\x16\x43\x51\x33\x65\x75\xda\x10\x83\xd6\x12\x36\xc3\x57\xbb\xa5\x06\x6b\xf5\xc2\xba\xb8\x33\x8a\x9c\xaa\xf9\x16\x04\x54\xc2\x50\x13\x9e\x16\x2e\xb4\x52\xed\xf2\x00\xbf\xe3\xb3\xcb\xf3\xeb\xd1\xa0\x07\xd2\x75\xfe\xf0\x9e\x08\x32\x7f\x51\x1a\x28\xd3\x29\x27\xb9\x0d\xe6\xbd\x68\xa0\x42\x9f\xcb\x84\x30\xb1\x3d\xab\x22\xfe\xf3\xab\x89\x87\x2b\xfc\x33\x15\x1a\xa8\x7f\x92\x69\xa4\x30\xcd\x61\xf9\xa3\x76\xb1\x88\xc5\xd6\x19\x9d\xe3\x8c\x64\xdc\xe8\x3a\x22\x25\xbc\x0a\x4c\x8f\xb8\x8c\x09\x7f\x15\xf5\xda\xb3\x8c\x97\x2f\xca\x6c\x34\x31\xdd\xc1\xaf\x0b\x13\x53\x58\x48\x4e\xb5\x55\x93\x19\x9b\x67\xca\x45\x0e\x17\xd0\xec\xfc\x6d\x8a\xd3\xd6\xbd\xda\xec\xd7\xc6\x83\xba\xb1\xfb\x6b\x07\xd0\xf0\x74\x8a\x1a\x16\xf2\xd6\x72\x3d\x96\x42\x58\x5f\x69\xa4\x0d\x58\x05\xca\x5a\x8c\x9e\xca\x32\xa3\x7b\x67\xc5\xe4\x82\x50\x89\x9d\x28\x84\x24\x33\x19\xe1\x3c\x07\xbc\xb3\x90\x6c\x85\xb5\xc8\xda\x49\x73\xb6\x44\xde\x30\x8e\x4d\xa3\xf7\x7d\xd1\xa9\x05\x76\xc1\x43\xc0\x64\xf2\x0e\xce\x2c\xf2\x19\x8b\xad\x0b\x3a\xcd\xcc\x42\x2a\x66\x72\x98\x59\x20\xab\x9c\x8d\x58\x9d\x2a\x6a\x8c\x33\x85\x81\x5c\xef\xa2\x63\x27\xab\x08\xae\xf1\x4b\xe6\xbc\x1b\x9b\x41\x66\x73\x68\x20\x70\xf3\x6e\x52\xf0\xd1\xc2\x34\xe2\x6e\x55\xae\x40\x34\x2a\xd3\x87\xec\x00\xbe\x41\x78\x5c\x12\xee\x74\xab\x20\x18\x8c\x6c\xa1\xf9\xe5\x08\x2e\xe2\xae\xee\x48\xf1\x45\x01\x6f\x03\x8b\xdb\x6f\x82\xc9\xd4\x96\x58\xeb\x9d\x5a\x83\x2a\x74\xf2\xa2\xd6\xb0\xca\x94\xca\x60\xd2\xb2\x72\x27\x0a\x0a\x20\xa2\x14\xc9\x1b\x60\x96\x98\xf7\x90\xea\x5b\x0f\xbd\x21\xd4\x25\xe6\x15\x51\x6e\x0a\xac\x65\xf7\xcf\x29\x4a\x15\x90\xd7\xd3\x38\x0c\xe6\xdc\x34\x18\xf4\xb8\x61\xb8\x54\x92\x86\xf1\xc0\xde\x41\xf3\xc6\x6b\x9d\xb6\x2b\xa7\xad\x17\xeb\xe0\x41\xbd\xb7\x4b\x95\x5c\x31\x8a\xf7\x3d\xf8\x52\xc8\xa9\x76\x6a\x57\x3c\x6f\x56\x17\x97\xde\x3b\x64\x4e\x7b\x99\xd0\x86\x88\x18\x9f\xdc\x9d\xda\xac\xef\x9c\xa9\x8e\x2a\x78\xee\xa1\xcb\xd0\xce\x14\xc6\x46\xaa\xdc\x6f\xfa\x96\x71\x0e\x29\x27\x31\x02\x6b\x10\xca\x7a\xd1\x75\xdc\x77\x51\xfe\x64\x45\xd4\x09\x67\xd3\x13\x8b\xe9\xd5\xc3\x7c\x47\x73\xbc\xef\x1b\xf7\x7b\xd9\xfb\xfd\xd0\xea\x37\xe1\xc4\xe5\xb6\x04\x44\xcd\xb3\xc4\xd6\x1d\x85\xc2\xd0\x50\xd8\xb5\xac\xec\x18\x3b\x65\x82\xa8\xdc\x97\xea\x2a\x13\x56\x3b\x18\x45\x57\x10\x11\xc3\x62\x48\x25\xdd\xc5\xb1\x46\x4d\x77\x6a\x82\xa8\x6c\xcc\x98\x9c\x5e\x75\x75\xb8\xe3\x8d\x29\xa0\xd1\xe8\x40\xe3\x24\xf3\x45\xd6\x29\x77\xda\x6a\xd8\x0a\x7d\xe3\xa8\x85\xc6\xa2\x74\x77\xb4\xda\xbd\x80\x66\x73\x61\x1d\x91\x75\x00\x2f\xef\xa6\x7d\xdb\xa4\x27\x83\x26\x95\x49\x3b\x58\xd4\x42\x83\x63\x5e\x95\x45\xa1\x8d\xf3\x67\x62\xd2\x2e\x37\x1f\xdc\x4c\x7f\x57\xdc\x32\x38\x43\x62\xeb\x5f\xbd\x23\xc1\x0e\x65\xe6\x1b\x0f\xed\xba\x2b\x8a\x7a\xff\x55\x60\x00\xb3\x20\xc6\x1b\xaa\x20\x53\x5e\x9b\x07\x4e\xf3\xd0\x03\xf0\xd5\x48\xdf\xa4\xdc\xe1\x7d\x4f\x5c\x65\x1c\x2f\x90\x66\x4d\x61\xdf\x13\x3c\x95\x92\x23\x11\x35\x10\x36\xde\x37\xc8\xb3\x55\xd4\x69\x07\x47\x47\xb5\x79\xb0\xa6\x68\x15\x3f\x10\x47\xbb\x2e\x39\x6d\xd2\xa6\x71\x4c\xab\x78\xb0\xa7\x23\x6c\x57\xf2\x79\x9a\xdd\x58\x80\xdd\x49\xc1\xcf\xe3\x0f\x16\xb2\xd2\xfd\x98\xa7\x99\xc3\x6f\xf3\xd3\x46\x1d\xea\xc0\xa0\x05\x19\xed\x1b\xe9\x97\x2d\x89\x66\xda\x29\x0c\xae\x58\xda\x36\xdc\x51\x43\x76\xc9\xd7\x9d\x46\xb0\xf4\x21\x01\xcd\x2c\x98\xa2\x63\xa2\x4c\xfe\xe7\x20\x19\x60\x95\x4a\x65\xda\x31\xcd\xa4\x4a\x88\x19\x01\x13\xe6\xfb\xbf\x75\x58\x93\x09\x83\x73\x54\x4f\xc6\xe7\xa1\xdf\xf4\xfe\x72\xd8\x01\xb0\x90\x72\xd9\xc0\xfb\x3e\xf9\xd9\x4e\x01\xec\xd8\x46\xd1\x3f\x7f\xf7\xcf\xfd\x1c\x32\x4b\x57\x7a\xbf\x99\x69\x36\xe5\x2c\xde\x77\x5d\xbd\x64\xe9\x99\x14\x9e\x51\xfb\x44\x84\x8e\x8c\xab\xf7\x87\x6d\x71\x99\x09\xc2\xd9\x1f\xa8\x76\x45\xe6\x37\x25\x60\xa8\x69\x65\x4a\xbe\x64\xe8\xce\x32\xad\x9f\xf4\xa7\x15\x3e\x38\x27\x99\x76\xfd\x74\x4c\x52\x93\xd7\x77\x0a\x53\x54\x09\x11\x28\x0c\xcf\x41\x61\x22\x57\x58\x74\xfc\x5d\x3b\x5f\x1b\xa9\xc8\x1c\xb7\xbd\x6e\x23\x93\xea\x37\x6b\x13\xb2\xa2\x00\x12\xee\x7f\x8a\xc2\xb0\x59\xee\x2b\xe7\x92\x7a\xa0\xcd\xf5\x5e\xd1\x2c\xe3\x6c\x86\x71\x1e\x73\x8c\xf6\x6b\x3a\xd6\xc9\x66\x99\x4d\x91\xa3\x79\xc1\xae\x67\x42\xe2\x85\x0d\x7c\xa3\x3d\x59\x1d\x32\xb6\xf7\x1e\x4d\xc1\xeb\xc4\x25\x4f\x05\x72\x1f\x5c\x1d\x19\x20\x67\x6d\x4c\xae\x63\xed\x2e\xdf\xdf\xec\xf4\x77\x1a\x0d\x27\x53\xe4\x8d\x36\xd7\xaf\xf0\xec\x10\x5b\x76\x3a\xe1\x94\x68\x3d\x5e\x28\xa2\x1b\xa3\x7f\x11\x77\xa6\xb9\xc1\x7d\xa9\xb6\xab\xdc\x4a\x45\xf7\x66\x5b\x5b\x78\xec\x12\x18\x77\x87\xc4\x54\xb1\x15\x31\xf8\x16\xf3\xa7\x64\x84\x21\x6d\x5d\xcb\x8a\x9a\x5f\xce\xdc\x51\x1a\x9b\x31\xa4\xc7\xde\x9d\x48\x8a\x07\x3a\xe0\x68\x2a\xe9\x76\x14\x74\x5b\x77\x11\x2c\x52\x7f\xaa\x79\x63\xf1\x3a\x57\x6b\x0c\xb1\x45\x88\xf5\x9a\x0b\xe2\xcd\xeb\x15\xce\x66\x18\x9b\x57\x2d\x89\x87\x14\x40\x44\x0e\xa9\xa4\xde\x27\x53\x89\x1a\x84\x34\x60\x24\x47\x45\x0c\x3a\x44\x6e\x95\xe8\x81\x69\x97\xdf\x4c\x7b\xbe\x54\xa1\xb4\x68\x62\x46\x8e\x66\x3f\xdd\x9f\xc9\xa3\xe7\xa7\xdd\x7d\x2a\xa9\x6e\x45\x09\x05\x61\xdb\x64\x39\x24\x11\x7c\x24\x9c\xd1\x80\xdf\xf7\x76\xae\x64\x51\xd0\x1d\xef\xc0\x3d\x56\x38\x43\xb5\x86\x77\x6d\xbd\x2b\x79\x71\x87\x71\x66\x30\x7a\x8c\x34\x73\x89\xf9\xde\x4c\xf3\x6c\x5a\x62\x6e\xd5\x62\x8a\x40\xd2\x94\xb3\x5d\x47\x04\xce\xa9\x39\x1d\x7b\x94\xfd\x1b\x96\xe0\x29\xa5\x6d\x49\xeb\xb6\x92\x17\x73\x36\x4e\xf6\xbd\xc8\x58\x82\x40\x0c\xdc\x2e\x58\xbc\xd8\xd1\x67\x2b\x6d\xd7\xdd\x7f\x21\x16\x5d\x04\x97\xce\x5a\xa4\xe0\x39\xdc\x2a\x66\x0c\xfa\x53\xb9\x52\x64\x3b\x2c\xb5\xea\x55\x28\x31\x38\xac\x5c\xbd\x79\x48\x39\x61\x53\xa3\x3e\x3c\x2a\xe5\xeb\x2f\x4f\xc4\x52\x29\xd4\xa9\x4d\x1f\xc5\xbc\xb8\x5e\xe2\x00\x76\xf0\x68\x89\x79\xf4\x3c\x25\xa0\xb7\xb1\x16\x80\x25\xe6\x0f\xaa\x11\x77\xb4\xb4\x32\x6d\xcb\x82\x04\xf7\x0c\x69\x6d\x24\x0e\xa1\xb6\xec\x1a\x42\x43\xbd\x35\x2c\x37\x33\xd8\xab\x7d\x56\x47\xa4\x40\x73\x2b\xd5\xf2\x65\x32\xc4\xb0\xf8\x39\xae\x58\xdc\xef\x1a\x42\x98\x59\xdf\x89\xa9\xa8\xfc\xd5\x1a\xb2\xd2\x89\x09\x18\x76\x77\x63\x5a\x76\x61\x79\x42\x0c\x13\xf3\x49\xae\x0d\x26\x3b\x76\xf2\x6b\x15\xba\xb2\x9b\x12\x93\xdd\x8a\x76\xe3\xbd\x76\x92\x92\x4c\xef\xe2\xc4\xd8\xc2\x0c\xfa\xd4\x97\x6d\x41\xba\x2e\x55\x0f\x5a\x92\x57\xba\xab\xc4\xf8\xab\x4b\xfe\xa2\x93\x75\x82\x8d\xc9\xfa\x03\x3a\xab\x09\xb9\x2b\x6e\x1c\xf9\x1b\x25\x57\x59\x52\x6f\xb2\xbb\xd3\xc8\x5d\x49\x64\x42\xee\xae\x24\xc5\xb1\xa4\x4f\xb8\x88\x5c\xa1\xd2\x92\xd3\x6b\xcb\xaf\x97\x6e\x8f\xb4\x0c\xfa\xee\xc5\xc6\x21\xc6\xc6\x75\xe2\x4e\xf9\xea\x9e\xf5\xad\xca\x84\x8d\xa2\x1d\x3c\xc0\xf5\x1a\xb2\x62\x73\x01\x83\xbb\x83\x26\x85\x8d\x7a\x3d\xcd\x5f\x87\x04\xee\x25\xaf\x16\x85\xfb\x53\xf5\x77\xd1\xb6\x4e\xa2\x02\xa4\x4d\x67\xd6\x77\x05\x0c\x10\xd0\x98\x12\x9b\xe3\x52\x70\xe3\x36\xbd\xd9\xb8\x9d\x55\x97\xd1\x32\x73\xa0\xd7\xc7\xce\x70\xcb\xcc\x02\xde\xd7\x18\x61\x2f\x8e\x1a\x14\x44\x98\xcb\xf3\x5e\xb1\xc0\xd4\xaa\x40\xcb\x84\x2c\x9d\x2b\x42\x77\xa9\xcd\x07\x0f\x55\xdc\xda\x2b\x66\x59\x3f\x14\xa3\xd6\xbd\x1d\x94\xa4\x5d\x0e\x0d\x8a\x55\x2d\xf8\xb1\x55\x57\x92\xf1\x22\x5b\x63\x1a\x4e\x33\x23\xf7\x3a\x30\xd0\xc6\x8a\x77\x9e\xf7\xd8\x42\x31\x25\x5c\x5b\x88\xf6\x3c\x70\x48\xc8\xdd\x07\xa1\x90\xd0\x96\x92\x84\x88\xfc\xd7\x59\x5b\x16\xd8\xad\xd5\x3d\xdc\xe3\x84\xfe\xc6\xf5\x94\xee\x58\x92\x25\x20\xb2\x64\x8a\xca\x3a\x04\x5b\x22\xfa\x22\x30\x26\xc2\x5d\xe0\xf5\x24\xb4\xdd\x66\x70\x7d\x3e\xe7\x59\x82\xaa\x44\xf0\xfa\x1b\x48\x90\x08\x0d\x84\x73\x8f\x53\xa0\x37\xbb\x29\x82\x43\x08\x64\x66\x5a\x28\x02\xc0\x15\xf3\xf7\x73\xbe\x7b\x5d\x62\x63\x73\x21\x15\x16\x55\xb8\x2e\x36\x17\x2a\x9c\x84\xe4\x30\x6d\x2b\x2f\x9c\xe9\x33\x01\x52\x20\xd8\x1c\x08\x95\xab\xdc\x8e\xed\x80\x6f\x11\xc4\x44\xe1\x2c\xe3\x3c\xff\x9f\xaa\x0a\xb6\x20\x65\x1a\x5e\x7f\xd3\x5c\x13\xdc\x0d\xd7\xef\x5d\x0c\x99\x30\x43\xa9\x86\x5e\x4c\x23\x30\x2a\xc3\x47\x8d\x49\xab\xba\xfb\xfc\xad\x46\x52\x9f\xa6\x0f\x4b\x87\x34\xd8\x56\xb3\x7b\x8f\xaa\x6f\x8c\xb4\x6e\xd1\xbf\xf7\xd1\xe1\xee\xb4\x83\xdb\x2c\x67\x37\x33\x28\x32\x95\x99\xbf\x90\xee\xf1\x6d\xa7\x54\xa4\x3e\x99\x6a\xb9\x5c\x4d\xa9\x42\xad\x77\xa6\x7b\xef\x42\x67\xbe\x84\xb7\xba\x1c\x2f\xc8\x94\x63\x51\x44\x36\xa6\x71\xbd\x7a\xbf\xa7\x7e\x01\xe7\x83\x09\x13\x55\x06\x14\x97\x21\xc2\x52\x07\xba\x29\x49\x52\xb5\xee\x7a\xb7\xef\x5a\xc8\xe6\x83\xf0\xc6\x57\x99\x5a\xd6\xfb\x93\x74\x40\x4d\xe3\xf1\x75\xc3\x1b\x36\x81\x24\x37\xf1\xd8\x79\x0d\x39\x83\xb1\xcb\xfb\x8e\xcb\xfb\x68\xe5\xeb\x58\x35\xa9\xac\x82\x4b\x51\x40\x45\x4f\x51\x47\x5b\x49\xf5\xab\xa4\xb7\x2c\xf8\x41\x55\xf4\xfe\x6f\x04\x24\xa9\x14\x58\xdb\x2a\xee\x65\x28\x67\x05\xa2\x4a\xe9\xb5\x8e\x67\xb1\x4c\x19\xfa\x5b\xb1\x24\x5e\xd4\x9f\x93\x94\x28\x42\xbf\xb4\xd0\x68\x7f\xe2\xb2\x8f\x01\x29\x4c\x39\x8b\x89\xee\xa3\x6d\x25\x25\xd7\x61\x72\x3d\x45\x8d\xca\x56\xa5\x74\xfd\x9a\x89\xfb\xb5\x07\x8d\x5d\x9b\xd5\x64\x45\x18\xb7\x1e\x70\x34\x78\xf8\xd9\x7e\xb7\x74\x27\xce\x94\x42\x61\x9e\x6f\xc1\xf0\xee\xce\xf3\x2d\x18\x5e\x9b\x7a\xae\x05\x77\x5f\x39\x2a\xa5\xdc\x08\x11\x84\xd2\x7c\x69\xc9\xf3\xb0\x71\x3c\x90\xfc\xa0\xcb\x4b\x4f\xe1\x62\x0b\x5b\x7e\x0e\x6f\xda\x72\x9d\xa1\xa7\x57\x0c\x88\xd6\x09\x04\x45\x43\x18\x5f\xdf\xa4\x0c\xe2\x5a\xaf\x59\xcb\xbb\xe2\xad\xce\x07\x9c\x25\x73\xa2\xcd\x58\xc9\x29\xde\xb0\xa4\x5b\xf8\x7d\x47\xb4\xf1\x87\x17\xb7\xae\xea\x9c\xda\xca\x61\x81\xeb\xad\x46\x83\x87\x1d\x35\x74\x38\xc4\xd6\xe6\x46\x11\xa1\x59\xf1\x82\x6f\xcf\x8d\x57\xb6\x0b\xa6\x44\x85\xd4\x5f\xc6\xb0\xb9\x84\xcf\x5e\x9b\x5b\xf8\x12\x88\x90\x66\xd1\xec\x9d\x1f\x8d\xdc\x04\xb5\x26\xf3\x6e\x34\xfe\x92\x25\x44\x0c\x6d\xd9\xe5\xb2\xde\x30\x15\x98\xa0\xee\x85\x0a\x31\x2f\x35\xcd\xe5\xe9\x8d\xe4\x71\xc7\xab\x92\x31\x7b\xa7\x8d\x0a\x89\xae\xab\x78\xea\xea\x7b\xc1\xbe\x64\x3e\x91\x1b\xde\x4a\x45\x8f\xd7\x6f\xa4\x06\x34\x6b\xeb\x28\x64\x77\xa0\x9f\x9c\x82\xba\xaa\xa8\xa9\x95\xe5\x0b\x9e\x70\x23\xa4\x2c\x7f\xee\x59\x07\x9c\xf9\xd2\xfd\x46\x65\x2d\x67\xb2\x6f\x08\xd7\x78\x0c\x1f\xc4\x52\xc8\xdb\xfd\x77\xdf\x39\xa9\x76\x6d\xc8\xb0\xf3\xe2\xec\xa1\x93\x55\x3f\xc8\x7b\x37\x1a\xd9\x63\xfb\x6e\xf7\xa9\x86\x5e\x39\x30\x97\xf1\xb2\x6e\xdb\x6d\x07\x13\x8d\x96\x5a\x61\xf5\x29\x2c\xac\x95\x42\x67\x2b\x85\xdb\x45\xde\x7e\x2c\x61\x25\xc7\xc2\x87\x12\x5a\x24\xd6\x76\x72\x25\xa9\xeb\xc1\xbe\x27\x7a\x39\x61\x7f\xd4\xd0\xd0\x9e\xd1\xb4\xe5\x31\xf7\x71\x5f\x8e\x57\x3f\x3c\x31\xfe\xbf\x3f\x26\x7e\xf7\x25\x8c\x8e\x47\x4d\x16\xb4\xd2\xce\x77\x93\x37\x4e\xf2\xac\x7c\xb4\x51\x59\x6c\x64\xbf\xb6\x7e\x93\x2b\xbd\xa7\x5b\x53\xc5\x70\xb6\xe1\x3a\x1f\x57\xb9\x5c\xdd\xd2\x73\xdb\x73\xa6\x8d\xca\x2f\xc7\xcf\x71\xfa\x12\xbe\xaf\xd0\x4d\x58\xc5\x87\x76\x2a\xa5\x5c\x91\x79\x95\x29\x75\xf8\x64\x85\x6b\xc5\xd6\x3a\xb8\x80\xe4\x4b\x26\x0d\x69\xeb\x04\xf5\x6d\xcd\x13\xce\x65\x4c\x4c\x73\xf1\xd6\xe7\xa0\xad\xb5\x95\xdd\xad\x91\xdd\xa9\x8d\x9d\x12\x63\x50\x89\x11\xfc\xeb\xf0\xb7\x6f\xbf\x0e\x8f\x7e\x3a\x3c\xfc\xf4\x7a\xf8\x8f\xcf\xdf\x1e\xfe\x16\xb9\x7f\xfe\xff\xe8\xa7\xa3\xaf\xc5\x8f\x6f\x8f\x8e\x0e\x0f\x3f\xbd\x7d\xff\xf3\xcd\xf8\xe2\x33\x3b\xfa\xfa\x49\x64\xc9\xd2\xff\xfa\x7a\xf8\x09\x2f\x3e\x77\x44\x72\x74\xf4\xd3\xff\x0d\x1e\xb1\xc3\x5b\xb5\xa9\xb5\x1c\xee\xdf\xe1\x29\x3f\x25\xe2\xda\x8e\x1b\x1f\x25\x6a\xbc\xbc\x45\x14\x6e\xa8\x96\xd5\x90\x70\x5c\xc7\xc4\xbc\xfa\xd2\xe1\x19\x49\x49\xcc\x4c\x1e\xed\x73\x6d\x3c\xe8\x4e\x53\xcd\xf8\x97\xe6\x3c\x8b\xe6\x14\x0e\xc6\x35\xa6\xfd\x37\x62\xd0\xf5\x78\x0e\x4b\xaf\x21\x48\x82\xc7\xf0\x25\x23\xc2\x30\x93\x1f\x35\xf2\x86\x29\xbd\x97\x22\xc4\x41\x8b\xfe\xd2\x83\x17\xd4\x83\xc2\x94\xb7\xae\x00\x4a\x43\x78\x83\x13\x89\x1e\xf5\xe8\x49\xa3\x4d\x0b\x89\xca\xcf\xf6\x6c\x09\x97\x08\x26\x6d\x27\xfb\xad\x08\xf6\x9b\xd7\xff\xd0\xac\x96\x0d\x5b\x0f\xdd\x6b\xaa\x74\x43\x82\xe1\x0d\x8c\xcd\x27\xd9\xb4\x94\x4d\xb1\x7e\x28\x07\xe1\xdf\xff\x19\xac\x2b\x43\x12\xc7\x98\x1a\xa4\x57\xf7\xbf\x81\xf7\xea\x55\xe5\x23\x77\xee\xe7\x46\x1b\x09\x3e\x7d\x1e\xf8\x85\x91\x7e\x2c\x3e\x59\x67\x1f\xfe\x37\x00\x00\xff\xff\x50\x2a\x45\x96\x38\x50\x00\x00"),
		},
		"/devops.k8s.io_machines.yaml": &vfsgen۰CompressedFileInfo{
			name:             "devops.k8s.io_machines.yaml",
			modTime:          time.Date(2020, 12, 29, 6, 52, 42, 26971806, time.UTC),
			uncompressedSize: 10540,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x3a\x5b\x6f\x1b\xbb\xd1\xef\xfa\x15\x83\x7c\x0f\xfe\x0a\x58\xab\x93\x9e\x06\x2d\xf4\xe6\x3a\x49\x63\x24\xce\x11\x2c\x27\x2f\x45\x11\x8c\x96\x23\x2d\x8f\xb8\xe4\x86\x1c\xda\xd1\x29\xfa\xdf\x0b\x92\xbb\xab\xdb\xae\x2e\x8e\x0d\x54\x4f\x5a\x5e\xe6\x3e\xc3\x99\x21\x07\xc3\xe1\x70\x80\x95\xfc\x4a\xd6\x49\xa3\xc7\x80\x95\xa4\x1f\x4c\x3a\x7c\xb9\x6c\xf9\x37\x97\x49\x33\x7a\x78\x3d\x58\x4a\x2d\xc6\x70\xed\x1d\x9b\xf2\x8e\x9c\xf1\x36\xa7\xb7\x34\x97\x5a\xb2\x34\x7a\x50\x12\xa3\x40\xc6\xf1\x00\x00\xb5\x36\x8c\x61\xd8\x85\x4f\x80\xdc\x68\xb6\x46\x29\xb2\xc3\x05\xe9\x6c\xe9\x67\x34\xf3\x52\x09\xb2\x11\x78\x83\xfa\xe1\x97\xec\x2f\xd9\x2f\x03\x80\xdc\x52\xdc\x7e\x2f\x4b\x72\x8c\x65\x35\x06\xed\x95\x1a\x00\x68\x2c\x69\x0c\x25\xe6\x85\xd4\xe4\x32\x41\x0f\xa6\x6a\xa8\x1c\xb8\x8a\xf2\x80\x70\x61\x8d\xaf\xc6\xb0\x3d\x99\xf6\xd6\x04\x25\x66\x6e\x13\x98\x38\xa2\xa4\xe3\x8f\x9b\xa3\x9f\xa4\xe3\x38\x53\x29\x6f\x51\xad\x91\xc6\x41\x57\x18\xcb\x9f\xd7\x00\x87\x50\xe6\x69\x42\xea\x85\x57\x68\xdb\xf5\x03\x00\x97\x9b\x8a\xc6\x10\x97\x57\x98\x93\x18\x00\xd4\x4c\xc7\xed\x43\x40\x21\xa2\x18\x51\x4d\xac\xd4\x4c\xf6\xda\x28\x5f\xea\x16\xb8\x20\x97\x5b\x59\x71\x14\xd3\x7d\x41\x90\x2b\xcf\x64\xa1\x2a\xd0\x51\x16\x17\x01\xfc\xee\x8c\x9e\x20\x17\x63\xc8\x1c\x23\x7b\x97\xc5\xe9\x7a\x36\x49\x6e\xf2\xe1\x6a\xfa\xae\x1e\xe1\x55\xa0\xca\xb1\x95\x7a\xd1\x85\xe7\xe2\x7a\x57\x0d\x20\x1d\x20\x70\xfb\x69\xa9\xb2\xe4\x48\xb3\xd4\x0b\xe0\x82\xc0\x91\x7d\x20\x1b\x57\xd4\x48\x00\x1e\x0b\xd2\xc0\x85\x74\x60\x66\xbf\x53\xce\xf0\x88\x2e\x69\x98\x44\x06\x17\xfb\xc4\x37\xa6\x94\xed\x99\xc1\x16\x2b\x57\xff\xd8\x66\x44\x20\x27\xa4\x69\xfa\xe1\x75\xd2\x47\x5e\x50\x89\xe3\x7a\xa5\xa9\x48\x5f\x4d\x6e\xbe\xfe\x3a\xdd\x1a\x86\x6d\xc6\x6b\x0b\x08\xdc\x06\xa6\xd2\x5a\x98\x1b\x1b\x3f\x9b\xd9\xab\xc9\x4d\xbb\xbd\xb2\xa6\x22\xcb\xb2\x31\x87\xf4\xdb\x70\xab\x8d\xd1\x1d\x64\x17\x81\x9e\xb4\x0a\x44\xf0\x27\x4a\x58\x6b\x03\x21\x51\xb3\x00\x66\x9e\xa4\xd8\x0a\x3d\xca\x66\x0b\x30\x84\x45\xa8\x6b\x41\x67\x30\x8d\xea\x70\xc1\x5a\xbd\x12\xc1\x0d\x1f\xc8\x32\x58\xca\xcd\x42\xcb\x3f\x5a\xd8\x0e\xd8\x44\xa4\x0a\x99\x6a\xb3\x5f\xff\xa2\x41\x6a\x54\xf0\x80\xca\xd3\x25\xa0\x16\x50\xe2\x0a\x2c\x45\x75\x7a\xbd\x01\x2f\x2e\x71\x19\xdc\x1a\x4b\x20\xf5\xdc\x8c\xa1\x60\xae\xdc\x78\x34\x5a\x48\x6e\xc2\x49\x6e\xca\xd2\x6b\xc9\xab\x51\x8c\x0c\x72\xe6\xd9\x58\x37\x12\xf4\x40\x6a\xe4\xe4\x62\x88\x36\x2f\x24\x53\xce\xde\xd2\x08\x2b\x39\x8c\xa4\xeb\x18\x52\xb2\x52\xfc\x9f\xad\x03\x90\xbb\xd8\xa2\x75\xcf\xa2\xd3\x2f\x7a\xfb\x01\x0d\x04\xbf\x4f\xa6\x9d\xb6\x26\x2e\xf6\xad\xfb\xee\xdd\xf4\x1e\x1a\xd4\x51\x19\xbb\xd2\x4f\x06\xde\x6e\x74\x6b\x15\x04\x81\x49\x3d\x0f\xce\x11\x94\x38\xb7\xa6\x8c\x30\x49\x8b\xca\x48\xcd\xf1\x23\x57\x92\xf4\xae\xf8\x9d\x9f\x95\x92\x83\xde\xbf\x7b\x72\x1c\x74\x95\xc1\x75\x8c\xb1\x30\x23\xf0\x95\x48\x9e\x74\xa3\xe1\x1a\x4b\x52\xd7\x21\x24\xbc\xb4\x02\x82\xa4\xdd\x30\x08\xf6\x34\x15\x6c\x1e\x0f\xbb\x8b\x93\xd4\x36\x26\x9a\x38\xde\xa3\xaf\xda\x01\xa7\x15\xe5\x49\x6b\x1b\xb3\xc1\x01\xea\xc0\x9b\x6d\x41\xe8\xf6\xd0\x78\x38\x29\xef\x98\x6c\x88\xce\xbb\x53\xbd\xec\x84\xdf\x9c\x30\x48\x67\x7f\x4f\x3f\xaa\xb8\x4d\xaa\xee\x09\x00\xc9\x54\xf6\x4c\x1d\x83\x5a\x8b\xc9\x71\xff\xe4\x41\x66\x36\x84\x6f\xf3\x9f\x84\x11\x0c\x55\x5a\x12\x7d\x60\x86\x81\xce\xde\x39\x67\xf3\xc1\x21\xd4\x7b\xd6\xb2\xbb\x00\xad\xc5\x55\xc7\x7c\x61\xcc\xb2\x47\x76\x9b\xc7\xef\x31\x29\x1f\x15\xc0\x11\x32\xdd\x52\x56\xd7\x46\x27\x84\x4f\x31\x84\x13\x09\xe8\x16\xc3\x01\xe2\xe6\x52\xa3\x92\x7f\x90\xed\xc0\xbc\xe5\x7f\xef\xdb\x85\xd1\xfd\x34\x98\x0a\xbf\x7b\x8a\x29\x54\xf0\xbf\x74\x06\x00\x17\xc8\x50\x7a\x17\xa3\x14\x95\x15\x77\x29\x85\x0d\x54\x64\x4b\xd4\xa4\x59\x85\x23\xa5\x34\x0f\xd4\xc4\xd1\x18\x24\x1d\x1b\x8b\x8b\x1d\x6f\x3e\x28\xa4\x6e\x62\x83\x7f\x37\x27\xba\x8e\xff\x45\x88\x67\xf3\x55\x88\xee\xb8\xe6\x1e\x84\xef\x95\x6c\x1d\x2a\x40\xc9\x39\xe5\xab\x5c\x75\x50\x75\x44\x3f\xfd\xba\xa9\xa3\xd6\x11\xd9\x5f\x27\x0a\x76\x32\x94\x12\x23\x59\x35\x88\x94\x46\xc8\x26\x1c\xd6\x44\x67\x67\xc6\x29\x59\x8d\x07\x4f\x30\x3f\x85\x33\x52\xff\x03\x6e\x56\xa1\x73\x93\xc2\xa2\xa3\x6e\x0c\x73\x63\x4b\xe4\x31\xcc\x56\x4c\x4f\xe1\x33\xc0\x7f\x34\x56\x3c\x49\x48\x95\xb1\x7c\x98\x2c\xa9\xf9\xd7\x3f\x1f\x00\x1d\x72\xb2\x05\xd9\x2e\xd8\x56\x3e\x20\xd3\x47\x5a\xbd\x0c\xe3\x8c\x52\x73\x8f\xda\xb6\x4c\xf5\x66\x1e\x0f\x72\x39\x97\x24\x2e\x93\xdb\x19\x41\x17\xae\x86\x90\x9d\x1f\xf9\xf6\xaa\xa0\x00\x30\xe5\x53\xf7\x01\x66\x0c\x47\xcc\x98\x17\x24\x42\x64\x29\x30\xb9\xc7\x2b\x9a\xcf\x29\xe7\x57\xbd\xc7\x9a\xd1\x80\x7a\x05\x95\x11\x29\x6a\x09\x43\x0e\x42\x7e\xc5\x46\x91\x45\xa6\x08\x26\xe2\xc8\x7e\xe2\x78\x4e\x64\x1c\x3a\x5d\xb7\x38\xbc\xab\xcf\xd1\x2c\xf2\x9a\x36\xa7\x2a\x80\x92\x0c\x03\xdd\x95\x11\x29\xd4\x1e\x80\x0a\x20\xcc\x3e\x3b\x11\x44\x06\x5f\x51\x49\x51\x43\x77\x80\x96\xe0\xb3\x09\x15\x8f\xf0\x8a\x2e\x0f\x02\x9d\x58\x9a\x93\x5d\xaf\x8e\x85\xc1\x67\xf3\xee\x07\xe5\x9e\x29\xfb\xd9\x44\x64\xd9\x67\xc1\x47\x45\x95\x84\xb3\xa4\x55\x30\x82\x19\x01\x56\x95\x92\xc9\x24\xf0\x20\x47\xc1\x9e\x7e\x9a\xee\x50\xfc\x5e\x09\xd1\x9f\xff\xec\x9b\x72\xb3\x63\xa3\x72\x48\x2a\x92\x25\x01\x32\x3c\x16\x32\x2f\xc2\xc8\x41\xea\x13\xdb\xa1\xba\xc6\x00\x2c\x83\x9b\xe8\x11\x46\xab\x15\x3c\x5a\xc9\x4c\x3a\x16\xb1\xad\x8a\x0e\x7a\xe2\x76\xb4\x08\x35\xc6\x70\xab\xac\x7f\xa2\x74\x62\x72\x70\xba\x64\x5a\x6d\xa6\x92\x2c\x37\xd6\x92\xab\x42\xfa\x14\x6a\x32\xb3\x36\xe4\x83\x92\x59\xd2\x2a\x7b\xe9\x9c\x36\x79\x50\xef\xf4\x92\x56\x2f\x93\xd6\x7a\x17\x8a\xf3\x92\x9e\x70\x10\xf5\x33\x35\x04\x59\x75\x0c\x86\x73\xab\x63\xb8\x21\xe1\x9c\x6c\xb3\x42\xef\x7a\xeb\xad\x99\x31\x8a\x70\xb7\xb7\xc1\xa4\x51\xf3\xcd\xdb\xb3\xaa\xb4\x38\x75\xfa\x86\x6e\x91\x0c\x37\x8b\xc4\x9d\x99\x00\xea\xa4\xa2\x36\xb6\xe4\x4e\x28\x6b\xe3\xba\xcd\x48\x10\xaa\xf8\xe0\x85\x21\x9f\xc3\x99\xf1\xa9\x57\x90\xe0\x81\x99\xef\x30\x87\xfa\xdc\x02\x18\x85\xb0\xe4\x1c\x1d\xcb\xfb\x3f\xd5\xf9\x7d\xbb\x1e\x2c\x61\x5e\xe0\x4c\x51\xe3\x8a\x9d\x98\x4f\x4f\xd6\x6b\x11\x5c\x25\x04\xb1\x6d\x8c\x52\x6f\x4b\xa0\x69\xc3\xd5\xa8\x2e\x5c\x5f\xaa\x19\x40\x64\x83\xf3\x4f\xea\x7a\xeb\xc9\x49\x48\x93\x75\x1f\x40\x79\x7a\x42\x7b\x0a\xd2\xdb\x6d\x84\x71\xe3\x25\x18\x4d\x41\x39\x13\x3f\x53\x32\xbf\x84\x77\x3f\x52\xd3\xee\x66\xd2\x9f\xf5\x58\xb8\xd1\xcd\xaa\x27\x92\x7d\x28\x2e\x0e\x1b\x0a\x3b\xe7\xf6\xfc\xe6\x84\x68\xd8\x1f\x09\xf3\x03\x15\xf5\x59\xb6\xd7\x96\xe6\x6b\xeb\x13\xc4\x28\x95\x6b\x2d\x2f\xf7\xd6\x92\xe6\x35\xce\x4e\xd1\x35\xed\xda\xdb\x3e\x97\x38\x6e\x89\x0a\x1d\x4f\xac\x99\x51\x48\x10\x4e\x32\x8d\x4f\xe8\x38\x65\x0d\x8f\x14\xc0\xcf\x42\xd6\x13\x48\x6e\x48\xed\x53\xf3\xa9\xe7\xfc\x51\x2b\x0e\x34\xdf\x5b\xd4\x4e\x36\x9d\xfb\x33\x09\xdf\x22\x17\xb8\x05\x45\x22\xf5\x03\x82\x9d\xa7\xd8\xd7\x9f\x81\x19\x40\x6d\xb8\xe8\x2a\x7a\x9f\x99\xdd\x92\x9c\xc3\xc5\x69\x3c\x7e\xf0\x25\xea\xa1\x25\x14\x31\x64\xd6\x5b\x41\x6a\x21\x73\x8c\x4d\xe6\xc6\xd2\x62\x94\xef\x65\x4f\x45\x59\xb5\x82\x79\x72\xc0\xb1\x84\x6e\xf7\x6a\xa2\x87\xf4\x2f\x5a\x7e\xf7\x29\xc8\x0c\x43\xd5\x7b\xb9\x6e\x35\xd7\x60\xd6\xde\xd1\xe8\xee\xc2\xbd\x38\x07\x5d\x67\x6a\x0f\x07\xf5\xb1\x5a\x37\x4c\xda\xc3\x73\xc7\x3b\xe0\x1a\x75\xa8\x18\xee\xad\x3f\x50\xfc\xbc\x47\xe5\xe8\x12\xbe\xe8\xa5\x36\x8f\xfa\xe5\x03\xfe\xfd\xaa\x6a\x5b\x3d\x61\xd3\x3e\xdd\x2f\x11\xbc\x7b\x9d\xec\xb9\x63\xb7\x32\xf9\xb2\x8b\x88\x43\xb9\x60\x7d\xe8\xde\xe8\xb9\x39\x92\xb5\x4c\x29\x26\x2d\x52\xb8\x91\xf7\x52\xc4\xab\x2e\x1f\xcd\x59\xad\xda\x1e\x60\xdb\x9e\x38\xb7\x4b\xb6\x79\x4f\x72\x42\x4f\x24\xe4\x0b\x57\x1b\x5b\x42\x9a\x67\x2c\x93\x80\xd9\x9a\x86\xa7\x74\x65\x66\xc6\x74\x66\xc6\x7b\x14\xfc\xdd\x18\x86\x9b\xb7\x9d\x88\xb3\xa7\x60\xae\x8f\x49\xb2\x77\x5e\x87\x48\xda\x79\xe3\xd9\xdd\xcb\xdc\xd9\x09\xcd\x35\xe8\xb3\xd1\xb6\x24\xab\x49\x9d\x4e\xd1\xc7\xb8\xfe\x05\xe8\xf0\x33\x9a\x58\xf3\x63\x75\x06\x29\xcd\x96\x97\xa1\x46\x11\x9f\x47\x8b\x22\x7e\x7e\x4a\x1a\x2f\x3e\xc5\x70\x2f\x6e\x9b\xc5\xdd\xf8\xe1\xbd\xb1\xb5\x63\x6f\x3c\xbd\xe8\xec\x31\x26\xa7\x8f\x87\xae\xd1\x20\x75\x7d\xf7\x9a\x7a\xfb\xe9\x7a\x56\x92\x8a\x57\xc2\x55\xec\x71\xc5\xce\xd2\x27\x42\xab\x7b\x40\x96\xc6\x52\x4a\x4f\x4a\xd4\xff\xff\xe6\x4f\x0d\x05\x43\x29\xd2\xfd\xeb\x78\x34\x2a\x51\xff\x35\x33\x76\x31\x52\x52\xfb\x1f\xe1\x73\x58\xe1\x82\x5c\xf8\xf7\x66\xb4\xde\x90\xbd\xc9\x0a\x2e\xd5\xc5\x53\x04\x1a\x42\x55\x4c\x25\xa6\x2b\xc7\x54\x9e\x18\x91\x7e\x6b\x76\x41\xda\xf6\x6c\x51\xc9\xb8\x9b\xb2\x37\x3b\xda\x22\xe3\xb7\x29\xc4\xa5\xcf\x67\x5b\x2e\xb2\xf2\xe5\xcb\x49\xc6\x35\x6d\x17\x3f\xb3\x71\xad\x8d\x76\xdb\x98\xee\xb7\xac\xac\xee\x93\xf7\x5e\x7c\x1a\xb8\x23\x01\x1f\x90\xa1\x30\x8e\x5d\x7b\xa3\x8f\x79\x1e\x2a\x4e\x4b\xa2\x40\xce\x72\x53\x8e\x84\xc9\x7d\xd9\xbc\x0d\x19\x91\x1e\x7e\x99\x8e\xee\x48\x7c\xfb\x80\xfc\x6d\xea\x67\x2d\xcb\xdf\x6e\x51\xe3\x82\xc2\xd2\xd1\xeb\x51\xb0\xb7\xd1\xdd\x87\xe9\xed\x68\x41\x1c\x0c\x61\x98\xa4\x37\x0c\x27\x66\xb4\xc6\xf3\x35\x70\x20\x19\xe8\x4d\x9a\xb7\x74\x72\x05\x45\x48\x98\xe1\xe4\x84\x19\x1e\x8b\xce\x2b\xc6\x8d\x1a\x5d\xba\xe4\xee\xd2\x1d\x4a\x9e\x0e\xf0\x15\x5f\x54\x1d\x21\xbc\xd6\xf9\x24\x2c\xdd\x7a\xd2\x13\x37\x6f\xbc\x50\x08\x34\x38\xb6\x3e\x67\x63\xcf\x21\xa2\x2f\x71\xdf\x11\xdf\xcc\x4a\x9a\x6f\x24\xea\xcf\x2b\xbf\x90\x1e\xd2\x19\xb2\xeb\xb4\x87\xbd\xc1\xf8\x80\x4c\x8c\x81\xad\x4f\x1e\x56\x5f\xff\x6e\x8e\xf8\x59\xfb\xfc\xa7\x91\x41\x5d\x08\xc0\xbf\xff\x33\x58\xd7\x04\xc1\x3b\x2a\x26\xf1\x79\xf7\xdd\xdf\xab\x57\x5b\x0f\xfb\xe2\xe7\x46\x03\x01\xfe\xf9\xaf\x41\x42\x4c\xe2\x6b\xf3\x4c\x2f\x0c\xfe\x37\x00\x00\xff\xff\x53\xc6\x77\x07\x2c\x29\x00\x00"),
		},
		"/workload.k8s.io_addons.yaml": &vfsgen۰CompressedFileInfo{
			name:             "workload.k8s.io_addons.yaml",
			modTime:          time.Date(2020, 12, 29, 6, 52, 42, 24913963, time.UTC),
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
