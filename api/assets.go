//nolint:gochecknoglobals
package api

import (
	"embed"
	"html/template"
	"io/fs"
	"strings"

	"go.xsfx.dev/twirp-example/tools"
)

//go:embed swagger/healthchecker/v1/healthchecker_v1.swagger.json
var swagger embed.FS

// Swagger has the generated swagger files.
var Swagger = tools.HTTPSub(swagger, "swagger")

func SwaggerJSON() map[string]template.URL {
	js := make(map[string]template.URL)

	fs.WalkDir(swagger, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".swagger.json") {
			fileNames := strings.Split(path, ".swagger.json")
			names := strings.Split(fileNames[0], "/")
			js[names[len(names)-1]] = template.URL(path)
		}

		return nil
	})

	return js
}
