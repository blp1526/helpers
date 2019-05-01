package ioutilh

import (
	"io/ioutil"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NopCloserKey = "NopCloser"

	ReadAllKey = "ReadAll"

	ReadDirKey = "ReadDir"

	ReadFileKey = "ReadFile"

	TempDirKey = "TempDir"

	TempFileKey = "TempFile"

	WriteFileKey = "WriteFile"
)

func New() hctx.Map {
	return hctx.Map{

		NopCloserKey: NopCloser,

		ReadAllKey: ReadAll,

		ReadDirKey: ReadDir,

		ReadFileKey: ReadFile,

		TempDirKey: TempDir,

		TempFileKey: TempFile,

		WriteFileKey: WriteFile,
	}
}

// NopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
var NopCloser = ioutil.NopCloser

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
var ReadAll = ioutil.ReadAll

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
var ReadDir = ioutil.ReadDir

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
var ReadFile = ioutil.ReadFile

// TempDir creates a new temporary directory in the directory dir
// with a name beginning with prefix and returns the path of the
// new directory. If dir is the empty string, TempDir uses the
// default directory for temporary files (see os.TempDir).
// Multiple programs calling TempDir simultaneously
// will not choose the same directory. It is the caller&#39;s responsibility
// to remove the directory when no longer needed.
var TempDir = ioutil.TempDir

// TempFile creates a new temporary file in the directory dir
// with a name beginning with prefix, opens the file for reading
// and writing, and returns the resulting *os.File.
// If dir is the empty string, TempFile uses the default directory
// for temporary files (see os.TempDir).
// Multiple programs calling TempFile simultaneously
// will not choose the same file. The caller can use f.Name()
// to find the pathname of the file. It is the caller&#39;s responsibility
// to remove the file when no longer needed.
var TempFile = ioutil.TempFile

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
var WriteFile = ioutil.WriteFile
