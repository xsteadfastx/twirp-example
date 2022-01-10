package tools

import (
	"io/fs"
	"net/http"

	"github.com/rs/zerolog/log"
)

func HTTPSub(f fs.FS, dir string) http.FileSystem {
	fsys, err := fs.Sub(f, dir)
	if err != nil {
		log.Error().Err(err).Str("dir", dir).Msg("could not sub into dir")

		return nil
	}

	return http.FS(fsys)
}
