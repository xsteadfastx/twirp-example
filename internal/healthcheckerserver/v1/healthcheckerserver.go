package healthcheckerserver

import (
	"context"

	"github.com/twitchtv/twirp"
	v1 "go.xsfx.dev/twirp-example/api/go/healthchecker/v1"
	"go.xsfx.dev/twirp-example/internal/config"
)

type Server struct{}

func (s Server) Check(context.Context, *v1.CheckRequest) (*v1.CheckResponse, error) {
	return &v1.CheckResponse{}, nil
}

func Handler() v1.TwirpServer { //nolint:ireturn
	h := v1.NewCheckServiceServer(Server{}, twirp.WithServerPathPrefix(config.Cfg.APIPrefix))

	return h
}
