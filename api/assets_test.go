package api_test

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
	"go.xsfx.dev/twirp-example/api"
)

func TestSwaggerJSON(t *testing.T) {
	sj := api.SwaggerJSON()

	require.Equal(
		t,
		map[string]template.URL{
			"healthchecker_v1": "swagger/healthchecker/v1/healthchecker_v1.swagger.json",
		},
		sj,
	)
}
