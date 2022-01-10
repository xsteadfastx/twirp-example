package main

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
	"go.xsfx.dev/twirp-example/api"
	healthcheckerserverv1 "go.xsfx.dev/twirp-example/internal/healthcheckerserver/v1"
	"go.xsfx.dev/twirp-example/third_party"
	"go.xsfx.dev/twirp-example/web"
)

func main() {
	mux := http.NewServeMux()

	hcHandlerV1 := healthcheckerserverv1.Handler()
	mux.Handle(hcHandlerV1.PathPrefix(), hcHandlerV1)

	// Swagger.
	mux.HandleFunc(
		"/swagger-ui/",
		func(w http.ResponseWriter, r *http.Request) {
			t, err := template.ParseFS(web.Templates, "template/swagger/index.html.tmpl")
			if err != nil {
				log.Error().Err(err).Msg("could not parse template")
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			if err := t.Execute(w, struct{ Swaggers map[string]template.URL }{api.SwaggerJSON()}); err != nil {
				log.Error().Err(err).Msg("could not execute template")
				w.WriteHeader(http.StatusInternalServerError)

				return
			}
		},
	)
	mux.Handle("/swagger-ui/vendor/", http.StripPrefix("/swagger-ui/vendor", http.FileServer(third_party.SwaggerUI)))
	mux.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(api.Swagger)))

	log.Info().Msg("starting server")

	log.Fatal().Err(http.ListenAndServe(":8080", mux)).Msg("goodbye")
}
