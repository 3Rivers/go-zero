package redis

import (
	"crypto/tls"
	"io"

	red "github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/core/syncx"
)

var clusterManager = syncx.NewResourceManager()

func getCluster(server, pass string, tlsFlag bool) (*red.ClusterClient, error) {
	val, err := clusterManager.GetResource(server, func() (io.Closer, error) {
		store := red.NewClusterClient(&red.ClusterOptions{
			Addrs:        []string{server},
			Password:     pass,
			MaxRetries:   maxRetries,
			MinIdleConns: idleConns,
			TLSConfig:    &tls.Config{InsecureSkipVerify: tlsFlag},
		})
		store.WrapProcess(process)

		return store, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*red.ClusterClient), nil
}
