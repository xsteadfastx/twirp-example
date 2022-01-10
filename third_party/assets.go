package third_party

import (
	"embed"

	"go.xsfx.dev/twirp-example/tools"
)

//go:embed swagger-ui/*
var swaggerUI embed.FS

var SwaggerUI = tools.HTTPSub(swaggerUI, "swagger-ui")
