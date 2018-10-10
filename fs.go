package strastic

import (
	"net/http"
	"os"
	"strings"
)

type FS struct {
	ServeFS http.FileSystem
	IsSPA bool
}

func (fs FS) Open(path string) (http.File, error) {
	f, err := fs.ServeFS.Open(path)
	if os.IsNotExist(err) && path != "/index.html" && fs.IsSPA {
		return fs.Open("/index.html")
	} else if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.ServeFS.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
