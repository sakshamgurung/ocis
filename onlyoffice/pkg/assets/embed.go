// Code generated by fileb0x at "2021-08-05 14:44:09.305719 +0200 CEST m=+0.018184459" from config file "embed.yml" DO NOT EDIT.
// modification hash(d966a8f90918312d5dd86f7544b38cc0.8058aec596c5fb73022d09bb97af796e)

package assets

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileOnlyofficeJs is "onlyoffice.js"
var FileOnlyofficeJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x90\xc1\x4a\x33\x31\x14\x46\xf7\xff\x53\x94\xd0\xc5\x0c\xa4\xed\x3e\xa5\xbf\x8b\x82\xd0\x85\x56\x44\x57\x22\x76\x4c\xbe\xd8\x60\xe6\xde\x30\xc9\x38\x53\x42\xde\x5d\x50\xda\xea\xc2\x7d\x77\x39\x10\x2e\xe7\x7c\x06\xd6\x11\xaa\xca\xf6\xa4\x93\x63\xaa\xea\x2c\xfa\x88\x49\x4c\x9d\xd3\x49\x2c\x3b\xa4\xbe\xa3\xdc\x84\xb0\x21\xcb\x2a\x53\xd3\x42\x89\x2d\xf9\xc3\xd6\x5a\xa7\x21\xa4\x33\x4a\x30\xf9\x03\x1f\x59\x33\x29\x31\xce\xbe\x79\x66\x58\xf7\x2d\x28\x09\x89\x31\x81\xa2\x63\x8a\xea\x29\x9f\x40\x09\xc3\x7a\x14\x72\xdf\x90\xf1\xe8\xd4\x49\xe4\xfc\x65\xcd\x64\xdd\x9b\x82\xb4\xce\xe3\xae\x49\x7b\x45\x5f\xcf\x8d\x51\x5c\xea\x3c\x38\x32\x3c\xcc\x39\x80\xaa\xdd\x34\x63\x1e\xd1\x7d\xa0\x2b\x8b\x26\x84\xb8\x38\xab\x2d\xa6\x99\xcb\xd5\xf1\xc6\x6a\x9a\x41\x9a\x0d\x1e\xef\x37\x6b\x6e\x03\x13\x28\x55\x54\x97\x9d\x14\x2f\xaf\xbe\xa1\x77\x51\x17\x49\x18\xae\x9d\xc7\x0d\xa8\x57\xb9\x05\xf5\x0f\x2e\x79\x28\xac\xfe\xa3\x12\xb7\x18\x26\xe7\x25\x26\xa7\xd2\xfa\xcf\x0d\x4a\x91\x3f\xcb\x47\x1f\x2f\xb5\xfc\xb7\x68\x08\xe9\x62\x45\x9f\x4b\x29\x75\xbd\xfc\xf7\x19\x00\x00\xff\xff\xf5\x4b\xd1\x2b\xca\x02\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileOnlyofficeJs)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "onlyoffice.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
