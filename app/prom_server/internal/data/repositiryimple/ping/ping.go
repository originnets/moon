package ping

import (
	"context"

	"prometheus-manager/app/prom_server/internal/biz/dobo"
	"prometheus-manager/app/prom_server/internal/biz/repository"
	"prometheus-manager/app/prom_server/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repository.PingRepo = (*pingRepoImpl)(nil)

type pingRepoImpl struct {
	repository.UnimplementedPingRepo
	data *data.Data
	log  *log.Helper
}

func (l *pingRepoImpl) Ping(ctx context.Context, g *dobo.Ping) (*dobo.Ping, error) {
	return &dobo.Ping{Hello: "world"}, nil
}

// NewPingRepo .
func NewPingRepo(data *data.Data, logger log.Logger) repository.PingRepo {
	return &pingRepoImpl{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/Ping")),
	}
}